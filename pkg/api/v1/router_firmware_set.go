package fleetdbapi

import (
	"context"
	"database/sql"
	"net/http"

	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"github.com/metal-automata/fleetdb/internal/metrics"
	"github.com/metal-automata/fleetdb/internal/models"
)

var (
	errComponentFirmwareSetRequest = errors.New("error in component firmware set request")
	errComponentFirmwareSetMap     = errors.New("error mapping firmware in set")
	errDBErr                       = errors.New("db error")
	ErrFwSetByVendorModel          = errors.New("error identifying firmware set by server vendor, model")

	// FleetDB attribute namespace for firmware set labels.
	FirmwareSetAttributeNS = "sh.hollow.firmware_set.labels"
)

// Firmware sets group firmware versions
//
// - firmware sets can only reference to unique firmware versions based on the vendor, model, component attributes.

func (r *Router) serverComponentFirmwareSetList(c *gin.Context) {
	// unmarshal query parameters
	var params ComponentFirmwareSetListParams
	var err error
	if err = c.ShouldBindQuery(&params); err != nil {
		badRequestResponse(c, "invalid filter payload: ComponentFirmwareSetListParams{}", err)
		return
	}

	// query parameters to query mods
	params.AttributeListParams = parseQueryAttributesListParams(c, "attr")
	if params.AttributeListParams, err = r.validateListParams(c, params); err != nil {
		badRequestResponse(
			c, "empty required attribute",
			errors.Wrap(ErrFwSetByVendorModel, err.Error()),
		)
		return
	}

	mods := params.queryMods(models.TableNames.ComponentFirmwareSet)
	mods = append(mods, qm.Load(models.ComponentFirmwareSetRels.FirmwareSetAttributesFirmwareSets))
	r.selectFirmware(c, mods)
}

func (r *Router) validateListParams(c *gin.Context, params ComponentFirmwareSetListParams) ([]AttributeListParams, error) {
	if len(params.AttributeListParams) != 0 {
		return params.AttributeListParams, nil
	}

	if params.Vendor == "" && params.Model == "" && params.Labels == "" {
		return params.AttributeListParams, nil
	}

	if params.Vendor == "" {
		return nil, errors.Wrap(ErrFwSetByVendorModel, "vendor")
	}

	if params.Model == "" {
		return nil, errors.Wrap(ErrFwSetByVendorModel, "model")
	}

	return r.serverComponentFirmwareSetsSelect(c, params.Vendor, params.Model, params.Labels)
}

func (r *Router) selectFirmware(c *gin.Context, mods []qm.QueryMod) {
	pager, err := parsePagination(c)
	if err != nil {
		badRequestResponse(c, "invalid pagination params", err)
		return
	}

	// count rows
	count, err := models.ComponentFirmwareSets(mods...).Count(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	pager.Preload = false

	// load firmware sets
	dbFirmwareSets, err := models.ComponentFirmwareSets(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	firmwareSets := make([]ComponentFirmwareSet, 0, count)

	// load firmware set mappings
	for _, dbFS := range dbFirmwareSets {
		f := ComponentFirmwareSet{}

		firmwares, err := r.queryFirmwareSetFirmware(c.Request.Context(), dbFS.ID)
		if err != nil {
			dbErrorResponse(c, err)
			return
		}

		if err := f.fromDBModel(dbFS, firmwares); err != nil {
			failedConvertingToVersioned(c, err)
			return
		}

		firmwareSets = append(firmwareSets, f)
	}

	pd := paginationData{
		pageCount:  len(firmwareSets),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, firmwareSets, pd)
}

func (r *Router) serverComponentFirmwareSetsSelect(_ *gin.Context, vendor, model, labels string) ([]AttributeListParams, error) {
	listParams := make([]AttributeListParams, 0)

	listParams = appendToQueryFirmwareSetsParams(listParams, "vendor", "eq", vendor)
	listParams = appendToQueryFirmwareSetsParams(listParams, "model", "eq", model)

	var additionalLabel bool
	arbitraryLabels := parseQueryFirmwareSetsLabels(labels)
	for k, v := range arbitraryLabels {
		additionalLabel = true
		listParams = appendToQueryFirmwareSetsParams(listParams, k, "eq", v)
	}
	if !additionalLabel {
		listParams = appendToQueryFirmwareSetsParams(listParams, "default", "eq", "true")
	}

	return listParams, nil
}

func (r *Router) serverComponentFirmwareSetGet(c *gin.Context) {
	setID := c.Param("uuid")
	if setID == "" || setID == uuid.Nil.String() {
		badRequestResponse(c, "expected a firmware set UUID, got none", errComponentFirmwareSetRequest)
		return
	}

	setIDParsed, err := uuid.Parse(setID)
	if err != nil {
		badRequestResponse(c, "invalid firmware set UUID: "+setID, err)
	}

	// query firmware set
	mods := []qm.QueryMod{
		qm.Where("id=?", setIDParsed),
		qm.Load(models.ComponentFirmwareSetRels.FirmwareSetAttributesFirmwareSets),
	}

	dbFirmwareSet, err := models.ComponentFirmwareSets(mods...).One(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	firmwares, err := r.queryFirmwareSetFirmware(c.Request.Context(), dbFirmwareSet.ID)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	// convert from db model type
	var firmwareSet ComponentFirmwareSet
	if err = firmwareSet.fromDBModel(dbFirmwareSet, firmwares); err != nil {
		failedConvertingToVersioned(c, err)
		return
	}

	itemResponse(c, firmwareSet)
}

func (r *Router) queryFirmwareSetFirmware(ctx context.Context, firmwareSetID string) ([]*models.ComponentFirmwareVersion, error) {
	mapMods := []qm.QueryMod{
		qm.Where("firmware_set_id=?", firmwareSetID),
		qm.Load(models.ComponentFirmwareSetMapRels.Firmware),
	}

	// query firmware set references
	dbFirmwareSetMap, err := models.ComponentFirmwareSetMaps(mapMods...).All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	firmwares := []*models.ComponentFirmwareVersion{}

	for _, m := range dbFirmwareSetMap {
		if m.R != nil && m.R.Firmware != nil {
			firmwares = append(firmwares, m.R.Firmware)
		}
	}

	return firmwares, nil
}

func (r *Router) serverComponentFirmwareSetCreate(c *gin.Context) {
	var firmwareSetPayload ComponentFirmwareSetRequest

	if err := c.ShouldBindJSON(&firmwareSetPayload); err != nil {
		badRequestResponse(c, "invalid payload: ComponentFirmwareSetCreate{}", err)
		return
	}

	if firmwareSetPayload.Name == "" {
		badRequestResponse(
			c,
			"invalid payload: ComponentFirmwareSetCreate{}",
			errors.Wrap(errSrvComponentPayload, "required attribute not set: Name"),
		)

		return
	}

	// vet and parse firmware uuids
	if len(firmwareSetPayload.ComponentFirmwareUUIDs) == 0 {
		err := errors.Wrap(errComponentFirmwareSetRequest, "expected one or more firmware UUIDs, got none")
		badRequestResponse(
			c,
			"",
			err,
		)

		return
	}

	firmwareUUIDs, err := r.firmwareSetVetFirmwareUUIDsForCreate(c, firmwareSetPayload.ComponentFirmwareUUIDs)
	if err != nil {
		if errors.Is(errDBErr, err) {
			dbErrorResponse(c, err)
			return
		}

		badRequestResponse(c, "", err)

		return
	}

	dbFirmwareSet, err := firmwareSetPayload.toDBModelFirmwareSet()
	if err != nil {
		badRequestResponse(c, "invalid db model: ComponentFirmwareSet", err)
		return
	}

	err = r.firmwareSetCreateTx(c.Request.Context(), dbFirmwareSet, firmwareSetPayload.Attributes, firmwareUUIDs)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, dbFirmwareSet.ID)
}

func (r *Router) firmwareSetVetFirmwareUUIDsForCreate(c *gin.Context, firmwareUUIDs []string) ([]uuid.UUID, error) {
	// validate and collect firmware UUIDs
	vetted := []uuid.UUID{}

	// unique is a map of keys to limit firmware sets to firmwares with unique vendor, version, component attributes.
	unique := map[string]bool{}

	for _, firmwareUUID := range firmwareUUIDs {
		// parse uuid
		firmwareUUIDParsed, err := uuid.Parse(firmwareUUID)
		if err != nil {
			return nil, errors.Wrap(errComponentFirmwareSetRequest, err.Error()+" invalid firmware UUID: "+firmwareUUID)
		}

		// validate component firmware version exists
		firmwareVersion, err := models.FindComponentFirmwareVersion(c.Request.Context(), r.DB, firmwareUUID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, errors.Wrap(err, "firmware object with given UUID does not exist: "+firmwareUUID)
			}

			return nil, errors.Wrap(errDBErr, err.Error())
		}

		// validate firmware is unique based on vendor, version, component attributes
		key := strings.ToLower(firmwareVersion.Vendor) + strings.ToLower(firmwareVersion.Version) + strings.ToLower(firmwareVersion.Component)

		_, exists := unique[key]
		if exists {
			return nil, errors.Wrap(
				errComponentFirmwareSetMap,
				"A firmware set can only reference unique firmware versions based on the vendor, version, component attributes",
			)
		}

		unique[key] = true

		vetted = append(vetted, firmwareUUIDParsed)
	}

	return vetted, nil
}

func (r *Router) firmwareSetCreateTx(ctx context.Context, dbFirmwareSet *models.ComponentFirmwareSet, attrs []Attributes, firmwareUUIDs []uuid.UUID) error {
	// being transaction to insert a new firmware set and its references
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer loggedRollback(r, tx)

	// insert set
	if errInsert := dbFirmwareSet.Insert(ctx, tx, boil.Infer()); errInsert != nil {
		return errInsert
	}

	// insert attributes
	for _, attributes := range attrs {
		dbAttributes := attributes.toDBModelAttributesFirmwareSet()
		dbAttributes.FirmwareSetID = null.StringFrom(dbFirmwareSet.ID)

		err = dbFirmwareSet.AddFirmwareSetAttributesFirmwareSets(ctx, tx, true, dbAttributes)
		if err != nil {
			return err
		}
	}

	// add firmware references
	for _, id := range firmwareUUIDs {
		m := models.ComponentFirmwareSetMap{FirmwareSetID: dbFirmwareSet.ID, FirmwareID: id.String()}

		err = m.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *Router) serverComponentFirmwareSetUpdate(c *gin.Context) {
	dbFirmware, err := r.componentFirmwareSetFromParams(c)
	if err != nil {
		badRequestResponse(c, "invalid payload: ComponentFirmwareSet{}", err)
		return
	}

	var newValues ComponentFirmwareSetRequest
	if errBind := c.ShouldBindJSON(&newValues); errBind != nil {
		badRequestResponse(c, "invalid payload: ComponentFirmwareSet{}", errBind)
		return
	}

	// firmware set ID is expected for updates
	if newValues.ID == uuid.Nil {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errComponentFirmwareSetRequest, "expected a valid firmware set ID, got none"),
		)

		return
	}

	dbFirmwareSet, err := newValues.toDBModelFirmwareSet()
	if err != nil {
		badRequestResponse(c, "invalid db model: ComponentFirmwareSet", err)
		return
	}

	dbAttributesFirmwareSet := make([]*models.AttributesFirmwareSet, 0, len(newValues.Attributes))

	for _, attributes := range newValues.Attributes {
		attr := attributes.toDBModelAttributesFirmwareSet()

		attr.FirmwareSetID = null.StringFrom(newValues.ID.String())
		dbAttributesFirmwareSet = append(dbAttributesFirmwareSet, attr)
	}

	// vet and parse firmware uuids
	var firmwareUUIDs []uuid.UUID

	if len(newValues.ComponentFirmwareUUIDs) > 0 {
		firmwareUUIDs, err = r.firmwareSetVetFirmwareUUIDsForUpdate(c.Request.Context(), dbFirmwareSet, newValues.ComponentFirmwareUUIDs)
		if err != nil {
			if errors.Is(errDBErr, err) {
				dbErrorResponse(c, err)
				return
			}

			badRequestResponse(c, "", err)

			return
		}
	}

	err = r.firmwareSetUpdateTx(c.Request.Context(), dbFirmwareSet, dbAttributesFirmwareSet, firmwareUUIDs)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	updatedResponse(c, dbFirmware.ID)
}

func (r *Router) firmwareSetVetFirmwareUUIDsForUpdate(ctx context.Context, firmwareSet *models.ComponentFirmwareSet, firmwareUUIDs []string) ([]uuid.UUID, error) {
	// firmware uuids are expected
	if len(firmwareUUIDs) == 0 {
		return nil, errors.Wrap(errComponentFirmwareSetRequest, "expected one or more firmware UUIDs, got none")
	}

	// validate and collect firmware UUIDs
	vetted := []uuid.UUID{}

	// unique is a map of keys to limit firmware sets to include only firmwares with,
	// unique vendor, version, component attributes.
	unique := map[string]bool{}

	if len(firmwareUUIDs) == 0 {
		return nil, errors.Wrap(errComponentFirmwareSetRequest, "expected one or more firmware UUIDs, got none")
	}

	for _, firmwareUUID := range firmwareUUIDs {
		// parse uuid
		firmwareUUIDParsed, err := uuid.Parse(firmwareUUID)
		if err != nil {
			return nil, errors.Wrap(errComponentFirmwareSetRequest, err.Error()+"invalid firmware UUID: "+firmwareUUID)
		}

		// validate firmware isn't part of set
		setMap, err := r.firmwareSetMap(ctx, firmwareSet, firmwareUUIDParsed)
		if err != nil {
			return nil, err
		}

		if len(setMap) > 0 {
			return nil, errors.Wrap(
				errComponentFirmwareSetRequest,
				fmt.Sprintf("firmware '%s' exists in firmware set '%s' ", firmwareUUID, firmwareSet.Name),
			)
		}

		// validate component firmware version exists
		firmwareVersion, err := models.FindComponentFirmwareVersion(ctx, r.DB, firmwareUUID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, errors.Wrap(errComponentFirmwareSetRequest, "firmware object with given UUID does not exist: "+firmwareUUID)
			}

			return nil, errors.Wrap(errDBErr, err.Error())
		}

		// validate firmware is unique based on vendor, version, component attributes
		key := strings.ToLower(firmwareVersion.Vendor) + strings.ToLower(firmwareVersion.Version) + strings.ToLower(firmwareVersion.Component)

		_, duplicateVendorModelComponent := unique[key]
		if duplicateVendorModelComponent {
			return nil, errors.Wrap(
				errComponentFirmwareSetMap,
				"A firmware set can only reference unique firmware versions based on the vendor, version, component attributes",
			)
		}

		unique[key] = true

		vetted = append(vetted, firmwareUUIDParsed)
	}

	return vetted, nil
}

func (r *Router) firmwareSetMap(ctx context.Context, firmwareSet *models.ComponentFirmwareSet, firmwareUUID uuid.UUID) ([]*models.ComponentFirmwareSetMap, error) {
	var m []*models.ComponentFirmwareSetMap

	// validate component firmware version does not already exist in map
	query := firmwareSet.FirmwareSetComponentFirmwareSetMaps(qm.Where("firmware_id=?", firmwareUUID))

	m, err := query.All(ctx, r.DB)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return m, errors.Wrap(errDBErr, err.Error())
	}

	return m, nil
}

func (r *Router) firmwareSetUpdateTx(ctx context.Context, fwSetUpdate *models.ComponentFirmwareSet, attrsUpdate models.AttributesFirmwareSetSlice, firmwareUUIDs []uuid.UUID) error {
	// being transaction to update a firmware set and its references
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer loggedRollback(r, tx)

	fwSetCurrent, err := models.FindComponentFirmwareSet(ctx, tx, fwSetUpdate.ID)
	if err != nil {
		return err
	}

	// update name column
	if fwSetUpdate.Name != "" && fwSetUpdate.Name != fwSetCurrent.Name {
		fwSetCurrent.Name = fwSetUpdate.Name
	}

	if _, err := fwSetCurrent.Update(ctx, tx, boil.Infer()); err != nil {
		return err
	}

	// update attributes if newer attributes were given
	if len(attrsUpdate) > 0 {
		if err := r.firmwareSetAttributesUpdate(ctx, tx, fwSetCurrent, attrsUpdate); err != nil {
			return err
		}
	}

	// add new firmware references into map
	for _, id := range firmwareUUIDs {
		m := models.ComponentFirmwareSetMap{FirmwareSetID: fwSetUpdate.ID, FirmwareID: id.String()}

		err := m.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *Router) firmwareSetAttributesUpdate(ctx context.Context, tx *sql.Tx, fwSet *models.ComponentFirmwareSet, attrsUpdate models.AttributesFirmwareSetSlice) error {
	// retrieve current firmware set attributes
	attrsCurrent, err := fwSet.FirmwareSetAttributesFirmwareSets().All(ctx, tx)
	if err != nil {
		return err
	}

	// In all cases a single *models.AttributesFirmwareSet obj holds all the firmware set attributes in the Data field,
	//
	// here we make sure that data is not overwritten by a 'null' or is set to empty.
	//
	// If the client needs to overwrite this data to be empty, Data should be set to an empty JSON - `{}`
	if len(attrsCurrent) == 1 && len(attrsUpdate) == 1 && len(attrsCurrent[0].Data) > 0 {
		if string(attrsUpdate[0].Data) == "null" || len(attrsUpdate[0].Data) == 0 {
			attrsUpdate[0].Data = attrsCurrent[0].Data
		}
	}

	// remove current referenced firmware set attributes
	_, err = attrsCurrent.DeleteAll(ctx, tx)
	if err != nil {
		return err
	}

	// add new firmware set attributes
	return fwSet.AddFirmwareSetAttributesFirmwareSets(ctx, tx, true, attrsUpdate...)
}

func (r *Router) serverComponentFirmwareSetRemoveFirmware(c *gin.Context) {
	firmwareSet, err := r.componentFirmwareSetFromParams(c)
	if err != nil {
		badRequestResponse(c, "invalid payload: ComponentFirmwareSet{}", err)
		return
	}

	var payload ComponentFirmwareSetRequest
	if errBind := c.ShouldBindJSON(&payload); errBind != nil {
		badRequestResponse(c, "invalid payload: ComponentFirmwareSet{}", errBind)
		return
	}

	// firmware set ID is expected for
	if payload.ID == uuid.Nil {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errComponentFirmwareSetRequest, "expected a valid firmware set ID in payload, got none"),
		)

		return
	}

	// firmware set ID URL param is expected to match payload firmware set ID
	if payload.ID.String() != firmwareSet.ID {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errComponentFirmwareSetRequest, "firmware set ID does not match payload ID attribute"),
		)

		return
	}

	// identify firmware set - firmware mappings for removal
	removeMappings := []*models.ComponentFirmwareSetMap{}

	for _, firmwareUUID := range payload.ComponentFirmwareUUIDs {
		// parse uuid
		firmwareUUIDParsed, errUUID := uuid.Parse(firmwareUUID)
		if errUUID != nil {
			badRequestResponse(
				c,
				"invalid firmware UUID: "+firmwareUUID,
				errors.Wrap(errComponentFirmwareSetRequest, errUUID.Error()),
			)

			return
		}

		// validate firmware is part of set
		setMap, errSet := r.firmwareSetMap(c.Request.Context(), firmwareSet, firmwareUUIDParsed)
		if errSet != nil {
			dbErrorResponse(c, errSet)

			return
		}

		if len(setMap) == 0 {
			badRequestResponse(
				c,
				"invalid firmware UUID: "+firmwareUUID,
				errors.Wrap(
					errComponentFirmwareSetRequest,

					fmt.Sprintf("firmware set '%s' does not contain firmware '%s'", firmwareSet.Name, firmwareUUID),
				),
			)

			return
		}

		removeMappings = append(removeMappings, setMap...)
	}

	err = r.firmwareSetDeleteMappingTx(c.Request.Context(), firmwareSet, removeMappings)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	deletedResponse(c)
}

func (r *Router) serverComponentFirmwareSetDelete(c *gin.Context) {
	dbFirmware, err := r.componentFirmwareSetFromParams(c)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	if _, err = dbFirmware.Delete(c.Request.Context(), r.DB); err != nil {
		dbErrorResponse(c, err)
		return
	}

	deletedResponse(c)
}

func (r *Router) componentFirmwareSetFromParams(c *gin.Context) (*models.ComponentFirmwareSet, error) {
	u, err := r.parseUUID(c.Param("uuid"))
	if err != nil {
		return nil, err
	}

	if u == uuid.Nil {
		return nil, errors.Wrap(errComponentFirmwareSetRequest, "expected a valid firmware set UUID")
	}

	firmwareSet, err := models.FindComponentFirmwareSet(c.Request.Context(), r.DB, u.String())
	if err != nil {
		return nil, err
	}

	return firmwareSet, nil
}

func (r *Router) firmwareSetDeleteMappingTx(ctx context.Context, _ *models.ComponentFirmwareSet, removeMappings []*models.ComponentFirmwareSetMap) error {
	// being transaction to insert a new firmware set and its mapping
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer loggedRollback(r, tx)

	for _, mapping := range removeMappings {
		if _, err := mapping.Delete(ctx, r.DB); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// We allow multiple calls for the same firmware-set and server-id because firmware sets are mutable.
// We will update any record in place, or create a new one as needed.
func (r *Router) validateFirmwareSet(c *gin.Context) {
	var payload FirmwareSetValidation
	if err := c.ShouldBindJSON(&payload); err != nil {
		badRequestResponse(c, "invalid validation payload", err)
		return
	}
	ctx := c.Request.Context()
	txn, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		dbErrorResponse2(c, "starting transaction", err)
		return
	}

	doRollback := false
	rollbackFn := func() {
		if doRollback {
			if rbErr := txn.Rollback(); rbErr != nil {
				r.Logger.With(
					zap.Error(rbErr),
				).Warn("rollback error on firmware validation")
				metrics.DBError("rollback firmware validation")
			}
		}
	}
	defer rollbackFn()

	fact := models.FirmwareSetValidationFact{
		TargetServerID: payload.TargetServer.String(),
		FirmwareSetID:  payload.FirmwareSet.String(),
		PerformedOn:    payload.PerformedOn,
	}

	existing, err := models.FirmwareSetValidationFacts(
		models.FirmwareSetValidationFactWhere.FirmwareSetID.EQ(payload.FirmwareSet.String()),
	).One(ctx, txn)

	switch {
	case err == nil:
		fact.ID = existing.ID
		_, updErr := fact.Update(ctx, txn, boil.Infer())
		if updErr != nil {
			r.Logger.With(
				zap.Error(updErr),
				zap.String("firmware.set", payload.FirmwareSet.String()),
				zap.String("target.server", payload.TargetServer.String()),
			).Warn("updating existing firmware validation record")
			metrics.DBError("update firmware validation")
			doRollback = true
			dbErrorResponse2(c, "update firmware validation", updErr)
			return
		}
	case errors.Is(err, sql.ErrNoRows):
		writeErr := fact.Insert(ctx, txn, boil.Infer())
		if writeErr != nil {
			r.Logger.With(
				zap.Error(writeErr),
				zap.String("firmware.set", payload.FirmwareSet.String()),
				zap.String("target.server", payload.TargetServer.String()),
			).Warn("inserting existing firmware validation record")
			metrics.DBError("insert firmware validation")
			doRollback = true
			dbErrorResponse2(c, "insert firmware validation", writeErr)
			return
		}
	default:
		dbErrorResponse2(c, "checking database for existing", err)
		return
	}

	if txErr := txn.Commit(); txErr != nil {
		r.Logger.With(
			zap.Error(txErr),
			zap.String("firmware.set", payload.FirmwareSet.String()),
			zap.String("target.server", payload.TargetServer.String()),
		).Warn("commit firmware validation record")
		doRollback = true
		metrics.DBError("commit firmware validation transaction")
		dbErrorResponse2(c, "commit firmware validation transaction", txErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
