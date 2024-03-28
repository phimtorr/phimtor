package repository

import (
	"context"
	"math"

	"github.com/phimtorr/phimtor/server/repository/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/admin/ui"
)

func (r AdminRepository) ListShows(ctx context.Context, page, pageSize int) ([]ui.Show, ui.Pagination, error) {
	dbShows, err := dbmodels.Shows(
		qm.OrderBy(dbmodels.ShowColumns.UpdatedAt+" DESC"),
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
	).All(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, err
	}

	count, err := dbmodels.Shows().Count(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, err
	}

	shows := make([]ui.Show, len(dbShows))
	for i, dbShow := range dbShows {
		shows[i] = ui.Show{
			ID:            dbShow.ID,
			Title:         dbShow.Title,
			OriginalTitle: dbShow.OriginalTitle,
			Poster:        dbShow.PosterLink,
		}
	}

	totalPages := int(math.Ceil(float64(count) / float64(pageSize)))

	return shows, ui.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}, nil

}
