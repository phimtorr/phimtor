package repository

import (
	"context"
	"math"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r TMDBRepository) ListLatestShows(ctx context.Context, page, pageSize int) ([]ui.LatestShow, ui.Pagination, error) {
	shows, err := dbmodels.LatestShows(
		qm.OrderBy(dbmodels.LatestShowColumns.UpdatedAt+" DESC"),
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
	).All(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, err
	}

	count, err := dbmodels.LatestShows().Count(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, err
	}

	totalPages := int(math.Ceil(float64(count) / float64(pageSize)))

	return toUILatestShows(shows), ui.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}, nil

}

func toUILatestShows(shows []*dbmodels.LatestShow) []ui.LatestShow {
	tvSeriesShows := make([]ui.LatestShow, 0, len(shows))
	for _, show := range shows {
		tvSeriesShows = append(tvSeriesShows, toUILatestShow(show))
	}
	return tvSeriesShows
}

func toUILatestShow(show *dbmodels.LatestShow) ui.LatestShow {
	airDateStr := ""
	if show.AirDate.Valid {
		airDateStr = show.AirDate.Time.String()
	}
	return ui.LatestShow{
		ID:            show.ID,
		Type:          show.Type.String(),
		ShowID:        show.ShowID,
		Title:         show.Title,
		OriginalTitle: show.OriginalTitle,
		PosterLink:    tmdb.GetImageURL(show.PosterPath, tmdb.W300),
		AirDate:       airDateStr,
		Runtime:       show.Runtime.Int,
		VoteAverage:   float64(show.VoteAverage),
		Quality:       show.Quality,
		SeasonNumber:  show.SeasonNumber.Int,
		EpisodeNumber: show.EpisodeNumber.Int,
		CreatedAt:     show.CreatedAt,
		UpdatedAt:     show.UpdatedAt,
	}
}
