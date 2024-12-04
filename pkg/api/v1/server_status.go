package fleetdbapi

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/metal-automata/fleetdb/internal/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ServerStatus struct {
	ID         uuid.UUID `json:"id"`
	ServerID   uuid.UUID `json:"server_id" binding:"required"`
	ServerName string    `json:"server_name"`
	Health     string    `json:"health" binding:"required"`
	State      string    `json:"state" binding:"required"`
	Info       string    `json:"info"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (t *ServerStatus) fromDBModel(dbT *models.ServerStatus) {
	t.ID = uuid.MustParse(dbT.ID)
	t.ServerID = uuid.MustParse(dbT.ServerID)
	t.Health = dbT.Health
	t.State = dbT.State
	t.Info = dbT.Info.String
	t.CreatedAt = dbT.CreatedAt.Time
	t.UpdatedAt = dbT.UpdatedAt.Time

	if dbT.R != nil {
		t.ServerName = dbT.R.Server.Name.String
	}
}

func (t *ServerStatus) toDBModel() *models.ServerStatus {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return &models.ServerStatus{
		ID:       t.ID.String(),
		ServerID: t.ServerID.String(),
		Health:   t.Health,
		State:    t.State,
		Info:     null.StringFrom(t.Info),
	}
}

func (r *Router) serverStatusSet(c *gin.Context) {
	var t ServerStatus
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid ServerStatus payload", err)
		return
	}

	// validate the server exists
	mod := qm.Where(models.ServerStatusColumns.ServerID+"=?", t.ServerID)
	exists, err := models.ServerStatuses(mod).Exists(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "Server Status lookup error", err)
		return
	}

	tx, err := r.DB.BeginTx(c.Request.Context(), nil)
	if err != nil {
		dbErrorResponse2(c, "Server Status set error", err)
		return
	}

	defer loggedRollback(r, tx)

	dbServerStatus := t.toDBModel()

	if exists {
		// update existing record
		_, err = models.ServerStatuses(mod).DeleteAll(c.Request.Context(), tx)
		if err != nil {
			dbErrorResponse2(c, "Server Status delete error", err)
			return
		}
	}

	if err := dbServerStatus.Insert(c.Request.Context(), tx, boil.Infer()); err != nil {
		dbErrorResponse2(c, "Server Status insert error", err)
		return
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse2(c, "Server Status set error", err)
		return
	}

	createdResponse(c, dbServerStatus.ID)
}

func (r *Router) serverStatusList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	mods := []qm.QueryMod{
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.Servers,
				models.ServerTableColumns.ID,
				models.ServerStatusTableColumns.ServerID,
			),
		),
		qm.Load(models.ServerStatusRels.Server),
	}

	dbServerStatuses, err := models.ServerStatuses(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count, err := models.ServerStatuses(mods...).Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	list := []ServerStatus{}
	for _, dbServerStatus := range dbServerStatuses {
		serverStatus := ServerStatus{}
		serverStatus.fromDBModel(dbServerStatus)
		list = append(list, serverStatus)
	}

	pd := paginationData{
		pageCount:  len(list),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, list, pd)
}

func (r *Router) serverStatusGet(c *gin.Context) {
	serverID := c.Param("serverID")

	serverUUID, err := uuid.Parse(serverID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid server UUID expected"))
		return
	}

	mods := []qm.QueryMod{
		qm.Where(models.ServerStatusColumns.ServerID+"=?", serverUUID.String()),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.Servers,
				models.ServerTableColumns.ID,
				models.ServerStatusTableColumns.ServerID,
			),
		),
		qm.Load(models.ServerStatusRels.Server),
	}

	dbServerStatus, err := models.ServerStatuses(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	serverStatus := &ServerStatus{}
	serverStatus.fromDBModel(dbServerStatus)
	itemResponse(c, serverStatus)
}

func (r *Router) serverStatusDelete(c *gin.Context) {
	serverID := c.Param("serverID")
	serverUUID, err := uuid.Parse(serverID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid server UUID expected"))
		return
	}

	mod := qm.Where(models.ServerStatusColumns.ServerID+"=?", serverUUID.String())
	_, err = models.ServerStatuses(mod).DeleteAll(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "Server Status delete error", err)
		return
	}

	deletedResponse(c)
}
