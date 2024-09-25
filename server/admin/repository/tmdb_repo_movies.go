package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r TMDBRepository) ListMovies(ctx context.Context, page, pageSize int) ([]ui.Movie, ui.Pagination, error) {
	dbMovies, err := dbmodels.Movies(
		qm.OrderBy(dbmodels.MovieColumns.UpdatedAt+" DESC"),
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
	).All(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, fmt.Errorf("list movies: %w", err)
	}

	count, err := dbmodels.Movies().Count(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, fmt.Errorf("count movies: %w", err)
	}

	totalPages := int(math.Ceil(float64(count) / float64(pageSize)))

	return toUIMovies(dbMovies), ui.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}, nil

}

func (r TMDBRepository) GetMovie(ctx context.Context, id int64) (ui.Movie, error) {
	dbMovie, err := dbmodels.FindMovie(ctx, r.db, id)
	if err != nil {
		return ui.Movie{}, fmt.Errorf("get movie: %w", err)
	}

	return toUIMovie(dbMovie), nil
}

func (r TMDBRepository) CreateMovieVideo(ctx context.Context, movieID int64) error {
	return withTx(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		dbMovie, err := dbmodels.FindMovie(ctx, tx, movieID)
		if err != nil {
			return fmt.Errorf("find movie: %w", err)
		}

		if dbMovie.VideoID != 0 {
			return nil
		}

		video := &dbmodels.Video{}

		if err := video.Insert(ctx, tx, boil.Infer()); err != nil {
			return fmt.Errorf("insert video: %w", err)
		}

		dbMovie.VideoID = video.ID

		if _, err := dbMovie.Update(ctx, tx, boil.Whitelist(dbmodels.MovieColumns.VideoID)); err != nil {
			return fmt.Errorf("update movie: %w", err)
		}

		return nil
	})
}

func toUIMovies(dbMovies []*dbmodels.Movie) []ui.Movie {
	movies := make([]ui.Movie, len(dbMovies))
	for i, dbMovie := range dbMovies {
		movies[i] = toUIMovie(dbMovie)
	}
	return movies
}

func toUIMovie(dbMovie *dbmodels.Movie) ui.Movie {
	return ui.Movie{
		ID:            dbMovie.ID,
		IMDBID:        dbMovie.ImdbID,
		Title:         dbMovie.Title,
		OriginalTitle: dbMovie.OriginalTitle,
		Status:        dbMovie.Status,
		Tagline:       dbMovie.Tagline,
		Genres:        string(dbMovie.Genres),
		Overview:      dbMovie.Overview,
		PosterLink:    tmdb.GetImageURL(dbMovie.PosterPath, tmdb.W300),
		BackdropLink:  tmdb.GetImageURL(dbMovie.BackdropPath, tmdb.W780),
		ReleaseDate:   dbMovie.ReleaseDate.Format("2006-01-02"),
		Runtime:       dbMovie.Runtime,
		VoteAverage:   float64(dbMovie.VoteAverage),
		VoteCount:     dbMovie.VoteCount,
		VideoID:       dbMovie.VideoID,
		CreatedAt:     dbMovie.CreatedAt,
		UpdatedAt:     dbMovie.UpdatedAt,
	}
}
