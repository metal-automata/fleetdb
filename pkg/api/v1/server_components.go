package fleetdbapi

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/metal-automata/fleetdb/internal/models"
)

var (
	errSrvComponentPayload = errors.New("error in server component payload")
	collectionMethods      = []string{"inband", "outofband"}
)

// ServerComponent represents a component of a server. These can be things like
// processors, NICs, hard drives, etc.
//
// Note: when setting validator struct tags, ensure no extra spaces are present between
//
//	comma separated values or validation will fail with a not so useful 500 error.
type ServerComponent struct {
	UUID                  uuid.UUID              `json:"uuid"`
	ServerUUID            uuid.UUID              `json:"server_uuid" binding:"required"`
	Name                  string                 `json:"name" binding:"required"` // name is the component slug
	Vendor                string                 `json:"vendor"`
	Model                 string                 `json:"model"`
	Serial                string                 `json:"serial" binding:"required"`
	Description           string                 `json:"description"`
	ServerComponentTypeID string                 `json:"server_component_type_id,omitempty"`
	Metadata              []*ComponentMetadata   `json:"metadata,omitempty"`
	Capabilities          []*ComponentCapability `json:"capabilities,omitempty"`
	InstalledFirmware     *InstalledFirmware     `json:"firmware,omitempty"`
	Status                *ComponentStatus       `json:"status,omitempty"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

// ServerComponentSlice is a slice of ServerComponent objects
type ServerComponentSlice []*ServerComponent

func componentKey(slug, serial string) string {
	return strings.ToLower(slug) + ":" + strings.ToLower(serial)
}

// asMap returns the slice as a map with the items keyed by the name:serial
func (s ServerComponentSlice) asMap() map[string]*ServerComponent {
	m := make(map[string]*ServerComponent)
	for _, curr := range s {
		m[componentKey(curr.Name, curr.Serial)] = curr
	}

	return m
}

func (s ServerComponentSlice) Compare(incomming ServerComponentSlice) (creates, updates, deletes ServerComponentSlice) {
	currentMap := s.asMap()
	incommingMap := incomming.asMap()

	// Find creates and updates
	for key, inc := range incommingMap {
		if current, exists := currentMap[key]; exists {
			if _, equal := current.Equals(inc); !equal {
				// uncomment for debugging comparison
				// fmt.Printf("component not equal, key: %s cause: %s\n", key, cause)
				updates = append(updates, inc)
			}
			delete(currentMap, key) // Remove from map to track deletes
		} else {
			creates = append(creates, inc)
		}
	}

	// Remaining items in map are deletes
	for _, curr := range currentMap {
		deletes = append(deletes, curr)
	}

	return creates, updates, deletes
}

func (s *ServerComponentSlice) fromDBModel(dbTS models.ServerComponentSlice) {
	for _, dbC := range dbTS {
		c := ServerComponent{}
		c.fromDBModel(dbC)

		*s = append(*s, &c)
	}
}

// Equals compares ServerComponents, excluding certain fields and timestamps
func (c *ServerComponent) Equals(b *ServerComponent) (string, bool) {
	if c == nil || b == nil {
		return "nil component", false
	}

	if c.Name != b.Name {
		return "name differs", false
	}

	if c.Serial != b.Serial {
		return "serial differs", false
	}

	if c.Vendor != b.Vendor {
		return "vendor differs", false
	}

	if c.Model != b.Model {
		return "model differs", false
	}

	// the first one in the list is the current and that is what we're interested in
	if b.InstalledFirmware != nil {
		if c.InstalledFirmware == nil {
			return "installed firmware differs", false
		}

		if !c.InstalledFirmware.Equals(b.InstalledFirmware) {
			return "installed firmware differs", false
		}
	}

	// compare status
	if b.Status != nil {
		if c.Status == nil {
			return "status differs", false
		}

		if !c.Status.Equals(b.Status) {
			return "status differs", false
		}
	}

	// compare metadata
	createsm, updatesm, deletesm := CompareMetadataSlices(c.Metadata, b.Metadata)
	if len(createsm) > 0 || len(updatesm) > 0 || len(deletesm) > 0 {
		return "metadata differs", false
	}

	// compare capabilities
	createsc, updatesc, deletesc := CompareCapabilitySlices(c.Capabilities, b.Capabilities)
	if len(createsc) > 0 || len(updatesc) > 0 || len(deletesc) > 0 {
		return "capability differs", false
	}

	return "", true
}

func ServerComponentFromModel(dbC *models.ServerComponent) *ServerComponent {
	c := &ServerComponent{}
	c.fromDBModel(dbC)
	return c
}

// fromDBModel populates the ServerComponent object fields based on values in the store
func (c *ServerComponent) fromDBModel(dbC *models.ServerComponent) {
	c.UUID = uuid.MustParse(dbC.ID)
	c.ServerUUID = uuid.MustParse(dbC.ServerID)

	c.Name = dbC.Name.String
	c.Description = dbC.Description.String
	c.Vendor = dbC.Vendor.String
	c.Model = dbC.Model.String
	c.Serial = dbC.Serial
	c.CreatedAt = dbC.CreatedAt.Time
	c.UpdatedAt = dbC.UpdatedAt.Time

	// set component type rel
	if dbC.R == nil {
		return
	}

	// set installed firmware fields
	if dbC.R.InstalledFirmware != nil {
		c.InstalledFirmware = &InstalledFirmware{Version: dbC.R.InstalledFirmware.Version}
	}

	// set component status fields
	if dbC.R.ComponentStatus != nil {
		componentStatus := &ComponentStatus{}
		componentStatus.fromDBModel(dbC.R.ComponentStatus)
		c.Status = componentStatus
	}

	// set component capability fields
	if dbC.R.ComponentCapabilities != nil {
		c.Capabilities = make([]*ComponentCapability, 0, len(dbC.R.ComponentCapabilities))
		for _, dbComponentCapability := range dbC.R.ComponentCapabilities {
			componentCapbility := &ComponentCapability{}
			componentCapbility.fromDBModel(dbComponentCapability)
			c.Capabilities = append(c.Capabilities, componentCapbility)
		}
	}

	// set component metadata fields
	if dbC.R.ComponentMetadata != nil {
		c.Metadata = make([]*ComponentMetadata, 0, len(dbC.R.ComponentMetadata))
		for _, dbComponentMetadata := range dbC.R.ComponentMetadata {
			componentMetadata := &ComponentMetadata{}
			componentMetadata.fromDBModel(dbComponentMetadata)
			c.Metadata = append(c.Metadata, componentMetadata)
		}
	}
}

// toDBModel converts a ServerComponent object to a model.ServerComponent object
func (c *ServerComponent) toDBModel(serverID string) *models.ServerComponent {
	return &models.ServerComponent{
		//ID:          c.UUID.String(),
		ServerID:              serverID,
		Name:                  null.StringFrom(c.Name),
		Description:           null.StringFrom(c.Description),
		Vendor:                null.StringFrom(c.Vendor),
		Model:                 null.StringFrom(c.Model),
		Serial:                c.Serial,
		ServerComponentTypeID: c.ServerComponentTypeID,
	}
}

// getServerComponents returns server components based on query parameters
func (r *Router) getServerComponents(c *gin.Context, params []ServerComponentListParams, pagination PaginationParams) (models.ServerComponentSlice, int64, error) {
	mods := []qm.QueryMod{}

	// for each parameter, setup the query modifiers
	for _, param := range params {
		mods = append(mods, param.queryMods(models.TableNames.ServerComponents))
	}

	count, err := models.ServerComponents(mods...).Count(c.Request.Context(), r.DB)
	if err != nil {
		return nil, 0, err
	}

	// add pagination
	mods = append(mods, pagination.serverComponentsQueryMods()...)

	sc, err := models.ServerComponents(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		return sc, 0, err
	}

	return sc, count, nil
}

// serverComponentList returns a response with the list of components that matched the params.
func (r *Router) serverComponentList(c *gin.Context) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	params, err := parseQueryServerComponentsListParams(c)
	if err != nil {
		badRequestResponse(c, "invalid server component list params", err)
		return
	}

	dbSC, count, err := r.getServerComponents(c, params, pager)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	serverComponents := ServerComponentSlice{}

	for _, dbSC := range dbSC {
		sc := ServerComponent{}
		sc.fromDBModel(dbSC)
		serverComponents = append(serverComponents, &sc)
	}

	pd := paginationData{
		pageCount:  len(serverComponents),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, serverComponents, pd)
}

// serverComponentGet returns a response with the list of components referenced by the server UUID.
func (r *Router) serverComponentGet(c *gin.Context) {
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

	params := &ServerComponentGetParams{}
	params.decode(c.Request.URL.Query())

	dbTS, err := srv.ServerComponents(params.queryMods(true)...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	serverComponents := make(ServerComponentSlice, 0, len(dbTS))
	serverComponents.fromDBModel(dbTS)

	pd := paginationData{
		pageCount:  len(serverComponents),
		totalCount: int64(0),
		pager:      pager,
	}

	listResponse(c, serverComponents, pd)
}

// serverComponentsInitCollection is the handler to to initialize the first component and related records for a server
//
// This will only create component records only if they don't already exist.
func (r *Router) serverComponentsInitCollection(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)
		return
	}

	cmethod := c.Param("collection-method")
	if cmethod == "" || !slices.Contains(collectionMethods, cmethod) {
		badRequestResponse(
			c,
			"expected a valid collection-method query param",
			nil,
		)
		return
	}

	// check server exists
	if server == nil {
		notFoundResponse(c, "server resource")
		return
	}

	// components payload
	var incoming ServerComponentSlice
	if errBind := c.ShouldBindJSON(&incoming); errBind != nil {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errSrvComponentPayload, "failed to unmarshal JSON as ServerComponentSlice: "+errBind.Error(),
			),
		)
		return
	}

	if len(incoming) == 0 {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errSrvComponentPayload, "ServerComponentSlice is empty"),
		)
		return
	}

	// check if there are any existing components for server
	mod := qm.Where(models.ServerComponentColumns.ServerID+"=?", server.ID)
	current, err := models.ServerComponents(mod).All(c.Request.Context(), r.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		dbErrorResponse2(c, "current Server component query error", err)
		return
	}

	if len(current) != 0 {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errSrvComponentPayload, "use the components/update or component-changes/report instead"),
		)
		return
	}

	// insert records
	if err := r.insertServerComponentsWithTx(c.Request.Context(), server, incoming); err != nil {
		if errors.Is(err, ErrComponentType) {
			badRequestResponse(
				c,
				"",
				errors.Wrap(
					errSrvComponentPayload, err.Error()),
			)
			return
		}

		dbErrorResponse2(c, "server component inventory insert error", err)
		return
	}

	createdResponse(c, "components added")
}

func (r *Router) insertServerComponentsWithTx(ctx context.Context, server *models.Server, components ServerComponentSlice) error {
	// init map of slug to component type ID
	slugMap, err := r.serverComponentTypeSlugMap(ctx)
	if err != nil {
		return errors.Wrap(err, "server component types query")
	}

	// component data is written in a transaction along with versioned attributes
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer loggedRollback(r, tx)

	for _, insert := range components {
		componentTypeID, exists := slugMap[insert.Name]
		if !exists {
			return errors.Wrap(ErrComponentType, "unsupported: "+insert.Name)
		}

		insert.ServerComponentTypeID = componentTypeID
		if err := r.componentAndRelationsUpsert(ctx, tx, server.ID, *insert); err != nil {
			return err
		}
	}

	if err := r.serverInventoryRefreshed(ctx, tx, server); err != nil {
		return errors.Wrap(err, "server component inventory ts update error")
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Router) upsertServerComponent(ctx context.Context, tx boil.ContextExecutor, serverID string, t ServerComponent) (string, error) {
	dbServerComponent := t.toDBModel(serverID)
	if err := dbServerComponent.Upsert(
		ctx,
		tx,
		true, // update on conflict
		// conflict columns
		[]string{
			models.ServerComponentColumns.ServerID,
			models.ServerComponentColumns.Serial,
			models.ServerComponentColumns.ServerComponentTypeID,
		},
		// Columns to update when its an UPDATE
		boil.Whitelist(
			models.ServerComponentColumns.Vendor,
			models.ServerComponentColumns.Model,
			models.ServerComponentColumns.Serial,
			models.ServerComponentColumns.Description,
		),
		// Columns to insert when its an INSERT
		boil.Infer(),
	); err != nil {
		return "", err
	}

	return dbServerComponent.ID, nil
}

// serverComponentUpdateCollection will update existing component record and its relations for existing server component records
//
// Note: this will not add or delete component records.
func (r *Router) serverComponentUpdateCollection(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse2(c, "server query error", err)
		return
	}

	// check server exists
	if server == nil {
		notFoundResponse(c, "server resource")
		return
	}

	cmethod := c.Param("collection-method")
	if cmethod == "" || !slices.Contains(collectionMethods, cmethod) {
		badRequestResponse(
			c,
			"expected a valid collection-method query param",
			nil,
		)
		return
	}

	// components payload
	var incoming ServerComponentSlice
	if errBind := c.ShouldBindJSON(&incoming); errBind != nil {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errSrvComponentPayload, errBind.Error()),
		)
		return
	}

	if len(incoming) == 0 {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errSrvComponentPayload, "ServerComponentSlice is empty"),
		)
		return
	}

	slugMap, err := r.serverComponentTypeSlugMap(c.Request.Context())
	if err != nil {
		dbErrorResponse2(c, "component type query error", err)
		return
	}

	if err := r.applyServerComponentUpdateWithTx(c.Request.Context(), server, incoming, slugMap); err != nil {
		if errors.Is(err, errDBErr) {
			dbErrorResponse2(c, "apply component update db error", err)
			return
		}

		badRequestResponse(
			c,
			"",
			errors.Wrap(errSrvComponentPayload, err.Error()),
		)
		return
	}

	updatedResponse(c, "component(s) updated")
}

// method applies one or more server component updates - records are updated within a tx
func (r *Router) applyServerComponentUpdateWithTx(ctx context.Context, server *models.Server, updates ServerComponentSlice, slugMap map[string]string) error {
	// select components by serverID, componentID's from the incoming payload
	//
	// https://github.com/volatiletech/sqlboiler/issues/227#issuecomment-348053252
	convertedIDs := make([]interface{}, 0, len(updates))
	for _, component := range updates {
		convertedIDs = append(convertedIDs, component.UUID.String())
	}

	mods := []qm.QueryMod{
		qm.Where(models.ServerComponentColumns.ServerID+"=?", server.ID),
		qm.WhereIn(
			fmt.Sprintf(
				"%s in ?",
				models.ServerComponentColumns.ID,
			),
			convertedIDs...,
		),
	}

	errComponentUpdate := errors.New("error applying component update")
	boil.DebugMode = true
	currentRecords, err := models.ServerComponents(mods...).All(ctx, r.DB)
	boil.DebugMode = false
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(errComponentUpdate, "no existing components for update - use the create endpoint")
		}

		return err
	}

	currentMap := map[string]*models.ServerComponent{}
	for _, c := range currentRecords {
		key := componentKey(c.Name.String, c.Serial)
		currentMap[key] = c
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(errDBErr, err.Error())
	}

	defer loggedRollback(r, tx)

	//	models.ServerComponent
	for key, update := range updates.asMap() {
		current, exists := currentMap[key]
		if !exists {
			return errors.Wrap(errComponentUpdate, "unknown component: "+key)
		}

		componentTypeID, exists := slugMap[update.Name]
		if !exists {
			return errors.Wrap(err, "unknown server component type: "+update.Name)
		}

		update.ServerComponentTypeID = componentTypeID
		if err := r.componentAndRelationsUpsert(ctx, tx, current.ServerID, *update); err != nil {
			return errors.Wrap(errDBErr, err.Error())
		}
	}

	if err := r.serverInventoryRefreshed(ctx, tx, server); err != nil {
		return errors.Wrap(errDBErr, err.Error())
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(errDBErr, err.Error())
	}

	return nil
}

func (r *Router) componentAndRelationsUpsert(ctx context.Context, tx boil.ContextExecutor, serverID string, component ServerComponent) error {
	// upsert component data
	componentID, err := r.upsertServerComponent(ctx, tx, serverID, component)
	if err != nil {
		return errors.Wrap(err, "component upsert error")
	}

	// upsert firmware installed data
	if component.InstalledFirmware != nil {
		if _, err := r.upsertInstalledFirmware(ctx, tx, componentID, *component.InstalledFirmware); err != nil {
			return errors.Wrap(err, "component installed firmware upsert error")
		}
	}

	// upsert status
	if component.Status != nil {
		if _, err := r.upsertComponentStatus(ctx, tx, componentID, *component.Status); err != nil {
			return errors.Wrap(err, "component status upsert error")
		}
	}

	// upsert capabilities
	if component.Capabilities != nil {
		if err := r.upsertComponentCapability(ctx, tx, componentID, component.Capabilities); err != nil {
			return errors.Wrap(err, "component capabilities upsert error")
		}
	}

	// upsert metadata
	if component.Metadata != nil {
		if err := r.upsertComponentMetadata(ctx, tx, componentID, component.Metadata); err != nil {
			return errors.Wrap(err, "component metadata upsert error")
		}
	}

	return nil
}

// serverComponentDelete deletes a server component.
func (r *Router) serverComponentDeleteAll(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c.Request.Context(), c.Param("uuid"))
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)
		return
	}

	if _, err := server.ServerComponents().DeleteAll(c.Request.Context(), r.DB); err != nil {
		dbErrorResponse(c, err)
		return
	}

	deletedResponse(c)
}
