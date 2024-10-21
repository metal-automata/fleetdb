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

func testBomInfos(t *testing.T) {
	t.Parallel()

	query := BomInfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBomInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
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

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBomInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BomInfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBomInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BomInfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBomInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BomInfoExists(ctx, tx, o.SerialNum)
	if err != nil {
		t.Errorf("Unable to check if BomInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BomInfoExists to return true, but got false.")
	}
}

func testBomInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bomInfoFound, err := FindBomInfo(ctx, tx, o.SerialNum)
	if err != nil {
		t.Error(err)
	}

	if bomInfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBomInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BomInfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBomInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BomInfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBomInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bomInfoOne := &BomInfo{}
	bomInfoTwo := &BomInfo{}
	if err = randomize.Struct(seed, bomInfoOne, bomInfoDBTypes, false, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, bomInfoTwo, bomInfoDBTypes, false, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bomInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bomInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BomInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBomInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bomInfoOne := &BomInfo{}
	bomInfoTwo := &BomInfo{}
	if err = randomize.Struct(seed, bomInfoOne, bomInfoDBTypes, false, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, bomInfoTwo, bomInfoDBTypes, false, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bomInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bomInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bomInfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func bomInfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func bomInfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func bomInfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func bomInfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func bomInfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func bomInfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func bomInfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func bomInfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BomInfo) error {
	*o = BomInfo{}
	return nil
}

func testBomInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BomInfo{}
	o := &BomInfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bomInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BomInfo object: %s", err)
	}

	AddBomInfoHook(boil.BeforeInsertHook, bomInfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bomInfoBeforeInsertHooks = []BomInfoHook{}

	AddBomInfoHook(boil.AfterInsertHook, bomInfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bomInfoAfterInsertHooks = []BomInfoHook{}

	AddBomInfoHook(boil.AfterSelectHook, bomInfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bomInfoAfterSelectHooks = []BomInfoHook{}

	AddBomInfoHook(boil.BeforeUpdateHook, bomInfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bomInfoBeforeUpdateHooks = []BomInfoHook{}

	AddBomInfoHook(boil.AfterUpdateHook, bomInfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bomInfoAfterUpdateHooks = []BomInfoHook{}

	AddBomInfoHook(boil.BeforeDeleteHook, bomInfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bomInfoBeforeDeleteHooks = []BomInfoHook{}

	AddBomInfoHook(boil.AfterDeleteHook, bomInfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bomInfoAfterDeleteHooks = []BomInfoHook{}

	AddBomInfoHook(boil.BeforeUpsertHook, bomInfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bomInfoBeforeUpsertHooks = []BomInfoHook{}

	AddBomInfoHook(boil.AfterUpsertHook, bomInfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bomInfoAfterUpsertHooks = []BomInfoHook{}
}

func testBomInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBomInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bomInfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBomInfoToManySerialNumAocMacAddresses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BomInfo
	var b, c AocMacAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, aocMacAddressDBTypes, false, aocMacAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, aocMacAddressDBTypes, false, aocMacAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.SerialNum = a.SerialNum
	c.SerialNum = a.SerialNum

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.SerialNumAocMacAddresses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SerialNum == b.SerialNum {
			bFound = true
		}
		if v.SerialNum == c.SerialNum {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BomInfoSlice{&a}
	if err = a.L.LoadSerialNumAocMacAddresses(ctx, tx, false, (*[]*BomInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SerialNumAocMacAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SerialNumAocMacAddresses = nil
	if err = a.L.LoadSerialNumAocMacAddresses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SerialNumAocMacAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBomInfoToManySerialNumBMCMacAddresses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BomInfo
	var b, c BMCMacAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, bmcMacAddressDBTypes, false, bmcMacAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bmcMacAddressDBTypes, false, bmcMacAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.SerialNum = a.SerialNum
	c.SerialNum = a.SerialNum

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.SerialNumBMCMacAddresses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SerialNum == b.SerialNum {
			bFound = true
		}
		if v.SerialNum == c.SerialNum {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BomInfoSlice{&a}
	if err = a.L.LoadSerialNumBMCMacAddresses(ctx, tx, false, (*[]*BomInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SerialNumBMCMacAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SerialNumBMCMacAddresses = nil
	if err = a.L.LoadSerialNumBMCMacAddresses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SerialNumBMCMacAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBomInfoToManyAddOpSerialNumAocMacAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BomInfo
	var b, c, d, e AocMacAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bomInfoDBTypes, false, strmangle.SetComplement(bomInfoPrimaryKeyColumns, bomInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AocMacAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, aocMacAddressDBTypes, false, strmangle.SetComplement(aocMacAddressPrimaryKeyColumns, aocMacAddressColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AocMacAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSerialNumAocMacAddresses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.SerialNum != first.SerialNum {
			t.Error("foreign key was wrong value", a.SerialNum, first.SerialNum)
		}
		if a.SerialNum != second.SerialNum {
			t.Error("foreign key was wrong value", a.SerialNum, second.SerialNum)
		}

		if first.R.SerialNumBomInfo != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.SerialNumBomInfo != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SerialNumAocMacAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SerialNumAocMacAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SerialNumAocMacAddresses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBomInfoToManyAddOpSerialNumBMCMacAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BomInfo
	var b, c, d, e BMCMacAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bomInfoDBTypes, false, strmangle.SetComplement(bomInfoPrimaryKeyColumns, bomInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BMCMacAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, bmcMacAddressDBTypes, false, strmangle.SetComplement(bmcMacAddressPrimaryKeyColumns, bmcMacAddressColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BMCMacAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSerialNumBMCMacAddresses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.SerialNum != first.SerialNum {
			t.Error("foreign key was wrong value", a.SerialNum, first.SerialNum)
		}
		if a.SerialNum != second.SerialNum {
			t.Error("foreign key was wrong value", a.SerialNum, second.SerialNum)
		}

		if first.R.SerialNumBomInfo != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.SerialNumBomInfo != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SerialNumBMCMacAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SerialNumBMCMacAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SerialNumBMCMacAddresses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBomInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
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

func testBomInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BomInfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBomInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BomInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bomInfoDBTypes = map[string]string{`SerialNum`: `text`, `AocMacAddress`: `text`, `BMCMacAddress`: `text`, `NumDefiPmi`: `text`, `NumDefPWD`: `text`, `Metro`: `text`}
	_              = bytes.MinRead
)

func testBomInfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bomInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bomInfoAllColumns) == len(bomInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBomInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bomInfoAllColumns) == len(bomInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BomInfo{}
	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bomInfoDBTypes, true, bomInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bomInfoAllColumns, bomInfoPrimaryKeyColumns) {
		fields = bomInfoAllColumns
	} else {
		fields = strmangle.SetComplement(
			bomInfoAllColumns,
			bomInfoPrimaryKeyColumns,
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

	slice := BomInfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBomInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(bomInfoAllColumns) == len(bomInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BomInfo{}
	if err = randomize.Struct(seed, &o, bomInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BomInfo: %s", err)
	}

	count, err := BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bomInfoDBTypes, false, bomInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BomInfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BomInfo: %s", err)
	}

	count, err = BomInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
