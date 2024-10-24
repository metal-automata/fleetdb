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

// BomInfo is an object representing the database table.
type BomInfo struct {
	SerialNum     string      `boil:"serial_num" json:"serial_num" toml:"serial_num" yaml:"serial_num"`
	AocMacAddress null.String `boil:"aoc_mac_address" json:"aoc_mac_address,omitempty" toml:"aoc_mac_address" yaml:"aoc_mac_address,omitempty"`
	BMCMacAddress null.String `boil:"bmc_mac_address" json:"bmc_mac_address,omitempty" toml:"bmc_mac_address" yaml:"bmc_mac_address,omitempty"`
	NumDefiPmi    null.String `boil:"num_defi_pmi" json:"num_defi_pmi,omitempty" toml:"num_defi_pmi" yaml:"num_defi_pmi,omitempty"`
	NumDefPWD     null.String `boil:"num_def_pwd" json:"num_def_pwd,omitempty" toml:"num_def_pwd" yaml:"num_def_pwd,omitempty"`
	Metro         null.String `boil:"metro" json:"metro,omitempty" toml:"metro" yaml:"metro,omitempty"`

	R *bomInfoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bomInfoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BomInfoColumns = struct {
	SerialNum     string
	AocMacAddress string
	BMCMacAddress string
	NumDefiPmi    string
	NumDefPWD     string
	Metro         string
}{
	SerialNum:     "serial_num",
	AocMacAddress: "aoc_mac_address",
	BMCMacAddress: "bmc_mac_address",
	NumDefiPmi:    "num_defi_pmi",
	NumDefPWD:     "num_def_pwd",
	Metro:         "metro",
}

var BomInfoTableColumns = struct {
	SerialNum     string
	AocMacAddress string
	BMCMacAddress string
	NumDefiPmi    string
	NumDefPWD     string
	Metro         string
}{
	SerialNum:     "bom_info.serial_num",
	AocMacAddress: "bom_info.aoc_mac_address",
	BMCMacAddress: "bom_info.bmc_mac_address",
	NumDefiPmi:    "bom_info.num_defi_pmi",
	NumDefPWD:     "bom_info.num_def_pwd",
	Metro:         "bom_info.metro",
}

// Generated where

var BomInfoWhere = struct {
	SerialNum     whereHelperstring
	AocMacAddress whereHelpernull_String
	BMCMacAddress whereHelpernull_String
	NumDefiPmi    whereHelpernull_String
	NumDefPWD     whereHelpernull_String
	Metro         whereHelpernull_String
}{
	SerialNum:     whereHelperstring{field: "\"bom_info\".\"serial_num\""},
	AocMacAddress: whereHelpernull_String{field: "\"bom_info\".\"aoc_mac_address\""},
	BMCMacAddress: whereHelpernull_String{field: "\"bom_info\".\"bmc_mac_address\""},
	NumDefiPmi:    whereHelpernull_String{field: "\"bom_info\".\"num_defi_pmi\""},
	NumDefPWD:     whereHelpernull_String{field: "\"bom_info\".\"num_def_pwd\""},
	Metro:         whereHelpernull_String{field: "\"bom_info\".\"metro\""},
}

// BomInfoRels is where relationship names are stored.
var BomInfoRels = struct {
	SerialNumAocMacAddresses string
	SerialNumBMCMacAddresses string
}{
	SerialNumAocMacAddresses: "SerialNumAocMacAddresses",
	SerialNumBMCMacAddresses: "SerialNumBMCMacAddresses",
}

// bomInfoR is where relationships are stored.
type bomInfoR struct {
	SerialNumAocMacAddresses AocMacAddressSlice `boil:"SerialNumAocMacAddresses" json:"SerialNumAocMacAddresses" toml:"SerialNumAocMacAddresses" yaml:"SerialNumAocMacAddresses"`
	SerialNumBMCMacAddresses BMCMacAddressSlice `boil:"SerialNumBMCMacAddresses" json:"SerialNumBMCMacAddresses" toml:"SerialNumBMCMacAddresses" yaml:"SerialNumBMCMacAddresses"`
}

// NewStruct creates a new relationship struct
func (*bomInfoR) NewStruct() *bomInfoR {
	return &bomInfoR{}
}

func (r *bomInfoR) GetSerialNumAocMacAddresses() AocMacAddressSlice {
	if r == nil {
		return nil
	}
	return r.SerialNumAocMacAddresses
}

func (r *bomInfoR) GetSerialNumBMCMacAddresses() BMCMacAddressSlice {
	if r == nil {
		return nil
	}
	return r.SerialNumBMCMacAddresses
}

// bomInfoL is where Load methods for each relationship are stored.
type bomInfoL struct{}

var (
	bomInfoAllColumns            = []string{"serial_num", "aoc_mac_address", "bmc_mac_address", "num_defi_pmi", "num_def_pwd", "metro"}
	bomInfoColumnsWithoutDefault = []string{"serial_num"}
	bomInfoColumnsWithDefault    = []string{"aoc_mac_address", "bmc_mac_address", "num_defi_pmi", "num_def_pwd", "metro"}
	bomInfoPrimaryKeyColumns     = []string{"serial_num"}
	bomInfoGeneratedColumns      = []string{}
)

type (
	// BomInfoSlice is an alias for a slice of pointers to BomInfo.
	// This should almost always be used instead of []BomInfo.
	BomInfoSlice []*BomInfo
	// BomInfoHook is the signature for custom BomInfo hook methods
	BomInfoHook func(context.Context, boil.ContextExecutor, *BomInfo) error

	bomInfoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bomInfoType                 = reflect.TypeOf(&BomInfo{})
	bomInfoMapping              = queries.MakeStructMapping(bomInfoType)
	bomInfoPrimaryKeyMapping, _ = queries.BindMapping(bomInfoType, bomInfoMapping, bomInfoPrimaryKeyColumns)
	bomInfoInsertCacheMut       sync.RWMutex
	bomInfoInsertCache          = make(map[string]insertCache)
	bomInfoUpdateCacheMut       sync.RWMutex
	bomInfoUpdateCache          = make(map[string]updateCache)
	bomInfoUpsertCacheMut       sync.RWMutex
	bomInfoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bomInfoAfterSelectHooks []BomInfoHook

var bomInfoBeforeInsertHooks []BomInfoHook
var bomInfoAfterInsertHooks []BomInfoHook

var bomInfoBeforeUpdateHooks []BomInfoHook
var bomInfoAfterUpdateHooks []BomInfoHook

var bomInfoBeforeDeleteHooks []BomInfoHook
var bomInfoAfterDeleteHooks []BomInfoHook

var bomInfoBeforeUpsertHooks []BomInfoHook
var bomInfoAfterUpsertHooks []BomInfoHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BomInfo) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BomInfo) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BomInfo) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BomInfo) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BomInfo) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BomInfo) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BomInfo) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BomInfo) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BomInfo) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bomInfoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBomInfoHook registers your hook function for all future operations.
func AddBomInfoHook(hookPoint boil.HookPoint, bomInfoHook BomInfoHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bomInfoAfterSelectHooks = append(bomInfoAfterSelectHooks, bomInfoHook)
	case boil.BeforeInsertHook:
		bomInfoBeforeInsertHooks = append(bomInfoBeforeInsertHooks, bomInfoHook)
	case boil.AfterInsertHook:
		bomInfoAfterInsertHooks = append(bomInfoAfterInsertHooks, bomInfoHook)
	case boil.BeforeUpdateHook:
		bomInfoBeforeUpdateHooks = append(bomInfoBeforeUpdateHooks, bomInfoHook)
	case boil.AfterUpdateHook:
		bomInfoAfterUpdateHooks = append(bomInfoAfterUpdateHooks, bomInfoHook)
	case boil.BeforeDeleteHook:
		bomInfoBeforeDeleteHooks = append(bomInfoBeforeDeleteHooks, bomInfoHook)
	case boil.AfterDeleteHook:
		bomInfoAfterDeleteHooks = append(bomInfoAfterDeleteHooks, bomInfoHook)
	case boil.BeforeUpsertHook:
		bomInfoBeforeUpsertHooks = append(bomInfoBeforeUpsertHooks, bomInfoHook)
	case boil.AfterUpsertHook:
		bomInfoAfterUpsertHooks = append(bomInfoAfterUpsertHooks, bomInfoHook)
	}
}

// One returns a single bomInfo record from the query.
func (q bomInfoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BomInfo, error) {
	o := &BomInfo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bom_info")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all BomInfo records from the query.
func (q bomInfoQuery) All(ctx context.Context, exec boil.ContextExecutor) (BomInfoSlice, error) {
	var o []*BomInfo

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BomInfo slice")
	}

	if len(bomInfoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all BomInfo records in the query.
func (q bomInfoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bom_info rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bomInfoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bom_info exists")
	}

	return count > 0, nil
}

// SerialNumAocMacAddresses retrieves all the aoc_mac_address's AocMacAddresses with an executor via serial_num column.
func (o *BomInfo) SerialNumAocMacAddresses(mods ...qm.QueryMod) aocMacAddressQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"aoc_mac_address\".\"serial_num\"=?", o.SerialNum),
	)

	return AocMacAddresses(queryMods...)
}

// SerialNumBMCMacAddresses retrieves all the bmc_mac_address's BMCMacAddresses with an executor via serial_num column.
func (o *BomInfo) SerialNumBMCMacAddresses(mods ...qm.QueryMod) bmcMacAddressQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"bmc_mac_address\".\"serial_num\"=?", o.SerialNum),
	)

	return BMCMacAddresses(queryMods...)
}

// LoadSerialNumAocMacAddresses allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bomInfoL) LoadSerialNumAocMacAddresses(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBomInfo interface{}, mods queries.Applicator) error {
	var slice []*BomInfo
	var object *BomInfo

	if singular {
		var ok bool
		object, ok = maybeBomInfo.(*BomInfo)
		if !ok {
			object = new(BomInfo)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBomInfo)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBomInfo))
			}
		}
	} else {
		s, ok := maybeBomInfo.(*[]*BomInfo)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBomInfo)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBomInfo))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bomInfoR{}
		}
		args = append(args, object.SerialNum)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bomInfoR{}
			}

			for _, a := range args {
				if a == obj.SerialNum {
					continue Outer
				}
			}

			args = append(args, obj.SerialNum)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`aoc_mac_address`),
		qm.WhereIn(`aoc_mac_address.serial_num in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load aoc_mac_address")
	}

	var resultSlice []*AocMacAddress
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice aoc_mac_address")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on aoc_mac_address")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for aoc_mac_address")
	}

	if len(aocMacAddressAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SerialNumAocMacAddresses = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &aocMacAddressR{}
			}
			foreign.R.SerialNumBomInfo = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.SerialNum == foreign.SerialNum {
				local.R.SerialNumAocMacAddresses = append(local.R.SerialNumAocMacAddresses, foreign)
				if foreign.R == nil {
					foreign.R = &aocMacAddressR{}
				}
				foreign.R.SerialNumBomInfo = local
				break
			}
		}
	}

	return nil
}

// LoadSerialNumBMCMacAddresses allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bomInfoL) LoadSerialNumBMCMacAddresses(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBomInfo interface{}, mods queries.Applicator) error {
	var slice []*BomInfo
	var object *BomInfo

	if singular {
		var ok bool
		object, ok = maybeBomInfo.(*BomInfo)
		if !ok {
			object = new(BomInfo)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBomInfo)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBomInfo))
			}
		}
	} else {
		s, ok := maybeBomInfo.(*[]*BomInfo)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBomInfo)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBomInfo))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bomInfoR{}
		}
		args = append(args, object.SerialNum)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bomInfoR{}
			}

			for _, a := range args {
				if a == obj.SerialNum {
					continue Outer
				}
			}

			args = append(args, obj.SerialNum)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`bmc_mac_address`),
		qm.WhereIn(`bmc_mac_address.serial_num in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load bmc_mac_address")
	}

	var resultSlice []*BMCMacAddress
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice bmc_mac_address")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on bmc_mac_address")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bmc_mac_address")
	}

	if len(bmcMacAddressAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SerialNumBMCMacAddresses = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &bmcMacAddressR{}
			}
			foreign.R.SerialNumBomInfo = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.SerialNum == foreign.SerialNum {
				local.R.SerialNumBMCMacAddresses = append(local.R.SerialNumBMCMacAddresses, foreign)
				if foreign.R == nil {
					foreign.R = &bmcMacAddressR{}
				}
				foreign.R.SerialNumBomInfo = local
				break
			}
		}
	}

	return nil
}

// AddSerialNumAocMacAddresses adds the given related objects to the existing relationships
// of the bom_info, optionally inserting them as new records.
// Appends related to o.R.SerialNumAocMacAddresses.
// Sets related.R.SerialNumBomInfo appropriately.
func (o *BomInfo) AddSerialNumAocMacAddresses(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AocMacAddress) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SerialNum = o.SerialNum
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"aoc_mac_address\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"serial_num"}),
				strmangle.WhereClause("\"", "\"", 2, aocMacAddressPrimaryKeyColumns),
			)
			values := []interface{}{o.SerialNum, rel.AocMacAddress}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.SerialNum = o.SerialNum
		}
	}

	if o.R == nil {
		o.R = &bomInfoR{
			SerialNumAocMacAddresses: related,
		}
	} else {
		o.R.SerialNumAocMacAddresses = append(o.R.SerialNumAocMacAddresses, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &aocMacAddressR{
				SerialNumBomInfo: o,
			}
		} else {
			rel.R.SerialNumBomInfo = o
		}
	}
	return nil
}

// AddSerialNumBMCMacAddresses adds the given related objects to the existing relationships
// of the bom_info, optionally inserting them as new records.
// Appends related to o.R.SerialNumBMCMacAddresses.
// Sets related.R.SerialNumBomInfo appropriately.
func (o *BomInfo) AddSerialNumBMCMacAddresses(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BMCMacAddress) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SerialNum = o.SerialNum
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"bmc_mac_address\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"serial_num"}),
				strmangle.WhereClause("\"", "\"", 2, bmcMacAddressPrimaryKeyColumns),
			)
			values := []interface{}{o.SerialNum, rel.BMCMacAddress}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.SerialNum = o.SerialNum
		}
	}

	if o.R == nil {
		o.R = &bomInfoR{
			SerialNumBMCMacAddresses: related,
		}
	} else {
		o.R.SerialNumBMCMacAddresses = append(o.R.SerialNumBMCMacAddresses, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &bmcMacAddressR{
				SerialNumBomInfo: o,
			}
		} else {
			rel.R.SerialNumBomInfo = o
		}
	}
	return nil
}

// BomInfos retrieves all the records using an executor.
func BomInfos(mods ...qm.QueryMod) bomInfoQuery {
	mods = append(mods, qm.From("\"bom_info\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"bom_info\".*"})
	}

	return bomInfoQuery{q}
}

// FindBomInfo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBomInfo(ctx context.Context, exec boil.ContextExecutor, serialNum string, selectCols ...string) (*BomInfo, error) {
	bomInfoObj := &BomInfo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bom_info\" where \"serial_num\"=$1", sel,
	)

	q := queries.Raw(query, serialNum)

	err := q.Bind(ctx, exec, bomInfoObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bom_info")
	}

	if err = bomInfoObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bomInfoObj, err
	}

	return bomInfoObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BomInfo) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bom_info provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bomInfoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bomInfoInsertCacheMut.RLock()
	cache, cached := bomInfoInsertCache[key]
	bomInfoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bomInfoAllColumns,
			bomInfoColumnsWithDefault,
			bomInfoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bomInfoType, bomInfoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bomInfoType, bomInfoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bom_info\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bom_info\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into bom_info")
	}

	if !cached {
		bomInfoInsertCacheMut.Lock()
		bomInfoInsertCache[key] = cache
		bomInfoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the BomInfo.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BomInfo) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bomInfoUpdateCacheMut.RLock()
	cache, cached := bomInfoUpdateCache[key]
	bomInfoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bomInfoAllColumns,
			bomInfoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bom_info, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bom_info\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bomInfoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bomInfoType, bomInfoMapping, append(wl, bomInfoPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update bom_info row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bom_info")
	}

	if !cached {
		bomInfoUpdateCacheMut.Lock()
		bomInfoUpdateCache[key] = cache
		bomInfoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bomInfoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bom_info")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bom_info")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BomInfoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bomInfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bom_info\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bomInfoPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bomInfo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bomInfo")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BomInfo) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bom_info provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bomInfoColumnsWithDefault, o)

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

	bomInfoUpsertCacheMut.RLock()
	cache, cached := bomInfoUpsertCache[key]
	bomInfoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bomInfoAllColumns,
			bomInfoColumnsWithDefault,
			bomInfoColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bomInfoAllColumns,
			bomInfoPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert bom_info, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bomInfoPrimaryKeyColumns))
			copy(conflict, bomInfoPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"bom_info\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bomInfoType, bomInfoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bomInfoType, bomInfoMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert bom_info")
	}

	if !cached {
		bomInfoUpsertCacheMut.Lock()
		bomInfoUpsertCache[key] = cache
		bomInfoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single BomInfo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BomInfo) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BomInfo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bomInfoPrimaryKeyMapping)
	sql := "DELETE FROM \"bom_info\" WHERE \"serial_num\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bom_info")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bom_info")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bomInfoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bomInfoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bom_info")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bom_info")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BomInfoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bomInfoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bomInfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bom_info\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bomInfoPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bomInfo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bom_info")
	}

	if len(bomInfoAfterDeleteHooks) != 0 {
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
func (o *BomInfo) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBomInfo(ctx, exec, o.SerialNum)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BomInfoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BomInfoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bomInfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bom_info\".* FROM \"bom_info\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bomInfoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BomInfoSlice")
	}

	*o = slice

	return nil
}

// BomInfoExists checks if the BomInfo row exists.
func BomInfoExists(ctx context.Context, exec boil.ContextExecutor, serialNum string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bom_info\" where \"serial_num\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, serialNum)
	}
	row := exec.QueryRowContext(ctx, sql, serialNum)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bom_info exists")
	}

	return exists, nil
}

// Exists checks if the BomInfo row exists.
func (o *BomInfo) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BomInfoExists(ctx, exec, o.SerialNum)
}
