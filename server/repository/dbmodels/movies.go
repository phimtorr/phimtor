// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Movie is an object representing the database table.
type Movie struct {
	ID            int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ImdbID        string     `boil:"imdb_id" json:"imdb_id" toml:"imdb_id" yaml:"imdb_id"`
	Title         string     `boil:"title" json:"title" toml:"title" yaml:"title"`
	OriginalTitle string     `boil:"original_title" json:"original_title" toml:"original_title" yaml:"original_title"`
	Status        string     `boil:"status" json:"status" toml:"status" yaml:"status"`
	Tagline       string     `boil:"tagline" json:"tagline" toml:"tagline" yaml:"tagline"`
	Genres        types.JSON `boil:"genres" json:"genres" toml:"genres" yaml:"genres"`
	Overview      string     `boil:"overview" json:"overview" toml:"overview" yaml:"overview"`
	PosterPath    string     `boil:"poster_path" json:"poster_path" toml:"poster_path" yaml:"poster_path"`
	BackdropPath  string     `boil:"backdrop_path" json:"backdrop_path" toml:"backdrop_path" yaml:"backdrop_path"`
	ReleaseDate   time.Time  `boil:"release_date" json:"release_date" toml:"release_date" yaml:"release_date"`
	Runtime       int        `boil:"runtime" json:"runtime" toml:"runtime" yaml:"runtime"`
	VoteAverage   float32    `boil:"vote_average" json:"vote_average" toml:"vote_average" yaml:"vote_average"`
	VoteCount     int        `boil:"vote_count" json:"vote_count" toml:"vote_count" yaml:"vote_count"`
	VideoID       int64      `boil:"video_id" json:"video_id" toml:"video_id" yaml:"video_id"`
	CreatedAt     time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     time.Time  `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *movieR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L movieL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MovieColumns = struct {
	ID            string
	ImdbID        string
	Title         string
	OriginalTitle string
	Status        string
	Tagline       string
	Genres        string
	Overview      string
	PosterPath    string
	BackdropPath  string
	ReleaseDate   string
	Runtime       string
	VoteAverage   string
	VoteCount     string
	VideoID       string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "id",
	ImdbID:        "imdb_id",
	Title:         "title",
	OriginalTitle: "original_title",
	Status:        "status",
	Tagline:       "tagline",
	Genres:        "genres",
	Overview:      "overview",
	PosterPath:    "poster_path",
	BackdropPath:  "backdrop_path",
	ReleaseDate:   "release_date",
	Runtime:       "runtime",
	VoteAverage:   "vote_average",
	VoteCount:     "vote_count",
	VideoID:       "video_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

var MovieTableColumns = struct {
	ID            string
	ImdbID        string
	Title         string
	OriginalTitle string
	Status        string
	Tagline       string
	Genres        string
	Overview      string
	PosterPath    string
	BackdropPath  string
	ReleaseDate   string
	Runtime       string
	VoteAverage   string
	VoteCount     string
	VideoID       string
	CreatedAt     string
	UpdatedAt     string
}{
	ID:            "movies.id",
	ImdbID:        "movies.imdb_id",
	Title:         "movies.title",
	OriginalTitle: "movies.original_title",
	Status:        "movies.status",
	Tagline:       "movies.tagline",
	Genres:        "movies.genres",
	Overview:      "movies.overview",
	PosterPath:    "movies.poster_path",
	BackdropPath:  "movies.backdrop_path",
	ReleaseDate:   "movies.release_date",
	Runtime:       "movies.runtime",
	VoteAverage:   "movies.vote_average",
	VoteCount:     "movies.vote_count",
	VideoID:       "movies.video_id",
	CreatedAt:     "movies.created_at",
	UpdatedAt:     "movies.updated_at",
}

// Generated where

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var MovieWhere = struct {
	ID            whereHelperint64
	ImdbID        whereHelperstring
	Title         whereHelperstring
	OriginalTitle whereHelperstring
	Status        whereHelperstring
	Tagline       whereHelperstring
	Genres        whereHelpertypes_JSON
	Overview      whereHelperstring
	PosterPath    whereHelperstring
	BackdropPath  whereHelperstring
	ReleaseDate   whereHelpertime_Time
	Runtime       whereHelperint
	VoteAverage   whereHelperfloat32
	VoteCount     whereHelperint
	VideoID       whereHelperint64
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpertime_Time
}{
	ID:            whereHelperint64{field: "`movies`.`id`"},
	ImdbID:        whereHelperstring{field: "`movies`.`imdb_id`"},
	Title:         whereHelperstring{field: "`movies`.`title`"},
	OriginalTitle: whereHelperstring{field: "`movies`.`original_title`"},
	Status:        whereHelperstring{field: "`movies`.`status`"},
	Tagline:       whereHelperstring{field: "`movies`.`tagline`"},
	Genres:        whereHelpertypes_JSON{field: "`movies`.`genres`"},
	Overview:      whereHelperstring{field: "`movies`.`overview`"},
	PosterPath:    whereHelperstring{field: "`movies`.`poster_path`"},
	BackdropPath:  whereHelperstring{field: "`movies`.`backdrop_path`"},
	ReleaseDate:   whereHelpertime_Time{field: "`movies`.`release_date`"},
	Runtime:       whereHelperint{field: "`movies`.`runtime`"},
	VoteAverage:   whereHelperfloat32{field: "`movies`.`vote_average`"},
	VoteCount:     whereHelperint{field: "`movies`.`vote_count`"},
	VideoID:       whereHelperint64{field: "`movies`.`video_id`"},
	CreatedAt:     whereHelpertime_Time{field: "`movies`.`created_at`"},
	UpdatedAt:     whereHelpertime_Time{field: "`movies`.`updated_at`"},
}

// MovieRels is where relationship names are stored.
var MovieRels = struct {
}{}

// movieR is where relationships are stored.
type movieR struct {
}

// NewStruct creates a new relationship struct
func (*movieR) NewStruct() *movieR {
	return &movieR{}
}

// movieL is where Load methods for each relationship are stored.
type movieL struct{}

var (
	movieAllColumns            = []string{"id", "imdb_id", "title", "original_title", "status", "tagline", "genres", "overview", "poster_path", "backdrop_path", "release_date", "runtime", "vote_average", "vote_count", "video_id", "created_at", "updated_at"}
	movieColumnsWithoutDefault = []string{"id", "imdb_id", "title", "original_title", "status", "tagline", "genres", "overview", "poster_path", "backdrop_path", "release_date", "runtime", "vote_average", "vote_count"}
	movieColumnsWithDefault    = []string{"video_id", "created_at", "updated_at"}
	moviePrimaryKeyColumns     = []string{"id"}
	movieGeneratedColumns      = []string{}
)

type (
	// MovieSlice is an alias for a slice of pointers to Movie.
	// This should almost always be used instead of []Movie.
	MovieSlice []*Movie
	// MovieHook is the signature for custom Movie hook methods
	MovieHook func(context.Context, boil.ContextExecutor, *Movie) error

	movieQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	movieType                 = reflect.TypeOf(&Movie{})
	movieMapping              = queries.MakeStructMapping(movieType)
	moviePrimaryKeyMapping, _ = queries.BindMapping(movieType, movieMapping, moviePrimaryKeyColumns)
	movieInsertCacheMut       sync.RWMutex
	movieInsertCache          = make(map[string]insertCache)
	movieUpdateCacheMut       sync.RWMutex
	movieUpdateCache          = make(map[string]updateCache)
	movieUpsertCacheMut       sync.RWMutex
	movieUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var movieAfterSelectMu sync.Mutex
var movieAfterSelectHooks []MovieHook

var movieBeforeInsertMu sync.Mutex
var movieBeforeInsertHooks []MovieHook
var movieAfterInsertMu sync.Mutex
var movieAfterInsertHooks []MovieHook

var movieBeforeUpdateMu sync.Mutex
var movieBeforeUpdateHooks []MovieHook
var movieAfterUpdateMu sync.Mutex
var movieAfterUpdateHooks []MovieHook

var movieBeforeDeleteMu sync.Mutex
var movieBeforeDeleteHooks []MovieHook
var movieAfterDeleteMu sync.Mutex
var movieAfterDeleteHooks []MovieHook

var movieBeforeUpsertMu sync.Mutex
var movieBeforeUpsertHooks []MovieHook
var movieAfterUpsertMu sync.Mutex
var movieAfterUpsertHooks []MovieHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Movie) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Movie) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Movie) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Movie) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Movie) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Movie) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Movie) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Movie) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Movie) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMovieHook registers your hook function for all future operations.
func AddMovieHook(hookPoint boil.HookPoint, movieHook MovieHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		movieAfterSelectMu.Lock()
		movieAfterSelectHooks = append(movieAfterSelectHooks, movieHook)
		movieAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		movieBeforeInsertMu.Lock()
		movieBeforeInsertHooks = append(movieBeforeInsertHooks, movieHook)
		movieBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		movieAfterInsertMu.Lock()
		movieAfterInsertHooks = append(movieAfterInsertHooks, movieHook)
		movieAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		movieBeforeUpdateMu.Lock()
		movieBeforeUpdateHooks = append(movieBeforeUpdateHooks, movieHook)
		movieBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		movieAfterUpdateMu.Lock()
		movieAfterUpdateHooks = append(movieAfterUpdateHooks, movieHook)
		movieAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		movieBeforeDeleteMu.Lock()
		movieBeforeDeleteHooks = append(movieBeforeDeleteHooks, movieHook)
		movieBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		movieAfterDeleteMu.Lock()
		movieAfterDeleteHooks = append(movieAfterDeleteHooks, movieHook)
		movieAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		movieBeforeUpsertMu.Lock()
		movieBeforeUpsertHooks = append(movieBeforeUpsertHooks, movieHook)
		movieBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		movieAfterUpsertMu.Lock()
		movieAfterUpsertHooks = append(movieAfterUpsertHooks, movieHook)
		movieAfterUpsertMu.Unlock()
	}
}

// One returns a single movie record from the query.
func (q movieQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Movie, error) {
	o := &Movie{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for movies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Movie records from the query.
func (q movieQuery) All(ctx context.Context, exec boil.ContextExecutor) (MovieSlice, error) {
	var o []*Movie

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to Movie slice")
	}

	if len(movieAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Movie records in the query.
func (q movieQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count movies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q movieQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if movies exists")
	}

	return count > 0, nil
}

// Movies retrieves all the records using an executor.
func Movies(mods ...qm.QueryMod) movieQuery {
	mods = append(mods, qm.From("`movies`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`movies`.*"})
	}

	return movieQuery{q}
}

// FindMovie retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMovie(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Movie, error) {
	movieObj := &Movie{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `movies` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, movieObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from movies")
	}

	if err = movieObj.doAfterSelectHooks(ctx, exec); err != nil {
		return movieObj, err
	}

	return movieObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Movie) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no movies provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(movieColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	movieInsertCacheMut.RLock()
	cache, cached := movieInsertCache[key]
	movieInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			movieAllColumns,
			movieColumnsWithDefault,
			movieColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(movieType, movieMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(movieType, movieMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `movies` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `movies` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `movies` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, moviePrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to insert into movies")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for movies")
	}

CacheNoHooks:
	if !cached {
		movieInsertCacheMut.Lock()
		movieInsertCache[key] = cache
		movieInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Movie.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Movie) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	movieUpdateCacheMut.RLock()
	cache, cached := movieUpdateCache[key]
	movieUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			movieAllColumns,
			moviePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update movies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `movies` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, moviePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(movieType, movieMapping, append(wl, moviePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update movies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for movies")
	}

	if !cached {
		movieUpdateCacheMut.Lock()
		movieUpdateCache[key] = cache
		movieUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q movieQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for movies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for movies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MovieSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moviePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `movies` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, moviePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in movie slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all movie")
	}
	return rowsAff, nil
}

var mySQLMovieUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Movie) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no movies provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(movieColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMovieUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	movieUpsertCacheMut.RLock()
	cache, cached := movieUpsertCache[key]
	movieUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			movieAllColumns,
			movieColumnsWithDefault,
			movieColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			movieAllColumns,
			moviePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert movies, could not build update column list")
		}

		ret := strmangle.SetComplement(movieAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`movies`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `movies` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(movieType, movieMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(movieType, movieMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert for movies")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(movieType, movieMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to retrieve unique values for movies")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for movies")
	}

CacheNoHooks:
	if !cached {
		movieUpsertCacheMut.Lock()
		movieUpsertCache[key] = cache
		movieUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Movie record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Movie) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no Movie provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), moviePrimaryKeyMapping)
	sql := "DELETE FROM `movies` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from movies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for movies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q movieQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no movieQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from movies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for movies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MovieSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(movieBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moviePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `movies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, moviePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from movie slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for movies")
	}

	if len(movieAfterDeleteHooks) != 0 {
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
func (o *Movie) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMovie(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MovieSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MovieSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moviePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `movies`.* FROM `movies` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, moviePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in MovieSlice")
	}

	*o = slice

	return nil
}

// MovieExists checks if the Movie row exists.
func MovieExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `movies` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if movies exists")
	}

	return exists, nil
}

// Exists checks if the Movie row exists.
func (o *Movie) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MovieExists(ctx, exec, o.ID)
}