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

// SolutionCostControlProduct is an object representing the database table.
type SolutionCostControlProduct struct { // コスト管理プロダクトID
	ID []byte `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ソリューションプロダクトID
	SolutionProductID []byte `boil:"solution_product_id" json:"solution_product_id" toml:"solution_product_id" yaml:"solution_product_id"`
	// データ整形拡張
	DataShapingExtension null.Bool `boil:"data_shaping_extension" json:"data_shaping_extension,omitempty" toml:"data_shaping_extension" yaml:"data_shaping_extension,omitempty"`
	// 他システムの連携
	CollaborationWithOtherSystems null.Bool `boil:"collaboration_with_other_systems" json:"collaboration_with_other_systems,omitempty" toml:"collaboration_with_other_systems" yaml:"collaboration_with_other_systems,omitempty"`
	// データ作成日
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// データ更新日
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *solutionCostControlProductR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L solutionCostControlProductL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SolutionCostControlProductColumns = struct {
	ID                            string
	SolutionProductID             string
	DataShapingExtension          string
	CollaborationWithOtherSystems string
	CreatedAt                     string
	UpdatedAt                     string
}{
	ID:                            "id",
	SolutionProductID:             "solution_product_id",
	DataShapingExtension:          "data_shaping_extension",
	CollaborationWithOtherSystems: "collaboration_with_other_systems",
	CreatedAt:                     "created_at",
	UpdatedAt:                     "updated_at",
}

var SolutionCostControlProductTableColumns = struct {
	ID                            string
	SolutionProductID             string
	DataShapingExtension          string
	CollaborationWithOtherSystems string
	CreatedAt                     string
	UpdatedAt                     string
}{
	ID:                            "solution_cost_control_products.id",
	SolutionProductID:             "solution_cost_control_products.solution_product_id",
	DataShapingExtension:          "solution_cost_control_products.data_shaping_extension",
	CollaborationWithOtherSystems: "solution_cost_control_products.collaboration_with_other_systems",
	CreatedAt:                     "solution_cost_control_products.created_at",
	UpdatedAt:                     "solution_cost_control_products.updated_at",
}

// Generated where

var SolutionCostControlProductWhere = struct {
	ID                            whereHelper__byte
	SolutionProductID             whereHelper__byte
	DataShapingExtension          whereHelpernull_Bool
	CollaborationWithOtherSystems whereHelpernull_Bool
	CreatedAt                     whereHelpertime_Time
	UpdatedAt                     whereHelpertime_Time
}{
	ID:                            whereHelper__byte{field: "\"solution_cost_control_products\".\"id\""},
	SolutionProductID:             whereHelper__byte{field: "\"solution_cost_control_products\".\"solution_product_id\""},
	DataShapingExtension:          whereHelpernull_Bool{field: "\"solution_cost_control_products\".\"data_shaping_extension\""},
	CollaborationWithOtherSystems: whereHelpernull_Bool{field: "\"solution_cost_control_products\".\"collaboration_with_other_systems\""},
	CreatedAt:                     whereHelpertime_Time{field: "\"solution_cost_control_products\".\"created_at\""},
	UpdatedAt:                     whereHelpertime_Time{field: "\"solution_cost_control_products\".\"updated_at\""},
}

// SolutionCostControlProductRels is where relationship names are stored.
var SolutionCostControlProductRels = struct {
	SolutionProduct string
}{
	SolutionProduct: "SolutionProduct",
}

// solutionCostControlProductR is where relationships are stored.
type solutionCostControlProductR struct {
	SolutionProduct *SolutionProduct `boil:"SolutionProduct" json:"SolutionProduct" toml:"SolutionProduct" yaml:"SolutionProduct"`
}

// NewStruct creates a new relationship struct
func (*solutionCostControlProductR) NewStruct() *solutionCostControlProductR {
	return &solutionCostControlProductR{}
}

func (r *solutionCostControlProductR) GetSolutionProduct() *SolutionProduct {
	if r == nil {
		return nil
	}
	return r.SolutionProduct
}

// solutionCostControlProductL is where Load methods for each relationship are stored.
type solutionCostControlProductL struct{}

var (
	solutionCostControlProductAllColumns            = []string{"id", "solution_product_id", "data_shaping_extension", "collaboration_with_other_systems", "created_at", "updated_at"}
	solutionCostControlProductColumnsWithoutDefault = []string{"solution_product_id"}
	solutionCostControlProductColumnsWithDefault    = []string{"id", "data_shaping_extension", "collaboration_with_other_systems", "created_at", "updated_at"}
	solutionCostControlProductPrimaryKeyColumns     = []string{"id"}
	solutionCostControlProductGeneratedColumns      = []string{}
)

type (
	// SolutionCostControlProductSlice is an alias for a slice of pointers to SolutionCostControlProduct.
	// This should almost always be used instead of []SolutionCostControlProduct.
	SolutionCostControlProductSlice []*SolutionCostControlProduct
	// SolutionCostControlProductHook is the signature for custom SolutionCostControlProduct hook methods
	SolutionCostControlProductHook func(context.Context, boil.ContextExecutor, *SolutionCostControlProduct) error

	solutionCostControlProductQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	solutionCostControlProductType                 = reflect.TypeOf(&SolutionCostControlProduct{})
	solutionCostControlProductMapping              = queries.MakeStructMapping(solutionCostControlProductType)
	solutionCostControlProductPrimaryKeyMapping, _ = queries.BindMapping(solutionCostControlProductType, solutionCostControlProductMapping, solutionCostControlProductPrimaryKeyColumns)
	solutionCostControlProductInsertCacheMut       sync.RWMutex
	solutionCostControlProductInsertCache          = make(map[string]insertCache)
	solutionCostControlProductUpdateCacheMut       sync.RWMutex
	solutionCostControlProductUpdateCache          = make(map[string]updateCache)
	solutionCostControlProductUpsertCacheMut       sync.RWMutex
	solutionCostControlProductUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var solutionCostControlProductAfterSelectHooks []SolutionCostControlProductHook

var solutionCostControlProductBeforeInsertHooks []SolutionCostControlProductHook
var solutionCostControlProductAfterInsertHooks []SolutionCostControlProductHook

var solutionCostControlProductBeforeUpdateHooks []SolutionCostControlProductHook
var solutionCostControlProductAfterUpdateHooks []SolutionCostControlProductHook

var solutionCostControlProductBeforeDeleteHooks []SolutionCostControlProductHook
var solutionCostControlProductAfterDeleteHooks []SolutionCostControlProductHook

var solutionCostControlProductBeforeUpsertHooks []SolutionCostControlProductHook
var solutionCostControlProductAfterUpsertHooks []SolutionCostControlProductHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SolutionCostControlProduct) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SolutionCostControlProduct) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SolutionCostControlProduct) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SolutionCostControlProduct) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SolutionCostControlProduct) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SolutionCostControlProduct) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SolutionCostControlProduct) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SolutionCostControlProduct) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SolutionCostControlProduct) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range solutionCostControlProductAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSolutionCostControlProductHook registers your hook function for all future operations.
func AddSolutionCostControlProductHook(hookPoint boil.HookPoint, solutionCostControlProductHook SolutionCostControlProductHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		solutionCostControlProductAfterSelectHooks = append(solutionCostControlProductAfterSelectHooks, solutionCostControlProductHook)
	case boil.BeforeInsertHook:
		solutionCostControlProductBeforeInsertHooks = append(solutionCostControlProductBeforeInsertHooks, solutionCostControlProductHook)
	case boil.AfterInsertHook:
		solutionCostControlProductAfterInsertHooks = append(solutionCostControlProductAfterInsertHooks, solutionCostControlProductHook)
	case boil.BeforeUpdateHook:
		solutionCostControlProductBeforeUpdateHooks = append(solutionCostControlProductBeforeUpdateHooks, solutionCostControlProductHook)
	case boil.AfterUpdateHook:
		solutionCostControlProductAfterUpdateHooks = append(solutionCostControlProductAfterUpdateHooks, solutionCostControlProductHook)
	case boil.BeforeDeleteHook:
		solutionCostControlProductBeforeDeleteHooks = append(solutionCostControlProductBeforeDeleteHooks, solutionCostControlProductHook)
	case boil.AfterDeleteHook:
		solutionCostControlProductAfterDeleteHooks = append(solutionCostControlProductAfterDeleteHooks, solutionCostControlProductHook)
	case boil.BeforeUpsertHook:
		solutionCostControlProductBeforeUpsertHooks = append(solutionCostControlProductBeforeUpsertHooks, solutionCostControlProductHook)
	case boil.AfterUpsertHook:
		solutionCostControlProductAfterUpsertHooks = append(solutionCostControlProductAfterUpsertHooks, solutionCostControlProductHook)
	}
}

// One returns a single solutionCostControlProduct record from the query.
func (q solutionCostControlProductQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SolutionCostControlProduct, error) {
	o := &SolutionCostControlProduct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for solution_cost_control_products")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SolutionCostControlProduct records from the query.
func (q solutionCostControlProductQuery) All(ctx context.Context, exec boil.ContextExecutor) (SolutionCostControlProductSlice, error) {
	var o []*SolutionCostControlProduct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to SolutionCostControlProduct slice")
	}

	if len(solutionCostControlProductAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SolutionCostControlProduct records in the query.
func (q solutionCostControlProductQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count solution_cost_control_products rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q solutionCostControlProductQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if solution_cost_control_products exists")
	}

	return count > 0, nil
}

// SolutionProduct pointed to by the foreign key.
func (o *SolutionCostControlProduct) SolutionProduct(mods ...qm.QueryMod) solutionProductQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SolutionProductID),
	}

	queryMods = append(queryMods, mods...)

	return SolutionProducts(queryMods...)
}

// LoadSolutionProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (solutionCostControlProductL) LoadSolutionProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSolutionCostControlProduct interface{}, mods queries.Applicator) error {
	var slice []*SolutionCostControlProduct
	var object *SolutionCostControlProduct

	if singular {
		var ok bool
		object, ok = maybeSolutionCostControlProduct.(*SolutionCostControlProduct)
		if !ok {
			object = new(SolutionCostControlProduct)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSolutionCostControlProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSolutionCostControlProduct))
			}
		}
	} else {
		s, ok := maybeSolutionCostControlProduct.(*[]*SolutionCostControlProduct)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSolutionCostControlProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSolutionCostControlProduct))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &solutionCostControlProductR{}
		}
		if !queries.IsNil(object.SolutionProductID) {
			args = append(args, object.SolutionProductID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &solutionCostControlProductR{}
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
		foreign.R.SolutionCostControlProduct = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SolutionProductID, foreign.ID) {
				local.R.SolutionProduct = foreign
				if foreign.R == nil {
					foreign.R = &solutionProductR{}
				}
				foreign.R.SolutionCostControlProduct = local
				break
			}
		}
	}

	return nil
}

// SetSolutionProduct of the solutionCostControlProduct to the related item.
// Sets o.R.SolutionProduct to related.
// Adds o to related.R.SolutionCostControlProduct.
func (o *SolutionCostControlProduct) SetSolutionProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *SolutionProduct) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"solution_cost_control_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"solution_product_id"}),
		strmangle.WhereClause("\"", "\"", 2, solutionCostControlProductPrimaryKeyColumns),
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
		o.R = &solutionCostControlProductR{
			SolutionProduct: related,
		}
	} else {
		o.R.SolutionProduct = related
	}

	if related.R == nil {
		related.R = &solutionProductR{
			SolutionCostControlProduct: o,
		}
	} else {
		related.R.SolutionCostControlProduct = o
	}

	return nil
}

// SolutionCostControlProducts retrieves all the records using an executor.
func SolutionCostControlProducts(mods ...qm.QueryMod) solutionCostControlProductQuery {
	mods = append(mods, qm.From("\"solution_cost_control_products\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"solution_cost_control_products\".*"})
	}

	return solutionCostControlProductQuery{q}
}

// FindSolutionCostControlProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSolutionCostControlProduct(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*SolutionCostControlProduct, error) {
	solutionCostControlProductObj := &SolutionCostControlProduct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"solution_cost_control_products\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, solutionCostControlProductObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from solution_cost_control_products")
	}

	if err = solutionCostControlProductObj.doAfterSelectHooks(ctx, exec); err != nil {
		return solutionCostControlProductObj, err
	}

	return solutionCostControlProductObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SolutionCostControlProduct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no solution_cost_control_products provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(solutionCostControlProductColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	solutionCostControlProductInsertCacheMut.RLock()
	cache, cached := solutionCostControlProductInsertCache[key]
	solutionCostControlProductInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			solutionCostControlProductAllColumns,
			solutionCostControlProductColumnsWithDefault,
			solutionCostControlProductColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(solutionCostControlProductType, solutionCostControlProductMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(solutionCostControlProductType, solutionCostControlProductMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"solution_cost_control_products\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"solution_cost_control_products\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "boiler: unable to insert into solution_cost_control_products")
	}

	if !cached {
		solutionCostControlProductInsertCacheMut.Lock()
		solutionCostControlProductInsertCache[key] = cache
		solutionCostControlProductInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SolutionCostControlProduct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SolutionCostControlProduct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	solutionCostControlProductUpdateCacheMut.RLock()
	cache, cached := solutionCostControlProductUpdateCache[key]
	solutionCostControlProductUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			solutionCostControlProductAllColumns,
			solutionCostControlProductPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update solution_cost_control_products, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"solution_cost_control_products\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, solutionCostControlProductPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(solutionCostControlProductType, solutionCostControlProductMapping, append(wl, solutionCostControlProductPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update solution_cost_control_products row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for solution_cost_control_products")
	}

	if !cached {
		solutionCostControlProductUpdateCacheMut.Lock()
		solutionCostControlProductUpdateCache[key] = cache
		solutionCostControlProductUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q solutionCostControlProductQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for solution_cost_control_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for solution_cost_control_products")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SolutionCostControlProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionCostControlProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"solution_cost_control_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, solutionCostControlProductPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in solutionCostControlProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all solutionCostControlProduct")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SolutionCostControlProduct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no solution_cost_control_products provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(solutionCostControlProductColumnsWithDefault, o)

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

	solutionCostControlProductUpsertCacheMut.RLock()
	cache, cached := solutionCostControlProductUpsertCache[key]
	solutionCostControlProductUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			solutionCostControlProductAllColumns,
			solutionCostControlProductColumnsWithDefault,
			solutionCostControlProductColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			solutionCostControlProductAllColumns,
			solutionCostControlProductPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert solution_cost_control_products, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(solutionCostControlProductPrimaryKeyColumns))
			copy(conflict, solutionCostControlProductPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"solution_cost_control_products\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(solutionCostControlProductType, solutionCostControlProductMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(solutionCostControlProductType, solutionCostControlProductMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert solution_cost_control_products")
	}

	if !cached {
		solutionCostControlProductUpsertCacheMut.Lock()
		solutionCostControlProductUpsertCache[key] = cache
		solutionCostControlProductUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SolutionCostControlProduct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SolutionCostControlProduct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no SolutionCostControlProduct provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), solutionCostControlProductPrimaryKeyMapping)
	sql := "DELETE FROM \"solution_cost_control_products\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from solution_cost_control_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for solution_cost_control_products")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q solutionCostControlProductQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no solutionCostControlProductQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from solution_cost_control_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for solution_cost_control_products")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SolutionCostControlProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(solutionCostControlProductBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionCostControlProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"solution_cost_control_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, solutionCostControlProductPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from solutionCostControlProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for solution_cost_control_products")
	}

	if len(solutionCostControlProductAfterDeleteHooks) != 0 {
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
func (o *SolutionCostControlProduct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSolutionCostControlProduct(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SolutionCostControlProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SolutionCostControlProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), solutionCostControlProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"solution_cost_control_products\".* FROM \"solution_cost_control_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, solutionCostControlProductPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in SolutionCostControlProductSlice")
	}

	*o = slice

	return nil
}

// SolutionCostControlProductExists checks if the SolutionCostControlProduct row exists.
func SolutionCostControlProductExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"solution_cost_control_products\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if solution_cost_control_products exists")
	}

	return exists, nil
}

// Exists checks if the SolutionCostControlProduct row exists.
func (o *SolutionCostControlProduct) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SolutionCostControlProductExists(ctx, exec, o.ID)
}
