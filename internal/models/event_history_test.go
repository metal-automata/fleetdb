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

func testEventHistories(t *testing.T) {
	t.Parallel()

	query := EventHistories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEventHistoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
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

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventHistoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EventHistories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventHistoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EventHistorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventHistoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EventHistoryExists(ctx, tx, o.EventID, o.EventType, o.TargetServer)
	if err != nil {
		t.Errorf("Unable to check if EventHistory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EventHistoryExists to return true, but got false.")
	}
}

func testEventHistoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	eventHistoryFound, err := FindEventHistory(ctx, tx, o.EventID, o.EventType, o.TargetServer)
	if err != nil {
		t.Error(err)
	}

	if eventHistoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEventHistoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EventHistories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEventHistoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EventHistories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEventHistoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	eventHistoryOne := &EventHistory{}
	eventHistoryTwo := &EventHistory{}
	if err = randomize.Struct(seed, eventHistoryOne, eventHistoryDBTypes, false, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, eventHistoryTwo, eventHistoryDBTypes, false, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = eventHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = eventHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EventHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEventHistoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	eventHistoryOne := &EventHistory{}
	eventHistoryTwo := &EventHistory{}
	if err = randomize.Struct(seed, eventHistoryOne, eventHistoryDBTypes, false, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, eventHistoryTwo, eventHistoryDBTypes, false, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = eventHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = eventHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func eventHistoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func eventHistoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func eventHistoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func eventHistoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func eventHistoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func eventHistoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func eventHistoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func eventHistoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func eventHistoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EventHistory) error {
	*o = EventHistory{}
	return nil
}

func testEventHistoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EventHistory{}
	o := &EventHistory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EventHistory object: %s", err)
	}

	AddEventHistoryHook(boil.BeforeInsertHook, eventHistoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	eventHistoryBeforeInsertHooks = []EventHistoryHook{}

	AddEventHistoryHook(boil.AfterInsertHook, eventHistoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	eventHistoryAfterInsertHooks = []EventHistoryHook{}

	AddEventHistoryHook(boil.AfterSelectHook, eventHistoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	eventHistoryAfterSelectHooks = []EventHistoryHook{}

	AddEventHistoryHook(boil.BeforeUpdateHook, eventHistoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	eventHistoryBeforeUpdateHooks = []EventHistoryHook{}

	AddEventHistoryHook(boil.AfterUpdateHook, eventHistoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	eventHistoryAfterUpdateHooks = []EventHistoryHook{}

	AddEventHistoryHook(boil.BeforeDeleteHook, eventHistoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	eventHistoryBeforeDeleteHooks = []EventHistoryHook{}

	AddEventHistoryHook(boil.AfterDeleteHook, eventHistoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	eventHistoryAfterDeleteHooks = []EventHistoryHook{}

	AddEventHistoryHook(boil.BeforeUpsertHook, eventHistoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	eventHistoryBeforeUpsertHooks = []EventHistoryHook{}

	AddEventHistoryHook(boil.AfterUpsertHook, eventHistoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	eventHistoryAfterUpsertHooks = []EventHistoryHook{}
}

func testEventHistoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventHistoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(eventHistoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventHistoryToOneServerUsingTargetServerServer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local EventHistory
	var foreign Server

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, eventHistoryDBTypes, false, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serverDBTypes, false, serverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Server struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TargetServer = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.TargetServerServer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddServerHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Server) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := EventHistorySlice{&local}
	if err = local.L.LoadTargetServerServer(ctx, tx, false, (*[]*EventHistory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TargetServerServer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TargetServerServer = nil
	if err = local.L.LoadTargetServerServer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TargetServerServer == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testEventHistoryToOneSetOpServerUsingTargetServerServer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EventHistory
	var b, c Server

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, eventHistoryDBTypes, false, strmangle.SetComplement(eventHistoryPrimaryKeyColumns, eventHistoryColumnsWithoutDefault)...); err != nil {
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
		err = a.SetTargetServerServer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TargetServerServer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TargetServerEventHistories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TargetServer != x.ID {
			t.Error("foreign key was wrong value", a.TargetServer)
		}

		if exists, err := EventHistoryExists(ctx, tx, a.EventID, a.EventType, a.TargetServer); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testEventHistoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
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

func testEventHistoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EventHistorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEventHistoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EventHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	eventHistoryDBTypes = map[string]string{`EventID`: `uuid`, `EventType`: `text`, `EventStart`: `timestamp with time zone`, `EventEnd`: `timestamp with time zone`, `TargetServer`: `uuid`, `Parameters`: `jsonb`, `FinalState`: `text`, `FinalStatus`: `jsonb`}
	_                   = bytes.MinRead
)

func testEventHistoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(eventHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(eventHistoryAllColumns) == len(eventHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEventHistoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(eventHistoryAllColumns) == len(eventHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EventHistory{}
	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, eventHistoryDBTypes, true, eventHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(eventHistoryAllColumns, eventHistoryPrimaryKeyColumns) {
		fields = eventHistoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			eventHistoryAllColumns,
			eventHistoryPrimaryKeyColumns,
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

	slice := EventHistorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEventHistoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(eventHistoryAllColumns) == len(eventHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EventHistory{}
	if err = randomize.Struct(seed, &o, eventHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EventHistory: %s", err)
	}

	count, err := EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, eventHistoryDBTypes, false, eventHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EventHistory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EventHistory: %s", err)
	}

	count, err = EventHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
