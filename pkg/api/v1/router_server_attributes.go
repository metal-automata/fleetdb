package fleetdbapi

import (
	"database/sql"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

func (r *Router) serverAttributesList(c *gin.Context) {
	srv, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	dbAttrs, err := srv.Attributes().All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count, err := srv.Attributes().Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	attrs, err := convertFromDBAttributes(dbAttrs)
	if err != nil {
		failedConvertingToVersioned(c, err)
		return
	}

	pd := paginationData{
		pageCount:  len(attrs),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, attrs, pd)
}

func (r *Router) serverAttributesGet(c *gin.Context) {
	srv, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	ns := c.Param("namespace")

	dbAttr, err := srv.Attributes(models.AttributeWhere.Namespace.EQ(ns)).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	attr := Attributes{}
	if err := attr.fromDBModel(dbAttr); err != nil {
		failedConvertingToVersioned(c, err)
		return
	}

	itemResponse(c, attr)
}

func (r *Router) serverAttributesCreate(c *gin.Context) {
	srv, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	var attr Attributes
	if err := c.ShouldBindJSON(&attr); err != nil {
		badRequestResponse(c, "invalid attributes", err)
		return
	}

	dbAttr, err := attr.toDBModel()
	if err != nil {
		badRequestResponse(c, "invalid attributes", err)
		return
	}

	if err := srv.AddAttributes(c.Request.Context(), r.DB, true, dbAttr); err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, dbAttr.Namespace)
}

func (r *Router) serverAttributesUpdate(c *gin.Context) {
	u, err := r.parseUUID(c.Param("uuid"))
	if err != nil {
		return
	}

	ns := c.Param("namespace")

	var attr Attributes
	if err := c.ShouldBindJSON(&attr); err != nil {
		badRequestResponse(c, "invalid attributes", err)
		return
	}

	ctx := c.Request.Context()

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	defer loggedRollback(r, tx)

	rows, err := models.Attributes(qm.Where("namespace = ?", ns), qm.Where("server_id = ?", u)).UpdateAll(ctx, tx, models.M{"data": attr.Data})
	if err != nil {
		dbErrorResponse(c, err)

		return
	}

	if rows == 0 {
		dbErrorResponse(c, errors.Join(err, sql.ErrNoRows))

		return
	}

	rows, err = models.Servers(qm.Where("id = ?", u)).UpdateAll(ctx, tx, models.M{"updated_at": time.Now()})
	if err != nil {
		dbErrorResponse(c, err)

		return
	}

	if rows == 0 {
		dbErrorResponse(c, errors.Join(err, sql.ErrNoRows))

		return
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse(c, err)

		return
	}

	updatedResponse(c, ns)
}

func (r *Router) serverAttributesDelete(c *gin.Context) {
	u := c.Param("uuid")
	ns := c.Param("namespace")

	rows, err := models.Attributes(qm.Where("namespace = ?", ns), qm.Where("server_id = ?", u)).DeleteAll(c.Request.Context(), r.DB)
	if rows == 0 && err == nil {
		err = sql.ErrNoRows
	}

	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	deletedResponse(c)
}
