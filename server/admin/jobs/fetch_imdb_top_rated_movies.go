package jobs

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/phimtorr/phimtor/server/admin/yts"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

type TMDBClient interface {
	ListTopRatedMovies(ctx context.Context, page int) (*tmdb.MovieTopRated, error)
	GetMovieDetails(_ context.Context, movieID int) (*tmdb.MovieDetails, error)
}

type YTSClient interface {
	GetMovieByIMDbID(ctx context.Context, imdbID string) (yts.Movie, error)
}

type Repository interface {
	UpdateMovie(ctx context.Context, movie *tmdb.MovieDetails) error
	UpdateYTSMovie(ctx context.Context, videoID int64, movie yts.Movie) error
	SyncMovie(ctx context.Context, movieID int64) error
}

type FetchIMDBTopRatedMoviesJob struct {
	tmdbClient TMDBClient
	ytsClient  YTSClient
	repo       Repository
	db         *sql.DB
}

func NewFetchIMDBTopRatedMoviesJob(tmdbClient TMDBClient, ytsClient YTSClient, repo Repository, db *sql.DB) FetchIMDBTopRatedMoviesJob {
	return FetchIMDBTopRatedMoviesJob{
		tmdbClient: tmdbClient,
		ytsClient:  ytsClient,
		repo:       repo,
		db:         db,
	}
}

func (j FetchIMDBTopRatedMoviesJob) Execute(ctx context.Context, startPage, endPage int) error {
	for page := startPage; page <= endPage; page++ {
		if err := j.processPage(ctx, page); err != nil {
			return fmt.Errorf("process page %d: %w", page, err)
		}
	}

	return nil
}

func (j FetchIMDBTopRatedMoviesJob) processPage(ctx context.Context, page int) error {
	log.Ctx(ctx).Info().Msgf("Processing page %d", page)
	resp, err := j.tmdbClient.ListTopRatedMovies(ctx, page)
	if err != nil {
		return fmt.Errorf("list top rated movies: %w", err)
	}

	for _, movie := range resp.Results {
		if err := j.processMovie(ctx, movie.ID); err != nil {
			return fmt.Errorf("process movie %d: %w", movie.ID, err)
		}
	}

	return nil
}

func (j FetchIMDBTopRatedMoviesJob) processMovie(ctx context.Context, movieID int64) error {
	log.Ctx(ctx).Info().Msgf("Processing movie %d", movieID)
	_, err := dbmodels.FindMovie(ctx, j.db, movieID)
	if errors.Is(err, sql.ErrNoRows) {
		if err := j.createMovie(ctx, movieID); err != nil {
			log.Ctx(ctx).Error().Err(err).Msgf("Create movie %d failed", movieID)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("find movie: %w", err)
	}
	log.Ctx(ctx).Info().Msgf("Movie %d already exists", movieID)
	return nil
}

func (j FetchIMDBTopRatedMoviesJob) createMovie(ctx context.Context, movieID int64) error {
	log.Ctx(ctx).Info().Msgf("Creating movie %d", movieID)

	tmdbMovie, err := j.tmdbClient.GetMovieDetails(ctx, int(movieID))
	if err != nil {
		return fmt.Errorf("get movie details: %w", err)
	}

	ytsMovie, err := j.ytsClient.GetMovieByIMDbID(ctx, tmdbMovie.IMDbID)
	if err != nil {
		return fmt.Errorf("get yts movie: %w", err)
	}

	// Create movie
	if err := j.repo.UpdateMovie(ctx, tmdbMovie); err != nil {
		return fmt.Errorf("update movie: %w", err)
	}

	// Create video
	video := dbmodels.Video{}
	if err := video.Insert(ctx, j.db, boil.Infer()); err != nil {
		return fmt.Errorf("insert video: %w", err)
	}

	movie, err := dbmodels.FindMovie(ctx, j.db, movieID)
	if err != nil {
		return fmt.Errorf("find movie: %w", err)
	}

	movie.VideoID = video.ID

	if _, err := movie.Update(ctx, j.db, boil.Whitelist(dbmodels.MovieColumns.VideoID)); err != nil {
		return fmt.Errorf("update movie: %w", err)
	}

	// Create yts movie
	if err := j.repo.UpdateYTSMovie(ctx, video.ID, ytsMovie); err != nil {
		return fmt.Errorf("update yts movie: %w", err)
	}

	// sync movie
	if err := j.repo.SyncMovie(ctx, movieID); err != nil {
		return fmt.Errorf("sync movie: %w", err)
	}

	return nil
}
