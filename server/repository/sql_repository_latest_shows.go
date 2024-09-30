package repository

import (
	"context"
	"fmt"
	"math"

	tmdb "github.com/cyruzin/golang-tmdb"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/http"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

type sortBy int

const (
	sortByNone = iota
	sortByAirDateDesc
	sortByCreatedAtDesc
)

func (r SQLRepository) ListRecentlyAddedMovies(ctx context.Context, params http.ListRecentlyAddedMoviesParams) ([]http.Show, http.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		dbmodels.LatestShowWhere.Type.EQ(dbmodels.LatestShowsTypeMovie),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, sortByCreatedAtDesc)
}

func (r SQLRepository) ListLatestEpisodes(ctx context.Context, params http.ListLatestEpisodesParams) ([]http.Show, http.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		dbmodels.LatestShowWhere.Type.EQ(dbmodels.LatestShowsTypeEpisode),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, sortByAirDateDesc)
}

func (r SQLRepository) ListLatestMovies(ctx context.Context, params http.ListLatestMoviesParams) ([]http.Show, http.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		dbmodels.LatestShowWhere.Type.EQ(dbmodels.LatestShowsTypeMovie),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, sortByAirDateDesc)
}

func (r SQLRepository) ListLatestTvSeries(ctx context.Context, params http.ListLatestTvSeriesParams) ([]http.Show, http.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		dbmodels.LatestShowWhere.Type.EQ(dbmodels.LatestShowsTypeTVSeries),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, sortByAirDateDesc)
}

func (r SQLRepository) SearchShows(ctx context.Context, params http.SearchShowsParams) ([]http.Show, http.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		qm.Where("MATCH(title, original_title) AGAINST (?)", params.Query),
		dbmodels.LatestShowWhere.Type.IN([]dbmodels.LatestShowsType{
			dbmodels.LatestShowsTypeMovie,
			dbmodels.LatestShowsTypeTVSeries,
		}),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, sortByNone)
}

func (r SQLRepository) queryShows(ctx context.Context, queryMods []qm.QueryMod, page, pageSize int, sort sortBy) ([]http.Show, http.Pagination, error) {
	pagingQueryMods := append(queryMods,
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
	)

	switch sort {
	case sortByAirDateDesc:
		pagingQueryMods = append(pagingQueryMods, qm.OrderBy(dbmodels.LatestShowColumns.AirDate+" DESC"))
	case sortByCreatedAtDesc:
		pagingQueryMods = append(pagingQueryMods, qm.OrderBy(dbmodels.LatestShowColumns.CreatedAt+" DESC"))
	default:
		// Do nothing
	}

	shows, err := dbmodels.LatestShows(pagingQueryMods...).
		All(ctx, r.db)
	if err != nil {
		return nil, http.Pagination{}, fmt.Errorf("get shows: %w", err)
	}

	count, err := dbmodels.LatestShows(queryMods...).Count(ctx, r.db)
	if err != nil {
		return nil, http.Pagination{}, fmt.Errorf("count shows: %w", err)
	}

	return toHTTP2Shows(shows), toHTTP2Pagination(page, pageSize, count), nil
}

func toPageAndPageSize(page, pageSize *int) (rPage, rPageSize int) {
	rPage = 1
	rPageSize = pageSizeDefault

	if page != nil && *page >= 1 {
		rPage = *page
	}

	if pageSize != nil && *pageSize >= 1 {
		rPageSize = *pageSize
	}

	if rPageSize > pageSizeMax {
		rPageSize = pageSizeMax
	}

	if rPageSize < pageSizeMin {
		rPageSize = pageSizeMin
	}

	return
}

func toHTTP2Shows(shows []*dbmodels.LatestShow) []http.Show {
	r := make([]http.Show, 0, len(shows))
	for _, show := range shows {
		r = append(r, toHTTP2Show(show))
	}
	return r
}

func toHTTP2Show(show *dbmodels.LatestShow) http.Show {
	return http.Show{
		AirDate:       openapiTypes.Date{Time: show.AirDate.Time},
		EpisodeNumber: show.EpisodeNumber.Int,
		Id:            show.ID,
		ShowId:        show.ShowID,
		OriginalTitle: show.OriginalTitle,
		PosterLink:    tmdb.GetImageURL(show.PosterPath, tmdb.W300),
		Quality:       show.Quality,
		HasViSub:      show.HasViSub,
		Runtime:       show.Runtime.Int,
		SeasonNumber:  show.SeasonNumber.Int,
		Title:         show.Title,
		Type:          http.ShowType(show.Type),
		VoteAverage:   show.VoteAverage,
	}
}

func toHTTP2Pagination(page, pageSize int, total int64) http.Pagination {
	return http.Pagination{
		Page:         page,
		TotalPages:   int(math.Ceil(float64(total) / float64(pageSize))),
		TotalResults: total,
	}
}
