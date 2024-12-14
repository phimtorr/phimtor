package handler

import (
	"context"

	tmdb "github.com/cyruzin/golang-tmdb"

	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/admin/yts"
)

type TMDBClient interface {
	GetMovieDetails(ctx context.Context, movieID int) (*tmdb.MovieDetails, error)
	GetTVSeriesDetails(ctx context.Context, tvID int) (*tmdb.TVDetails, []*tmdb.TVSeasonDetails, error)
}

type TMDBRepository interface {
	UpdateMovie(ctx context.Context, movie *tmdb.MovieDetails) error
	UpdateTVSeries(ctx context.Context, tv *tmdb.TVDetails, seasons []*tmdb.TVSeasonDetails) error
	UpdateYTSMovie(ctx context.Context, videoID int64, movie yts.Movie) error

	ListLatestShows(ctx context.Context, page, pageSize int) ([]ui.LatestShow, ui.Pagination, error)

	ListMovies(ctx context.Context, page, pageSize int) ([]ui.Movie, ui.Pagination, error)
	GetMovie(ctx context.Context, id int64) (ui.Movie, error)
	CreateMovieVideo(ctx context.Context, movieID int64) error

	ListTVSeriesShows(ctx context.Context, page, pageSize int) ([]ui.TVSeriesShow, ui.Pagination, error)
	GetTVSeriesShow(ctx context.Context, showID int64) (ui.TVSeriesShow, []ui.TVSeason, error)
	GetTVSeason(ctx context.Context, showID int64, seasonNumber int) (ui.TVSeason, []ui.TVEpisode, error)
	GetTVEpisode(ctx context.Context, showID int64, seasonNumber, episodeNumber int) (ui.TVEpisode, error)
	CreateTVEpisodeVideo(ctx context.Context, showID int64, seasonNumber, episodeNumber int) error

	SyncMovie(ctx context.Context, movieID int64) error
	SyncTVSeries(ctx context.Context, tvID int64) error
}

type YTSClient interface {
	GetMovieByIMDbID(ctx context.Context, imdbID string) (yts.Movie, error)
}

type TMDBHandler struct {
	tmdbClient TMDBClient
	repo       TMDBRepository
	ytsClient  YTSClient
}

func NewTMDBHandler(tmdbClient TMDBClient, repo TMDBRepository, ytsClient YTSClient) *TMDBHandler {
	if tmdbClient == nil {
		panic("tmdbClient is required")
	}

	if repo == nil {
		panic("repo is required")
	}

	if ytsClient == nil {
		panic("ytsClient is required")
	}

	return &TMDBHandler{
		tmdbClient: tmdbClient,
		repo:       repo,
		ytsClient:  ytsClient,
	}
}
