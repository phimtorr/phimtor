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
	"github.com/volatiletech/strmangle"
)

// TorrentLink is an object representing the database table.
type TorrentLink struct {
	ID              int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	VideoID         int64     `boil:"video_id" json:"video_id" toml:"video_id" yaml:"video_id"`
	Resolution      int       `boil:"resolution" json:"resolution" toml:"resolution" yaml:"resolution"`
	Type            string    `boil:"type" json:"type" toml:"type" yaml:"type"`
	Codec           string    `boil:"codec" json:"codec" toml:"codec" yaml:"codec"`
	Source          string    `boil:"source" json:"source" toml:"source" yaml:"source"`
	Link            string    `boil:"link" json:"link" toml:"link" yaml:"link"`
	FileIndex       int       `boil:"file_index" json:"file_index" toml:"file_index" yaml:"file_index"`
	Priority        int       `boil:"priority" json:"priority" toml:"priority" yaml:"priority"`
	RequiredPremium bool      `boil:"required_premium" json:"required_premium" toml:"required_premium" yaml:"required_premium"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *torrentLinkR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L torrentLinkL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TorrentLinkColumns = struct {
	ID              string
	VideoID         string
	Resolution      string
	Type            string
	Codec           string
	Source          string
	Link            string
	FileIndex       string
	Priority        string
	RequiredPremium string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	VideoID:         "video_id",
	Resolution:      "resolution",
	Type:            "type",
	Codec:           "codec",
	Source:          "source",
	Link:            "link",
	FileIndex:       "file_index",
	Priority:        "priority",
	RequiredPremium: "required_premium",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

var TorrentLinkTableColumns = struct {
	ID              string
	VideoID         string
	Resolution      string
	Type            string
	Codec           string
	Source          string
	Link            string
	FileIndex       string
	Priority        string
	RequiredPremium string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "torrent_links.id",
	VideoID:         "torrent_links.video_id",
	Resolution:      "torrent_links.resolution",
	Type:            "torrent_links.type",
	Codec:           "torrent_links.codec",
	Source:          "torrent_links.source",
	Link:            "torrent_links.link",
	FileIndex:       "torrent_links.file_index",
	Priority:        "torrent_links.priority",
	RequiredPremium: "torrent_links.required_premium",
	CreatedAt:       "torrent_links.created_at",
	UpdatedAt:       "torrent_links.updated_at",
}

// Generated where

var TorrentLinkWhere = struct {
	ID              whereHelperint64
	VideoID         whereHelperint64
	Resolution      whereHelperint
	Type            whereHelperstring
	Codec           whereHelperstring
	Source          whereHelperstring
	Link            whereHelperstring
	FileIndex       whereHelperint
	Priority        whereHelperint
	RequiredPremium whereHelperbool
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
}{
	ID:              whereHelperint64{field: "`torrent_links`.`id`"},
	VideoID:         whereHelperint64{field: "`torrent_links`.`video_id`"},
	Resolution:      whereHelperint{field: "`torrent_links`.`resolution`"},
	Type:            whereHelperstring{field: "`torrent_links`.`type`"},
	Codec:           whereHelperstring{field: "`torrent_links`.`codec`"},
	Source:          whereHelperstring{field: "`torrent_links`.`source`"},
	Link:            whereHelperstring{field: "`torrent_links`.`link`"},
	FileIndex:       whereHelperint{field: "`torrent_links`.`file_index`"},
	Priority:        whereHelperint{field: "`torrent_links`.`priority`"},
	RequiredPremium: whereHelperbool{field: "`torrent_links`.`required_premium`"},
	CreatedAt:       whereHelpertime_Time{field: "`torrent_links`.`created_at`"},
	UpdatedAt:       whereHelpertime_Time{field: "`torrent_links`.`updated_at`"},
}

// TorrentLinkRels is where relationship names are stored.
var TorrentLinkRels = struct {
	Video string
}{
	Video: "Video",
}

// torrentLinkR is where relationships are stored.
type torrentLinkR struct {
	Video *Video `boil:"Video" json:"Video" toml:"Video" yaml:"Video"`
}

// NewStruct creates a new relationship struct
func (*torrentLinkR) NewStruct() *torrentLinkR {
	return &torrentLinkR{}
}

func (r *torrentLinkR) GetVideo() *Video {
	if r == nil {
		return nil
	}
	return r.Video
}

// torrentLinkL is where Load methods for each relationship are stored.
type torrentLinkL struct{}

var (
	torrentLinkAllColumns            = []string{"id", "video_id", "resolution", "type", "codec", "source", "link", "file_index", "priority", "required_premium", "created_at", "updated_at"}
	torrentLinkColumnsWithoutDefault = []string{"video_id", "resolution", "type", "codec", "source", "link"}
	torrentLinkColumnsWithDefault    = []string{"id", "file_index", "priority", "required_premium", "created_at", "updated_at"}
	torrentLinkPrimaryKeyColumns     = []string{"id"}
	torrentLinkGeneratedColumns      = []string{}
)

type (
	// TorrentLinkSlice is an alias for a slice of pointers to TorrentLink.
	// This should almost always be used instead of []TorrentLink.
	TorrentLinkSlice []*TorrentLink
	// TorrentLinkHook is the signature for custom TorrentLink hook methods
	TorrentLinkHook func(context.Context, boil.ContextExecutor, *TorrentLink) error

	torrentLinkQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	torrentLinkType                 = reflect.TypeOf(&TorrentLink{})
	torrentLinkMapping              = queries.MakeStructMapping(torrentLinkType)
	torrentLinkPrimaryKeyMapping, _ = queries.BindMapping(torrentLinkType, torrentLinkMapping, torrentLinkPrimaryKeyColumns)
	torrentLinkInsertCacheMut       sync.RWMutex
	torrentLinkInsertCache          = make(map[string]insertCache)
	torrentLinkUpdateCacheMut       sync.RWMutex
	torrentLinkUpdateCache          = make(map[string]updateCache)
	torrentLinkUpsertCacheMut       sync.RWMutex
	torrentLinkUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var torrentLinkAfterSelectMu sync.Mutex
var torrentLinkAfterSelectHooks []TorrentLinkHook

var torrentLinkBeforeInsertMu sync.Mutex
var torrentLinkBeforeInsertHooks []TorrentLinkHook
var torrentLinkAfterInsertMu sync.Mutex
var torrentLinkAfterInsertHooks []TorrentLinkHook

var torrentLinkBeforeUpdateMu sync.Mutex
var torrentLinkBeforeUpdateHooks []TorrentLinkHook
var torrentLinkAfterUpdateMu sync.Mutex
var torrentLinkAfterUpdateHooks []TorrentLinkHook

var torrentLinkBeforeDeleteMu sync.Mutex
var torrentLinkBeforeDeleteHooks []TorrentLinkHook
var torrentLinkAfterDeleteMu sync.Mutex
var torrentLinkAfterDeleteHooks []TorrentLinkHook

var torrentLinkBeforeUpsertMu sync.Mutex
var torrentLinkBeforeUpsertHooks []TorrentLinkHook
var torrentLinkAfterUpsertMu sync.Mutex
var torrentLinkAfterUpsertHooks []TorrentLinkHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TorrentLink) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TorrentLink) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TorrentLink) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TorrentLink) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TorrentLink) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TorrentLink) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TorrentLink) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TorrentLink) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TorrentLink) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentLinkAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTorrentLinkHook registers your hook function for all future operations.
func AddTorrentLinkHook(hookPoint boil.HookPoint, torrentLinkHook TorrentLinkHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		torrentLinkAfterSelectMu.Lock()
		torrentLinkAfterSelectHooks = append(torrentLinkAfterSelectHooks, torrentLinkHook)
		torrentLinkAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		torrentLinkBeforeInsertMu.Lock()
		torrentLinkBeforeInsertHooks = append(torrentLinkBeforeInsertHooks, torrentLinkHook)
		torrentLinkBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		torrentLinkAfterInsertMu.Lock()
		torrentLinkAfterInsertHooks = append(torrentLinkAfterInsertHooks, torrentLinkHook)
		torrentLinkAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		torrentLinkBeforeUpdateMu.Lock()
		torrentLinkBeforeUpdateHooks = append(torrentLinkBeforeUpdateHooks, torrentLinkHook)
		torrentLinkBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		torrentLinkAfterUpdateMu.Lock()
		torrentLinkAfterUpdateHooks = append(torrentLinkAfterUpdateHooks, torrentLinkHook)
		torrentLinkAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		torrentLinkBeforeDeleteMu.Lock()
		torrentLinkBeforeDeleteHooks = append(torrentLinkBeforeDeleteHooks, torrentLinkHook)
		torrentLinkBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		torrentLinkAfterDeleteMu.Lock()
		torrentLinkAfterDeleteHooks = append(torrentLinkAfterDeleteHooks, torrentLinkHook)
		torrentLinkAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		torrentLinkBeforeUpsertMu.Lock()
		torrentLinkBeforeUpsertHooks = append(torrentLinkBeforeUpsertHooks, torrentLinkHook)
		torrentLinkBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		torrentLinkAfterUpsertMu.Lock()
		torrentLinkAfterUpsertHooks = append(torrentLinkAfterUpsertHooks, torrentLinkHook)
		torrentLinkAfterUpsertMu.Unlock()
	}
}

// One returns a single torrentLink record from the query.
func (q torrentLinkQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TorrentLink, error) {
	o := &TorrentLink{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for torrent_links")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TorrentLink records from the query.
func (q torrentLinkQuery) All(ctx context.Context, exec boil.ContextExecutor) (TorrentLinkSlice, error) {
	var o []*TorrentLink

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to TorrentLink slice")
	}

	if len(torrentLinkAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TorrentLink records in the query.
func (q torrentLinkQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count torrent_links rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q torrentLinkQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if torrent_links exists")
	}

	return count > 0, nil
}

// Video pointed to by the foreign key.
func (o *TorrentLink) Video(mods ...qm.QueryMod) videoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.VideoID),
	}

	queryMods = append(queryMods, mods...)

	return Videos(queryMods...)
}

// LoadVideo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (torrentLinkL) LoadVideo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTorrentLink interface{}, mods queries.Applicator) error {
	var slice []*TorrentLink
	var object *TorrentLink

	if singular {
		var ok bool
		object, ok = maybeTorrentLink.(*TorrentLink)
		if !ok {
			object = new(TorrentLink)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTorrentLink)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTorrentLink))
			}
		}
	} else {
		s, ok := maybeTorrentLink.(*[]*TorrentLink)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTorrentLink)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTorrentLink))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &torrentLinkR{}
		}
		args[object.VideoID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &torrentLinkR{}
			}

			args[obj.VideoID] = struct{}{}

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
		qm.From(`videos`),
		qm.WhereIn(`videos.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Video")
	}

	var resultSlice []*Video
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Video")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for videos")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for videos")
	}

	if len(videoAfterSelectHooks) != 0 {
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
		object.R.Video = foreign
		if foreign.R == nil {
			foreign.R = &videoR{}
		}
		foreign.R.TorrentLinks = append(foreign.R.TorrentLinks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.VideoID == foreign.ID {
				local.R.Video = foreign
				if foreign.R == nil {
					foreign.R = &videoR{}
				}
				foreign.R.TorrentLinks = append(foreign.R.TorrentLinks, local)
				break
			}
		}
	}

	return nil
}

// SetVideo of the torrentLink to the related item.
// Sets o.R.Video to related.
// Adds o to related.R.TorrentLinks.
func (o *TorrentLink) SetVideo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Video) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `torrent_links` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"video_id"}),
		strmangle.WhereClause("`", "`", 0, torrentLinkPrimaryKeyColumns),
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

	o.VideoID = related.ID
	if o.R == nil {
		o.R = &torrentLinkR{
			Video: related,
		}
	} else {
		o.R.Video = related
	}

	if related.R == nil {
		related.R = &videoR{
			TorrentLinks: TorrentLinkSlice{o},
		}
	} else {
		related.R.TorrentLinks = append(related.R.TorrentLinks, o)
	}

	return nil
}

// TorrentLinks retrieves all the records using an executor.
func TorrentLinks(mods ...qm.QueryMod) torrentLinkQuery {
	mods = append(mods, qm.From("`torrent_links`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`torrent_links`.*"})
	}

	return torrentLinkQuery{q}
}

// FindTorrentLink retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTorrentLink(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TorrentLink, error) {
	torrentLinkObj := &TorrentLink{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `torrent_links` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, torrentLinkObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from torrent_links")
	}

	if err = torrentLinkObj.doAfterSelectHooks(ctx, exec); err != nil {
		return torrentLinkObj, err
	}

	return torrentLinkObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TorrentLink) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no torrent_links provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(torrentLinkColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	torrentLinkInsertCacheMut.RLock()
	cache, cached := torrentLinkInsertCache[key]
	torrentLinkInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			torrentLinkAllColumns,
			torrentLinkColumnsWithDefault,
			torrentLinkColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(torrentLinkType, torrentLinkMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(torrentLinkType, torrentLinkMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `torrent_links` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `torrent_links` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `torrent_links` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, torrentLinkPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to insert into torrent_links")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == torrentLinkMapping["id"] {
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
		return errors.Wrap(err, "dbmodels: unable to populate default values for torrent_links")
	}

CacheNoHooks:
	if !cached {
		torrentLinkInsertCacheMut.Lock()
		torrentLinkInsertCache[key] = cache
		torrentLinkInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TorrentLink.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TorrentLink) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	torrentLinkUpdateCacheMut.RLock()
	cache, cached := torrentLinkUpdateCache[key]
	torrentLinkUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			torrentLinkAllColumns,
			torrentLinkPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update torrent_links, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `torrent_links` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, torrentLinkPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(torrentLinkType, torrentLinkMapping, append(wl, torrentLinkPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update torrent_links row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for torrent_links")
	}

	if !cached {
		torrentLinkUpdateCacheMut.Lock()
		torrentLinkUpdateCache[key] = cache
		torrentLinkUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q torrentLinkQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for torrent_links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for torrent_links")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TorrentLinkSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), torrentLinkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `torrent_links` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, torrentLinkPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in torrentLink slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all torrentLink")
	}
	return rowsAff, nil
}

var mySQLTorrentLinkUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TorrentLink) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no torrent_links provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(torrentLinkColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTorrentLinkUniqueColumns, o)

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

	torrentLinkUpsertCacheMut.RLock()
	cache, cached := torrentLinkUpsertCache[key]
	torrentLinkUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			torrentLinkAllColumns,
			torrentLinkColumnsWithDefault,
			torrentLinkColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			torrentLinkAllColumns,
			torrentLinkPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert torrent_links, could not build update column list")
		}

		ret := strmangle.SetComplement(torrentLinkAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`torrent_links`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `torrent_links` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(torrentLinkType, torrentLinkMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(torrentLinkType, torrentLinkMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert for torrent_links")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == torrentLinkMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(torrentLinkType, torrentLinkMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to retrieve unique values for torrent_links")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to populate default values for torrent_links")
	}

CacheNoHooks:
	if !cached {
		torrentLinkUpsertCacheMut.Lock()
		torrentLinkUpsertCache[key] = cache
		torrentLinkUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TorrentLink record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TorrentLink) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no TorrentLink provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), torrentLinkPrimaryKeyMapping)
	sql := "DELETE FROM `torrent_links` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from torrent_links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for torrent_links")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q torrentLinkQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no torrentLinkQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from torrent_links")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for torrent_links")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TorrentLinkSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(torrentLinkBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), torrentLinkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `torrent_links` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, torrentLinkPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from torrentLink slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for torrent_links")
	}

	if len(torrentLinkAfterDeleteHooks) != 0 {
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
func (o *TorrentLink) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTorrentLink(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TorrentLinkSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TorrentLinkSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), torrentLinkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `torrent_links`.* FROM `torrent_links` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, torrentLinkPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in TorrentLinkSlice")
	}

	*o = slice

	return nil
}

// TorrentLinkExists checks if the TorrentLink row exists.
func TorrentLinkExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `torrent_links` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if torrent_links exists")
	}

	return exists, nil
}

// Exists checks if the TorrentLink row exists.
func (o *TorrentLink) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TorrentLinkExists(ctx, exec, o.ID)
}
