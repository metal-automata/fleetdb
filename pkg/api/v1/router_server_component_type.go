package fleetdbapi

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/metal-automata/fleetdb/internal/models"
)

func (r *Router) serverComponentTypeCreate(c *gin.Context) {
	var t ServerComponentType
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid server component type", err)
		return
	}

	dbT, err := t.toDBModel()
	if err != nil {
		badRequestResponse(c, "invalid server component type", err)
		return
	}

	if err := dbT.Insert(c.Request.Context(), r.DB, boil.Infer()); err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, dbT.Slug)
}

func (r *Router) serverComponentTypeList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	// dbFilter := &gormdb.ServerComponentTypeFilter{
	// 	Name: c.Query("name"),
	// }

	dbTypes, err := models.ServerComponentTypes().All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count, err := models.ServerComponentTypes().Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	types := []ServerComponentType{}

	for _, dbT := range dbTypes {
		t := ServerComponentType{}
		if err := t.fromDBModel(dbT); err != nil {
			failedConvertingToVersioned(c, err)
			return
		}

		types = append(types, t)
	}

	pd := paginationData{
		pageCount:  len(types),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, types, pd)
}

// returns a map of component slug/name to component type ID for lookups
func (r *Router) serverComponentTypeSlugMap(ctx context.Context) (map[string]string, error) {
	dbComponentTypes, err := models.ServerComponentTypes().All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	slugMap := make(map[string]string, len(dbComponentTypes))
	for _, dbCcomponentType := range dbComponentTypes {
		slugMap[dbCcomponentType.Slug] = dbCcomponentType.ID
	}

	return slugMap, nil
}
