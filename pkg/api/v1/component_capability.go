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

type ComponentCapability struct {
	ID                uuid.UUID `json:"id,omitempty"`
	ServerComponentID uuid.UUID `json:"server_component_id,omitempty" binding:"required"`
	ComponentName     string    `json:"component_name"`
	Name              string    `json:"name" binding:"required"`
	Description       string    `json:"description" binding:"required"`
	Enabled           bool      `json:"enabled"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// Equals compares two ComponentCapabilities focusing on required fields
func (c *ComponentCapability) Equals(b *ComponentCapability) bool {
	if c.Name != b.Name {
		return false
	}

	if c.Description != b.Description {
		return false
	}

	if c.Enabled != b.Enabled {
		return false
	}

	return true
}

func (c *ComponentCapability) fromDBModel(dbT *models.ComponentCapability) {
	c.ID = uuid.MustParse(dbT.ID)
	c.ServerComponentID = uuid.MustParse(dbT.ServerComponentID)
	c.Name = dbT.Name
	c.Description = dbT.Description.String
	c.Enabled = dbT.Enabled.Bool
	c.CreatedAt = dbT.CreatedAt.Time
	c.UpdatedAt = dbT.UpdatedAt.Time

	if dbT.R != nil {
		c.ComponentName = dbT.R.ServerComponent.Name.String
	}
}

func (c *ComponentCapability) toDBModel(componentID string) (*models.ComponentCapability, error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	if componentID == "" && c.ServerComponentID != uuid.Nil {
		componentID = c.ServerComponentID.String()
	}

	if _, err := uuid.Parse(componentID); err != nil {
		return nil, errors.Wrap(ErrValidatePayload, "invalid componentID"+err.Error())
	}

	return &models.ComponentCapability{
		ID:                c.ID.String(),
		ServerComponentID: componentID,
		Name:              c.Name,
		Description:       null.StringFrom(c.Description),
		Enabled:           null.BoolFrom(c.Enabled),
	}, nil
}

// insert/update component capability - the caller needs to invoke this within a transaction.
func (r *Router) upsertComponentCapability(ctx context.Context, tx boil.ContextExecutor, componentID string, t []*ComponentCapability) error {
	for _, cap := range t {
		dbComponentCapability, err := cap.toDBModel(componentID) // componentID passed by parameter for validation
		if err != nil {
			return err
		}

		if err := dbComponentCapability.Upsert(
			ctx,
			tx,
			true, // update on conflict
			// conflict columns
			[]string{
				models.ComponentCapabilityColumns.ServerComponentID,
				models.ComponentCapabilityColumns.Name,
			},
			// Columns to update when its an UPDATE
			boil.Whitelist(
				models.ComponentCapabilityColumns.Enabled,
				models.ComponentCapabilityColumns.Description,
			),
			// Columns to insert when its an INSERT
			boil.Infer(),
		); err != nil {
			return errors.Wrap(ErrDBQuery, err.Error())
		}
	}

	return nil
}

func (r *Router) componentCapabilitySet(c *gin.Context) {
	var t []*ComponentCapability
	if err := c.ShouldBindJSON(&t); err != nil {
		badRequestResponse(c, "invalid ComponentCapability payload", err)
		return
	}

	ctx := c.Request.Context()
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	defer loggedRollback(r, tx)

	if err := r.upsertComponentCapability(ctx, r.DB, c.Param("componentID"), t); err != nil {
		if errors.Is(err, ErrDBQuery) {
			dbErrorResponse2(c, "component capability upsert error", err)
			return
		}

		badRequestResponse(c, "invalid ComponentCapability payload", err)
		return
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse2(c, "component capability upsert error", err)
		return
	}

	createdResponse(c, "component capabilities set")
}

func (r *Router) componentCapabilityGet(c *gin.Context) {
	componentID := c.Param("componentID")
	componentUUID, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	mods := []qm.QueryMod{
		qm.Where("server_component_id = ?", componentUUID.String()),
		qm.InnerJoin(
			fmt.Sprintf(
				"%s on %s = %s",
				models.TableNames.ServerComponents,
				models.ServerComponentTableColumns.ID,
				models.ComponentCapabilityTableColumns.ServerComponentID,
			),
		),
		qm.Load(models.ComponentCapabilityRels.ServerComponent),
	}

	dbComponentCapability, err := models.ComponentCapabilities(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	componentCapability := &ComponentCapability{}
	componentCapability.fromDBModel(dbComponentCapability)
	itemResponse(c, componentCapability)
}

func (r *Router) componentCapabilityDelete(c *gin.Context) {
	componentID := c.Param("componentID")
	componentUUID, err := uuid.Parse(componentID)
	if err != nil {
		badRequestResponse(c, "", errors.Wrap(err, "valid component UUID expected"))
		return
	}

	mod := qm.Where("server_component_id = ?", componentUUID.String())
	_, err = models.ComponentCapabilities(mod).DeleteAll(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse2(c, "Component Capability delete error", err)
		return
	}

	deletedResponse(c)
}

// CompareCapabilitySlices compares slices of ComponentCapability using Name as the key
// and checks Description and Enabled fields for changes
func CompareCapabilitySlices(existing, incoming []*ComponentCapability) (creates, updates, deletes []*ComponentCapability) {
	// Create map for efficient lookup of existing capabilities by name
	existingMap := make(map[string]*ComponentCapability)
	for _, e := range existing {
		if e != nil && e.Name != "" {
			existingMap[e.Name] = e
		}
	}

	// Check incoming capabilities for creates and updates
	for _, inc := range incoming {
		if inc == nil || inc.Name == "" {
			continue
		}

		if ex, exists := existingMap[inc.Name]; exists {
			// Found matching name, check if update needed
			if ex.Description != inc.Description || ex.Enabled != inc.Enabled {
				updates = append(updates, inc)
			}

			// Remove from map to track deletes
			delete(existingMap, inc.Name)
		} else {
			// No matching name found, needs to be created
			creates = append(creates, inc)
		}
	}

	// Remaining items in existingMap need to be deleted
	for _, ex := range existingMap {
		deletes = append(deletes, ex)
	}

	return creates, updates, deletes
}
