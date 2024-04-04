package repository

import (
	"context"

	"github.com/friendsofgo/errors"

	"github.com/phimtorr/phimtor/server/http"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r Repository) GetMovie(ctx context.Context, id int64) (http.Movie, error) {
	dbMovie, err := dbmodels.Shows(
		dbmodels.ShowWhere.ID.EQ(id),
		dbmodels.ShowWhere.Type.EQ(dbmodels.ShowsTypeMovie),
	).One(ctx, r.db)
	if err != nil {
		return http.Movie{}, errors.Wrap(err, "get movie")
	}

	return toHTTPMovie(dbMovie), nil
}

func toHTTPMovie(show *dbmodels.Show) http.Movie {
	return http.Movie{
		Description:       show.Description,
		DurationInMinutes: show.DurationInMinutes,
		Id:                show.ID,
		OriginalTitle:     show.OriginalTitle,
		PosterLink:        show.PosterLink,
		Quantity:          show.Quantity,
		ReleaseYear:       show.ReleaseYear,
		Score:             float32(show.Score),
		Title:             show.Title,
		VideoId:           show.VideoID.Int64,
	}
}
