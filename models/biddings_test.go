// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testBiddings(t *testing.T) {
	t.Parallel()

	query := Biddings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBiddingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
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

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBiddingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Biddings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBiddingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BiddingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBiddingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BiddingExists(ctx, tx, o.BiddingID)
	if err != nil {
		t.Errorf("Unable to check if Bidding exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BiddingExists to return true, but got false.")
	}
}

func testBiddingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	biddingFound, err := FindBidding(ctx, tx, o.BiddingID)
	if err != nil {
		t.Error(err)
	}

	if biddingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBiddingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Biddings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBiddingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Biddings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBiddingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	biddingOne := &Bidding{}
	biddingTwo := &Bidding{}
	if err = randomize.Struct(seed, biddingOne, biddingDBTypes, false, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}
	if err = randomize.Struct(seed, biddingTwo, biddingDBTypes, false, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = biddingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = biddingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Biddings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBiddingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	biddingOne := &Bidding{}
	biddingTwo := &Bidding{}
	if err = randomize.Struct(seed, biddingOne, biddingDBTypes, false, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}
	if err = randomize.Struct(seed, biddingTwo, biddingDBTypes, false, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = biddingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = biddingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func biddingBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func biddingAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func biddingAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func biddingBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func biddingAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func biddingBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func biddingAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func biddingBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func biddingAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Bidding) error {
	*o = Bidding{}
	return nil
}

func testBiddingsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Bidding{}
	o := &Bidding{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, biddingDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Bidding object: %s", err)
	}

	AddBiddingHook(boil.BeforeInsertHook, biddingBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	biddingBeforeInsertHooks = []BiddingHook{}

	AddBiddingHook(boil.AfterInsertHook, biddingAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	biddingAfterInsertHooks = []BiddingHook{}

	AddBiddingHook(boil.AfterSelectHook, biddingAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	biddingAfterSelectHooks = []BiddingHook{}

	AddBiddingHook(boil.BeforeUpdateHook, biddingBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	biddingBeforeUpdateHooks = []BiddingHook{}

	AddBiddingHook(boil.AfterUpdateHook, biddingAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	biddingAfterUpdateHooks = []BiddingHook{}

	AddBiddingHook(boil.BeforeDeleteHook, biddingBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	biddingBeforeDeleteHooks = []BiddingHook{}

	AddBiddingHook(boil.AfterDeleteHook, biddingAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	biddingAfterDeleteHooks = []BiddingHook{}

	AddBiddingHook(boil.BeforeUpsertHook, biddingBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	biddingBeforeUpsertHooks = []BiddingHook{}

	AddBiddingHook(boil.AfterUpsertHook, biddingAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	biddingAfterUpsertHooks = []BiddingHook{}
}

func testBiddingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBiddingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(biddingColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBiddingToManyAutoOffers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bidding
	var b, c AutoOffer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, autoOfferDBTypes, false, autoOfferColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, autoOfferDBTypes, false, autoOfferColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.BiddingID, a.BiddingID)
	queries.Assign(&c.BiddingID, a.BiddingID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.AutoOffers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.BiddingID, b.BiddingID) {
			bFound = true
		}
		if queries.Equal(v.BiddingID, c.BiddingID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BiddingSlice{&a}
	if err = a.L.LoadAutoOffers(ctx, tx, false, (*[]*Bidding)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AutoOffers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AutoOffers = nil
	if err = a.L.LoadAutoOffers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AutoOffers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBiddingToManyAddOpAutoOffers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bidding
	var b, c, d, e AutoOffer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, biddingDBTypes, false, strmangle.SetComplement(biddingPrimaryKeyColumns, biddingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AutoOffer{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, autoOfferDBTypes, false, strmangle.SetComplement(autoOfferPrimaryKeyColumns, autoOfferColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AutoOffer{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAutoOffers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.BiddingID, first.BiddingID) {
			t.Error("foreign key was wrong value", a.BiddingID, first.BiddingID)
		}
		if !queries.Equal(a.BiddingID, second.BiddingID) {
			t.Error("foreign key was wrong value", a.BiddingID, second.BiddingID)
		}

		if first.R.Bidding != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Bidding != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AutoOffers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AutoOffers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AutoOffers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBiddingToOneAutoPartsCategoryUsingCategory(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Bidding
	var foreign AutoPartsCategory

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, biddingDBTypes, false, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, autoPartsCategoryDBTypes, true, autoPartsCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutoPartsCategory struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.CategoryID, foreign.CategoryID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Category().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.CategoryID, foreign.CategoryID) {
		t.Errorf("want: %v, got %v", foreign.CategoryID, check.CategoryID)
	}

	ranAfterSelectHook := false
	AddAutoPartsCategoryHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *AutoPartsCategory) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := BiddingSlice{&local}
	if err = local.L.LoadCategory(ctx, tx, false, (*[]*Bidding)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Category == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Category = nil
	if err = local.L.LoadCategory(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Category == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testBiddingToOneCompanyUsingCompany(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Bidding
	var foreign Company

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, biddingDBTypes, false, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, companyDBTypes, true, companyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Company struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.CompanyID, foreign.CompanyID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Company().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.CompanyID, foreign.CompanyID) {
		t.Errorf("want: %v, got %v", foreign.CompanyID, check.CompanyID)
	}

	ranAfterSelectHook := false
	AddCompanyHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Company) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := BiddingSlice{&local}
	if err = local.L.LoadCompany(ctx, tx, false, (*[]*Bidding)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Company == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Company = nil
	if err = local.L.LoadCompany(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Company == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testBiddingToOneSetOpAutoPartsCategoryUsingCategory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bidding
	var b, c AutoPartsCategory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, biddingDBTypes, false, strmangle.SetComplement(biddingPrimaryKeyColumns, biddingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, autoPartsCategoryDBTypes, false, strmangle.SetComplement(autoPartsCategoryPrimaryKeyColumns, autoPartsCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, autoPartsCategoryDBTypes, false, strmangle.SetComplement(autoPartsCategoryPrimaryKeyColumns, autoPartsCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AutoPartsCategory{&b, &c} {
		err = a.SetCategory(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Category != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CategoryBiddings[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.CategoryID, x.CategoryID) {
			t.Error("foreign key was wrong value", a.CategoryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CategoryID))
		reflect.Indirect(reflect.ValueOf(&a.CategoryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.CategoryID, x.CategoryID) {
			t.Error("foreign key was wrong value", a.CategoryID, x.CategoryID)
		}
	}
}
func testBiddingToOneSetOpCompanyUsingCompany(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Bidding
	var b, c Company

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, biddingDBTypes, false, strmangle.SetComplement(biddingPrimaryKeyColumns, biddingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Company{&b, &c} {
		err = a.SetCompany(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Company != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Biddings[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.CompanyID, x.CompanyID) {
			t.Error("foreign key was wrong value", a.CompanyID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CompanyID))
		reflect.Indirect(reflect.ValueOf(&a.CompanyID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.CompanyID, x.CompanyID) {
			t.Error("foreign key was wrong value", a.CompanyID, x.CompanyID)
		}
	}
}

func testBiddingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
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

func testBiddingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BiddingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBiddingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Biddings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	biddingDBTypes = map[string]string{`BiddingID`: `INTEGER`, `CompanyID`: `INTEGER`, `Title`: `TEXT`, `Description`: `TEXT`, `StartDate`: `DATE`, `EndDate`: `DATE`, `CategoryID`: `INTEGER`, `CreatedAt`: `DATETIME`, `UpdatedAt`: `DATETIME`}
	_              = bytes.MinRead
)

func testBiddingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(biddingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(biddingAllColumns) == len(biddingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBiddingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(biddingAllColumns) == len(biddingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Bidding{}
	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, biddingDBTypes, true, biddingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(biddingAllColumns, biddingPrimaryKeyColumns) {
		fields = biddingAllColumns
	} else {
		fields = strmangle.SetComplement(
			biddingAllColumns,
			biddingPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, biddingGeneratedColumns)
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

	slice := BiddingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBiddingsUpsert(t *testing.T) {
	t.Parallel()
	if len(biddingAllColumns) == len(biddingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Bidding{}
	if err = randomize.Struct(seed, &o, biddingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Bidding: %s", err)
	}

	count, err := Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, biddingDBTypes, false, biddingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Bidding struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Bidding: %s", err)
	}

	count, err = Biddings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
