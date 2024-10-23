package fleetdbapi

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

type HardwareVendor struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
}

func (t *HardwareVendor) fromDBModel(dbT *models.HardwareVendor) {
	t.ID = dbT.ID
	t.Name = dbT.Name
}

func (t *HardwareVendor) toDBModel() *models.HardwareVendor {
	if t.ID == uuid.Nil.String() || t.ID == "" {
		t.ID = uuid.NewString()
	}

	return &models.HardwareVendor{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (r *Router) hardwareVendorCreate(c *gin.Context) {
	var t HardwareVendor
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid HardwareVendor payload", err)
		return
	}

	hv := t.toDBModel()

	if err := hv.Insert(c.Request.Context(), r.DB, boil.Infer()); err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, t.Name)
}

func (r *Router) hardwareVendorList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	dbHardwareVendors, err := models.HardwareVendors().All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count, err := models.HardwareVendors().Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	list := []HardwareVendor{}
	for _, dbhwv := range dbHardwareVendors {
		hw := HardwareVendor{}
		hw.fromDBModel(dbhwv)
		list = append(list, hw)
	}

	pd := paginationData{
		pageCount:  len(list),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, list, pd)
}

func (r *Router) hardwareVendorGet(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		badRequestResponse(c, "", errors.New("vendor name slug expected"))
		return
	}

	dbHardwareVendor, err := r.hardwareVendorBySlug(c.Request.Context(), slug)
	if err != nil {
		dbErrorResponse2(c, "hardware vendor lookup error", err)
		return
	}

	var hwv HardwareVendor
	hwv.fromDBModel(dbHardwareVendor)

	itemResponse(c, hwv)
}

func (r *Router) hardwareVendorBySlug(ctx context.Context, slug string) (*models.HardwareVendor, error) {
	mods := []qm.QueryMod{
		qm.Where("name=?", slug),
	}

	return models.HardwareVendors(mods...).One(ctx, r.DB)
}

func (r *Router) hardwareVendorDelete(c *gin.Context) {
	mods := []qm.QueryMod{
		qm.Where("name=?", c.Param("slug")),
	}

	_, err := models.HardwareVendors(mods...).DeleteAll(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "hardware vendor delete error", err)
		return
	}

	deletedResponse(c)
}
