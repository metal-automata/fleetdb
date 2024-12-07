package fleetdbapi

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

type InstalledFirmware struct {
	ID                  uuid.UUID `json:"id,omitempty"`
	ServerComponentID   uuid.UUID `json:"server_component_id,omitempty"`
	ServerComponentName string    `json:"server_component_name,omitempty"`
	Version             string    `json:"version" binding:"required"`
	Current             bool      `json:"current"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (t *InstalledFirmware) Equals(b *InstalledFirmware) bool {
	if t == nil {
		return false
	}

	return t.Version == b.Version
}

func (t *InstalledFirmware) fromDBModel(dbT *models.InstalledFirmware) {
	t.ID = uuid.MustParse(dbT.ID)
	t.Version = dbT.Version
	t.CreatedAt = dbT.CreatedAt.Time
	t.UpdatedAt = dbT.UpdatedAt.Time
	t.ServerComponentID = uuid.MustParse(dbT.ServerComponentID)

	if dbT.R != nil {
		t.ServerComponentName = dbT.R.ServerComponent.Name.String
	}
}

func (t *InstalledFirmware) toDBModel(componentID string) (*models.InstalledFirmware, error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	if _, err := uuid.Parse(componentID); err != nil {
		return nil, errors.Wrap(ErrValidatePayload, "invalid componentID"+err.Error())
	}

	return &models.InstalledFirmware{
		ID:                t.ID.String(),
		ServerComponentID: componentID,
		Version:           t.Version,
	}, nil
}

func (r *Router) upsertInstalledFirmware(ctx context.Context, tx boil.ContextExecutor, componentID string, t InstalledFirmware) (string, error) {
	dbInstalledFirmware, err := t.toDBModel(componentID)
	if err != nil {
		return "", err
	}

	if err := dbInstalledFirmware.Upsert(
		ctx,
		tx,
		true, // update on conflict
		// conflict columns
		[]string{models.InstalledFirmwareColumns.ServerComponentID},
		// Columns to update when its an UPDATE
		boil.Whitelist(
			models.InstalledFirmwareColumns.Version,
		),
		// Columns to insert when its an INSERT
		boil.Infer(),
	); err != nil {
		return "", errors.Wrap(ErrDBQuery, err.Error())
	}

	return dbInstalledFirmware.ID, nil
}

// installedFirmwareSet will create a new record for the server component firmware,
//
// an existing record  is soft deleted for archival purposes and a new record is created
func (r *Router) installedFirmwareSet(c *gin.Context) {
	var t InstalledFirmware
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid InstalledFirmware payload", err)
		return
	}

	id, err := r.upsertInstalledFirmware(c.Request.Context(), r.DB, t.ServerComponentID.String(), t)
	if err != nil {
		if errors.Is(err, ErrDBQuery) {
			dbErrorResponse2(c, "installed firmware upsert error", err)
			return
		}

		badRequestResponse(c, "invalid InstalledFirmware payload", err)
		return
	}

	createdResponse(c, id)
}

// TODO:
// Once the filtering system is updated based on,
// https://github.com/metal-automata/docs/blob/main/design/data-store/001-fleetdb.md#api-query-filter
// add filtering on serverID or component names or firmware version
func (r *Router) installedFirmwareList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	mods := []qm.QueryMod{
		// join server components
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ServerComponents,
				models.ServerComponentTableColumns.ID,
				models.InstalledFirmwareTableColumns.ServerComponentID,
			),
		),
		// Load N-1 relationship in db model struct field R
		qm.Load(models.InstalledFirmwareRels.ServerComponent),
	}

	dbInstalledFirmwares, err := models.InstalledFirmwares(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count, err := models.InstalledFirmwares(mods...).Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	list := []InstalledFirmware{}
	for _, dbInstalledFirmware := range dbInstalledFirmwares {
		installedFirmware := InstalledFirmware{}
		installedFirmware.fromDBModel(dbInstalledFirmware)
		list = append(list, installedFirmware)
	}

	pd := paginationData{
		pageCount:  len(list),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, list, pd)
}

func (r *Router) installedFirmwareGet(c *gin.Context) {
	componentID := c.Param("componentID")

	componentUUID, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	mods := []qm.QueryMod{
		qm.Where(models.InstalledFirmwareColumns.ServerComponentID+"=?", componentUUID.String()),
		// join server components
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ServerComponents,
				models.ServerComponentTableColumns.ID,
				models.InstalledFirmwareTableColumns.ServerComponentID,
			),
		),
		// Load N-1 relationship in db model struct field R
		qm.Load(models.InstalledFirmwareRels.ServerComponent),
	}

	dbInstalledFirmware, err := models.InstalledFirmwares(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	installedFirmware := &InstalledFirmware{}
	installedFirmware.fromDBModel(dbInstalledFirmware)
	itemResponse(c, installedFirmware)
}

func (r *Router) installedFirmwareDelete(c *gin.Context) {
	componentID := c.Param("componentID")

	_, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	mod := qm.Where(models.InstalledFirmwareColumns.ServerComponentID+"=?", componentID)
	if _, err = models.InstalledFirmwares(mod).DeleteAll(c.Request.Context(), r.DB); err != nil {
		dbErrorResponse2(c, "Installed Firmware soft delete error", err)
		return
	}

	deletedResponse(c)
}
