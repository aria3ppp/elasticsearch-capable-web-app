// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testPostsAudits(t *testing.T) {
	t.Parallel()

	query := PostsAudits()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPostsAuditsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
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

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsAuditsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PostsAudits().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsAuditsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PostsAuditSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPostsAuditsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PostsAuditExists(ctx, tx, o.ID, o.ContributedBy, o.ContributedAt)
	if err != nil {
		t.Errorf("Unable to check if PostsAudit exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PostsAuditExists to return true, but got false.")
	}
}

func testPostsAuditsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	postsAuditFound, err := FindPostsAudit(ctx, tx, o.ID, o.ContributedBy, o.ContributedAt)
	if err != nil {
		t.Error(err)
	}

	if postsAuditFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPostsAuditsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PostsAudits().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPostsAuditsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PostsAudits().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPostsAuditsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	postsAuditOne := &PostsAudit{}
	postsAuditTwo := &PostsAudit{}
	if err = randomize.Struct(seed, postsAuditOne, postsAuditDBTypes, false, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}
	if err = randomize.Struct(seed, postsAuditTwo, postsAuditDBTypes, false, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = postsAuditOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = postsAuditTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PostsAudits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPostsAuditsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	postsAuditOne := &PostsAudit{}
	postsAuditTwo := &PostsAudit{}
	if err = randomize.Struct(seed, postsAuditOne, postsAuditDBTypes, false, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}
	if err = randomize.Struct(seed, postsAuditTwo, postsAuditDBTypes, false, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = postsAuditOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = postsAuditTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func postsAuditBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func postsAuditAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func postsAuditAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func postsAuditBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func postsAuditAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func postsAuditBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func postsAuditAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func postsAuditBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func postsAuditAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PostsAudit) error {
	*o = PostsAudit{}
	return nil
}

func testPostsAuditsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PostsAudit{}
	o := &PostsAudit{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, postsAuditDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PostsAudit object: %s", err)
	}

	AddPostsAuditHook(boil.BeforeInsertHook, postsAuditBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	postsAuditBeforeInsertHooks = []PostsAuditHook{}

	AddPostsAuditHook(boil.AfterInsertHook, postsAuditAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	postsAuditAfterInsertHooks = []PostsAuditHook{}

	AddPostsAuditHook(boil.AfterSelectHook, postsAuditAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	postsAuditAfterSelectHooks = []PostsAuditHook{}

	AddPostsAuditHook(boil.BeforeUpdateHook, postsAuditBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	postsAuditBeforeUpdateHooks = []PostsAuditHook{}

	AddPostsAuditHook(boil.AfterUpdateHook, postsAuditAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	postsAuditAfterUpdateHooks = []PostsAuditHook{}

	AddPostsAuditHook(boil.BeforeDeleteHook, postsAuditBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	postsAuditBeforeDeleteHooks = []PostsAuditHook{}

	AddPostsAuditHook(boil.AfterDeleteHook, postsAuditAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	postsAuditAfterDeleteHooks = []PostsAuditHook{}

	AddPostsAuditHook(boil.BeforeUpsertHook, postsAuditBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	postsAuditBeforeUpsertHooks = []PostsAuditHook{}

	AddPostsAuditHook(boil.AfterUpsertHook, postsAuditAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	postsAuditAfterUpsertHooks = []PostsAuditHook{}
}

func testPostsAuditsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPostsAuditsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(postsAuditColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPostsAuditsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
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

func testPostsAuditsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PostsAuditSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPostsAuditsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PostsAudits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	postsAuditDBTypes = map[string]string{`ID`: `integer`, `Title`: `character varying`, `Body`: `text`, `ContributedBy`: `integer`, `ContributedAt`: `timestamp with time zone`, `Deleted`: `boolean`}
	_                 = bytes.MinRead
)

func testPostsAuditsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(postsAuditPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(postsAuditAllColumns) == len(postsAuditPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPostsAuditsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(postsAuditAllColumns) == len(postsAuditPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PostsAudit{}
	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, postsAuditDBTypes, true, postsAuditPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(postsAuditAllColumns, postsAuditPrimaryKeyColumns) {
		fields = postsAuditAllColumns
	} else {
		fields = strmangle.SetComplement(
			postsAuditAllColumns,
			postsAuditPrimaryKeyColumns,
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

	slice := PostsAuditSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPostsAuditsUpsert(t *testing.T) {
	t.Parallel()

	if len(postsAuditAllColumns) == len(postsAuditPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PostsAudit{}
	if err = randomize.Struct(seed, &o, postsAuditDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PostsAudit: %s", err)
	}

	count, err := PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, postsAuditDBTypes, false, postsAuditPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PostsAudit struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PostsAudit: %s", err)
	}

	count, err = PostsAudits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
