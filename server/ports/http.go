package ports

import (
	"database/sql"
	"math"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

const (
	pageSize = 15
)

type HttpServer struct {
	db *sql.DB
}

func NewHttpServer(db *sql.DB) HttpServer {
	return HttpServer{db: db}
}

func (h HttpServer) GetMovie(ctx echo.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (h HttpServer) ListShows(ctx echo.Context, params ListShowsParams) error {
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
		All(ctx.Request().Context(), h.db)
	if err != nil {
		return errors.Wrap(err, "get shows")
	}

	count, err := dbmodels.Shows(queryMods...).Count(ctx.Request().Context(), h.db)
	if err != nil {
		return errors.Wrap(err, "count shows")
	}

	return ctx.JSON(200, map[string]interface{}{
		"shows":      toHTTPBasicInfos(shows),
		"pagination": toHTTPPagination(page, pageSize, count),
	})
}

func unPointer[T any](v *T) T {
	var r T
	if v == nil {
		return r
	}
	return *v
}

func toHTTPBasicInfos(shows []*dbmodels.Show) []BasicInfo {
	var res []BasicInfo
	for _, show := range shows {
		res = append(res, toHTTPBasicInfo(show))
	}
	return res
}

func toHTTPBasicInfo(show *dbmodels.Show) BasicInfo {
	return BasicInfo{
		CurrentEpisode:    show.CurrentEpisode,
		DurationInMinutes: show.DurationInMinutes,
		Id:                show.ID,
		OriginalTitle:     show.OriginalTitle,
		PosterLink:        show.PosterLink,
		Quantity:          show.Quantity,
		ReleaseYear:       show.ReleaseYear,
		Score:             show.Score,
		ShowType:          BasicInfoShowType(show.Type),
		Title:             show.Title,
		TotalEpisodes:     show.TotalEpisodes,
	}
}

func toHTTPPagination(page, pageSize int, total int64) Pagination {
	return Pagination{
		Page:         page,
		TotalPages:   int(math.Ceil(float64(total) / float64(pageSize))),
		TotalResults: total,
	}
}
