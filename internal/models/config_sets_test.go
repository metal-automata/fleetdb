// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

func testConfigSetsUpsert(t *testing.T) {
	t.Parallel()

	if len(configSetAllColumns) == len(configSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ConfigSet{}
	if err = randomize.Struct(seed, &o, configSetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ConfigSet: %s", err)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, configSetDBTypes, false, configSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ConfigSet: %s", err)
	}

	count, err = ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testConfigSets(t *testing.T) {
	t.Parallel()

	query := ConfigSets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testConfigSetsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testConfigSetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ConfigSets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testConfigSetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ConfigSetSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testConfigSetsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ConfigSetExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ConfigSet exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ConfigSetExists to return true, but got false.")
	}
}

func testConfigSetsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	configSetFound, err := FindConfigSet(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if configSetFound == nil {
		t.Error("want a record, got nil")
	}
}

func testConfigSetsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ConfigSets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testConfigSetsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ConfigSets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testConfigSetsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	configSetOne := &ConfigSet{}
	configSetTwo := &ConfigSet{}
	if err = randomize.Struct(seed, configSetOne, configSetDBTypes, false, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}
	if err = randomize.Struct(seed, configSetTwo, configSetDBTypes, false, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = configSetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = configSetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ConfigSets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testConfigSetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	configSetOne := &ConfigSet{}
	configSetTwo := &ConfigSet{}
	if err = randomize.Struct(seed, configSetOne, configSetDBTypes, false, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}
	if err = randomize.Struct(seed, configSetTwo, configSetDBTypes, false, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = configSetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = configSetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func configSetBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func configSetAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func configSetAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func configSetBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func configSetAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func configSetBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func configSetAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func configSetBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func configSetAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ConfigSet) error {
	*o = ConfigSet{}
	return nil
}

func testConfigSetsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ConfigSet{}
	o := &ConfigSet{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, configSetDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ConfigSet object: %s", err)
	}

	AddConfigSetHook(boil.BeforeInsertHook, configSetBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	configSetBeforeInsertHooks = []ConfigSetHook{}

	AddConfigSetHook(boil.AfterInsertHook, configSetAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	configSetAfterInsertHooks = []ConfigSetHook{}

	AddConfigSetHook(boil.AfterSelectHook, configSetAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	configSetAfterSelectHooks = []ConfigSetHook{}

	AddConfigSetHook(boil.BeforeUpdateHook, configSetBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	configSetBeforeUpdateHooks = []ConfigSetHook{}

	AddConfigSetHook(boil.AfterUpdateHook, configSetAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	configSetAfterUpdateHooks = []ConfigSetHook{}

	AddConfigSetHook(boil.BeforeDeleteHook, configSetBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	configSetBeforeDeleteHooks = []ConfigSetHook{}

	AddConfigSetHook(boil.AfterDeleteHook, configSetAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	configSetAfterDeleteHooks = []ConfigSetHook{}

	AddConfigSetHook(boil.BeforeUpsertHook, configSetBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	configSetBeforeUpsertHooks = []ConfigSetHook{}

	AddConfigSetHook(boil.AfterUpsertHook, configSetAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	configSetAfterUpsertHooks = []ConfigSetHook{}
}

func testConfigSetsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testConfigSetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(configSetColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testConfigSetToManyFKConfigSetConfigComponents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ConfigSet
	var b, c ConfigComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, configComponentDBTypes, false, configComponentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, configComponentDBTypes, false, configComponentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.FKConfigSetID = a.ID
	c.FKConfigSetID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.FKConfigSetConfigComponents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.FKConfigSetID == b.FKConfigSetID {
			bFound = true
		}
		if v.FKConfigSetID == c.FKConfigSetID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ConfigSetSlice{&a}
	if err = a.L.LoadFKConfigSetConfigComponents(ctx, tx, false, (*[]*ConfigSet)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FKConfigSetConfigComponents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FKConfigSetConfigComponents = nil
	if err = a.L.LoadFKConfigSetConfigComponents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FKConfigSetConfigComponents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testConfigSetToManyAddOpFKConfigSetConfigComponents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ConfigSet
	var b, c, d, e ConfigComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, configSetDBTypes, false, strmangle.SetComplement(configSetPrimaryKeyColumns, configSetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ConfigComponent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, configComponentDBTypes, false, strmangle.SetComplement(configComponentPrimaryKeyColumns, configComponentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ConfigComponent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFKConfigSetConfigComponents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.FKConfigSetID {
			t.Error("foreign key was wrong value", a.ID, first.FKConfigSetID)
		}
		if a.ID != second.FKConfigSetID {
			t.Error("foreign key was wrong value", a.ID, second.FKConfigSetID)
		}

		if first.R.FKConfigSet != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.FKConfigSet != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FKConfigSetConfigComponents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FKConfigSetConfigComponents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FKConfigSetConfigComponents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testConfigSetsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testConfigSetsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ConfigSetSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testConfigSetsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ConfigSets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	configSetDBTypes = map[string]string{`ID`: `uuid`, `Name`: `string`, `Version`: `string`, `CreatedAt`: `timestamptz`, `UpdatedAt`: `timestamptz`}
	_                = bytes.MinRead
)

func testConfigSetsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(configSetPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(configSetAllColumns) == len(configSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testConfigSetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(configSetAllColumns) == len(configSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ConfigSet{}
	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, configSetDBTypes, true, configSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ConfigSet struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(configSetAllColumns, configSetPrimaryKeyColumns) {
		fields = configSetAllColumns
	} else {
		fields = strmangle.SetComplement(
			configSetAllColumns,
			configSetPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ConfigSetSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
