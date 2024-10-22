package fleetdbapi

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

type HardwareModel struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name" binding:"required"`
	HardwareVendorName string    `json:"hardware_vendor_name" binding:"required" boil:"hardware_vendor_name"`
	HardwareVendorID   string    `json:"hardware_vendor_id"`
}

func (t *HardwareModel) fromDBModel(dbT *models.HardwareModel) {
	t.ID = uuid.MustParse(dbT.ID)
	t.Name = dbT.Name

	t.HardwareVendorID = dbT.HardwareVendorID
	t.HardwareVendorName = dbT.R.HardwareVendor.Name
}

func (t *HardwareModel) toDBModel() *models.HardwareModel {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return &models.HardwareModel{
		ID:               t.ID.String(),
		Name:             t.Name,
		HardwareVendorID: t.HardwareVendorID,
	}
}

func (r *Router) hardwareModelCreate(c *gin.Context) {
	var t HardwareModel
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid HardwareModel payload", err)
		return
	}

	if t.HardwareVendorName == "" {
		badRequestResponse(
			c,
			"",
			errors.New("invalid HardwareModel payload: hardware_vendor_name expected"),
		)
		return
	}

	dbHardwareVendor, err := r.hardwareVendorBySlug(c.Request.Context(), t.HardwareVendorName)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "hardware_vendor not identified: "+t.HardwareVendorName))
		return
	}

	t.HardwareVendorID = dbHardwareVendor.ID
	hv := t.toDBModel()

	if err := hv.Insert(c.Request.Context(), r.DB, boil.Infer()); err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, t.Name)
}

func (r *Router) hardwareModelList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	mods := []qm.QueryMod{
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s.%s",
				models.TableNames.HardwareVendors,
				models.HardwareVendorTableColumns.ID,
				models.TableNames.HardwareModels,
				models.HardwareModelColumns.HardwareVendorID,
			),
		),
		// load N-1 relationship
		qm.Load(models.HardwareModelRels.HardwareVendor),
	}

	dbHardwareModels, err := models.HardwareModels(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count, err := models.HardwareModels().Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	list := []HardwareModel{}
	for _, dbhwm := range dbHardwareModels {
		hwm := HardwareModel{}
		hwm.fromDBModel(dbhwm)
		spew.Dump(hwm)
		list = append(list, hwm)
	}

	pd := paginationData{
		pageCount:  len(list),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, list, pd)
}

func (r *Router) hardwareModelGet(c *gin.Context) {
	mods := []qm.QueryMod{
		qm.Where("name=?", c.Param("slug")),
		qm.Load(models.HardwareModelRels.HardwareVendor),
	}

	dbHardwareModel, err := models.HardwareModels(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "hardware model lookup error", err)
		return
	}

	var hwv HardwareModel
	hwv.fromDBModel(dbHardwareModel)

	itemResponse(c, hwv)
}

func (r *Router) hardwareModelDelete(c *gin.Context) {
	mods := []qm.QueryMod{
		qm.Where("name=?", c.Param("slug")),
	}

	_, err := models.HardwareModels(mods...).DeleteAll(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "hardware model delete error", err)
		return
	}

	deletedResponse(c)
}
