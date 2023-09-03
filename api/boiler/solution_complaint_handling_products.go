// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiler

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// SolutionComplaintHandlingProduct is an object representing the database table.
type SolutionComplaintHandlingProduct struct { // クレーム対応プロダクトID
	ID []byte `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ソリューションプロダクトID
	SolutionProductID []byte `boil:"solution_product_id" json:"solution_product_id" toml:"solution_product_id" yaml:"solution_product_id"`
	// データ作成日
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// データ更新日
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *solutionComplaintHandlingProductR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L solutionComplaintHandlingProductL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SolutionComplaintHandlingProductColumns = struct {
	ID                string
	SolutionProductID string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	SolutionProductID: "solution_product_id",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

var SolutionComplaintHandlingProductTableColumns = struct {
	ID                string
	SolutionProductID string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "solution_complaint_handling_products.id",
	SolutionProductID: "solution_complaint_handling_products.solution_product_id",
	CreatedAt:         "solution_complaint_handling_products.created_at",
	UpdatedAt:         "solution_complaint_handling_products.updated_at",
}

// Generated where

var SolutionComplaintHandlingProductWhere = struct {
	ID                whereHelper__byte
	SolutionProductID whereHelper__byte
	CreatedAt         whereHelpertime_Time
	UpdatedAt         whereHelpertime_Time
}{
	ID:                whereHelper__byte{field: "\"solution_complaint_handling_products\".\"id\""},
	SolutionProductID: whereHelper__byte{field: "\"solution_complaint_handling_products\".\"solution_product_id\""},
	CreatedAt:         whereHelpertime_Time{field: "\"solution_complaint_handling_products\".\"created_at\""},
	UpdatedAt:         whereHelpertime_Time{field: "\"solution_complaint_handling_products\".\"updated_at\""},
}

// SolutionComplaintHandlingProductRels is where relationship names are stored.
var SolutionComplaintHandlingProductRels = struct {
	SolutionProduct string
}{
	SolutionProduct: "SolutionProduct",
}

// solutionComplaintHandlingProductR is where relationships are stored.
type solutionComplaintHandlingProductR struct {
	SolutionProduct *SolutionProduct `boil:"SolutionProduct" json:"SolutionProduct" toml:"SolutionProduct" yaml:"SolutionProduct"`
}

// NewStruct creates a new relationship struct
func (*solutionComplaintHandlingProductR) NewStruct() *solutionComplaintHandlingProductR {
	return &solutionComplaintHandlingProductR{}
}

func (r *solutionComplaintHandlingProductR) GetSolutionProduct() *SolutionProduct {
	if r == nil {
		return nil
	}
	return r.SolutionProduct
}

// solutionComplaintHandlingProductL is where Load methods for each relationship are stored.
type solutionComplaintHandlingProductL struct{}

var (
	solutionComplaintHandlingProductAllColumns            = []string{"id", "solution_product_id", "created_at", "updated_at"}
	solutionComplaintHandlingProductColumnsWithoutDefault = []string{"solution_product_id"}
	solutionComplaintHandlingProductColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	solutionComplaintHandlingProductPrimaryKeyColumns     = []string{"id"}
	solutionComplaintHandlingProductGeneratedColumns      = []string{}
)

type (
	// SolutionComplaintHandlingProductSlice is an alias for a slice of pointers to SolutionComplaintHandlingProduct.
	// This should almost always be used instead of []SolutionComplaintHandlingProduct.
	SolutionComplaintHandlingProductSlice []*SolutionComplaintHandlingProduct
	// SolutionComplaintHandlingProductHook is the signature for custom SolutionComplaintHandlingProduct hook methods
	SolutionComplaintHandlingProductHook func(context.Context, boil.ContextExecutor, *SolutionComplaintHandlingProduct) error

	solutionComplaintHandlingProductQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	solutionComplaintHandlingProductType                 = reflect.TypeOf(&SolutionComplaintHandlingProduct{})
	solutionComplaintHandlingProductMapping              = queries.MakeStructMapping(solutionComplaintHandlingProductType)
	solutionComplaintHandlingProductPrimaryKeyMapping, _ = queries.BindMapping(solutionComplaintHandlingProductType, solutionComplaintHandlingProductMapping, solutionComplaintHandlingProductPrimaryKeyColumns)
	solutionComplaintHandlingProductInsertCacheMut       sync.RWMutex
	solutionComplaintHandlingProductInsertCache          = make(map[string]insertCache)
	solutionComplaintHandlingProductUpdateCacheMut       sync.RWMutex
	solutionComplaintHandlingProductUpdateCache          = make(map[string]updateCache)
	solutionComplaintHandlingProductUpsertCacheMut       sync.RWMutex
	solutionComplaintHandlingProductUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var solutionComplaintHandlingProductAfterSelectHooks []SolutionComplaintHandlingProductHook

var solutionComplaintHandlingProductBeforeInsertHooks []SolutionComplaintHandlingProductHook
var solutionComplaintHandlingProductAfterInsertHooks []SolutionComplaintHandlingProductHook

var solutionComplaintHandlingProductBeforeUpdateHooks []SolutionComplaintHandlingProductHook
var solutionComplaintHandlingProductAfterUpdateHooks []SolutionComplaintHandlingProductHook

var solutionComplaintHandlingProductBeforeDeleteHooks []SolutionComplaintHandlingProductHook
var solutionComplaintHandlingProductAfterDeleteHooks []SolutionComplaintHandlingProductHook

var solutionComplaintHandlingProductBeforeUpsertHooks []SolutionComplaintHandlingProductHook
var solutionComplaintHandlingProductAfterUpsertHooks []SolutionComplaintHandlingProductHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SolutionComplaintHandlingProduct) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SolutionComplaintHandlingProduct) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SolutionComplaintHandlingProduct) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SolutionComplaintHandlingProduct) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SolutionComplaintHandlingProduct) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SolutionComplaintHandlingProduct) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SolutionComplaintHandlingProduct) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SolutionComplaintHandlingProduct) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SolutionComplaintHandlingProduct) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionComplaintHandlingProductAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSolutionComplaintHandlingProductHook registers your hook function for all future operations.
func AddSolutionComplaintHandlingProductHook(hookPoint boil.HookPoint, solutionComplaintHandlingProductHook SolutionComplaintHandlingProductHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		solutionComplaintHandlingProductAfterSelectHooks = append(solutionComplaintHandlingProductAfterSelectHooks, solutionComplaintHandlingProductHook)
	case boil.BeforeInsertHook:
		solutionComplaintHandlingProductBeforeInsertHooks = append(solutionComplaintHandlingProductBeforeInsertHooks, solutionComplaintHandlingProductHook)
	case boil.AfterInsertHook:
		solutionComplaintHandlingProductAfterInsertHooks = append(solutionComplaintHandlingProductAfterInsertHooks, solutionComplaintHandlingProductHook)
	case boil.BeforeUpdateHook:
		solutionComplaintHandlingProductBeforeUpdateHooks = append(solutionComplaintHandlingProductBeforeUpdateHooks, solutionComplaintHandlingProductHook)
	case boil.AfterUpdateHook:
		solutionComplaintHandlingProductAfterUpdateHooks = append(solutionComplaintHandlingProductAfterUpdateHooks, solutionComplaintHandlingProductHook)
	case boil.BeforeDeleteHook:
		solutionComplaintHandlingProductBeforeDeleteHooks = append(solutionComplaintHandlingProductBeforeDeleteHooks, solutionComplaintHandlingProductHook)
	case boil.AfterDeleteHook:
		solutionComplaintHandlingProductAfterDeleteHooks = append(solutionComplaintHandlingProductAfterDeleteHooks, solutionComplaintHandlingProductHook)
	case boil.BeforeUpsertHook:
		solutionComplaintHandlingProductBeforeUpsertHooks = append(solutionComplaintHandlingProductBeforeUpsertHooks, solutionComplaintHandlingProductHook)
	case boil.AfterUpsertHook:
		solutionComplaintHandlingProductAfterUpsertHooks = append(solutionComplaintHandlingProductAfterUpsertHooks, solutionComplaintHandlingProductHook)
	}
}

// One returns a single solutionComplaintHandlingProduct record from the query.
func (q solutionComplaintHandlingProductQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SolutionComplaintHandlingProduct, error) {
	o := &SolutionComplaintHandlingProduct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for solution_complaint_handling_products")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SolutionComplaintHandlingProduct records from the query.
func (q solutionComplaintHandlingProductQuery) All(ctx context.Context, exec boil.ContextExecutor) (SolutionComplaintHandlingProductSlice, error) {
	var o []*SolutionComplaintHandlingProduct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to SolutionComplaintHandlingProduct slice")
	}

	if len(solutionComplaintHandlingProductAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SolutionComplaintHandlingProduct records in the query.
func (q solutionComplaintHandlingProductQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count solution_complaint_handling_products rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q solutionComplaintHandlingProductQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if solution_complaint_handling_products exists")
	}

	return count > 0, nil
}

// SolutionProduct pointed to by the foreign key.
func (o *SolutionComplaintHandlingProduct) SolutionProduct(mods ...qm.QueryMod) solutionProductQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SolutionProductID),
	}

	queryMods = append(queryMods, mods...)

	return SolutionProducts(queryMods...)
}

// LoadSolutionProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (solutionComplaintHandlingProductL) LoadSolutionProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSolutionComplaintHandlingProduct interface{}, mods queries.Applicator) error {
	var slice []*SolutionComplaintHandlingProduct
	var object *SolutionComplaintHandlingProduct

	if singular {
		var ok bool
		object, ok = maybeSolutionComplaintHandlingProduct.(*SolutionComplaintHandlingProduct)
		if !ok {
			object = new(SolutionComplaintHandlingProduct)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSolutionComplaintHandlingProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSolutionComplaintHandlingProduct))
			}
		}
	} else {
		s, ok := maybeSolutionComplaintHandlingProduct.(*[]*SolutionComplaintHandlingProduct)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSolutionComplaintHandlingProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSolutionComplaintHandlingProduct))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &solutionComplaintHandlingProductR{}
		}
		if !queries.IsNil(object.SolutionProductID) {
			args = append(args, object.SolutionProductID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &solutionComplaintHandlingProductR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.SolutionProductID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.SolutionProductID) {
				args = append(args, obj.SolutionProductID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`solution_products`),
		qm.WhereIn(`solution_products.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load SolutionProduct")
	}

	var resultSlice []*SolutionProduct
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice SolutionProduct")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for solution_products")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for solution_products")
	}

	if len(solutionProductAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.SolutionProduct = foreign
		if foreign.R == nil {
			foreign.R = &solutionProductR{}
		}
		foreign.R.SolutionComplaintHandlingProduct = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SolutionProductID, foreign.ID) {
				local.R.SolutionProduct = foreign
				if foreign.R == nil {
					foreign.R = &solutionProductR{}
				}
				foreign.R.SolutionComplaintHandlingProduct = local
				break
			}
		}
	}

	return nil
}

// SetSolutionProduct of the solutionComplaintHandlingProduct to the related item.
// Sets o.R.SolutionProduct to related.
// Adds o to related.R.SolutionComplaintHandlingProduct.
func (o *SolutionComplaintHandlingProduct) SetSolutionProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *SolutionProduct) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"solution_complaint_handling_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"solution_product_id"}),
		strmangle.WhereClause("\"", "\"", 2, solutionComplaintHandlingProductPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.SolutionProductID, related.ID)
	if o.R == nil {
		o.R = &solutionComplaintHandlingProductR{
			SolutionProduct: related,
		}
	} else {
		o.R.SolutionProduct = related
	}

	if related.R == nil {
		related.R = &solutionProductR{
			SolutionComplaintHandlingProduct: o,
		}
	} else {
		related.R.SolutionComplaintHandlingProduct = o
	}

	return nil
}

// SolutionComplaintHandlingProducts retrieves all the records using an executor.
func SolutionComplaintHandlingProducts(mods ...qm.QueryMod) solutionComplaintHandlingProductQuery {
	mods = append(mods, qm.From("\"solution_complaint_handling_products\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"solution_complaint_handling_products\".*"})
	}

	return solutionComplaintHandlingProductQuery{q}
}

// FindSolutionComplaintHandlingProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSolutionComplaintHandlingProduct(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*SolutionComplaintHandlingProduct, error) {
	solutionComplaintHandlingProductObj := &SolutionComplaintHandlingProduct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"solution_complaint_handling_products\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, solutionComplaintHandlingProductObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from solution_complaint_handling_products")
	}

	if err = solutionComplaintHandlingProductObj.doAfterSelectHooks(ctx, exec); err != nil {
		return solutionComplaintHandlingProductObj, err
	}

	return solutionComplaintHandlingProductObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SolutionComplaintHandlingProduct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no solution_complaint_handling_products provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(solutionComplaintHandlingProductColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	solutionComplaintHandlingProductInsertCacheMut.RLock()
	cache, cached := solutionComplaintHandlingProductInsertCache[key]
	solutionComplaintHandlingProductInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			solutionComplaintHandlingProductAllColumns,
			solutionComplaintHandlingProductColumnsWithDefault,
			solutionComplaintHandlingProductColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(solutionComplaintHandlingProductType, solutionComplaintHandlingProductMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(solutionComplaintHandlingProductType, solutionComplaintHandlingProductMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"solution_complaint_handling_products\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"solution_complaint_handling_products\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "boiler: unable to insert into solution_complaint_handling_products")
	}

	if !cached {
		solutionComplaintHandlingProductInsertCacheMut.Lock()
		solutionComplaintHandlingProductInsertCache[key] = cache
		solutionComplaintHandlingProductInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SolutionComplaintHandlingProduct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SolutionComplaintHandlingProduct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	solutionComplaintHandlingProductUpdateCacheMut.RLock()
	cache, cached := solutionComplaintHandlingProductUpdateCache[key]
	solutionComplaintHandlingProductUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			solutionComplaintHandlingProductAllColumns,
			solutionComplaintHandlingProductPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update solution_complaint_handling_products, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"solution_complaint_handling_products\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, solutionComplaintHandlingProductPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(solutionComplaintHandlingProductType, solutionComplaintHandlingProductMapping, append(wl, solutionComplaintHandlingProductPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update solution_complaint_handling_products row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for solution_complaint_handling_products")
	}

	if !cached {
		solutionComplaintHandlingProductUpdateCacheMut.Lock()
		solutionComplaintHandlingProductUpdateCache[key] = cache
		solutionComplaintHandlingProductUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q solutionComplaintHandlingProductQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for solution_complaint_handling_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for solution_complaint_handling_products")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SolutionComplaintHandlingProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boiler: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionComplaintHandlingProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"solution_complaint_handling_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, solutionComplaintHandlingProductPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in solutionComplaintHandlingProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all solutionComplaintHandlingProduct")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SolutionComplaintHandlingProduct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no solution_complaint_handling_products provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(solutionComplaintHandlingProductColumnsWithDefault, o)

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

	solutionComplaintHandlingProductUpsertCacheMut.RLock()
	cache, cached := solutionComplaintHandlingProductUpsertCache[key]
	solutionComplaintHandlingProductUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			solutionComplaintHandlingProductAllColumns,
			solutionComplaintHandlingProductColumnsWithDefault,
			solutionComplaintHandlingProductColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			solutionComplaintHandlingProductAllColumns,
			solutionComplaintHandlingProductPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert solution_complaint_handling_products, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(solutionComplaintHandlingProductPrimaryKeyColumns))
			copy(conflict, solutionComplaintHandlingProductPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"solution_complaint_handling_products\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(solutionComplaintHandlingProductType, solutionComplaintHandlingProductMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(solutionComplaintHandlingProductType, solutionComplaintHandlingProductMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert solution_complaint_handling_products")
	}

	if !cached {
		solutionComplaintHandlingProductUpsertCacheMut.Lock()
		solutionComplaintHandlingProductUpsertCache[key] = cache
		solutionComplaintHandlingProductUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SolutionComplaintHandlingProduct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SolutionComplaintHandlingProduct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no SolutionComplaintHandlingProduct provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), solutionComplaintHandlingProductPrimaryKeyMapping)
	sql := "DELETE FROM \"solution_complaint_handling_products\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from solution_complaint_handling_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for solution_complaint_handling_products")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q solutionComplaintHandlingProductQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no solutionComplaintHandlingProductQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from solution_complaint_handling_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for solution_complaint_handling_products")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SolutionComplaintHandlingProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(solutionComplaintHandlingProductBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionComplaintHandlingProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"solution_complaint_handling_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, solutionComplaintHandlingProductPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from solutionComplaintHandlingProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for solution_complaint_handling_products")
	}

	if len(solutionComplaintHandlingProductAfterDeleteHooks) != 0 {
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
func (o *SolutionComplaintHandlingProduct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSolutionComplaintHandlingProduct(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SolutionComplaintHandlingProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SolutionComplaintHandlingProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionComplaintHandlingProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"solution_complaint_handling_products\".* FROM \"solution_complaint_handling_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, solutionComplaintHandlingProductPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in SolutionComplaintHandlingProductSlice")
	}

	*o = slice

	return nil
}

// SolutionComplaintHandlingProductExists checks if the SolutionComplaintHandlingProduct row exists.
func SolutionComplaintHandlingProductExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"solution_complaint_handling_products\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if solution_complaint_handling_products exists")
	}

	return exists, nil
}

// Exists checks if the SolutionComplaintHandlingProduct row exists.
func (o *SolutionComplaintHandlingProduct) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SolutionComplaintHandlingProductExists(ctx, exec, o.ID)
}