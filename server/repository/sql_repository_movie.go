package repository

import (
	"context"
	"encoding/json"
	"fmt"

	tmdb "github.com/cyruzin/golang-tmdb"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/phimtorr/phimtor/server/http"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r SQLRepository) GetMovie(ctx context.Context, id int64) (http.Movie, error) {
	dbMovie, err := dbmodels.Movies(
		dbmodels.MovieWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		return http.Movie{}, fmt.Errorf("get movie: %w", err)
	}

	return toHTTP2Movie(dbMovie), nil
}

func toHTTP2Movie(dbMovie *dbmodels.Movie) http.Movie {
	return http.Movie{
		BackdropLink:  tmdb.GetImageURL(dbMovie.BackdropPath, tmdb.Original),
		Genres:        toHTTP2Genres(dbMovie.Genres),
		Id:            dbMovie.ID,
		OriginalTitle: dbMovie.OriginalTitle,
		Overview:      dbMovie.Overview,
		PosterLink:    tmdb.GetImageURL(dbMovie.PosterPath, tmdb.W300),
		ReleaseDate:   openapiTypes.Date{Time: dbMovie.ReleaseDate},
		Runtime:       dbMovie.Runtime,
		Status:        dbMovie.Status,
		Tagline:       dbMovie.Tagline,
		Title:         dbMovie.Title,
		VideoID:       dbMovie.VideoID,
		VoteAverage:   dbMovie.VoteAverage,
	}
}

func toHTTP2Genres(dbGenres types.JSON) []http.Genre {
	var genres []http.Genre
	_ = json.Unmarshal(dbGenres, &genres)

	return genres
}
