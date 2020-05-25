// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package nixie

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testContactForms(t *testing.T) {
	t.Parallel()

	query := ContactForms()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testContactFormsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
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

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContactFormsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ContactForms().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContactFormsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContactFormSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContactFormsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ContactFormExists(ctx, tx, o.ContactFormID)
	if err != nil {
		t.Errorf("Unable to check if ContactForm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContactFormExists to return true, but got false.")
	}
}

func testContactFormsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	contactFormFound, err := FindContactForm(ctx, tx, o.ContactFormID)
	if err != nil {
		t.Error(err)
	}

	if contactFormFound == nil {
		t.Error("want a record, got nil")
	}
}

func testContactFormsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ContactForms().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testContactFormsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ContactForms().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContactFormsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contactFormOne := &ContactForm{}
	contactFormTwo := &ContactForm{}
	if err = randomize.Struct(seed, contactFormOne, contactFormDBTypes, false, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}
	if err = randomize.Struct(seed, contactFormTwo, contactFormDBTypes, false, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = contactFormOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = contactFormTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ContactForms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContactFormsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	contactFormOne := &ContactForm{}
	contactFormTwo := &ContactForm{}
	if err = randomize.Struct(seed, contactFormOne, contactFormDBTypes, false, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}
	if err = randomize.Struct(seed, contactFormTwo, contactFormDBTypes, false, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = contactFormOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = contactFormTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func contactFormBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func contactFormAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func contactFormAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func contactFormBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func contactFormAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func contactFormBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func contactFormAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func contactFormBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func contactFormAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ContactForm) error {
	*o = ContactForm{}
	return nil
}

func testContactFormsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ContactForm{}
	o := &ContactForm{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, contactFormDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ContactForm object: %s", err)
	}

	AddContactFormHook(boil.BeforeInsertHook, contactFormBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	contactFormBeforeInsertHooks = []ContactFormHook{}

	AddContactFormHook(boil.AfterInsertHook, contactFormAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	contactFormAfterInsertHooks = []ContactFormHook{}

	AddContactFormHook(boil.AfterSelectHook, contactFormAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	contactFormAfterSelectHooks = []ContactFormHook{}

	AddContactFormHook(boil.BeforeUpdateHook, contactFormBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	contactFormBeforeUpdateHooks = []ContactFormHook{}

	AddContactFormHook(boil.AfterUpdateHook, contactFormAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	contactFormAfterUpdateHooks = []ContactFormHook{}

	AddContactFormHook(boil.BeforeDeleteHook, contactFormBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	contactFormBeforeDeleteHooks = []ContactFormHook{}

	AddContactFormHook(boil.AfterDeleteHook, contactFormAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	contactFormAfterDeleteHooks = []ContactFormHook{}

	AddContactFormHook(boil.BeforeUpsertHook, contactFormBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	contactFormBeforeUpsertHooks = []ContactFormHook{}

	AddContactFormHook(boil.AfterUpsertHook, contactFormAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	contactFormAfterUpsertHooks = []ContactFormHook{}
}

func testContactFormsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContactFormsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(contactFormColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContactFormsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
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

func testContactFormsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContactFormSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testContactFormsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ContactForms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	contactFormDBTypes = map[string]string{`ContactFormID`: `bigint`, `Email`: `text`, `Description`: `text`, `Domain`: `text`}
	_                  = bytes.MinRead
)

func testContactFormsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(contactFormPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(contactFormAllColumns) == len(contactFormPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testContactFormsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(contactFormAllColumns) == len(contactFormPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ContactForm{}
	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, contactFormDBTypes, true, contactFormPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(contactFormAllColumns, contactFormPrimaryKeyColumns) {
		fields = contactFormAllColumns
	} else {
		fields = strmangle.SetComplement(
			contactFormAllColumns,
			contactFormPrimaryKeyColumns,
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

	slice := ContactFormSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testContactFormsUpsert(t *testing.T) {
	t.Parallel()

	if len(contactFormAllColumns) == len(contactFormPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ContactForm{}
	if err = randomize.Struct(seed, &o, contactFormDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ContactForm: %s", err)
	}

	count, err := ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, contactFormDBTypes, false, contactFormPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContactForm struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ContactForm: %s", err)
	}

	count, err = ContactForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}