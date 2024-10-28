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
	ID                  uuid.UUID `json:"id"`
	ServerComponentID   uuid.UUID `json:"server_component_id" binding:"required"`
	ServerComponentName string    `json:"server_component_name"`
	Version             string    `json:"version" binding:"required"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
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

func (t *InstalledFirmware) toDBModel() *models.InstalledFirmware {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return &models.InstalledFirmware{
		ID:                t.ID.String(),
		ServerComponentID: t.ServerComponentID.String(),
		Version:           t.Version,
	}
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

	// validate the component exists
	mod := qm.Where(models.InstalledFirmwareColumns.ServerComponentID+"=?", t.ServerComponentID)
	exists, err := models.InstalledFirmwares(mod).Exists(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "Installed Firmware lookup error", err)
		return
	}

	tx, err := r.DB.BeginTx(c.Request.Context(), nil)
	if err != nil {
		dbErrorResponse2(c, "Installed Firmware set error", err)
		return
	}

	defer loggedRollback(r, tx)

	if exists {
		// soft delete existing record
		errDelete := r.softDeleteInstalledFirwmare(c.Request.Context(), tx, t.ServerComponentID)
		if errDelete != nil {
			dbErrorResponse2(c, "Installed Firmware soft delete error", errDelete)
			return
		}
	}

	// insert new record
	dbInstalledFirmware := t.toDBModel()
	errInsert := dbInstalledFirmware.Insert(c.Request.Context(), tx, boil.Infer())
	if errInsert != nil {
		dbErrorResponse2(c, "Installed Firmware delete error", errInsert)
		return
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse2(c, "Installed Firmware set error", err)
		return
	}

	createdResponse(c, dbInstalledFirmware.ID)
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
	componentUUID, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	err = r.softDeleteInstalledFirwmare(c.Request.Context(), r.DB, componentUUID)
	if err != nil {
		dbErrorResponse2(c, "Installed Firmware soft delete error", err)
		return
	}

	deletedResponse(c)
}

func (r *Router) softDeleteInstalledFirwmare(ctx context.Context, btx boil.ContextExecutor, componentID uuid.UUID) error {
	mod := qm.Where(models.InstalledFirmwareColumns.ServerComponentID+"=?", componentID.String())
	_, err := models.InstalledFirmwares(mod).DeleteAll(ctx, btx, false)
	return err
}
