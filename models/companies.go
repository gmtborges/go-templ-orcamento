// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Company is an object representing the database table.
type Company struct {
	CompanyID   null.Int64  `boil:"company_id" json:"company_id,omitempty" toml:"company_id" yaml:"company_id,omitempty"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Address     null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	ContactInfo null.String `boil:"contact_info" json:"contact_info,omitempty" toml:"contact_info" yaml:"contact_info,omitempty"`
	CreatedAt   null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *companyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L companyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompanyColumns = struct {
	CompanyID   string
	Name        string
	Address     string
	ContactInfo string
	CreatedAt   string
	UpdatedAt   string
}{
	CompanyID:   "company_id",
	Name:        "name",
	Address:     "address",
	ContactInfo: "contact_info",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var CompanyTableColumns = struct {
	CompanyID   string
	Name        string
	Address     string
	ContactInfo string
	CreatedAt   string
	UpdatedAt   string
}{
	CompanyID:   "companies.company_id",
	Name:        "companies.name",
	Address:     "companies.address",
	ContactInfo: "companies.contact_info",
	CreatedAt:   "companies.created_at",
	UpdatedAt:   "companies.updated_at",
}

// Generated where

var CompanyWhere = struct {
	CompanyID   whereHelpernull_Int64
	Name        whereHelperstring
	Address     whereHelpernull_String
	ContactInfo whereHelpernull_String
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
}{
	CompanyID:   whereHelpernull_Int64{field: "\"companies\".\"company_id\""},
	Name:        whereHelperstring{field: "\"companies\".\"name\""},
	Address:     whereHelpernull_String{field: "\"companies\".\"address\""},
	ContactInfo: whereHelpernull_String{field: "\"companies\".\"contact_info\""},
	CreatedAt:   whereHelpernull_Time{field: "\"companies\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"companies\".\"updated_at\""},
}

// CompanyRels is where relationship names are stored.
var CompanyRels = struct {
	Biddings string
}{
	Biddings: "Biddings",
}

// companyR is where relationships are stored.
type companyR struct {
	Biddings BiddingSlice `boil:"Biddings" json:"Biddings" toml:"Biddings" yaml:"Biddings"`
}

// NewStruct creates a new relationship struct
func (*companyR) NewStruct() *companyR {
	return &companyR{}
}

func (r *companyR) GetBiddings() BiddingSlice {
	if r == nil {
		return nil
	}
	return r.Biddings
}

// companyL is where Load methods for each relationship are stored.
type companyL struct{}

var (
	companyAllColumns            = []string{"company_id", "name", "address", "contact_info", "created_at", "updated_at"}
	companyColumnsWithoutDefault = []string{"name"}
	companyColumnsWithDefault    = []string{"company_id", "address", "contact_info", "created_at", "updated_at"}
	companyPrimaryKeyColumns     = []string{"company_id"}
	companyGeneratedColumns      = []string{"company_id"}
)

type (
	// CompanySlice is an alias for a slice of pointers to Company.
	// This should almost always be used instead of []Company.
	CompanySlice []*Company
	// CompanyHook is the signature for custom Company hook methods
	CompanyHook func(context.Context, boil.ContextExecutor, *Company) error

	companyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	companyType                 = reflect.TypeOf(&Company{})
	companyMapping              = queries.MakeStructMapping(companyType)
	companyPrimaryKeyMapping, _ = queries.BindMapping(companyType, companyMapping, companyPrimaryKeyColumns)
	companyInsertCacheMut       sync.RWMutex
	companyInsertCache          = make(map[string]insertCache)
	companyUpdateCacheMut       sync.RWMutex
	companyUpdateCache          = make(map[string]updateCache)
	companyUpsertCacheMut       sync.RWMutex
	companyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var companyAfterSelectMu sync.Mutex
var companyAfterSelectHooks []CompanyHook

var companyBeforeInsertMu sync.Mutex
var companyBeforeInsertHooks []CompanyHook
var companyAfterInsertMu sync.Mutex
var companyAfterInsertHooks []CompanyHook

var companyBeforeUpdateMu sync.Mutex
var companyBeforeUpdateHooks []CompanyHook
var companyAfterUpdateMu sync.Mutex
var companyAfterUpdateHooks []CompanyHook

var companyBeforeDeleteMu sync.Mutex
var companyBeforeDeleteHooks []CompanyHook
var companyAfterDeleteMu sync.Mutex
var companyAfterDeleteHooks []CompanyHook

var companyBeforeUpsertMu sync.Mutex
var companyBeforeUpsertHooks []CompanyHook
var companyAfterUpsertMu sync.Mutex
var companyAfterUpsertHooks []CompanyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Company) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Company) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Company) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Company) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Company) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Company) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Company) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Company) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Company) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range companyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCompanyHook registers your hook function for all future operations.
func AddCompanyHook(hookPoint boil.HookPoint, companyHook CompanyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		companyAfterSelectMu.Lock()
		companyAfterSelectHooks = append(companyAfterSelectHooks, companyHook)
		companyAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		companyBeforeInsertMu.Lock()
		companyBeforeInsertHooks = append(companyBeforeInsertHooks, companyHook)
		companyBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		companyAfterInsertMu.Lock()
		companyAfterInsertHooks = append(companyAfterInsertHooks, companyHook)
		companyAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		companyBeforeUpdateMu.Lock()
		companyBeforeUpdateHooks = append(companyBeforeUpdateHooks, companyHook)
		companyBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		companyAfterUpdateMu.Lock()
		companyAfterUpdateHooks = append(companyAfterUpdateHooks, companyHook)
		companyAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		companyBeforeDeleteMu.Lock()
		companyBeforeDeleteHooks = append(companyBeforeDeleteHooks, companyHook)
		companyBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		companyAfterDeleteMu.Lock()
		companyAfterDeleteHooks = append(companyAfterDeleteHooks, companyHook)
		companyAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		companyBeforeUpsertMu.Lock()
		companyBeforeUpsertHooks = append(companyBeforeUpsertHooks, companyHook)
		companyBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		companyAfterUpsertMu.Lock()
		companyAfterUpsertHooks = append(companyAfterUpsertHooks, companyHook)
		companyAfterUpsertMu.Unlock()
	}
}

// One returns a single company record from the query.
func (q companyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Company, error) {
	o := &Company{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for companies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Company records from the query.
func (q companyQuery) All(ctx context.Context, exec boil.ContextExecutor) (CompanySlice, error) {
	var o []*Company

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Company slice")
	}

	if len(companyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Company records in the query.
func (q companyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count companies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q companyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if companies exists")
	}

	return count > 0, nil
}

// Biddings retrieves all the bidding's Biddings with an executor.
func (o *Company) Biddings(mods ...qm.QueryMod) biddingQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"biddings\".\"company_id\"=?", o.CompanyID),
	)

	return Biddings(queryMods...)
}

// LoadBiddings allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (companyL) LoadBiddings(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCompany interface{}, mods queries.Applicator) error {
	var slice []*Company
	var object *Company

	if singular {
		var ok bool
		object, ok = maybeCompany.(*Company)
		if !ok {
			object = new(Company)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCompany)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCompany))
			}
		}
	} else {
		s, ok := maybeCompany.(*[]*Company)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCompany)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCompany))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &companyR{}
		}
		args[object.CompanyID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &companyR{}
			}
			args[obj.CompanyID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`biddings`),
		qm.WhereIn(`biddings.company_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load biddings")
	}

	var resultSlice []*Bidding
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice biddings")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on biddings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for biddings")
	}

	if len(biddingAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Biddings = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &biddingR{}
			}
			foreign.R.Company = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.CompanyID, foreign.CompanyID) {
				local.R.Biddings = append(local.R.Biddings, foreign)
				if foreign.R == nil {
					foreign.R = &biddingR{}
				}
				foreign.R.Company = local
				break
			}
		}
	}

	return nil
}

// AddBiddings adds the given related objects to the existing relationships
// of the company, optionally inserting them as new records.
// Appends related to o.R.Biddings.
// Sets related.R.Company appropriately.
func (o *Company) AddBiddings(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Bidding) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.CompanyID, o.CompanyID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"biddings\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"company_id"}),
				strmangle.WhereClause("\"", "\"", 0, biddingPrimaryKeyColumns),
			)
			values := []interface{}{o.CompanyID, rel.BiddingID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.CompanyID, o.CompanyID)
		}
	}

	if o.R == nil {
		o.R = &companyR{
			Biddings: related,
		}
	} else {
		o.R.Biddings = append(o.R.Biddings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &biddingR{
				Company: o,
			}
		} else {
			rel.R.Company = o
		}
	}
	return nil
}

// Companies retrieves all the records using an executor.
func Companies(mods ...qm.QueryMod) companyQuery {
	mods = append(mods, qm.From("\"companies\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"companies\".*"})
	}

	return companyQuery{q}
}

// FindCompany retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompany(ctx context.Context, exec boil.ContextExecutor, companyID null.Int64, selectCols ...string) (*Company, error) {
	companyObj := &Company{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"companies\" where \"company_id\"=?", sel,
	)

	q := queries.Raw(query, companyID)

	err := q.Bind(ctx, exec, companyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from companies")
	}

	if err = companyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return companyObj, err
	}

	return companyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Company) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no companies provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(companyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	companyInsertCacheMut.RLock()
	cache, cached := companyInsertCache[key]
	companyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			companyAllColumns,
			companyColumnsWithDefault,
			companyColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, companyGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(companyType, companyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(companyType, companyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"companies\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"companies\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into companies")
	}

	if !cached {
		companyInsertCacheMut.Lock()
		companyInsertCache[key] = cache
		companyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Company.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Company) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	companyUpdateCacheMut.RLock()
	cache, cached := companyUpdateCache[key]
	companyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			companyAllColumns,
			companyPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, companyGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update companies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"companies\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, companyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(companyType, companyMapping, append(wl, companyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update companies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for companies")
	}

	if !cached {
		companyUpdateCacheMut.Lock()
		companyUpdateCache[key] = cache
		companyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q companyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for companies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompanySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"companies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, companyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in company slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all company")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Company) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no companies provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(companyColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	companyUpsertCacheMut.RLock()
	cache, cached := companyUpsertCache[key]
	companyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			companyAllColumns,
			companyColumnsWithDefault,
			companyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			companyAllColumns,
			companyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert companies, could not build update column list")
		}

		ret := strmangle.SetComplement(companyAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(companyPrimaryKeyColumns))
			copy(conflict, companyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"companies\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(companyType, companyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(companyType, companyMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert companies")
	}

	if !cached {
		companyUpsertCacheMut.Lock()
		companyUpsertCache[key] = cache
		companyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Company record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Company) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Company provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), companyPrimaryKeyMapping)
	sql := "DELETE FROM \"companies\" WHERE \"company_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for companies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q companyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no companyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from companies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for companies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompanySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(companyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"companies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, companyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from company slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for companies")
	}

	if len(companyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Company) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCompany(ctx, exec, o.CompanyID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompanySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CompanySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), companyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"companies\".* FROM \"companies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, companyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompanySlice")
	}

	*o = slice

	return nil
}

// CompanyExists checks if the Company row exists.
func CompanyExists(ctx context.Context, exec boil.ContextExecutor, companyID null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"companies\" where \"company_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, companyID)
	}
	row := exec.QueryRowContext(ctx, sql, companyID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if companies exists")
	}

	return exists, nil
}

// Exists checks if the Company row exists.
func (o *Company) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CompanyExists(ctx, exec, o.CompanyID)
}
