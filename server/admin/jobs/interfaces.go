package jobs

import (
	"context"

	tmdb "github.com/cyruzin/golang-tmdb"

	"github.com/phimtorr/phimtor/server/admin/yts"
)

type TMDBClient interface {
	ListTopRatedMovies(ctx context.Context, page int) (*tmdb.MovieTopRated, error)
	GetMovieDetails(ctx context.Context, movieID int) (*tmdb.MovieDetails, error)
	GetMovieDetailsByIMDbID(_ context.Context, imdbID string) (*tmdb.MovieDetails, error)
}

type YTSClient interface {
	GetMovieByIMDbID(ctx context.Context, imdbID string) (yts.Movie, error)
	GetMovieByID(ctx context.Context, id int64) (yts.Movie, error)
}

type Repository interface {
	UpdateMovie(ctx context.Context, movie *tmdb.MovieDetails) error
	UpdateYTSMovie(ctx context.Context, videoID int64, movie yts.Movie) error
	SyncMovie(ctx context.Context, movieID int64) error
	CreateMovieVideo(ctx context.Context, movieID int64) error
}
