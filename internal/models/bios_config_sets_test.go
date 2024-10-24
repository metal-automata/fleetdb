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

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBiosConfigSets(t *testing.T) {
	t.Parallel()

	query := BiosConfigSets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBiosConfigSetsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
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

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBiosConfigSetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BiosConfigSets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBiosConfigSetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BiosConfigSetSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBiosConfigSetsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BiosConfigSetExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BiosConfigSet exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BiosConfigSetExists to return true, but got false.")
	}
}

func testBiosConfigSetsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	biosConfigSetFound, err := FindBiosConfigSet(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if biosConfigSetFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBiosConfigSetsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BiosConfigSets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBiosConfigSetsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BiosConfigSets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBiosConfigSetsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	biosConfigSetOne := &BiosConfigSet{}
	biosConfigSetTwo := &BiosConfigSet{}
	if err = randomize.Struct(seed, biosConfigSetOne, biosConfigSetDBTypes, false, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}
	if err = randomize.Struct(seed, biosConfigSetTwo, biosConfigSetDBTypes, false, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = biosConfigSetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = biosConfigSetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BiosConfigSets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBiosConfigSetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	biosConfigSetOne := &BiosConfigSet{}
	biosConfigSetTwo := &BiosConfigSet{}
	if err = randomize.Struct(seed, biosConfigSetOne, biosConfigSetDBTypes, false, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}
	if err = randomize.Struct(seed, biosConfigSetTwo, biosConfigSetDBTypes, false, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = biosConfigSetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = biosConfigSetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func biosConfigSetBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func biosConfigSetAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func biosConfigSetAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func biosConfigSetBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func biosConfigSetAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func biosConfigSetBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func biosConfigSetAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func biosConfigSetBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func biosConfigSetAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BiosConfigSet) error {
	*o = BiosConfigSet{}
	return nil
}

func testBiosConfigSetsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BiosConfigSet{}
	o := &BiosConfigSet{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet object: %s", err)
	}

	AddBiosConfigSetHook(boil.BeforeInsertHook, biosConfigSetBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	biosConfigSetBeforeInsertHooks = []BiosConfigSetHook{}

	AddBiosConfigSetHook(boil.AfterInsertHook, biosConfigSetAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	biosConfigSetAfterInsertHooks = []BiosConfigSetHook{}

	AddBiosConfigSetHook(boil.AfterSelectHook, biosConfigSetAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	biosConfigSetAfterSelectHooks = []BiosConfigSetHook{}

	AddBiosConfigSetHook(boil.BeforeUpdateHook, biosConfigSetBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	biosConfigSetBeforeUpdateHooks = []BiosConfigSetHook{}

	AddBiosConfigSetHook(boil.AfterUpdateHook, biosConfigSetAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	biosConfigSetAfterUpdateHooks = []BiosConfigSetHook{}

	AddBiosConfigSetHook(boil.BeforeDeleteHook, biosConfigSetBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	biosConfigSetBeforeDeleteHooks = []BiosConfigSetHook{}

	AddBiosConfigSetHook(boil.AfterDeleteHook, biosConfigSetAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	biosConfigSetAfterDeleteHooks = []BiosConfigSetHook{}

	AddBiosConfigSetHook(boil.BeforeUpsertHook, biosConfigSetBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	biosConfigSetBeforeUpsertHooks = []BiosConfigSetHook{}

	AddBiosConfigSetHook(boil.AfterUpsertHook, biosConfigSetAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	biosConfigSetAfterUpsertHooks = []BiosConfigSetHook{}
}

func testBiosConfigSetsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBiosConfigSetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(biosConfigSetColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBiosConfigSetToManyFKBiosConfigSetBiosConfigComponents(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BiosConfigSet
	var b, c BiosConfigComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, biosConfigComponentDBTypes, false, biosConfigComponentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, biosConfigComponentDBTypes, false, biosConfigComponentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.FKBiosConfigSetID = a.ID
	c.FKBiosConfigSetID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.FKBiosConfigSetBiosConfigComponents().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.FKBiosConfigSetID == b.FKBiosConfigSetID {
			bFound = true
		}
		if v.FKBiosConfigSetID == c.FKBiosConfigSetID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BiosConfigSetSlice{&a}
	if err = a.L.LoadFKBiosConfigSetBiosConfigComponents(ctx, tx, false, (*[]*BiosConfigSet)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FKBiosConfigSetBiosConfigComponents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FKBiosConfigSetBiosConfigComponents = nil
	if err = a.L.LoadFKBiosConfigSetBiosConfigComponents(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FKBiosConfigSetBiosConfigComponents); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBiosConfigSetToManyAddOpFKBiosConfigSetBiosConfigComponents(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BiosConfigSet
	var b, c, d, e BiosConfigComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, biosConfigSetDBTypes, false, strmangle.SetComplement(biosConfigSetPrimaryKeyColumns, biosConfigSetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BiosConfigComponent{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, biosConfigComponentDBTypes, false, strmangle.SetComplement(biosConfigComponentPrimaryKeyColumns, biosConfigComponentColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BiosConfigComponent{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFKBiosConfigSetBiosConfigComponents(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.FKBiosConfigSetID {
			t.Error("foreign key was wrong value", a.ID, first.FKBiosConfigSetID)
		}
		if a.ID != second.FKBiosConfigSetID {
			t.Error("foreign key was wrong value", a.ID, second.FKBiosConfigSetID)
		}

		if first.R.FKBiosConfigSet != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.FKBiosConfigSet != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FKBiosConfigSetBiosConfigComponents[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FKBiosConfigSetBiosConfigComponents[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FKBiosConfigSetBiosConfigComponents().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBiosConfigSetsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
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

func testBiosConfigSetsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BiosConfigSetSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBiosConfigSetsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BiosConfigSets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	biosConfigSetDBTypes = map[string]string{`ID`: `uuid`, `Name`: `text`, `Version`: `text`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                    = bytes.MinRead
)

func testBiosConfigSetsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(biosConfigSetPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(biosConfigSetAllColumns) == len(biosConfigSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBiosConfigSetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(biosConfigSetAllColumns) == len(biosConfigSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BiosConfigSet{}
	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, biosConfigSetDBTypes, true, biosConfigSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(biosConfigSetAllColumns, biosConfigSetPrimaryKeyColumns) {
		fields = biosConfigSetAllColumns
	} else {
		fields = strmangle.SetComplement(
			biosConfigSetAllColumns,
			biosConfigSetPrimaryKeyColumns,
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

	slice := BiosConfigSetSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBiosConfigSetsUpsert(t *testing.T) {
	t.Parallel()

	if len(biosConfigSetAllColumns) == len(biosConfigSetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BiosConfigSet{}
	if err = randomize.Struct(seed, &o, biosConfigSetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BiosConfigSet: %s", err)
	}

	count, err := BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, biosConfigSetDBTypes, false, biosConfigSetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BiosConfigSet struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BiosConfigSet: %s", err)
	}

	count, err = BiosConfigSets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
