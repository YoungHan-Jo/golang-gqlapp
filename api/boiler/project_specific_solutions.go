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

// ProjectSpecificSolution is an object representing the database table.
type ProjectSpecificSolution struct { // 個別ソリューションID
	ID []byte `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ヒアリングプロジェクトID
	HearingProjectID []byte `boil:"hearing_project_id" json:"hearing_project_id" toml:"hearing_project_id" yaml:"hearing_project_id"`
	// ソリューション名
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`
	// データ作成日
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// データ更新日
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *projectSpecificSolutionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectSpecificSolutionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectSpecificSolutionColumns = struct {
	ID               string
	HearingProjectID string
	Name             string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	HearingProjectID: "hearing_project_id",
	Name:             "name",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

var ProjectSpecificSolutionTableColumns = struct {
	ID               string
	HearingProjectID string
	Name             string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "project_specific_solutions.id",
	HearingProjectID: "project_specific_solutions.hearing_project_id",
	Name:             "project_specific_solutions.name",
	CreatedAt:        "project_specific_solutions.created_at",
	UpdatedAt:        "project_specific_solutions.updated_at",
}

// Generated where

var ProjectSpecificSolutionWhere = struct {
	ID               whereHelper__byte
	HearingProjectID whereHelper__byte
	Name             whereHelperstring
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
}{
	ID:               whereHelper__byte{field: "\"project_specific_solutions\".\"id\""},
	HearingProjectID: whereHelper__byte{field: "\"project_specific_solutions\".\"hearing_project_id\""},
	Name:             whereHelperstring{field: "\"project_specific_solutions\".\"name\""},
	CreatedAt:        whereHelpertime_Time{field: "\"project_specific_solutions\".\"created_at\""},
	UpdatedAt:        whereHelpertime_Time{field: "\"project_specific_solutions\".\"updated_at\""},
}

// ProjectSpecificSolutionRels is where relationship names are stored.
var ProjectSpecificSolutionRels = struct {
	HearingProject string
}{
	HearingProject: "HearingProject",
}

// projectSpecificSolutionR is where relationships are stored.
type projectSpecificSolutionR struct {
	HearingProject *HearingProject `boil:"HearingProject" json:"HearingProject" toml:"HearingProject" yaml:"HearingProject"`
}

// NewStruct creates a new relationship struct
func (*projectSpecificSolutionR) NewStruct() *projectSpecificSolutionR {
	return &projectSpecificSolutionR{}
}

func (r *projectSpecificSolutionR) GetHearingProject() *HearingProject {
	if r == nil {
		return nil
	}
	return r.HearingProject
}

// projectSpecificSolutionL is where Load methods for each relationship are stored.
type projectSpecificSolutionL struct{}

var (
	projectSpecificSolutionAllColumns            = []string{"id", "hearing_project_id", "name", "created_at", "updated_at"}
	projectSpecificSolutionColumnsWithoutDefault = []string{"hearing_project_id", "name"}
	projectSpecificSolutionColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	projectSpecificSolutionPrimaryKeyColumns     = []string{"id"}
	projectSpecificSolutionGeneratedColumns      = []string{}
)

type (
	// ProjectSpecificSolutionSlice is an alias for a slice of pointers to ProjectSpecificSolution.
	// This should almost always be used instead of []ProjectSpecificSolution.
	ProjectSpecificSolutionSlice []*ProjectSpecificSolution
	// ProjectSpecificSolutionHook is the signature for custom ProjectSpecificSolution hook methods
	ProjectSpecificSolutionHook func(context.Context, boil.ContextExecutor, *ProjectSpecificSolution) error

	projectSpecificSolutionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectSpecificSolutionType                 = reflect.TypeOf(&ProjectSpecificSolution{})
	projectSpecificSolutionMapping              = queries.MakeStructMapping(projectSpecificSolutionType)
	projectSpecificSolutionPrimaryKeyMapping, _ = queries.BindMapping(projectSpecificSolutionType, projectSpecificSolutionMapping, projectSpecificSolutionPrimaryKeyColumns)
	projectSpecificSolutionInsertCacheMut       sync.RWMutex
	projectSpecificSolutionInsertCache          = make(map[string]insertCache)
	projectSpecificSolutionUpdateCacheMut       sync.RWMutex
	projectSpecificSolutionUpdateCache          = make(map[string]updateCache)
	projectSpecificSolutionUpsertCacheMut       sync.RWMutex
	projectSpecificSolutionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var projectSpecificSolutionAfterSelectHooks []ProjectSpecificSolutionHook

var projectSpecificSolutionBeforeInsertHooks []ProjectSpecificSolutionHook
var projectSpecificSolutionAfterInsertHooks []ProjectSpecificSolutionHook

var projectSpecificSolutionBeforeUpdateHooks []ProjectSpecificSolutionHook
var projectSpecificSolutionAfterUpdateHooks []ProjectSpecificSolutionHook

var projectSpecificSolutionBeforeDeleteHooks []ProjectSpecificSolutionHook
var projectSpecificSolutionAfterDeleteHooks []ProjectSpecificSolutionHook

var projectSpecificSolutionBeforeUpsertHooks []ProjectSpecificSolutionHook
var projectSpecificSolutionAfterUpsertHooks []ProjectSpecificSolutionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProjectSpecificSolution) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProjectSpecificSolution) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProjectSpecificSolution) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProjectSpecificSolution) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProjectSpecificSolution) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProjectSpecificSolution) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProjectSpecificSolution) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProjectSpecificSolution) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProjectSpecificSolution) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectSpecificSolutionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProjectSpecificSolutionHook registers your hook function for all future operations.
func AddProjectSpecificSolutionHook(hookPoint boil.HookPoint, projectSpecificSolutionHook ProjectSpecificSolutionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		projectSpecificSolutionAfterSelectHooks = append(projectSpecificSolutionAfterSelectHooks, projectSpecificSolutionHook)
	case boil.BeforeInsertHook:
		projectSpecificSolutionBeforeInsertHooks = append(projectSpecificSolutionBeforeInsertHooks, projectSpecificSolutionHook)
	case boil.AfterInsertHook:
		projectSpecificSolutionAfterInsertHooks = append(projectSpecificSolutionAfterInsertHooks, projectSpecificSolutionHook)
	case boil.BeforeUpdateHook:
		projectSpecificSolutionBeforeUpdateHooks = append(projectSpecificSolutionBeforeUpdateHooks, projectSpecificSolutionHook)
	case boil.AfterUpdateHook:
		projectSpecificSolutionAfterUpdateHooks = append(projectSpecificSolutionAfterUpdateHooks, projectSpecificSolutionHook)
	case boil.BeforeDeleteHook:
		projectSpecificSolutionBeforeDeleteHooks = append(projectSpecificSolutionBeforeDeleteHooks, projectSpecificSolutionHook)
	case boil.AfterDeleteHook:
		projectSpecificSolutionAfterDeleteHooks = append(projectSpecificSolutionAfterDeleteHooks, projectSpecificSolutionHook)
	case boil.BeforeUpsertHook:
		projectSpecificSolutionBeforeUpsertHooks = append(projectSpecificSolutionBeforeUpsertHooks, projectSpecificSolutionHook)
	case boil.AfterUpsertHook:
		projectSpecificSolutionAfterUpsertHooks = append(projectSpecificSolutionAfterUpsertHooks, projectSpecificSolutionHook)
	}
}

// One returns a single projectSpecificSolution record from the query.
func (q projectSpecificSolutionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProjectSpecificSolution, error) {
	o := &ProjectSpecificSolution{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for project_specific_solutions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProjectSpecificSolution records from the query.
func (q projectSpecificSolutionQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProjectSpecificSolutionSlice, error) {
	var o []*ProjectSpecificSolution

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to ProjectSpecificSolution slice")
	}

	if len(projectSpecificSolutionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProjectSpecificSolution records in the query.
func (q projectSpecificSolutionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count project_specific_solutions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectSpecificSolutionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if project_specific_solutions exists")
	}

	return count > 0, nil
}

// HearingProject pointed to by the foreign key.
func (o *ProjectSpecificSolution) HearingProject(mods ...qm.QueryMod) hearingProjectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.HearingProjectID),
	}

	queryMods = append(queryMods, mods...)

	return HearingProjects(queryMods...)
}

// LoadHearingProject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectSpecificSolutionL) LoadHearingProject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectSpecificSolution interface{}, mods queries.Applicator) error {
	var slice []*ProjectSpecificSolution
	var object *ProjectSpecificSolution

	if singular {
		var ok bool
		object, ok = maybeProjectSpecificSolution.(*ProjectSpecificSolution)
		if !ok {
			object = new(ProjectSpecificSolution)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeProjectSpecificSolution)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeProjectSpecificSolution))
			}
		}
	} else {
		s, ok := maybeProjectSpecificSolution.(*[]*ProjectSpecificSolution)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeProjectSpecificSolution)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeProjectSpecificSolution))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectSpecificSolutionR{}
		}
		if !queries.IsNil(object.HearingProjectID) {
			args = append(args, object.HearingProjectID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectSpecificSolutionR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.HearingProjectID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.HearingProjectID) {
				args = append(args, obj.HearingProjectID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`hearing_projects`),
		qm.WhereIn(`hearing_projects.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load HearingProject")
	}

	var resultSlice []*HearingProject
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice HearingProject")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for hearing_projects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for hearing_projects")
	}

	if len(hearingProjectAfterSelectHooks) != 0 {
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
		object.R.HearingProject = foreign
		if foreign.R == nil {
			foreign.R = &hearingProjectR{}
		}
		foreign.R.ProjectSpecificSolutions = append(foreign.R.ProjectSpecificSolutions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.HearingProjectID, foreign.ID) {
				local.R.HearingProject = foreign
				if foreign.R == nil {
					foreign.R = &hearingProjectR{}
				}
				foreign.R.ProjectSpecificSolutions = append(foreign.R.ProjectSpecificSolutions, local)
				break
			}
		}
	}

	return nil
}

// SetHearingProject of the projectSpecificSolution to the related item.
// Sets o.R.HearingProject to related.
// Adds o to related.R.ProjectSpecificSolutions.
func (o *ProjectSpecificSolution) SetHearingProject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *HearingProject) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"project_specific_solutions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"hearing_project_id"}),
		strmangle.WhereClause("\"", "\"", 2, projectSpecificSolutionPrimaryKeyColumns),
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

	queries.Assign(&o.HearingProjectID, related.ID)
	if o.R == nil {
		o.R = &projectSpecificSolutionR{
			HearingProject: related,
		}
	} else {
		o.R.HearingProject = related
	}

	if related.R == nil {
		related.R = &hearingProjectR{
			ProjectSpecificSolutions: ProjectSpecificSolutionSlice{o},
		}
	} else {
		related.R.ProjectSpecificSolutions = append(related.R.ProjectSpecificSolutions, o)
	}

	return nil
}

// ProjectSpecificSolutions retrieves all the records using an executor.
func ProjectSpecificSolutions(mods ...qm.QueryMod) projectSpecificSolutionQuery {
	mods = append(mods, qm.From("\"project_specific_solutions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"project_specific_solutions\".*"})
	}

	return projectSpecificSolutionQuery{q}
}

// FindProjectSpecificSolution retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProjectSpecificSolution(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*ProjectSpecificSolution, error) {
	projectSpecificSolutionObj := &ProjectSpecificSolution{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"project_specific_solutions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, projectSpecificSolutionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from project_specific_solutions")
	}

	if err = projectSpecificSolutionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return projectSpecificSolutionObj, err
	}

	return projectSpecificSolutionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProjectSpecificSolution) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no project_specific_solutions provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(projectSpecificSolutionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectSpecificSolutionInsertCacheMut.RLock()
	cache, cached := projectSpecificSolutionInsertCache[key]
	projectSpecificSolutionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectSpecificSolutionAllColumns,
			projectSpecificSolutionColumnsWithDefault,
			projectSpecificSolutionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectSpecificSolutionType, projectSpecificSolutionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectSpecificSolutionType, projectSpecificSolutionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"project_specific_solutions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"project_specific_solutions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "boiler: unable to insert into project_specific_solutions")
	}

	if !cached {
		projectSpecificSolutionInsertCacheMut.Lock()
		projectSpecificSolutionInsertCache[key] = cache
		projectSpecificSolutionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProjectSpecificSolution.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProjectSpecificSolution) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	projectSpecificSolutionUpdateCacheMut.RLock()
	cache, cached := projectSpecificSolutionUpdateCache[key]
	projectSpecificSolutionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectSpecificSolutionAllColumns,
			projectSpecificSolutionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update project_specific_solutions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"project_specific_solutions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, projectSpecificSolutionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectSpecificSolutionType, projectSpecificSolutionMapping, append(wl, projectSpecificSolutionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "boiler: unable to update project_specific_solutions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for project_specific_solutions")
	}

	if !cached {
		projectSpecificSolutionUpdateCacheMut.Lock()
		projectSpecificSolutionUpdateCache[key] = cache
		projectSpecificSolutionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q projectSpecificSolutionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for project_specific_solutions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for project_specific_solutions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectSpecificSolutionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectSpecificSolutionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"project_specific_solutions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, projectSpecificSolutionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in projectSpecificSolution slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all projectSpecificSolution")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProjectSpecificSolution) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no project_specific_solutions provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(projectSpecificSolutionColumnsWithDefault, o)

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

	projectSpecificSolutionUpsertCacheMut.RLock()
	cache, cached := projectSpecificSolutionUpsertCache[key]
	projectSpecificSolutionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectSpecificSolutionAllColumns,
			projectSpecificSolutionColumnsWithDefault,
			projectSpecificSolutionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			projectSpecificSolutionAllColumns,
			projectSpecificSolutionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert project_specific_solutions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(projectSpecificSolutionPrimaryKeyColumns))
			copy(conflict, projectSpecificSolutionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"project_specific_solutions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(projectSpecificSolutionType, projectSpecificSolutionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectSpecificSolutionType, projectSpecificSolutionMapping, ret)
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
		return errors.Wrap(err, "boiler: unable to upsert project_specific_solutions")
	}

	if !cached {
		projectSpecificSolutionUpsertCacheMut.Lock()
		projectSpecificSolutionUpsertCache[key] = cache
		projectSpecificSolutionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProjectSpecificSolution record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProjectSpecificSolution) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no ProjectSpecificSolution provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectSpecificSolutionPrimaryKeyMapping)
	sql := "DELETE FROM \"project_specific_solutions\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from project_specific_solutions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for project_specific_solutions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectSpecificSolutionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no projectSpecificSolutionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from project_specific_solutions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for project_specific_solutions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectSpecificSolutionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(projectSpecificSolutionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectSpecificSolutionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"project_specific_solutions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectSpecificSolutionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from projectSpecificSolution slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for project_specific_solutions")
	}

	if len(projectSpecificSolutionAfterDeleteHooks) != 0 {
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
func (o *ProjectSpecificSolution) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProjectSpecificSolution(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectSpecificSolutionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectSpecificSolutionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectSpecificSolutionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"project_specific_solutions\".* FROM \"project_specific_solutions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectSpecificSolutionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in ProjectSpecificSolutionSlice")
	}

	*o = slice

	return nil
}

// ProjectSpecificSolutionExists checks if the ProjectSpecificSolution row exists.
func ProjectSpecificSolutionExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"project_specific_solutions\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if project_specific_solutions exists")
	}

	return exists, nil
}

// Exists checks if the ProjectSpecificSolution row exists.
func (o *ProjectSpecificSolution) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProjectSpecificSolutionExists(ctx, exec, o.ID)
}