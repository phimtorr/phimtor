package repository

import (
	"context"
	"fmt"
	"math"

	tmdb "github.com/cyruzin/golang-tmdb"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/http2"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r SQLRepo2) GetLatestEpisodes(ctx context.Context, params http2.GetLatestEpisodesParams) ([]http2.Show, http2.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		dbmodels.ViewShowWhere.Type.EQ(dbmodels.ViewShowsTypeEpisode),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, true)
}

func (r SQLRepo2) GetLatestMovies(ctx context.Context, params http2.GetLatestMoviesParams) ([]http2.Show, http2.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		dbmodels.ViewShowWhere.Type.EQ(dbmodels.ViewShowsTypeMovie),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, true)
}

func (r SQLRepo2) GetLatestTvSeries(ctx context.Context, params http2.GetLatestTvSeriesParams) ([]http2.Show, http2.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		dbmodels.ViewShowWhere.Type.EQ(dbmodels.ViewShowsTypeEpisode),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, true)
}

func (r SQLRepo2) SearchShows(ctx context.Context, params http2.SearchShowsParams) ([]http2.Show, http2.Pagination, error) {
	page, pageSize := toPageAndPageSize(params.Page, params.PageSize)

	queryMods := []qm.QueryMod{
		qm.Where("MATCH(title, original_title) AGAINST (?)", params.Query),
	}

	return r.queryShows(ctx, queryMods, page, pageSize, false)
}

func (r SQLRepo2) queryShows(ctx context.Context, queryMods []qm.QueryMod, page, pageSize int, sort bool) ([]http2.Show, http2.Pagination, error) {
	pagingQueryMods := append(queryMods,
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
	)

	if sort {
		pagingQueryMods = append(pagingQueryMods, qm.OrderBy(dbmodels.ViewShowColumns.AirDate+" DESC"))
	}

	shows, err := dbmodels.ViewShows(pagingQueryMods...).
		All(ctx, r.db)
	if err != nil {
		return nil, http2.Pagination{}, fmt.Errorf("get shows: %w", err)
	}

	count, err := dbmodels.ViewShows(queryMods...).Count(ctx, r.db)
	if err != nil {
		return nil, http2.Pagination{}, fmt.Errorf("count shows: %w", err)
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

func toHTTP2Shows(shows []*dbmodels.ViewShow) []http2.Show {
	r := make([]http2.Show, 0, len(shows))
	for _, show := range shows {
		r = append(r, toHTTP2Show(show))
	}
	return r
}

func toHTTP2Show(show *dbmodels.ViewShow) http2.Show {
	return http2.Show{
		AirDate:       openapiTypes.Date{Time: show.AirDate.Time},
		EpisodeNumber: show.EpisodeNumber.Int,
		Id:            show.ID,
		MovieID:       show.MovieID.Int64,
		OriginalTitle: show.OriginalTitle,
		PosterLink:    tmdb.GetImageURL(show.PosterPath, tmdb.W300),
		Quality:       http2.ShowQuality(show.Quality),
		Runtime:       show.Runtime.Int,
		SeasonNumber:  show.SeasonNumber.Int,
		Title:         show.Title,
		TvSeriesID:    show.TVSeriesID.Int64,
		Type:          http2.ShowType(show.Type),
		VoteAverage:   show.VoteAverage,
	}
}

func toHTTP2Pagination(page, pageSize int, total int64) http2.Pagination {
	return http2.Pagination{
		Page:         page,
		TotalPages:   int(math.Ceil(float64(total) / float64(pageSize))),
		TotalResults: total,
	}
}
