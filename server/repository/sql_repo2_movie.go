package repository

import (
	"context"
	"encoding/json"
	"fmt"

	tmdb "github.com/cyruzin/golang-tmdb"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/phimtorr/phimtor/server/http2"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r SQLRepo2) GetMovie(ctx context.Context, id int64) (http2.Movie, error) {
	dbMovie, err := dbmodels.Movies(
		dbmodels.MovieWhere.ID.EQ(id),
	).One(ctx, r.db)
	if err != nil {
		return http2.Movie{}, fmt.Errorf("get movie: %w", err)
	}

	return toHTTP2Movie(dbMovie), nil
}

func toHTTP2Movie(dbMovie *dbmodels.Movie) http2.Movie {
	return http2.Movie{
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

func toHTTP2Genres(dbGenres types.JSON) []http2.Genre {
	var genres []http2.Genre
	_ = json.Unmarshal(dbGenres, &genres)

	return genres
}
