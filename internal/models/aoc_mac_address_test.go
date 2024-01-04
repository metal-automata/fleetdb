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

func testAocMacAddressesUpsert(t *testing.T) {
	t.Parallel()

	if len(aocMacAddressAllColumns) == len(aocMacAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AocMacAddress{}
	if err = randomize.Struct(seed, &o, aocMacAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AocMacAddress: %s", err)
	}

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, aocMacAddressDBTypes, false, aocMacAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AocMacAddress: %s", err)
	}

	count, err = AocMacAddresses().Count(ctx, tx)
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

func testAocMacAddresses(t *testing.T) {
	t.Parallel()

	query := AocMacAddresses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAocMacAddressesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
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

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAocMacAddressesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AocMacAddresses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAocMacAddressesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AocMacAddressSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAocMacAddressesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AocMacAddressExists(ctx, tx, o.AocMacAddress)
	if err != nil {
		t.Errorf("Unable to check if AocMacAddress exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AocMacAddressExists to return true, but got false.")
	}
}

func testAocMacAddressesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	aocMacAddressFound, err := FindAocMacAddress(ctx, tx, o.AocMacAddress)
	if err != nil {
		t.Error(err)
	}

	if aocMacAddressFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAocMacAddressesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AocMacAddresses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAocMacAddressesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AocMacAddresses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAocMacAddressesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	aocMacAddressOne := &AocMacAddress{}
	aocMacAddressTwo := &AocMacAddress{}
	if err = randomize.Struct(seed, aocMacAddressOne, aocMacAddressDBTypes, false, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, aocMacAddressTwo, aocMacAddressDBTypes, false, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = aocMacAddressOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = aocMacAddressTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AocMacAddresses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAocMacAddressesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	aocMacAddressOne := &AocMacAddress{}
	aocMacAddressTwo := &AocMacAddress{}
	if err = randomize.Struct(seed, aocMacAddressOne, aocMacAddressDBTypes, false, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, aocMacAddressTwo, aocMacAddressDBTypes, false, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = aocMacAddressOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = aocMacAddressTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func aocMacAddressBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func aocMacAddressAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func aocMacAddressAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func aocMacAddressBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func aocMacAddressAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func aocMacAddressBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func aocMacAddressAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func aocMacAddressBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func aocMacAddressAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AocMacAddress) error {
	*o = AocMacAddress{}
	return nil
}

func testAocMacAddressesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AocMacAddress{}
	o := &AocMacAddress{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AocMacAddress object: %s", err)
	}

	AddAocMacAddressHook(boil.BeforeInsertHook, aocMacAddressBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	aocMacAddressBeforeInsertHooks = []AocMacAddressHook{}

	AddAocMacAddressHook(boil.AfterInsertHook, aocMacAddressAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	aocMacAddressAfterInsertHooks = []AocMacAddressHook{}

	AddAocMacAddressHook(boil.AfterSelectHook, aocMacAddressAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	aocMacAddressAfterSelectHooks = []AocMacAddressHook{}

	AddAocMacAddressHook(boil.BeforeUpdateHook, aocMacAddressBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	aocMacAddressBeforeUpdateHooks = []AocMacAddressHook{}

	AddAocMacAddressHook(boil.AfterUpdateHook, aocMacAddressAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	aocMacAddressAfterUpdateHooks = []AocMacAddressHook{}

	AddAocMacAddressHook(boil.BeforeDeleteHook, aocMacAddressBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	aocMacAddressBeforeDeleteHooks = []AocMacAddressHook{}

	AddAocMacAddressHook(boil.AfterDeleteHook, aocMacAddressAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	aocMacAddressAfterDeleteHooks = []AocMacAddressHook{}

	AddAocMacAddressHook(boil.BeforeUpsertHook, aocMacAddressBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	aocMacAddressBeforeUpsertHooks = []AocMacAddressHook{}

	AddAocMacAddressHook(boil.AfterUpsertHook, aocMacAddressAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	aocMacAddressAfterUpsertHooks = []AocMacAddressHook{}
}

func testAocMacAddressesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAocMacAddressesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(aocMacAddressColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAocMacAddressToOneBomInfoUsingSerialNumBomInfo(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AocMacAddress
	var foreign BomInfo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, aocMacAddressDBTypes, false, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bomInfoDBTypes, false, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SerialNum = foreign.SerialNum
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.SerialNumBomInfo().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.SerialNum != foreign.SerialNum {
		t.Errorf("want: %v, got %v", foreign.SerialNum, check.SerialNum)
	}

	ranAfterSelectHook := false
	AddBomInfoHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := AocMacAddressSlice{&local}
	if err = local.L.LoadSerialNumBomInfo(ctx, tx, false, (*[]*AocMacAddress)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SerialNumBomInfo == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SerialNumBomInfo = nil
	if err = local.L.LoadSerialNumBomInfo(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SerialNumBomInfo == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testAocMacAddressToOneSetOpBomInfoUsingSerialNumBomInfo(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AocMacAddress
	var b, c BomInfo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, aocMacAddressDBTypes, false, strmangle.SetComplement(aocMacAddressPrimaryKeyColumns, aocMacAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bomInfoDBTypes, false, strmangle.SetComplement(bomInfoPrimaryKeyColumns, bomInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bomInfoDBTypes, false, strmangle.SetComplement(bomInfoPrimaryKeyColumns, bomInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BomInfo{&b, &c} {
		err = a.SetSerialNumBomInfo(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SerialNumBomInfo != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SerialNumAocMacAddresses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SerialNum != x.SerialNum {
			t.Error("foreign key was wrong value", a.SerialNum)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SerialNum))
		reflect.Indirect(reflect.ValueOf(&a.SerialNum)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SerialNum != x.SerialNum {
			t.Error("foreign key was wrong value", a.SerialNum, x.SerialNum)
		}
	}
}

func testAocMacAddressesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
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

func testAocMacAddressesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AocMacAddressSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAocMacAddressesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AocMacAddresses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	aocMacAddressDBTypes = map[string]string{`AocMacAddress`: `string`, `SerialNum`: `string`}
	_                    = bytes.MinRead
)

func testAocMacAddressesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(aocMacAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(aocMacAddressAllColumns) == len(aocMacAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAocMacAddressesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(aocMacAddressAllColumns) == len(aocMacAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AocMacAddress{}
	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AocMacAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, aocMacAddressDBTypes, true, aocMacAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AocMacAddress struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(aocMacAddressAllColumns, aocMacAddressPrimaryKeyColumns) {
		fields = aocMacAddressAllColumns
	} else {
		fields = strmangle.SetComplement(
			aocMacAddressAllColumns,
			aocMacAddressPrimaryKeyColumns,
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

	slice := AocMacAddressSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}