package repository

import (
	"context"
	"math"

	commonErrors "github.com/phimtorr/phimtor/common/errors"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/http"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

const (
	pageSizeDefault = 18
	pageSizeMin     = 6
	pageSizeMax     = 24
)

func (r Repository) ListShows(ctx context.Context, params http.ListShowsParams) ([]http.Show, http.Pagination, error) {
	page := unPointer(params.Page)
	pageSize := unPointer(params.PageSize)
	showType := unPointer(params.Type)

	page = max(page, 1)
	pageSize = max(pageSize, pageSizeMin)
	pageSize = min(pageSize, pageSizeMax)

	var queryMods []qm.QueryMod
	if showType != "" {
		queryMods = append(queryMods, dbmodels.ShowWhere.Type.EQ(dbmodels.ShowsType(showType)))
	}

	pagingQueryMods := append(queryMods,
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
		qm.OrderBy(dbmodels.ShowColumns.CreatedAt+" DESC"),
	)

	shows, err := dbmodels.Shows(pagingQueryMods...).
		All(ctx, r.db)
	if err != nil {
		return nil, http.Pagination{}, errors.Wrap(err, "get shows")
	}

	count, err := dbmodels.Shows(queryMods...).Count(ctx, r.db)
	if err != nil {
		return nil, http.Pagination{}, errors.Wrap(err, "count shows")
	}

	return toHTTPShows(shows), toHTTPPagination(page, pageSize, count), nil
}

func (r Repository) SearchShow(ctx context.Context, params http.SearchShowsParams) ([]http.Show, http.Pagination, error) {
	query := params.Query
	page := unPointer(params.Page)

	page = max(page, 1)
	pageSize := pageSizeDefault

	if query == "" {
		return nil, http.Pagination{}, commonErrors.NewIncorrectInputError("empty-query", "query is empty")
	}

	var queryMods []qm.QueryMod
	queryMods = append(queryMods, qm.Where("MATCH(title, original_title) AGAINST (?)", query))

	pagingQueryMods := append(queryMods,
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
	)

	shows, err := dbmodels.Shows(pagingQueryMods...).All(ctx, r.db)
	if err != nil {
		return nil, http.Pagination{}, errors.Wrap(err, "get shows")
	}

	count, err := dbmodels.Shows(queryMods...).Count(ctx, r.db)
	if err != nil {
		return nil, http.Pagination{}, errors.Wrap(err, "count shows")
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

func toHTTPShows(shows []*dbmodels.Show) []http.Show {
	var res []http.Show
	for _, show := range shows {
		res = append(res, toHTTPBasicInfo(show))
	}
	return res
}

func toHTTPBasicInfo(show *dbmodels.Show) http.Show {
	return http.Show{
		CurrentEpisode:    show.CurrentEpisode,
		DurationInMinutes: show.DurationInMinutes,
		Id:                show.ID,
		OriginalTitle:     show.OriginalTitle,
		PosterLink:        show.PosterLink,
		Quantity:          show.Quantity,
		ReleaseYear:       show.ReleaseYear,
		Score:             float32(show.Score),
		Type:              http.ShowType(show.Type),
		Title:             show.Title,
		TotalEpisodes:     show.TotalEpisodes,
	}
}

func toHTTPPagination(page, pageSize int, total int64) http.Pagination {
	return http.Pagination{
		Page:         page,
		TotalPages:   int(math.Ceil(float64(total) / float64(pageSize))),
		TotalResults: total,
	}
}
