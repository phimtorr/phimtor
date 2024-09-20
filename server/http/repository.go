package http

import (
	"context"

	"github.com/phimtorr/phimtor/server/auth"
)

type Repository interface {
	ListRecentlyAddedMovies(ctx context.Context, params ListRecentlyAddedMoviesParams) ([]Show, Pagination, error)
	ListLatestEpisodes(ctx context.Context, params ListLatestEpisodesParams) ([]Show, Pagination, error)
	ListLatestMovies(ctx context.Context, params ListLatestMoviesParams) ([]Show, Pagination, error)
	ListLatestTvSeries(ctx context.Context, params ListLatestTvSeriesParams) ([]Show, Pagination, error)
	SearchShows(ctx context.Context, params SearchShowsParams) ([]Show, Pagination, error)

	GetMovie(ctx context.Context, id int64) (Movie, error)
	GetTvSeries(ctx context.Context, id int64) (TvSeries, error)
	GetTvSeason(ctx context.Context, showID int64, seasonNumber int) (TVSeason, error)

	GetVideo(ctx context.Context, user auth.User, id int64) (Video, error)
}
