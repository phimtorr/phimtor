package ports

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Repository interface {
	ListShows(ctx context.Context, params ListShowsParams) ([]Show, Pagination, error)
	GetMovie(ctx context.Context, id int64) (Movie, error)
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
	http.Error(w, msg+": "+err.Error(), status)
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("HX-Redirect", url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
