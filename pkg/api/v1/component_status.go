package fleetdbapi

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

type ComponentStatus struct {
	ID                uuid.UUID `json:"id,omitempty"`
	ServerComponentID uuid.UUID `json:"server_component_id,omitempty" binding:"required"`
	ComponentName     string    `json:"component_name,omitempty"`
	Health            string    `json:"health" binding:"required"`
	State             string    `json:"state" binding:"required"`
	Info              string    `json:"info"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// Equals compares two ComponentStatus
func (c *ComponentStatus) Equals(b *ComponentStatus) bool {
	if c.Health != b.Health {
		return false
	}

	if c.State != b.State {
		return false
	}

	if c.Info != b.Info {
		return false
	}

	return true
}

func (c *ComponentStatus) fromDBModel(dbT *models.ComponentStatus) {
	c.ID = uuid.MustParse(dbT.ID)
	c.ServerComponentID = uuid.MustParse(dbT.ServerComponentID)
	c.Health = dbT.Health
	c.State = dbT.State
	c.Info = dbT.Info.String
	c.CreatedAt = dbT.CreatedAt.Time
	c.UpdatedAt = dbT.UpdatedAt.Time

	if dbT.R != nil {
		c.ComponentName = dbT.R.ServerComponent.Name.String
	}
}

func (c *ComponentStatus) toDBModel(componentID string) (*models.ComponentStatus, error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	if componentID == "" && c.ServerComponentID != uuid.Nil {
		componentID = c.ServerComponentID.String()
	}

	if _, err := uuid.Parse(componentID); err != nil {
		return nil, errors.Wrap(ErrValidatePayload, "invalid componentID"+err.Error())
	}

	return &models.ComponentStatus{
		ID:                c.ID.String(),
		ServerComponentID: componentID,
		Health:            c.Health,
		State:             c.State,
		Info:              null.StringFrom(c.Info),
	}, nil
}

func (r *Router) upsertComponentStatus(ctx context.Context, tx boil.ContextExecutor, componentID string, t ComponentStatus) (string, error) {
	dbComponentStatus, err := t.toDBModel(componentID)
	if err != nil {
		return "", err
	}

	if err := dbComponentStatus.Upsert(
		ctx,
		tx,
		true, // update on conflict
		// conflict columns
		[]string{models.ComponentStatusColumns.ServerComponentID},
		// Columns to update when its an UPDATE
		boil.Whitelist(
			models.ComponentStatusColumns.Health,
			models.ComponentStatusColumns.State,
			models.ComponentStatusColumns.Info,
		),
		// Columns to insert when its an INSERT
		boil.Infer(),
	); err != nil {
		return "", errors.Wrap(ErrDBQuery, err.Error())
	}

	return dbComponentStatus.ID, nil
}

func (r *Router) componentStatusSet(c *gin.Context) {
	var t ComponentStatus
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid ComponentStatus payload", err)
		return
	}

	id, err := r.upsertComponentStatus(c.Request.Context(), r.DB, t.ServerComponentID.String(), t)
	if err != nil {
		if errors.Is(err, ErrDBQuery) {
			dbErrorResponse2(c, "component status upsert error", err)
			return
		}

		badRequestResponse(c, "invalid ComponentStatus payload", err)
		return
	}

	createdResponse(c, id)
}

func componentStatusQueryMods() []qm.QueryMod {
	return []qm.QueryMod{
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ServerComponents,
				models.ServerComponentTableColumns.ID,
				models.ComponentStatusTableColumns.ServerComponentID,
			),
		),
		qm.Load(models.ComponentStatusRels.ServerComponent),
	}
}

func (r *Router) componentStatusList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	mods := componentStatusQueryMods()
	dbComponentStatuses, err := models.ComponentStatuses(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "component status list query error", err)
		return
	}

	count, err := models.ComponentStatuses(mods...).Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "component status list count query error", err)
		return
	}

	list := []ComponentStatus{}
	for _, dbComponentStatus := range dbComponentStatuses {
		componentStatus := ComponentStatus{}
		componentStatus.fromDBModel(dbComponentStatus)
		list = append(list, componentStatus)
	}

	pd := paginationData{
		pageCount:  len(list),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, list, pd)
}

func (r *Router) componentStatusGet(c *gin.Context) {
	componentID := c.Param("componentID")

	componentUUID, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	mods := []qm.QueryMod{
		qm.Where(models.ComponentStatusColumns.ServerComponentID+"=?", componentUUID.String()),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ServerComponents,
				models.ServerComponentTableColumns.ID,
				models.ComponentStatusTableColumns.ServerComponentID,
			),
		),
		qm.Load(models.ComponentStatusRels.ServerComponent),
	}

	dbComponentStatus, err := models.ComponentStatuses(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "component status query error", err)
		return
	}

	componentStatus := &ComponentStatus{}
	componentStatus.fromDBModel(dbComponentStatus)
	itemResponse(c, componentStatus)
}

func (r *Router) componentStatusDelete(c *gin.Context) {
	componentID := c.Param("componentID")
	componentUUID, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	mod := qm.Where(models.ComponentStatusColumns.ServerComponentID+"=?", componentUUID.String())
	_, err = models.ComponentStatuses(mod).DeleteAll(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "Component Status delete error", err)
		return
	}

	deletedResponse(c)
}
