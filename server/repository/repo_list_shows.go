package repository

import (
	"context"
	"math"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/ports"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

const (
	pageSize = 15
)

func (r Repository) ListShows(ctx context.Context, params ports.ListShowsParams) ([]ports.Show, ports.Pagination, error) {
	name := unPointer(params.Name)
	page := unPointer(params.Page)

	if page <= 0 {
		page = 1
	}

	var queryMods []qm.QueryMod
	if name != "" {
		queryMods = append(queryMods, dbmodels.ShowWhere.Title.LIKE("%"+name+"%"))
		queryMods = append(queryMods, qm.Or2(dbmodels.ShowWhere.OriginalTitle.LIKE("%"+name+"%")))
	}

	pagingQueryMods := append(queryMods,
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
		qm.OrderBy(dbmodels.ShowColumns.CreatedAt+" DESC"),
	)

	shows, err := dbmodels.Shows(pagingQueryMods...).
		All(ctx, r.db)
	if err != nil {
		return nil, ports.Pagination{}, errors.Wrap(err, "get shows")
	}

	count, err := dbmodels.Shows(queryMods...).Count(ctx, r.db)
	if err != nil {
		return nil, ports.Pagination{}, errors.Wrap(err, "count shows")
	}

	return toHTTPShows(shows), toHTTPPagination(page, pageSize, count), nil
}

func unPointer[T any](v *T) T {
	var r T
	if v == nil {
		return r
	}
	return *v
}

func toHTTPShows(shows []*dbmodels.Show) []ports.Show {
	var res []ports.Show
	for _, show := range shows {
		res = append(res, toHTTPBasicInfo(show))
	}
	return res
}

func toHTTPBasicInfo(show *dbmodels.Show) ports.Show {
	return ports.Show{
		CurrentEpisode:    show.CurrentEpisode,
		DurationInMinutes: show.DurationInMinutes,
		Id:                show.ID,
		OriginalTitle:     show.OriginalTitle,
		PosterLink:        show.PosterLink,
		Quantity:          show.Quantity,
		ReleaseYear:       show.ReleaseYear,
		Score:             float32(show.Score),
		Type:              ports.ShowType(show.Type),
		Title:             show.Title,
		TotalEpisodes:     show.TotalEpisodes,
	}
}

func toHTTPPagination(page, pageSize int, total int64) ports.Pagination {
	return ports.Pagination{
		Page:         page,
		TotalPages:   int(math.Ceil(float64(total) / float64(pageSize))),
		TotalResults: total,
	}
}
