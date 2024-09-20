package http

import (
	"context"

	"github.com/phimtorr/phimtor/server/auth"
)

type Repository interface {
	GetLatestEpisodes(ctx context.Context, params GetLatestEpisodesParams) ([]Show, Pagination, error)
	GetLatestMovies(ctx context.Context, params GetLatestMoviesParams) ([]Show, Pagination, error)
	GetLatestTvSeries(ctx context.Context, params GetLatestTvSeriesParams) ([]Show, Pagination, error)
	SearchShows(ctx context.Context, params SearchShowsParams) ([]Show, Pagination, error)

	GetMovie(ctx context.Context, id int64) (Movie, error)
	GetTvSeries(ctx context.Context, id int64) (TvSeries, error)
	GetTvSeason(ctx context.Context, showID int64, seasonNumber int) (TVSeason, error)

	GetVideo(ctx context.Context, user auth.User, id int64) (Video, error)
}
