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

// ConfigComponentSetting is an object representing the database table.
type ConfigComponentSetting struct {
	ID            string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	FKComponentID string    `boil:"fk_component_id" json:"fk_component_id" toml:"fk_component_id" yaml:"fk_component_id"`
	SettingsKey   string    `boil:"settings_key" json:"settings_key" toml:"settings_key" yaml:"settings_key"`
	SettingsValue string    `boil:"settings_value" json:"settings_value" toml:"settings_value" yaml:"settings_value"`
	Custom        null.JSON `boil:"custom" json:"custom,omitempty" toml:"custom" yaml:"custom,omitempty"`
	CreatedAt     null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *configComponentSettingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L configComponentSettingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ConfigComponentSettingColumns = struct {
	ID            string
	FKComponentID string
	SettingsKey   string
	SettingsValue string
	Custom        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	FKComponentID: "fk_component_id",
	SettingsKey:   "settings_key",
	SettingsValue: "settings_value",
	Custom:        "custom",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var ConfigComponentSettingTableColumns = struct {
	ID            string
	FKComponentID string
	SettingsKey   string
	SettingsValue string
	Custom        string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "config_component_settings.id",
	FKComponentID: "config_component_settings.fk_component_id",
	SettingsKey:   "config_component_settings.settings_key",
	SettingsValue: "config_component_settings.settings_value",
	Custom:        "config_component_settings.custom",
	CreatedAt:     "config_component_settings.created_at",
	UpdatedAt:     "config_component_settings.updated_at",
}

// Generated where

type whereHelpernull_JSON struct{ field string }

func (w whereHelpernull_JSON) EQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_JSON) NEQ(x null.JSON) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_JSON) LT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_JSON) LTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_JSON) GT(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_JSON) GTE(x null.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_JSON) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_JSON) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ConfigComponentSettingWhere = struct {
	ID            whereHelperstring
	FKComponentID whereHelperstring
	SettingsKey   whereHelperstring
	SettingsValue whereHelperstring
	Custom        whereHelpernull_JSON
	CreatedAt     whereHelpernull_Time
	UpdatedAt     whereHelpernull_Time
}{
	ID:            whereHelperstring{field: "\"config_component_settings\".\"id\""},
	FKComponentID: whereHelperstring{field: "\"config_component_settings\".\"fk_component_id\""},
	SettingsKey:   whereHelperstring{field: "\"config_component_settings\".\"settings_key\""},
	SettingsValue: whereHelperstring{field: "\"config_component_settings\".\"settings_value\""},
	Custom:        whereHelpernull_JSON{field: "\"config_component_settings\".\"custom\""},
	CreatedAt:     whereHelpernull_Time{field: "\"config_component_settings\".\"created_at\""},
	UpdatedAt:     whereHelpernull_Time{field: "\"config_component_settings\".\"updated_at\""},
}

// ConfigComponentSettingRels is where relationship names are stored.
var ConfigComponentSettingRels = struct {
	FKComponent string
}{
	FKComponent: "FKComponent",
}

// configComponentSettingR is where relationships are stored.
type configComponentSettingR struct {
	FKComponent *ConfigComponent `boil:"FKComponent" json:"FKComponent" toml:"FKComponent" yaml:"FKComponent"`
}

// NewStruct creates a new relationship struct
func (*configComponentSettingR) NewStruct() *configComponentSettingR {
	return &configComponentSettingR{}
}

func (r *configComponentSettingR) GetFKComponent() *ConfigComponent {
	if r == nil {
		return nil
	}
	return r.FKComponent
}

// configComponentSettingL is where Load methods for each relationship are stored.
type configComponentSettingL struct{}

var (
	configComponentSettingAllColumns            = []string{"id", "fk_component_id", "settings_key", "settings_value", "custom", "created_at", "updated_at"}
	configComponentSettingColumnsWithoutDefault = []string{"fk_component_id", "settings_key", "settings_value"}
	configComponentSettingColumnsWithDefault    = []string{"id", "custom", "created_at", "updated_at"}
	configComponentSettingPrimaryKeyColumns     = []string{"id"}
	configComponentSettingGeneratedColumns      = []string{}
)

type (
	// ConfigComponentSettingSlice is an alias for a slice of pointers to ConfigComponentSetting.
	// This should almost always be used instead of []ConfigComponentSetting.
	ConfigComponentSettingSlice []*ConfigComponentSetting
	// ConfigComponentSettingHook is the signature for custom ConfigComponentSetting hook methods
	ConfigComponentSettingHook func(context.Context, boil.ContextExecutor, *ConfigComponentSetting) error

	configComponentSettingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	configComponentSettingType                 = reflect.TypeOf(&ConfigComponentSetting{})
	configComponentSettingMapping              = queries.MakeStructMapping(configComponentSettingType)
	configComponentSettingPrimaryKeyMapping, _ = queries.BindMapping(configComponentSettingType, configComponentSettingMapping, configComponentSettingPrimaryKeyColumns)
	configComponentSettingInsertCacheMut       sync.RWMutex
	configComponentSettingInsertCache          = make(map[string]insertCache)
	configComponentSettingUpdateCacheMut       sync.RWMutex
	configComponentSettingUpdateCache          = make(map[string]updateCache)
	configComponentSettingUpsertCacheMut       sync.RWMutex
	configComponentSettingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var configComponentSettingAfterSelectHooks []ConfigComponentSettingHook

var configComponentSettingBeforeInsertHooks []ConfigComponentSettingHook
var configComponentSettingAfterInsertHooks []ConfigComponentSettingHook

var configComponentSettingBeforeUpdateHooks []ConfigComponentSettingHook
var configComponentSettingAfterUpdateHooks []ConfigComponentSettingHook

var configComponentSettingBeforeDeleteHooks []ConfigComponentSettingHook
var configComponentSettingAfterDeleteHooks []ConfigComponentSettingHook

var configComponentSettingBeforeUpsertHooks []ConfigComponentSettingHook
var configComponentSettingAfterUpsertHooks []ConfigComponentSettingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ConfigComponentSetting) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ConfigComponentSetting) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ConfigComponentSetting) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ConfigComponentSetting) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ConfigComponentSetting) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ConfigComponentSetting) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ConfigComponentSetting) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ConfigComponentSetting) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ConfigComponentSetting) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range configComponentSettingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddConfigComponentSettingHook registers your hook function for all future operations.
func AddConfigComponentSettingHook(hookPoint boil.HookPoint, configComponentSettingHook ConfigComponentSettingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		configComponentSettingAfterSelectHooks = append(configComponentSettingAfterSelectHooks, configComponentSettingHook)
	case boil.BeforeInsertHook:
		configComponentSettingBeforeInsertHooks = append(configComponentSettingBeforeInsertHooks, configComponentSettingHook)
	case boil.AfterInsertHook:
		configComponentSettingAfterInsertHooks = append(configComponentSettingAfterInsertHooks, configComponentSettingHook)
	case boil.BeforeUpdateHook:
		configComponentSettingBeforeUpdateHooks = append(configComponentSettingBeforeUpdateHooks, configComponentSettingHook)
	case boil.AfterUpdateHook:
		configComponentSettingAfterUpdateHooks = append(configComponentSettingAfterUpdateHooks, configComponentSettingHook)
	case boil.BeforeDeleteHook:
		configComponentSettingBeforeDeleteHooks = append(configComponentSettingBeforeDeleteHooks, configComponentSettingHook)
	case boil.AfterDeleteHook:
		configComponentSettingAfterDeleteHooks = append(configComponentSettingAfterDeleteHooks, configComponentSettingHook)
	case boil.BeforeUpsertHook:
		configComponentSettingBeforeUpsertHooks = append(configComponentSettingBeforeUpsertHooks, configComponentSettingHook)
	case boil.AfterUpsertHook:
		configComponentSettingAfterUpsertHooks = append(configComponentSettingAfterUpsertHooks, configComponentSettingHook)
	}
}

// One returns a single configComponentSetting record from the query.
func (q configComponentSettingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ConfigComponentSetting, error) {
	o := &ConfigComponentSetting{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for config_component_settings")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ConfigComponentSetting records from the query.
func (q configComponentSettingQuery) All(ctx context.Context, exec boil.ContextExecutor) (ConfigComponentSettingSlice, error) {
	var o []*ConfigComponentSetting

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ConfigComponentSetting slice")
	}

	if len(configComponentSettingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ConfigComponentSetting records in the query.
func (q configComponentSettingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count config_component_settings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q configComponentSettingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if config_component_settings exists")
	}

	return count > 0, nil
}

// FKComponent pointed to by the foreign key.
func (o *ConfigComponentSetting) FKComponent(mods ...qm.QueryMod) configComponentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.FKComponentID),
	}

	queryMods = append(queryMods, mods...)

	return ConfigComponents(queryMods...)
}

// LoadFKComponent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (configComponentSettingL) LoadFKComponent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeConfigComponentSetting interface{}, mods queries.Applicator) error {
	var slice []*ConfigComponentSetting
	var object *ConfigComponentSetting

	if singular {
		var ok bool
		object, ok = maybeConfigComponentSetting.(*ConfigComponentSetting)
		if !ok {
			object = new(ConfigComponentSetting)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeConfigComponentSetting)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeConfigComponentSetting))
			}
		}
	} else {
		s, ok := maybeConfigComponentSetting.(*[]*ConfigComponentSetting)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeConfigComponentSetting)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeConfigComponentSetting))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &configComponentSettingR{}
		}
		args = append(args, object.FKComponentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &configComponentSettingR{}
			}

			for _, a := range args {
				if a == obj.FKComponentID {
					continue Outer
				}
			}

			args = append(args, obj.FKComponentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`config_components`),
		qm.WhereIn(`config_components.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ConfigComponent")
	}

	var resultSlice []*ConfigComponent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ConfigComponent")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for config_components")
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

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.FKComponent = foreign
		if foreign.R == nil {
			foreign.R = &configComponentR{}
		}
		foreign.R.FKComponentConfigComponentSettings = append(foreign.R.FKComponentConfigComponentSettings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FKComponentID == foreign.ID {
				local.R.FKComponent = foreign
				if foreign.R == nil {
					foreign.R = &configComponentR{}
				}
				foreign.R.FKComponentConfigComponentSettings = append(foreign.R.FKComponentConfigComponentSettings, local)
				break
			}
		}
	}

	return nil
}

// SetFKComponent of the configComponentSetting to the related item.
// Sets o.R.FKComponent to related.
// Adds o to related.R.FKComponentConfigComponentSettings.
func (o *ConfigComponentSetting) SetFKComponent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ConfigComponent) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"config_component_settings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"fk_component_id"}),
		strmangle.WhereClause("\"", "\"", 2, configComponentSettingPrimaryKeyColumns),
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

	o.FKComponentID = related.ID
	if o.R == nil {
		o.R = &configComponentSettingR{
			FKComponent: related,
		}
	} else {
		o.R.FKComponent = related
	}

	if related.R == nil {
		related.R = &configComponentR{
			FKComponentConfigComponentSettings: ConfigComponentSettingSlice{o},
		}
	} else {
		related.R.FKComponentConfigComponentSettings = append(related.R.FKComponentConfigComponentSettings, o)
	}

	return nil
}

// ConfigComponentSettings retrieves all the records using an executor.
func ConfigComponentSettings(mods ...qm.QueryMod) configComponentSettingQuery {
	mods = append(mods, qm.From("\"config_component_settings\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"config_component_settings\".*"})
	}

	return configComponentSettingQuery{q}
}

// FindConfigComponentSetting retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindConfigComponentSetting(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ConfigComponentSetting, error) {
	configComponentSettingObj := &ConfigComponentSetting{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"config_component_settings\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, configComponentSettingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from config_component_settings")
	}

	if err = configComponentSettingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return configComponentSettingObj, err
	}

	return configComponentSettingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ConfigComponentSetting) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no config_component_settings provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(configComponentSettingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	configComponentSettingInsertCacheMut.RLock()
	cache, cached := configComponentSettingInsertCache[key]
	configComponentSettingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			configComponentSettingAllColumns,
			configComponentSettingColumnsWithDefault,
			configComponentSettingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(configComponentSettingType, configComponentSettingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(configComponentSettingType, configComponentSettingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"config_component_settings\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"config_component_settings\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into config_component_settings")
	}

	if !cached {
		configComponentSettingInsertCacheMut.Lock()
		configComponentSettingInsertCache[key] = cache
		configComponentSettingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ConfigComponentSetting.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ConfigComponentSetting) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	configComponentSettingUpdateCacheMut.RLock()
	cache, cached := configComponentSettingUpdateCache[key]
	configComponentSettingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			configComponentSettingAllColumns,
			configComponentSettingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update config_component_settings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"config_component_settings\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, configComponentSettingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(configComponentSettingType, configComponentSettingMapping, append(wl, configComponentSettingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update config_component_settings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for config_component_settings")
	}

	if !cached {
		configComponentSettingUpdateCacheMut.Lock()
		configComponentSettingUpdateCache[key] = cache
		configComponentSettingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q configComponentSettingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for config_component_settings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for config_component_settings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ConfigComponentSettingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configComponentSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"config_component_settings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, configComponentSettingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in configComponentSetting slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all configComponentSetting")
	}
	return rowsAff, nil
}

// Delete deletes a single ConfigComponentSetting record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ConfigComponentSetting) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ConfigComponentSetting provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), configComponentSettingPrimaryKeyMapping)
	sql := "DELETE FROM \"config_component_settings\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from config_component_settings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for config_component_settings")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q configComponentSettingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no configComponentSettingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from config_component_settings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for config_component_settings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ConfigComponentSettingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(configComponentSettingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configComponentSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"config_component_settings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, configComponentSettingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from configComponentSetting slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for config_component_settings")
	}

	if len(configComponentSettingAfterDeleteHooks) != 0 {
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
func (o *ConfigComponentSetting) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindConfigComponentSetting(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ConfigComponentSettingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ConfigComponentSettingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), configComponentSettingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"config_component_settings\".* FROM \"config_component_settings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, configComponentSettingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ConfigComponentSettingSlice")
	}

	*o = slice

	return nil
}

// ConfigComponentSettingExists checks if the ConfigComponentSetting row exists.
func ConfigComponentSettingExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"config_component_settings\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if config_component_settings exists")
	}

	return exists, nil
}

// Exists checks if the ConfigComponentSetting row exists.
func (o *ConfigComponentSetting) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ConfigComponentSettingExists(ctx, exec, o.ID)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ConfigComponentSetting) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no config_component_settings provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(configComponentSettingColumnsWithDefault, o)

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

	configComponentSettingUpsertCacheMut.RLock()
	cache, cached := configComponentSettingUpsertCache[key]
	configComponentSettingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			configComponentSettingAllColumns,
			configComponentSettingColumnsWithDefault,
			configComponentSettingColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			configComponentSettingAllColumns,
			configComponentSettingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert config_component_settings, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(configComponentSettingPrimaryKeyColumns))
			copy(conflict, configComponentSettingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"config_component_settings\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(configComponentSettingType, configComponentSettingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(configComponentSettingType, configComponentSettingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert config_component_settings")
	}

	if !cached {
		configComponentSettingUpsertCacheMut.Lock()
		configComponentSettingUpsertCache[key] = cache
		configComponentSettingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}