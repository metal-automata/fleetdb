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

// ConfigSet is an object representing the database table.
type ConfigSet struct {
	ID        string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Version   null.String `boil:"version" json:"version,omitempty" toml:"version" yaml:"version,omitempty"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *configSetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L configSetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ConfigSetColumns = struct {
	ID        string
	Name      string
	Version   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Version:   "version",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var ConfigSetTableColumns = struct {
	ID        string
	Name      string
	Version   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "config_sets.id",
	Name:      "config_sets.name",
	Version:   "config_sets.version",
	CreatedAt: "config_sets.created_at",
	UpdatedAt: "config_sets.updated_at",
}

// Generated where

var ConfigSetWhere = struct {
	ID        whereHelperstring
	Name      whereHelperstring
	Version   whereHelpernull_String
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"config_sets\".\"id\""},
	Name:      whereHelperstring{field: "\"config_sets\".\"name\""},
	Version:   whereHelpernull_String{field: "\"config_sets\".\"version\""},
	CreatedAt: whereHelpernull_Time{field: "\"config_sets\".\"created_at\""},
	UpdatedAt: whereHelpernull_Time{field: "\"config_sets\".\"updated_at\""},
}

// ConfigSetRels is where relationship names are stored.
var ConfigSetRels = struct {
	FKConfigSetConfigComponents string
}{
	FKConfigSetConfigComponents: "FKConfigSetConfigComponents",
}

// configSetR is where relationships are stored.
type configSetR struct {
	FKConfigSetConfigComponents ConfigComponentSlice `boil:"FKConfigSetConfigComponents" json:"FKConfigSetConfigComponents" toml:"FKConfigSetConfigComponents" yaml:"FKConfigSetConfigComponents"`
}

// NewStruct creates a new relationship struct
func (*configSetR) NewStruct() *configSetR {
	return &configSetR{}
}

func (r *configSetR) GetFKConfigSetConfigComponents() ConfigComponentSlice {
	if r == nil {
		return nil
	}
	return r.FKConfigSetConfigComponents
}

// configSetL is where Load methods for each relationship are stored.
type configSetL struct{}

var (
	configSetAllColumns            = []string{"id", "name", "version", "created_at", "updated_at"}
	configSetColumnsWithoutDefault = []string{"name"}
	configSetColumnsWithDefault    = []string{"id", "version", "created_at", "updated_at"}
	configSetPrimaryKeyColumns     = []string{"id"}
	configSetGeneratedColumns      = []string{}
)

type (
	// ConfigSetSlice is an alias for a slice of pointers to ConfigSet.
	// This should almost always be used instead of []ConfigSet.
	ConfigSetSlice []*ConfigSet
	// ConfigSetHook is the signature for custom ConfigSet hook methods
	ConfigSetHook func(context.Context, boil.ContextExecutor, *ConfigSet) error

	configSetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	configSetType                 = reflect.TypeOf(&ConfigSet{})
	configSetMapping              = queries.MakeStructMapping(configSetType)
	configSetPrimaryKeyMapping, _ = queries.BindMapping(configSetType, configSetMapping, configSetPrimaryKeyColumns)
	configSetInsertCacheMut       sync.RWMutex
	configSetInsertCache          = make(map[string]insertCache)
	configSetUpdateCacheMut       sync.RWMutex
	configSetUpdateCache          = make(map[string]updateCache)
	configSetUpsertCacheMut       sync.RWMutex
	configSetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var configSetAfterSelectHooks []ConfigSetHook

var configSetBeforeInsertHooks []ConfigSetHook
var configSetAfterInsertHooks []ConfigSetHook

var configSetBeforeUpdateHooks []ConfigSetHook
var configSetAfterUpdateHooks []ConfigSetHook

var configSetBeforeDeleteHooks []ConfigSetHook
var configSetAfterDeleteHooks []ConfigSetHook

var configSetBeforeUpsertHooks []ConfigSetHook
var configSetAfterUpsertHooks []ConfigSetHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ConfigSet) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ConfigSet) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ConfigSet) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ConfigSet) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ConfigSet) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ConfigSet) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ConfigSet) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ConfigSet) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ConfigSet) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configSetAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddConfigSetHook registers your hook function for all future operations.
func AddConfigSetHook(hookPoint boil.HookPoint, configSetHook ConfigSetHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		configSetAfterSelectHooks = append(configSetAfterSelectHooks, configSetHook)
	case boil.BeforeInsertHook:
		configSetBeforeInsertHooks = append(configSetBeforeInsertHooks, configSetHook)
	case boil.AfterInsertHook:
		configSetAfterInsertHooks = append(configSetAfterInsertHooks, configSetHook)
	case boil.BeforeUpdateHook:
		configSetBeforeUpdateHooks = append(configSetBeforeUpdateHooks, configSetHook)
	case boil.AfterUpdateHook:
		configSetAfterUpdateHooks = append(configSetAfterUpdateHooks, configSetHook)
	case boil.BeforeDeleteHook:
		configSetBeforeDeleteHooks = append(configSetBeforeDeleteHooks, configSetHook)
	case boil.AfterDeleteHook:
		configSetAfterDeleteHooks = append(configSetAfterDeleteHooks, configSetHook)
	case boil.BeforeUpsertHook:
		configSetBeforeUpsertHooks = append(configSetBeforeUpsertHooks, configSetHook)
	case boil.AfterUpsertHook:
		configSetAfterUpsertHooks = append(configSetAfterUpsertHooks, configSetHook)
	}
}

// One returns a single configSet record from the query.
func (q configSetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ConfigSet, error) {
	o := &ConfigSet{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for config_sets")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ConfigSet records from the query.
func (q configSetQuery) All(ctx context.Context, exec boil.ContextExecutor) (ConfigSetSlice, error) {
	var o []*ConfigSet

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ConfigSet slice")
	}

	if len(configSetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ConfigSet records in the query.
func (q configSetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count config_sets rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q configSetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if config_sets exists")
	}

	return count > 0, nil
}

// FKConfigSetConfigComponents retrieves all the config_component's ConfigComponents with an executor via fk_config_set_id column.
func (o *ConfigSet) FKConfigSetConfigComponents(mods ...qm.QueryMod) configComponentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"config_components\".\"fk_config_set_id\"=?", o.ID),
	)

	return ConfigComponents(queryMods...)
}

// LoadFKConfigSetConfigComponents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (configSetL) LoadFKConfigSetConfigComponents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeConfigSet interface{}, mods queries.Applicator) error {
	var slice []*ConfigSet
	var object *ConfigSet

	if singular {
		var ok bool
		object, ok = maybeConfigSet.(*ConfigSet)
		if !ok {
			object = new(ConfigSet)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeConfigSet)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeConfigSet))
			}
		}
	} else {
		s, ok := maybeConfigSet.(*[]*ConfigSet)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeConfigSet)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeConfigSet))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &configSetR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &configSetR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`config_components`),
		qm.WhereIn(`config_components.fk_config_set_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load config_components")
	}

	var resultSlice []*ConfigComponent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice config_components")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on config_components")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for config_components")
	}

	if len(configComponentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.FKConfigSetConfigComponents = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &configComponentR{}
			}
			foreign.R.FKConfigSet = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.FKConfigSetID {
				local.R.FKConfigSetConfigComponents = append(local.R.FKConfigSetConfigComponents, foreign)
				if foreign.R == nil {
					foreign.R = &configComponentR{}
				}
				foreign.R.FKConfigSet = local
				break
			}
		}
	}

	return nil
}

// AddFKConfigSetConfigComponents adds the given related objects to the existing relationships
// of the config_set, optionally inserting them as new records.
// Appends related to o.R.FKConfigSetConfigComponents.
// Sets related.R.FKConfigSet appropriately.
func (o *ConfigSet) AddFKConfigSetConfigComponents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ConfigComponent) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FKConfigSetID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"config_components\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"fk_config_set_id"}),
				strmangle.WhereClause("\"", "\"", 2, configComponentPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FKConfigSetID = o.ID
		}
	}

	if o.R == nil {
		o.R = &configSetR{
			FKConfigSetConfigComponents: related,
		}
	} else {
		o.R.FKConfigSetConfigComponents = append(o.R.FKConfigSetConfigComponents, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &configComponentR{
				FKConfigSet: o,
			}
		} else {
			rel.R.FKConfigSet = o
		}
	}
	return nil
}

// ConfigSets retrieves all the records using an executor.
func ConfigSets(mods ...qm.QueryMod) configSetQuery {
	mods = append(mods, qm.From("\"config_sets\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"config_sets\".*"})
	}

	return configSetQuery{q}
}

// FindConfigSet retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindConfigSet(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ConfigSet, error) {
	configSetObj := &ConfigSet{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"config_sets\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, configSetObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from config_sets")
	}

	if err = configSetObj.doAfterSelectHooks(ctx, exec); err != nil {
		return configSetObj, err
	}

	return configSetObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ConfigSet) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no config_sets provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(configSetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	configSetInsertCacheMut.RLock()
	cache, cached := configSetInsertCache[key]
	configSetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			configSetAllColumns,
			configSetColumnsWithDefault,
			configSetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(configSetType, configSetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(configSetType, configSetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"config_sets\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"config_sets\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into config_sets")
	}

	if !cached {
		configSetInsertCacheMut.Lock()
		configSetInsertCache[key] = cache
		configSetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ConfigSet.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ConfigSet) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	configSetUpdateCacheMut.RLock()
	cache, cached := configSetUpdateCache[key]
	configSetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			configSetAllColumns,
			configSetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update config_sets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"config_sets\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, configSetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(configSetType, configSetMapping, append(wl, configSetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update config_sets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for config_sets")
	}

	if !cached {
		configSetUpdateCacheMut.Lock()
		configSetUpdateCache[key] = cache
		configSetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q configSetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for config_sets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for config_sets")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ConfigSetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configSetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"config_sets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, configSetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in configSet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all configSet")
	}
	return rowsAff, nil
}

// Delete deletes a single ConfigSet record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ConfigSet) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ConfigSet provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), configSetPrimaryKeyMapping)
	sql := "DELETE FROM \"config_sets\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from config_sets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for config_sets")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q configSetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no configSetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from config_sets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for config_sets")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ConfigSetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(configSetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configSetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"config_sets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, configSetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from configSet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for config_sets")
	}

	if len(configSetAfterDeleteHooks) != 0 {
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
func (o *ConfigSet) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindConfigSet(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ConfigSetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ConfigSetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configSetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"config_sets\".* FROM \"config_sets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, configSetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ConfigSetSlice")
	}

	*o = slice

	return nil
}

// ConfigSetExists checks if the ConfigSet row exists.
func ConfigSetExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"config_sets\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if config_sets exists")
	}

	return exists, nil
}

// Exists checks if the ConfigSet row exists.
func (o *ConfigSet) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ConfigSetExists(ctx, exec, o.ID)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ConfigSet) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no config_sets provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(configSetColumnsWithDefault, o)

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

	configSetUpsertCacheMut.RLock()
	cache, cached := configSetUpsertCache[key]
	configSetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			configSetAllColumns,
			configSetColumnsWithDefault,
			configSetColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			configSetAllColumns,
			configSetPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert config_sets, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(configSetPrimaryKeyColumns))
			copy(conflict, configSetPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"config_sets\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(configSetType, configSetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(configSetType, configSetMapping, ret)
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

	if boil.DebugMode {
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert config_sets")
	}

	if !cached {
		configSetUpsertCacheMut.Lock()
		configSetUpsertCache[key] = cache
		configSetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
