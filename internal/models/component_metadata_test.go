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

func testComponentMetadata(t *testing.T) {
	t.Parallel()

	query := ComponentMetadata()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testComponentMetadataDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
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

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComponentMetadataQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ComponentMetadata().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComponentMetadataSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ComponentMetadatumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComponentMetadataExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ComponentMetadatumExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ComponentMetadatum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ComponentMetadatumExists to return true, but got false.")
	}
}

func testComponentMetadataFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	componentMetadatumFound, err := FindComponentMetadatum(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if componentMetadatumFound == nil {
		t.Error("want a record, got nil")
	}
}

func testComponentMetadataBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ComponentMetadata().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testComponentMetadataOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ComponentMetadata().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testComponentMetadataAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	componentMetadatumOne := &ComponentMetadatum{}
	componentMetadatumTwo := &ComponentMetadatum{}
	if err = randomize.Struct(seed, componentMetadatumOne, componentMetadatumDBTypes, false, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}
	if err = randomize.Struct(seed, componentMetadatumTwo, componentMetadatumDBTypes, false, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = componentMetadatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = componentMetadatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ComponentMetadata().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testComponentMetadataCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	componentMetadatumOne := &ComponentMetadatum{}
	componentMetadatumTwo := &ComponentMetadatum{}
	if err = randomize.Struct(seed, componentMetadatumOne, componentMetadatumDBTypes, false, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}
	if err = randomize.Struct(seed, componentMetadatumTwo, componentMetadatumDBTypes, false, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = componentMetadatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = componentMetadatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func componentMetadatumBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func componentMetadatumAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func componentMetadatumAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func componentMetadatumBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func componentMetadatumAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func componentMetadatumBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func componentMetadatumAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func componentMetadatumBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func componentMetadatumAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ComponentMetadatum) error {
	*o = ComponentMetadatum{}
	return nil
}

func testComponentMetadataHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ComponentMetadatum{}
	o := &ComponentMetadatum{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum object: %s", err)
	}

	AddComponentMetadatumHook(boil.BeforeInsertHook, componentMetadatumBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	componentMetadatumBeforeInsertHooks = []ComponentMetadatumHook{}

	AddComponentMetadatumHook(boil.AfterInsertHook, componentMetadatumAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	componentMetadatumAfterInsertHooks = []ComponentMetadatumHook{}

	AddComponentMetadatumHook(boil.AfterSelectHook, componentMetadatumAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	componentMetadatumAfterSelectHooks = []ComponentMetadatumHook{}

	AddComponentMetadatumHook(boil.BeforeUpdateHook, componentMetadatumBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	componentMetadatumBeforeUpdateHooks = []ComponentMetadatumHook{}

	AddComponentMetadatumHook(boil.AfterUpdateHook, componentMetadatumAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	componentMetadatumAfterUpdateHooks = []ComponentMetadatumHook{}

	AddComponentMetadatumHook(boil.BeforeDeleteHook, componentMetadatumBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	componentMetadatumBeforeDeleteHooks = []ComponentMetadatumHook{}

	AddComponentMetadatumHook(boil.AfterDeleteHook, componentMetadatumAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	componentMetadatumAfterDeleteHooks = []ComponentMetadatumHook{}

	AddComponentMetadatumHook(boil.BeforeUpsertHook, componentMetadatumBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	componentMetadatumBeforeUpsertHooks = []ComponentMetadatumHook{}

	AddComponentMetadatumHook(boil.AfterUpsertHook, componentMetadatumAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	componentMetadatumAfterUpsertHooks = []ComponentMetadatumHook{}
}

func testComponentMetadataInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testComponentMetadataInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(componentMetadatumColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testComponentMetadatumToOneServerComponentUsingServerComponent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ComponentMetadatum
	var foreign ServerComponent

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, componentMetadatumDBTypes, false, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serverComponentDBTypes, false, serverComponentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerComponent struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ServerComponentID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ServerComponent().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddServerComponentHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *ServerComponent) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ComponentMetadatumSlice{&local}
	if err = local.L.LoadServerComponent(ctx, tx, false, (*[]*ComponentMetadatum)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServerComponent == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ServerComponent = nil
	if err = local.L.LoadServerComponent(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServerComponent == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testComponentMetadatumToOneSetOpServerComponentUsingServerComponent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ComponentMetadatum
	var b, c ServerComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, componentMetadatumDBTypes, false, strmangle.SetComplement(componentMetadatumPrimaryKeyColumns, componentMetadatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serverComponentDBTypes, false, strmangle.SetComplement(serverComponentPrimaryKeyColumns, serverComponentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serverComponentDBTypes, false, strmangle.SetComplement(serverComponentPrimaryKeyColumns, serverComponentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServerComponent{&b, &c} {
		err = a.SetServerComponent(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ServerComponent != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ComponentMetadata[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ServerComponentID != x.ID {
			t.Error("foreign key was wrong value", a.ServerComponentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ServerComponentID))
		reflect.Indirect(reflect.ValueOf(&a.ServerComponentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ServerComponentID != x.ID {
			t.Error("foreign key was wrong value", a.ServerComponentID, x.ID)
		}
	}
}

func testComponentMetadataReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
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

func testComponentMetadataReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ComponentMetadatumSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testComponentMetadataSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ComponentMetadata().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	componentMetadatumDBTypes = map[string]string{`ID`: `uuid`, `ServerComponentID`: `uuid`, `Namespace`: `text`, `Data`: `jsonb`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                         = bytes.MinRead
)

func testComponentMetadataUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(componentMetadatumPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(componentMetadatumAllColumns) == len(componentMetadatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testComponentMetadataSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(componentMetadatumAllColumns) == len(componentMetadatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ComponentMetadatum{}
	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, componentMetadatumDBTypes, true, componentMetadatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(componentMetadatumAllColumns, componentMetadatumPrimaryKeyColumns) {
		fields = componentMetadatumAllColumns
	} else {
		fields = strmangle.SetComplement(
			componentMetadatumAllColumns,
			componentMetadatumPrimaryKeyColumns,
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

	slice := ComponentMetadatumSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testComponentMetadataUpsert(t *testing.T) {
	t.Parallel()

	if len(componentMetadatumAllColumns) == len(componentMetadatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ComponentMetadatum{}
	if err = randomize.Struct(seed, &o, componentMetadatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ComponentMetadatum: %s", err)
	}

	count, err := ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, componentMetadatumDBTypes, false, componentMetadatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ComponentMetadatum struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ComponentMetadatum: %s", err)
	}

	count, err = ComponentMetadata().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
