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

func testAttributes(t *testing.T) {
	t.Parallel()

	query := Attributes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAttributesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
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

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAttributesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Attributes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAttributesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AttributeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAttributesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AttributeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Attribute exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AttributeExists to return true, but got false.")
	}
}

func testAttributesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	attributeFound, err := FindAttribute(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if attributeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAttributesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Attributes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAttributesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Attributes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAttributesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	attributeOne := &Attribute{}
	attributeTwo := &Attribute{}
	if err = randomize.Struct(seed, attributeOne, attributeDBTypes, false, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}
	if err = randomize.Struct(seed, attributeTwo, attributeDBTypes, false, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = attributeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = attributeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Attributes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAttributesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	attributeOne := &Attribute{}
	attributeTwo := &Attribute{}
	if err = randomize.Struct(seed, attributeOne, attributeDBTypes, false, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}
	if err = randomize.Struct(seed, attributeTwo, attributeDBTypes, false, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = attributeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = attributeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func attributeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func attributeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func attributeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func attributeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func attributeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func attributeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func attributeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func attributeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func attributeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Attribute) error {
	*o = Attribute{}
	return nil
}

func testAttributesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Attribute{}
	o := &Attribute{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, attributeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Attribute object: %s", err)
	}

	AddAttributeHook(boil.BeforeInsertHook, attributeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	attributeBeforeInsertHooks = []AttributeHook{}

	AddAttributeHook(boil.AfterInsertHook, attributeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	attributeAfterInsertHooks = []AttributeHook{}

	AddAttributeHook(boil.AfterSelectHook, attributeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	attributeAfterSelectHooks = []AttributeHook{}

	AddAttributeHook(boil.BeforeUpdateHook, attributeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	attributeBeforeUpdateHooks = []AttributeHook{}

	AddAttributeHook(boil.AfterUpdateHook, attributeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	attributeAfterUpdateHooks = []AttributeHook{}

	AddAttributeHook(boil.BeforeDeleteHook, attributeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	attributeBeforeDeleteHooks = []AttributeHook{}

	AddAttributeHook(boil.AfterDeleteHook, attributeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	attributeAfterDeleteHooks = []AttributeHook{}

	AddAttributeHook(boil.BeforeUpsertHook, attributeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	attributeBeforeUpsertHooks = []AttributeHook{}

	AddAttributeHook(boil.AfterUpsertHook, attributeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	attributeAfterUpsertHooks = []AttributeHook{}
}

func testAttributesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAttributesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(attributeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAttributeToOneServerComponentUsingServerComponent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Attribute
	var foreign ServerComponent

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serverComponentDBTypes, false, serverComponentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServerComponent struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ServerComponentID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ServerComponent().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddServerComponentHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *ServerComponent) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := AttributeSlice{&local}
	if err = local.L.LoadServerComponent(ctx, tx, false, (*[]*Attribute)(&slice), nil); err != nil {
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

func testAttributeToOneServerUsingServer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Attribute
	var foreign Server

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serverDBTypes, false, serverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Server struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ServerID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Server().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddServerHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Server) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := AttributeSlice{&local}
	if err = local.L.LoadServer(ctx, tx, false, (*[]*Attribute)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Server == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Server = nil
	if err = local.L.LoadServer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Server == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testAttributeToOneSetOpServerComponentUsingServerComponent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Attribute
	var b, c ServerComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, attributeDBTypes, false, strmangle.SetComplement(attributePrimaryKeyColumns, attributeColumnsWithoutDefault)...); err != nil {
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

		if x.R.Attributes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ServerComponentID, x.ID) {
			t.Error("foreign key was wrong value", a.ServerComponentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ServerComponentID))
		reflect.Indirect(reflect.ValueOf(&a.ServerComponentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ServerComponentID, x.ID) {
			t.Error("foreign key was wrong value", a.ServerComponentID, x.ID)
		}
	}
}

func testAttributeToOneRemoveOpServerComponentUsingServerComponent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Attribute
	var b ServerComponent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, attributeDBTypes, false, strmangle.SetComplement(attributePrimaryKeyColumns, attributeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serverComponentDBTypes, false, strmangle.SetComplement(serverComponentPrimaryKeyColumns, serverComponentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetServerComponent(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveServerComponent(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ServerComponent().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ServerComponent != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ServerComponentID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Attributes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testAttributeToOneSetOpServerUsingServer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Attribute
	var b, c Server

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, attributeDBTypes, false, strmangle.SetComplement(attributePrimaryKeyColumns, attributeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serverDBTypes, false, strmangle.SetComplement(serverPrimaryKeyColumns, serverColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serverDBTypes, false, strmangle.SetComplement(serverPrimaryKeyColumns, serverColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Server{&b, &c} {
		err = a.SetServer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Server != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Attributes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ServerID, x.ID) {
			t.Error("foreign key was wrong value", a.ServerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ServerID))
		reflect.Indirect(reflect.ValueOf(&a.ServerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ServerID, x.ID) {
			t.Error("foreign key was wrong value", a.ServerID, x.ID)
		}
	}
}

func testAttributeToOneRemoveOpServerUsingServer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Attribute
	var b Server

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, attributeDBTypes, false, strmangle.SetComplement(attributePrimaryKeyColumns, attributeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serverDBTypes, false, strmangle.SetComplement(serverPrimaryKeyColumns, serverColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetServer(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveServer(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Server().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Server != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ServerID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Attributes) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testAttributesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
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

func testAttributesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AttributeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAttributesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Attributes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	attributeDBTypes = map[string]string{`ID`: `uuid`, `ServerID`: `uuid`, `ServerComponentID`: `uuid`, `Namespace`: `text`, `Data`: `jsonb`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                = bytes.MinRead
)

func testAttributesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(attributePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(attributeAllColumns) == len(attributePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAttributesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(attributeAllColumns) == len(attributePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Attribute{}
	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, attributeDBTypes, true, attributePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(attributeAllColumns, attributePrimaryKeyColumns) {
		fields = attributeAllColumns
	} else {
		fields = strmangle.SetComplement(
			attributeAllColumns,
			attributePrimaryKeyColumns,
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

	slice := AttributeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAttributesUpsert(t *testing.T) {
	t.Parallel()

	if len(attributeAllColumns) == len(attributePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Attribute{}
	if err = randomize.Struct(seed, &o, attributeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Attribute: %s", err)
	}

	count, err := Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, attributeDBTypes, false, attributePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Attribute struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Attribute: %s", err)
	}

	count, err = Attributes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
