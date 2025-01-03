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

func testServerCredentialTypes(t *testing.T) {
	t.Parallel()

	query := ServerCredentialTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServerCredentialTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
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

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServerCredentialTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServerCredentialTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServerCredentialTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServerCredentialTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServerCredentialTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServerCredentialTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServerCredentialType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServerCredentialTypeExists to return true, but got false.")
	}
}

func testServerCredentialTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serverCredentialTypeFound, err := FindServerCredentialType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serverCredentialTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServerCredentialTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServerCredentialTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServerCredentialTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServerCredentialTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServerCredentialTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serverCredentialTypeOne := &ServerCredentialType{}
	serverCredentialTypeTwo := &ServerCredentialType{}
	if err = randomize.Struct(seed, serverCredentialTypeOne, serverCredentialTypeDBTypes, false, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}
	if err = randomize.Struct(seed, serverCredentialTypeTwo, serverCredentialTypeDBTypes, false, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serverCredentialTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serverCredentialTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServerCredentialTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServerCredentialTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serverCredentialTypeOne := &ServerCredentialType{}
	serverCredentialTypeTwo := &ServerCredentialType{}
	if err = randomize.Struct(seed, serverCredentialTypeOne, serverCredentialTypeDBTypes, false, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}
	if err = randomize.Struct(seed, serverCredentialTypeTwo, serverCredentialTypeDBTypes, false, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serverCredentialTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serverCredentialTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serverCredentialTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func serverCredentialTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func serverCredentialTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func serverCredentialTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func serverCredentialTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func serverCredentialTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func serverCredentialTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func serverCredentialTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func serverCredentialTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServerCredentialType) error {
	*o = ServerCredentialType{}
	return nil
}

func testServerCredentialTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServerCredentialType{}
	o := &ServerCredentialType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType object: %s", err)
	}

	AddServerCredentialTypeHook(boil.BeforeInsertHook, serverCredentialTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeBeforeInsertHooks = []ServerCredentialTypeHook{}

	AddServerCredentialTypeHook(boil.AfterInsertHook, serverCredentialTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeAfterInsertHooks = []ServerCredentialTypeHook{}

	AddServerCredentialTypeHook(boil.AfterSelectHook, serverCredentialTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeAfterSelectHooks = []ServerCredentialTypeHook{}

	AddServerCredentialTypeHook(boil.BeforeUpdateHook, serverCredentialTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeBeforeUpdateHooks = []ServerCredentialTypeHook{}

	AddServerCredentialTypeHook(boil.AfterUpdateHook, serverCredentialTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeAfterUpdateHooks = []ServerCredentialTypeHook{}

	AddServerCredentialTypeHook(boil.BeforeDeleteHook, serverCredentialTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeBeforeDeleteHooks = []ServerCredentialTypeHook{}

	AddServerCredentialTypeHook(boil.AfterDeleteHook, serverCredentialTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeAfterDeleteHooks = []ServerCredentialTypeHook{}

	AddServerCredentialTypeHook(boil.BeforeUpsertHook, serverCredentialTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeBeforeUpsertHooks = []ServerCredentialTypeHook{}

	AddServerCredentialTypeHook(boil.AfterUpsertHook, serverCredentialTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serverCredentialTypeAfterUpsertHooks = []ServerCredentialTypeHook{}
}

func testServerCredentialTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServerCredentialTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serverCredentialTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServerCredentialTypeToManyServerCredentials(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServerCredentialType
	var b, c ServerCredential

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, serverCredentialDBTypes, false, serverCredentialColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serverCredentialDBTypes, false, serverCredentialColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ServerCredentialTypeID = a.ID
	c.ServerCredentialTypeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ServerCredentials().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ServerCredentialTypeID == b.ServerCredentialTypeID {
			bFound = true
		}
		if v.ServerCredentialTypeID == c.ServerCredentialTypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ServerCredentialTypeSlice{&a}
	if err = a.L.LoadServerCredentials(ctx, tx, false, (*[]*ServerCredentialType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ServerCredentials); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ServerCredentials = nil
	if err = a.L.LoadServerCredentials(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ServerCredentials); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testServerCredentialTypeToManyAddOpServerCredentials(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServerCredentialType
	var b, c, d, e ServerCredential

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serverCredentialTypeDBTypes, false, strmangle.SetComplement(serverCredentialTypePrimaryKeyColumns, serverCredentialTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ServerCredential{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, serverCredentialDBTypes, false, strmangle.SetComplement(serverCredentialPrimaryKeyColumns, serverCredentialColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*ServerCredential{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddServerCredentials(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ServerCredentialTypeID {
			t.Error("foreign key was wrong value", a.ID, first.ServerCredentialTypeID)
		}
		if a.ID != second.ServerCredentialTypeID {
			t.Error("foreign key was wrong value", a.ID, second.ServerCredentialTypeID)
		}

		if first.R.ServerCredentialType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ServerCredentialType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ServerCredentials[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ServerCredentials[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ServerCredentials().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testServerCredentialTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
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

func testServerCredentialTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServerCredentialTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServerCredentialTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServerCredentialTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serverCredentialTypeDBTypes = map[string]string{`ID`: `uuid`, `Name`: `text`, `Slug`: `text`, `Builtin`: `boolean`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                           = bytes.MinRead
)

func testServerCredentialTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serverCredentialTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serverCredentialTypeAllColumns) == len(serverCredentialTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServerCredentialTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serverCredentialTypeAllColumns) == len(serverCredentialTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServerCredentialType{}
	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serverCredentialTypeDBTypes, true, serverCredentialTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serverCredentialTypeAllColumns, serverCredentialTypePrimaryKeyColumns) {
		fields = serverCredentialTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			serverCredentialTypeAllColumns,
			serverCredentialTypePrimaryKeyColumns,
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

	slice := ServerCredentialTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServerCredentialTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(serverCredentialTypeAllColumns) == len(serverCredentialTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServerCredentialType{}
	if err = randomize.Struct(seed, &o, serverCredentialTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServerCredentialType: %s", err)
	}

	count, err := ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serverCredentialTypeDBTypes, false, serverCredentialTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServerCredentialType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServerCredentialType: %s", err)
	}

	count, err = ServerCredentialTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
