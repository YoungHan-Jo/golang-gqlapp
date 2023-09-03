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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// SolutionOcrProduct is an object representing the database table.
type SolutionOcrProduct struct { // OCRプロダクトID
	ID []byte `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ソリューションプロダクトID
	SolutionProductID []byte `boil:"solution_product_id" json:"solution_product_id" toml:"solution_product_id" yaml:"solution_product_id"`
	// 手書き
	Handwriting null.Bool `boil:"handwriting" json:"handwriting,omitempty" toml:"handwriting" yaml:"handwriting,omitempty"`
	// 請求書
	Bill null.Bool `boil:"bill" json:"bill,omitempty" toml:"bill" yaml:"bill,omitempty"`
	// 明細項目
	DetailedItem null.Bool `boil:"detailed_item" json:"detailed_item,omitempty" toml:"detailed_item" yaml:"detailed_item,omitempty"`
	// データ作成日
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// データ更新日
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *solutionOcrProductR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L solutionOcrProductL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SolutionOcrProductColumns = struct {
	ID                string
	SolutionProductID string
	Handwriting       string
	Bill              string
	DetailedItem      string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	SolutionProductID: "solution_product_id",
	Handwriting:       "handwriting",
	Bill:              "bill",
	DetailedItem:      "detailed_item",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

var SolutionOcrProductTableColumns = struct {
	ID                string
	SolutionProductID string
	Handwriting       string
	Bill              string
	DetailedItem      string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "solution_ocr_products.id",
	SolutionProductID: "solution_ocr_products.solution_product_id",
	Handwriting:       "solution_ocr_products.handwriting",
	Bill:              "solution_ocr_products.bill",
	DetailedItem:      "solution_ocr_products.detailed_item",
	CreatedAt:         "solution_ocr_products.created_at",
	UpdatedAt:         "solution_ocr_products.updated_at",
}

// Generated where

var SolutionOcrProductWhere = struct {
	ID                whereHelper__byte
	SolutionProductID whereHelper__byte
	Handwriting       whereHelpernull_Bool
	Bill              whereHelpernull_Bool
	DetailedItem      whereHelpernull_Bool
	CreatedAt         whereHelpertime_Time
	UpdatedAt         whereHelpertime_Time
}{
	ID:                whereHelper__byte{field: "\"solution_ocr_products\".\"id\""},
	SolutionProductID: whereHelper__byte{field: "\"solution_ocr_products\".\"solution_product_id\""},
	Handwriting:       whereHelpernull_Bool{field: "\"solution_ocr_products\".\"handwriting\""},
	Bill:              whereHelpernull_Bool{field: "\"solution_ocr_products\".\"bill\""},
	DetailedItem:      whereHelpernull_Bool{field: "\"solution_ocr_products\".\"detailed_item\""},
	CreatedAt:         whereHelpertime_Time{field: "\"solution_ocr_products\".\"created_at\""},
	UpdatedAt:         whereHelpertime_Time{field: "\"solution_ocr_products\".\"updated_at\""},
}

// SolutionOcrProductRels is where relationship names are stored.
var SolutionOcrProductRels = struct {
	SolutionProduct string
}{
	SolutionProduct: "SolutionProduct",
}

// solutionOcrProductR is where relationships are stored.
type solutionOcrProductR struct {
	SolutionProduct *SolutionProduct `boil:"SolutionProduct" json:"SolutionProduct" toml:"SolutionProduct" yaml:"SolutionProduct"`
}

// NewStruct creates a new relationship struct
func (*solutionOcrProductR) NewStruct() *solutionOcrProductR {
	return &solutionOcrProductR{}
}

func (r *solutionOcrProductR) GetSolutionProduct() *SolutionProduct {
	if r == nil {
		return nil
	}
	return r.SolutionProduct
}

// solutionOcrProductL is where Load methods for each relationship are stored.
type solutionOcrProductL struct{}

var (
	solutionOcrProductAllColumns            = []string{"id", "solution_product_id", "handwriting", "bill", "detailed_item", "created_at", "updated_at"}
	solutionOcrProductColumnsWithoutDefault = []string{"solution_product_id"}
	solutionOcrProductColumnsWithDefault    = []string{"id", "handwriting", "bill", "detailed_item", "created_at", "updated_at"}
	solutionOcrProductPrimaryKeyColumns     = []string{"id"}
	solutionOcrProductGeneratedColumns      = []string{}
)

type (
	// SolutionOcrProductSlice is an alias for a slice of pointers to SolutionOcrProduct.
	// This should almost always be used instead of []SolutionOcrProduct.
	SolutionOcrProductSlice []*SolutionOcrProduct
	// SolutionOcrProductHook is the signature for custom SolutionOcrProduct hook methods
	SolutionOcrProductHook func(context.Context, boil.ContextExecutor, *SolutionOcrProduct) error

	solutionOcrProductQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	solutionOcrProductType                 = reflect.TypeOf(&SolutionOcrProduct{})
	solutionOcrProductMapping              = queries.MakeStructMapping(solutionOcrProductType)
	solutionOcrProductPrimaryKeyMapping, _ = queries.BindMapping(solutionOcrProductType, solutionOcrProductMapping, solutionOcrProductPrimaryKeyColumns)
	solutionOcrProductInsertCacheMut       sync.RWMutex
	solutionOcrProductInsertCache          = make(map[string]insertCache)
	solutionOcrProductUpdateCacheMut       sync.RWMutex
	solutionOcrProductUpdateCache          = make(map[string]updateCache)
	solutionOcrProductUpsertCacheMut       sync.RWMutex
	solutionOcrProductUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var solutionOcrProductAfterSelectHooks []SolutionOcrProductHook

var solutionOcrProductBeforeInsertHooks []SolutionOcrProductHook
var solutionOcrProductAfterInsertHooks []SolutionOcrProductHook

var solutionOcrProductBeforeUpdateHooks []SolutionOcrProductHook
var solutionOcrProductAfterUpdateHooks []SolutionOcrProductHook

var solutionOcrProductBeforeDeleteHooks []SolutionOcrProductHook
var solutionOcrProductAfterDeleteHooks []SolutionOcrProductHook

var solutionOcrProductBeforeUpsertHooks []SolutionOcrProductHook
var solutionOcrProductAfterUpsertHooks []SolutionOcrProductHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SolutionOcrProduct) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SolutionOcrProduct) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SolutionOcrProduct) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SolutionOcrProduct) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SolutionOcrProduct) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SolutionOcrProduct) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SolutionOcrProduct) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SolutionOcrProduct) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SolutionOcrProduct) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionOcrProductAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSolutionOcrProductHook registers your hook function for all future operations.
func AddSolutionOcrProductHook(hookPoint boil.HookPoint, solutionOcrProductHook SolutionOcrProductHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		solutionOcrProductAfterSelectHooks = append(solutionOcrProductAfterSelectHooks, solutionOcrProductHook)
	case boil.BeforeInsertHook:
		solutionOcrProductBeforeInsertHooks = append(solutionOcrProductBeforeInsertHooks, solutionOcrProductHook)
	case boil.AfterInsertHook:
		solutionOcrProductAfterInsertHooks = append(solutionOcrProductAfterInsertHooks, solutionOcrProductHook)
	case boil.BeforeUpdateHook:
		solutionOcrProductBeforeUpdateHooks = append(solutionOcrProductBeforeUpdateHooks, solutionOcrProductHook)
	case boil.AfterUpdateHook:
		solutionOcrProductAfterUpdateHooks = append(solutionOcrProductAfterUpdateHooks, solutionOcrProductHook)
	case boil.BeforeDeleteHook:
		solutionOcrProductBeforeDeleteHooks = append(solutionOcrProductBeforeDeleteHooks, solutionOcrProductHook)
	case boil.AfterDeleteHook:
		solutionOcrProductAfterDeleteHooks = append(solutionOcrProductAfterDeleteHooks, solutionOcrProductHook)
	case boil.BeforeUpsertHook:
		solutionOcrProductBeforeUpsertHooks = append(solutionOcrProductBeforeUpsertHooks, solutionOcrProductHook)
	case boil.AfterUpsertHook:
		solutionOcrProductAfterUpsertHooks = append(solutionOcrProductAfterUpsertHooks, solutionOcrProductHook)
	}
}

// One returns a single solutionOcrProduct record from the query.
func (q solutionOcrProductQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SolutionOcrProduct, error) {
	o := &SolutionOcrProduct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for solution_ocr_products")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SolutionOcrProduct records from the query.
func (q solutionOcrProductQuery) All(ctx context.Context, exec boil.ContextExecutor) (SolutionOcrProductSlice, error) {
	var o []*SolutionOcrProduct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to SolutionOcrProduct slice")
	}

	if len(solutionOcrProductAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SolutionOcrProduct records in the query.
func (q solutionOcrProductQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count solution_ocr_products rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q solutionOcrProductQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if solution_ocr_products exists")
	}

	return count > 0, nil
}

// SolutionProduct pointed to by the foreign key.
func (o *SolutionOcrProduct) SolutionProduct(mods ...qm.QueryMod) solutionProductQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SolutionProductID),
	}

	queryMods = append(queryMods, mods...)

	return SolutionProducts(queryMods...)
}

// LoadSolutionProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (solutionOcrProductL) LoadSolutionProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSolutionOcrProduct interface{}, mods queries.Applicator) error {
	var slice []*SolutionOcrProduct
	var object *SolutionOcrProduct

	if singular {
		var ok bool
		object, ok = maybeSolutionOcrProduct.(*SolutionOcrProduct)
		if !ok {
			object = new(SolutionOcrProduct)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSolutionOcrProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSolutionOcrProduct))
			}
		}
	} else {
		s, ok := maybeSolutionOcrProduct.(*[]*SolutionOcrProduct)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSolutionOcrProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSolutionOcrProduct))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &solutionOcrProductR{}
		}
		if !queries.IsNil(object.SolutionProductID) {
			args = append(args, object.SolutionProductID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &solutionOcrProductR{}
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
		foreign.R.SolutionOcrProduct = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SolutionProductID, foreign.ID) {
				local.R.SolutionProduct = foreign
				if foreign.R == nil {
					foreign.R = &solutionProductR{}
				}
				foreign.R.SolutionOcrProduct = local
				break
			}
		}
	}

	return nil
}

// SetSolutionProduct of the solutionOcrProduct to the related item.
// Sets o.R.SolutionProduct to related.
// Adds o to related.R.SolutionOcrProduct.
func (o *SolutionOcrProduct) SetSolutionProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *SolutionProduct) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"solution_ocr_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"solution_product_id"}),
		strmangle.WhereClause("\"", "\"", 2, solutionOcrProductPrimaryKeyColumns),
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
		o.R = &solutionOcrProductR{
			SolutionProduct: related,
		}
	} else {
		o.R.SolutionProduct = related
	}

	if related.R == nil {
		related.R = &solutionProductR{
			SolutionOcrProduct: o,
		}
	} else {
		related.R.SolutionOcrProduct = o
	}

	return nil
}

// SolutionOcrProducts retrieves all the records using an executor.
func SolutionOcrProducts(mods ...qm.QueryMod) solutionOcrProductQuery {
	mods = append(mods, qm.From("\"solution_ocr_products\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"solution_ocr_products\".*"})
	}

	return solutionOcrProductQuery{q}
}

// FindSolutionOcrProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSolutionOcrProduct(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*SolutionOcrProduct, error) {
	solutionOcrProductObj := &SolutionOcrProduct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"solution_ocr_products\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, solutionOcrProductObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from solution_ocr_products")
	}

	if err = solutionOcrProductObj.doAfterSelectHooks(ctx, exec); err != nil {
		return solutionOcrProductObj, err
	}

	return solutionOcrProductObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SolutionOcrProduct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no solution_ocr_products provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(solutionOcrProductColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	solutionOcrProductInsertCacheMut.RLock()
	cache, cached := solutionOcrProductInsertCache[key]
	solutionOcrProductInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			solutionOcrProductAllColumns,
			solutionOcrProductColumnsWithDefault,
			solutionOcrProductColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(solutionOcrProductType, solutionOcrProductMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(solutionOcrProductType, solutionOcrProductMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"solution_ocr_products\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"solution_ocr_products\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "boiler: unable to insert into solution_ocr_products")
	}

	if !cached {
		solutionOcrProductInsertCacheMut.Lock()
		solutionOcrProductInsertCache[key] = cache
		solutionOcrProductInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SolutionOcrProduct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SolutionOcrProduct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	solutionOcrProductUpdateCacheMut.RLock()
	cache, cached := solutionOcrProductUpdateCache[key]
	solutionOcrProductUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			solutionOcrProductAllColumns,
			solutionOcrProductPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update solution_ocr_products, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"solution_ocr_products\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, solutionOcrProductPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(solutionOcrProductType, solutionOcrProductMapping, append(wl, solutionOcrProductPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update solution_ocr_products row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for solution_ocr_products")
	}

	if !cached {
		solutionOcrProductUpdateCacheMut.Lock()
		solutionOcrProductUpdateCache[key] = cache
		solutionOcrProductUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q solutionOcrProductQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for solution_ocr_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for solution_ocr_products")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SolutionOcrProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionOcrProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"solution_ocr_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, solutionOcrProductPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in solutionOcrProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all solutionOcrProduct")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SolutionOcrProduct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no solution_ocr_products provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(solutionOcrProductColumnsWithDefault, o)

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

	solutionOcrProductUpsertCacheMut.RLock()
	cache, cached := solutionOcrProductUpsertCache[key]
	solutionOcrProductUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			solutionOcrProductAllColumns,
			solutionOcrProductColumnsWithDefault,
			solutionOcrProductColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			solutionOcrProductAllColumns,
			solutionOcrProductPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert solution_ocr_products, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(solutionOcrProductPrimaryKeyColumns))
			copy(conflict, solutionOcrProductPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"solution_ocr_products\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(solutionOcrProductType, solutionOcrProductMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(solutionOcrProductType, solutionOcrProductMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert solution_ocr_products")
	}

	if !cached {
		solutionOcrProductUpsertCacheMut.Lock()
		solutionOcrProductUpsertCache[key] = cache
		solutionOcrProductUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SolutionOcrProduct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SolutionOcrProduct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no SolutionOcrProduct provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), solutionOcrProductPrimaryKeyMapping)
	sql := "DELETE FROM \"solution_ocr_products\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from solution_ocr_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for solution_ocr_products")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q solutionOcrProductQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no solutionOcrProductQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from solution_ocr_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for solution_ocr_products")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SolutionOcrProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(solutionOcrProductBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionOcrProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"solution_ocr_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, solutionOcrProductPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from solutionOcrProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for solution_ocr_products")
	}

	if len(solutionOcrProductAfterDeleteHooks) != 0 {
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
func (o *SolutionOcrProduct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSolutionOcrProduct(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SolutionOcrProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SolutionOcrProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionOcrProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"solution_ocr_products\".* FROM \"solution_ocr_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, solutionOcrProductPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in SolutionOcrProductSlice")
	}

	*o = slice

	return nil
}

// SolutionOcrProductExists checks if the SolutionOcrProduct row exists.
func SolutionOcrProductExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"solution_ocr_products\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if solution_ocr_products exists")
	}

	return exists, nil
}

// Exists checks if the SolutionOcrProduct row exists.
func (o *SolutionOcrProduct) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SolutionOcrProductExists(ctx, exec, o.ID)
}