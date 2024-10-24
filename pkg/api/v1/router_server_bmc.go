package fleetdbapi

import (
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

type ServerBMC struct {
	ID                 uuid.UUID `json:"id"`
	ServerID           uuid.UUID `json:"server_id" binding:"required,uuid"` // Note: binding attributes should not have spaces
	HardwareVendorName string    `json:"hardware_vendor_name" binding:"required"`
	HardwareVendorID   string    `json:"-"`
	HardwareModelName  string    `json:"hardware_model_name"`
	HardwareModelID    string    `json:"-"`
	Username           string    `json:"username" binding:"required"`
	IPAddress          string    `json:"ipaddress" binding:"required,ip"`
	MacAddress         string    `json:"macaddress" binding:"mac"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (t *ServerBMC) fromDBModel(dbT *models.ServerBMC) {
	t.ID = uuid.MustParse(dbT.ID)
	t.ServerID = uuid.MustParse(dbT.ServerID)
	t.Username = dbT.Username
	t.IPAddress = dbT.IPAddress
	t.MacAddress = dbT.MacAddress.String
	t.CreatedAt = dbT.CreatedAt.Time
	t.UpdatedAt = dbT.UpdatedAt.Time

	if dbT.R != nil && dbT.R.HardwareModel != nil {
		t.HardwareModelName = dbT.R.HardwareModel.Name
	}
	if dbT.R != nil && dbT.R.HardwareVendor != nil {
		t.HardwareVendorName = dbT.R.HardwareVendor.Name
	}
}

func (t *ServerBMC) toDBModel() *models.ServerBMC {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return &models.ServerBMC{
		ID:               t.ID.String(),
		ServerID:         t.ServerID.String(),
		HardwareVendorID: t.HardwareVendorID,
		HardwareModelID:  t.HardwareModelID,
		Username:         t.Username,
		IPAddress:        t.IPAddress,
		MacAddress:       null.StringFrom(t.MacAddress),
	}
}

func (r *Router) serverBMCCreate(c *gin.Context) {
	var t ServerBMC
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid ServerBMC payload", err)
		return
	}

	bmc := t.toDBModel()

	// identify hardware vendor id
	dbHardwareVendor, err := r.hardwareVendorBySlug(c.Request.Context(), t.HardwareVendorName)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "hardware_vendor not identified: "+t.HardwareVendorName))
		return
	}

	bmc.HardwareVendorID = dbHardwareVendor.ID

	// identify hardware model id
	if t.HardwareModelName != "" {
		mod := qm.Where("name=?", t.HardwareModelName)
		dbHardwareModel, err := models.HardwareModels(mod).One(c.Request.Context(), r.DB)
		if err != nil {
			dbErrorResponse2(c, "hardware model lookup error", err)
			return
		}

		bmc.HardwareModelID = dbHardwareModel.ID
	}

	if err := bmc.Insert(c.Request.Context(), r.DB, boil.Infer()); err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, bmc.ID)
}

func (r *Router) serverBMCList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	mods := []qm.QueryMod{
		// join hardware vendor
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.HardwareVendors,
				models.HardwareVendorTableColumns.ID,
				models.ServerBMCTableColumns.HardwareVendorID,
			),
		),
		// Load N-1 relationship in db model struct field R
		qm.Load(models.ServerBMCRels.HardwareVendor),

		// join hardware model
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.HardwareModels,
				models.HardwareModelTableColumns.ID,
				models.ServerBMCTableColumns.HardwareModelID,
			),
		),
		// Load N-1 relationship in db model struct field R
		qm.Load(models.ServerBMCRels.HardwareModel),
	}

	dbServerBMCs, err := models.ServerBMCS(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count, err := models.ServerBMCS(mods...).Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	list := []ServerBMC{}
	for _, dbServerBMC := range dbServerBMCs {
		serverBMC := ServerBMC{}
		serverBMC.fromDBModel(dbServerBMC)
		list = append(list, serverBMC)
	}

	pd := paginationData{
		pageCount:  len(list),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, list, pd)
}

func (r *Router) serverBMCGet(c *gin.Context) {
	serverID := c.Param("serverID")

	serverUUID, err := uuid.Parse(serverID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid server UUID expected"))
		return
	}

	mods := []qm.QueryMod{
		qm.Where("server_id=?", serverUUID.String()),
		// join hardware vendor
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.HardwareVendors,
				models.HardwareVendorTableColumns.ID,
				models.ServerBMCTableColumns.HardwareVendorID,
			),
		),
		// Load N-1 relationship in db model struct field R
		qm.Load(models.ServerBMCRels.HardwareVendor),

		// join hardware model
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.HardwareModels,
				models.HardwareModelTableColumns.ID,
				models.ServerBMCTableColumns.HardwareModelID,
			),
		),
		// Load N-1 relationship in db model struct field R
		qm.Load(models.ServerBMCRels.HardwareModel),
	}

	dbServerBMC, err := models.ServerBMCS(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	serverBMC := &ServerBMC{}
	serverBMC.fromDBModel(dbServerBMC)
	itemResponse(c, serverBMC)
}

func (r *Router) serverBMCDelete(c *gin.Context) {
	serverID := c.Param("serverID")

	serverUUID, err := uuid.Parse(serverID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid server UUID expected"))
		return
	}

	mod := qm.Where("server_id=?", serverUUID)

	_, err = models.ServerBMCS(mod).DeleteAll(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "server BMC delete error", err)
		return
	}

	deletedResponse(c)
}
