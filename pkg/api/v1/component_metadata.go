package fleetdbapi

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/metal-automata/fleetdb/internal/models"
)

const (
	ComponentMetadataGenericNS = "metadata.generic"
)

type ComponentMetadata struct {
	ID                uuid.UUID  `json:"id,omitempty"`
	ServerComponentID uuid.UUID  `json:"server_component_id,omitempty" binding:"required"`
	ComponentName     string     `json:"component_name,omitempty"`
	Namespace         string     `json:"namespace" binding:"required"`
	Data              types.JSON `json:"data" binding:"required"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

// Equals compares two ComponentMetadata
func (c *ComponentMetadata) Equals(b *ComponentMetadata) bool {
	if c.Namespace != b.Namespace {
		return false
	}

	var m1, m2 map[string]string
	if err := json.Unmarshal(c.Data, &m1); err != nil {
		return false
	}
	if err := json.Unmarshal(b.Data, &m2); err != nil {
		return false
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}

	return true
}

func fromComponentMetadatumDBSlice(sl models.ComponentMetadatumSlice) []*ComponentMetadata {
	list := make([]*ComponentMetadata, 0, len(sl))
	for _, dbMeta := range sl {
		metadata := &ComponentMetadata{}
		metadata.fromDBModel(dbMeta)
		list = append(list, metadata)
	}

	return list
}

func (c *ComponentMetadata) fromDBModel(dbT *models.ComponentMetadatum) {
	c.ID = uuid.MustParse(dbT.ID)
	c.ServerComponentID = uuid.MustParse(dbT.ServerComponentID)
	c.Namespace = dbT.Namespace
	c.Data = dbT.Data
	c.CreatedAt = dbT.CreatedAt.Time
	c.UpdatedAt = dbT.UpdatedAt.Time

	if dbT.R != nil {
		c.ComponentName = dbT.R.ServerComponent.Name.String
	}
}

func (c *ComponentMetadata) toDBModel(componentID string) (*models.ComponentMetadatum, error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	if componentID == "" && c.ServerComponentID != uuid.Nil {
		componentID = c.ServerComponentID.String()
	}

	if _, err := uuid.Parse(componentID); err != nil {
		return nil, errors.Wrap(ErrValidatePayload, "invalid componentID, "+err.Error())
	}

	if c.Namespace != ComponentMetadataGenericNS {
		return nil, errors.Wrap(ErrValidatePayload, "unsupported metadata namespace: "+c.Namespace)
	}

	return &models.ComponentMetadatum{
		ID:                c.ID.String(),
		ServerComponentID: componentID,
		Namespace:         c.Namespace,
		Data:              c.Data,
	}, nil
}

// insert/update component metadata - the caller needs to invoke this within a transaction.
func (r *Router) upsertComponentMetadata(ctx context.Context, tx boil.ContextExecutor, componentID string, cm []*ComponentMetadata) error {
	for _, metadata := range cm {
		dbComponentMetadata, err := metadata.toDBModel(componentID) // component ID passed by parameter for validation
		if err != nil {
			return err
		}

		if err := dbComponentMetadata.Upsert(
			ctx,
			tx,
			true, // update on conflict
			// conflict columns
			[]string{
				models.ComponentMetadatumColumns.ServerComponentID,
				models.ComponentMetadatumColumns.Namespace,
			},
			// Columns to update when its an UPDATE
			boil.Whitelist(
				models.ComponentMetadatumColumns.Data,
			),
			// Columns to insert when its an INSERT
			boil.Infer(),
		); err != nil {
			return errors.Wrap(ErrDBQuery, err.Error())
		}
	}

	return nil
}

func (r *Router) componentMetadataSet(c *gin.Context) {
	var t []*ComponentMetadata
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid ComponentMetadata payload", err)
		return
	}

	ctx := c.Request.Context()
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	defer loggedRollback(r, tx)

	if err := r.upsertComponentMetadata(c.Request.Context(), r.DB, c.Param("componentID"), t); err != nil {
		if errors.Is(err, ErrDBQuery) {
			dbErrorResponse2(c, "component metadata upsert error", err)
			return
		}

		badRequestResponse(c, "invalid ComponentMetadata payload", err)
		return
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse2(c, "component metadata upsert error", err)
		return
	}

	createdResponse(c, "component metadata set")
}

func (r *Router) componentMetadataList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	mods := []qm.QueryMod{
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ServerComponents,
				models.ServerComponentTableColumns.ID,
				models.ComponentMetadatumTableColumns.ServerComponentID,
			),
		),
		qm.Load(models.ComponentMetadatumRels.ServerComponent),
	}

	// Add filter for component ID if provided
	if componentID := c.Param("componentID"); componentID != "" {
		mods = append(mods, qm.Where(models.ComponentMetadatumColumns.ServerComponentID+"=?", componentID))
	}

	// Add filter for namespace if provided
	if namespace := c.Param("namespace"); namespace != "" {
		mods = append(mods, qm.Where(models.ComponentMetadatumColumns.Namespace+"= ?", namespace))
	}

	dbComponentMetadata, err := models.ComponentMetadata(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count, err := models.ComponentMetadata(mods...).Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	list := fromComponentMetadatumDBSlice(dbComponentMetadata)

	pd := paginationData{
		pageCount:  len(list),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, list, pd)
}

func (r *Router) componentMetadataGet(c *gin.Context) {
	componentID := c.Param("componentID")
	componentUUID, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	namespace := c.Param("namespace")
	if namespace == "" {
		badRequestResponse(c, "", errors.Wrap(err, "valid component metadata namespace expected"))
		return
	}

	mods := []qm.QueryMod{
		qm.Where("server_component_id = ? AND namespace = ?", componentUUID.String(), namespace),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ServerComponents,
				models.ServerComponentTableColumns.ID,
				models.ComponentMetadatumTableColumns.ServerComponentID,
			),
		),
		qm.Load(models.ComponentMetadatumRels.ServerComponent),
	}

	dbComponentMetadata, err := models.ComponentMetadata(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	metadata := &ComponentMetadata{}
	metadata.fromDBModel(dbComponentMetadata)
	itemResponse(c, metadata)
}

func (r *Router) componentMetadataDelete(c *gin.Context) {
	componentID := c.Param("componentID")

	componentUUID, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	namespace := c.Param("namespace")
	if namespace == "" {
		badRequestResponse(c, "", errors.Wrap(err, "valid component metadata namespace expected"))
		return
	}

	mod := qm.Where("server_component_id = ? AND namespace = ?", componentUUID.String(), namespace)
	_, err = models.ComponentMetadata(mod).DeleteAll(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "Component Metadata delete error", err)
		return
	}

	deletedResponse(c)
}

// Compare metdata data slices for changes
func CompareMetadataSlices(existing, incoming []*ComponentMetadata) (creates, updates, deletes []*ComponentMetadata) {
	// Create map for efficient lookup of existing metadata by namespace
	existingMap := make(map[string]*ComponentMetadata)
	for _, e := range existing {
		if e != nil {
			existingMap[e.Namespace] = e
		}
	}

	// Check incoming metadata for creates and updates
	for _, inc := range incoming {
		if inc == nil {
			continue
		}

		if ex, exists := existingMap[inc.Namespace]; exists {
			// Found matching namespace, check if update needed
			if !ex.Equals(inc) {
				updates = append(updates, inc)
			}

			// Remove from map to track deletes
			delete(existingMap, inc.Namespace)
		} else {
			// No matching namespace found, needs to be created
			creates = append(creates, inc)
		}
	}

	// Remaining items in existingMap need to be deleted
	for _, ex := range existingMap {
		deletes = append(deletes, ex)
	}

	return creates, updates, deletes
}
