package ports

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type Repository interface {
	ListShows(ctx context.Context, params ListShowsParams) ([]Show, Pagination, error)
	GetMovie(ctx context.Context, id int64) (Movie, error)
	GetSeries(ctx context.Context, id int64) (Series, error)
	GetVideo(ctx context.Context, id int64) (Video, error)
}

type HttpServer struct {
	repo Repository
}

func NewHttpServer(repository Repository) HttpServer {
	if repository == nil {
		panic("repository is nil")
	}
	return HttpServer{repo: repository}
}

func handleError(w http.ResponseWriter, r *http.Request, msg string, err error, status int) {
	log.Ctx(r.Context()).Error().Err(err).Msg(msg)
	code := "internal-error"
	if status == http.StatusBadRequest {
		code = "bad-request"
	}

	render.Status(r, status)
	render.JSON(w, r, ErrorResponse{
		Code:    code,
		Message: msg,
	})
}
