// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ComponentStatus is an object representing the database table.
type ComponentStatus struct {
	ID                string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ServerComponentID string      `boil:"server_component_id" json:"server_component_id" toml:"server_component_id" yaml:"server_component_id"`
	Health            string      `boil:"health" json:"health" toml:"health" yaml:"health"`
	State             string      `boil:"state" json:"state" toml:"state" yaml:"state"`
	Info              null.String `boil:"info" json:"info,omitempty" toml:"info" yaml:"info,omitempty"`
	CreatedAt         null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt         null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *componentStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L componentStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ComponentStatusColumns = struct {
	ID                string
	ServerComponentID string
	Health            string
	State             string
	Info              string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	ServerComponentID: "server_component_id",
	Health:            "health",
	State:             "state",
	Info:              "info",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

var ComponentStatusTableColumns = struct {
	ID                string
	ServerComponentID string
	Health            string
	State             string
	Info              string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "component_status.id",
	ServerComponentID: "component_status.server_component_id",
	Health:            "component_status.health",
	State:             "component_status.state",
	Info:              "component_status.info",
	CreatedAt:         "component_status.created_at",
	UpdatedAt:         "component_status.updated_at",
}

// Generated where

var ComponentStatusWhere = struct {
	ID                whereHelperstring
	ServerComponentID whereHelperstring
	Health            whereHelperstring
	State             whereHelperstring
	Info              whereHelpernull_String
	CreatedAt         whereHelpernull_Time
	UpdatedAt         whereHelpernull_Time
}{
	ID:                whereHelperstring{field: "\"component_status\".\"id\""},
	ServerComponentID: whereHelperstring{field: "\"component_status\".\"server_component_id\""},
	Health:            whereHelperstring{field: "\"component_status\".\"health\""},
	State:             whereHelperstring{field: "\"component_status\".\"state\""},
	Info:              whereHelpernull_String{field: "\"component_status\".\"info\""},
	CreatedAt:         whereHelpernull_Time{field: "\"component_status\".\"created_at\""},
	UpdatedAt:         whereHelpernull_Time{field: "\"component_status\".\"updated_at\""},
}

// ComponentStatusRels is where relationship names are stored.
var ComponentStatusRels = struct {
	ServerComponent string
}{
	ServerComponent: "ServerComponent",
}

// componentStatusR is where relationships are stored.
type componentStatusR struct {
	ServerComponent *ServerComponent `boil:"ServerComponent" json:"ServerComponent" toml:"ServerComponent" yaml:"ServerComponent"`
}

// NewStruct creates a new relationship struct
func (*componentStatusR) NewStruct() *componentStatusR {
	return &componentStatusR{}
}

func (r *componentStatusR) GetServerComponent() *ServerComponent {
	if r == nil {
		return nil
	}
	return r.ServerComponent
}

// componentStatusL is where Load methods for each relationship are stored.
type componentStatusL struct{}

var (
	componentStatusAllColumns            = []string{"id", "server_component_id", "health", "state", "info", "created_at", "updated_at"}
	componentStatusColumnsWithoutDefault = []string{"server_component_id", "health", "state"}
	componentStatusColumnsWithDefault    = []string{"id", "info", "created_at", "updated_at"}
	componentStatusPrimaryKeyColumns     = []string{"id"}
	componentStatusGeneratedColumns      = []string{}
)

type (
	// ComponentStatusSlice is an alias for a slice of pointers to ComponentStatus.
	// This should almost always be used instead of []ComponentStatus.
	ComponentStatusSlice []*ComponentStatus
	// ComponentStatusHook is the signature for custom ComponentStatus hook methods
	ComponentStatusHook func(context.Context, boil.ContextExecutor, *ComponentStatus) error

	componentStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	componentStatusType                 = reflect.TypeOf(&ComponentStatus{})
	componentStatusMapping              = queries.MakeStructMapping(componentStatusType)
	componentStatusPrimaryKeyMapping, _ = queries.BindMapping(componentStatusType, componentStatusMapping, componentStatusPrimaryKeyColumns)
	componentStatusInsertCacheMut       sync.RWMutex
	componentStatusInsertCache          = make(map[string]insertCache)
	componentStatusUpdateCacheMut       sync.RWMutex
	componentStatusUpdateCache          = make(map[string]updateCache)
	componentStatusUpsertCacheMut       sync.RWMutex
	componentStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var componentStatusAfterSelectHooks []ComponentStatusHook

var componentStatusBeforeInsertHooks []ComponentStatusHook
var componentStatusAfterInsertHooks []ComponentStatusHook

var componentStatusBeforeUpdateHooks []ComponentStatusHook
var componentStatusAfterUpdateHooks []ComponentStatusHook

var componentStatusBeforeDeleteHooks []ComponentStatusHook
var componentStatusAfterDeleteHooks []ComponentStatusHook

var componentStatusBeforeUpsertHooks []ComponentStatusHook
var componentStatusAfterUpsertHooks []ComponentStatusHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ComponentStatus) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ComponentStatus) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ComponentStatus) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ComponentStatus) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ComponentStatus) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ComponentStatus) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ComponentStatus) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ComponentStatus) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ComponentStatus) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range componentStatusAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddComponentStatusHook registers your hook function for all future operations.
func AddComponentStatusHook(hookPoint boil.HookPoint, componentStatusHook ComponentStatusHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		componentStatusAfterSelectHooks = append(componentStatusAfterSelectHooks, componentStatusHook)
	case boil.BeforeInsertHook:
		componentStatusBeforeInsertHooks = append(componentStatusBeforeInsertHooks, componentStatusHook)
	case boil.AfterInsertHook:
		componentStatusAfterInsertHooks = append(componentStatusAfterInsertHooks, componentStatusHook)
	case boil.BeforeUpdateHook:
		componentStatusBeforeUpdateHooks = append(componentStatusBeforeUpdateHooks, componentStatusHook)
	case boil.AfterUpdateHook:
		componentStatusAfterUpdateHooks = append(componentStatusAfterUpdateHooks, componentStatusHook)
	case boil.BeforeDeleteHook:
		componentStatusBeforeDeleteHooks = append(componentStatusBeforeDeleteHooks, componentStatusHook)
	case boil.AfterDeleteHook:
		componentStatusAfterDeleteHooks = append(componentStatusAfterDeleteHooks, componentStatusHook)
	case boil.BeforeUpsertHook:
		componentStatusBeforeUpsertHooks = append(componentStatusBeforeUpsertHooks, componentStatusHook)
	case boil.AfterUpsertHook:
		componentStatusAfterUpsertHooks = append(componentStatusAfterUpsertHooks, componentStatusHook)
	}
}

// One returns a single componentStatus record from the query.
func (q componentStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ComponentStatus, error) {
	o := &ComponentStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for component_status")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ComponentStatus records from the query.
func (q componentStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (ComponentStatusSlice, error) {
	var o []*ComponentStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ComponentStatus slice")
	}

	if len(componentStatusAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ComponentStatus records in the query.
func (q componentStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count component_status rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q componentStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if component_status exists")
	}

	return count > 0, nil
}

// ServerComponent pointed to by the foreign key.
func (o *ComponentStatus) ServerComponent(mods ...qm.QueryMod) serverComponentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ServerComponentID),
	}

	queryMods = append(queryMods, mods...)

	return ServerComponents(queryMods...)
}

// LoadServerComponent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (componentStatusL) LoadServerComponent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeComponentStatus interface{}, mods queries.Applicator) error {
	var slice []*ComponentStatus
	var object *ComponentStatus

	if singular {
		var ok bool
		object, ok = maybeComponentStatus.(*ComponentStatus)
		if !ok {
			object = new(ComponentStatus)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeComponentStatus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeComponentStatus))
			}
		}
	} else {
		s, ok := maybeComponentStatus.(*[]*ComponentStatus)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeComponentStatus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeComponentStatus))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &componentStatusR{}
		}
		args = append(args, object.ServerComponentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &componentStatusR{}
			}

			for _, a := range args {
				if a == obj.ServerComponentID {
					continue Outer
				}
			}

			args = append(args, obj.ServerComponentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`server_components`),
		qm.WhereIn(`server_components.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ServerComponent")
	}

	var resultSlice []*ServerComponent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ServerComponent")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for server_components")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for server_components")
	}

	if len(serverComponentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ServerComponent = foreign
		if foreign.R == nil {
			foreign.R = &serverComponentR{}
		}
		foreign.R.ComponentStatus = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ServerComponentID == foreign.ID {
				local.R.ServerComponent = foreign
				if foreign.R == nil {
					foreign.R = &serverComponentR{}
				}
				foreign.R.ComponentStatus = local
				break
			}
		}
	}

	return nil
}

// SetServerComponent of the componentStatus to the related item.
// Sets o.R.ServerComponent to related.
// Adds o to related.R.ComponentStatus.
func (o *ComponentStatus) SetServerComponent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ServerComponent) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"component_status\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"server_component_id"}),
		strmangle.WhereClause("\"", "\"", 2, componentStatusPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ServerComponentID = related.ID
	if o.R == nil {
		o.R = &componentStatusR{
			ServerComponent: related,
		}
	} else {
		o.R.ServerComponent = related
	}

	if related.R == nil {
		related.R = &serverComponentR{
			ComponentStatus: o,
		}
	} else {
		related.R.ComponentStatus = o
	}

	return nil
}

// ComponentStatuses retrieves all the records using an executor.
func ComponentStatuses(mods ...qm.QueryMod) componentStatusQuery {
	mods = append(mods, qm.From("\"component_status\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"component_status\".*"})
	}

	return componentStatusQuery{q}
}

// FindComponentStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindComponentStatus(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ComponentStatus, error) {
	componentStatusObj := &ComponentStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"component_status\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, componentStatusObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from component_status")
	}

	if err = componentStatusObj.doAfterSelectHooks(ctx, exec); err != nil {
		return componentStatusObj, err
	}

	return componentStatusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ComponentStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no component_status provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(componentStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	componentStatusInsertCacheMut.RLock()
	cache, cached := componentStatusInsertCache[key]
	componentStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			componentStatusAllColumns,
			componentStatusColumnsWithDefault,
			componentStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(componentStatusType, componentStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(componentStatusType, componentStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"component_status\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"component_status\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into component_status")
	}

	if !cached {
		componentStatusInsertCacheMut.Lock()
		componentStatusInsertCache[key] = cache
		componentStatusInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ComponentStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ComponentStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	componentStatusUpdateCacheMut.RLock()
	cache, cached := componentStatusUpdateCache[key]
	componentStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			componentStatusAllColumns,
			componentStatusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update component_status, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"component_status\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, componentStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(componentStatusType, componentStatusMapping, append(wl, componentStatusPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update component_status row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for component_status")
	}

	if !cached {
		componentStatusUpdateCacheMut.Lock()
		componentStatusUpdateCache[key] = cache
		componentStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q componentStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for component_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for component_status")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ComponentStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), componentStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"component_status\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, componentStatusPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in componentStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all componentStatus")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ComponentStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no component_status provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(componentStatusColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	componentStatusUpsertCacheMut.RLock()
	cache, cached := componentStatusUpsertCache[key]
	componentStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			componentStatusAllColumns,
			componentStatusColumnsWithDefault,
			componentStatusColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			componentStatusAllColumns,
			componentStatusPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert component_status, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(componentStatusPrimaryKeyColumns))
			copy(conflict, componentStatusPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"component_status\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(componentStatusType, componentStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(componentStatusType, componentStatusMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert component_status")
	}

	if !cached {
		componentStatusUpsertCacheMut.Lock()
		componentStatusUpsertCache[key] = cache
		componentStatusUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ComponentStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ComponentStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ComponentStatus provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), componentStatusPrimaryKeyMapping)
	sql := "DELETE FROM \"component_status\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from component_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for component_status")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q componentStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no componentStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from component_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for component_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ComponentStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(componentStatusBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), componentStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"component_status\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, componentStatusPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from componentStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for component_status")
	}

	if len(componentStatusAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ComponentStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindComponentStatus(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ComponentStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ComponentStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), componentStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"component_status\".* FROM \"component_status\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, componentStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ComponentStatusSlice")
	}

	*o = slice

	return nil
}

// ComponentStatusExists checks if the ComponentStatus row exists.
func ComponentStatusExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"component_status\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if component_status exists")
	}

	return exists, nil
}

// Exists checks if the ComponentStatus row exists.
func (o *ComponentStatus) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ComponentStatusExists(ctx, exec, o.ID)
}
