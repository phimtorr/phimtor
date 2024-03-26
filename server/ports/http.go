package ports

import (
	"net/http"

	"github.com/friendsofgo/errors"

	"github.com/go-chi/render"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/rs/zerolog/log"
)

type HttpServer struct {
	repo Repository
}

func NewHttpServer(repository Repository) HttpServer {
	if repository == nil {
		panic("repository is nil")
	}
	return HttpServer{repo: repository}
}

func respondError(w http.ResponseWriter, r *http.Request, err error) {
	var slugErr commonErrors.SlugError
	if !errors.As(err, &slugErr) {
		internalError(w, r, "internal-server-error", err)
		return
	}

	switch slugErr.ErrorType() {
	case commonErrors.ErrorTypeAuthorization:
		unauthorizedError(w, r, slugErr.Slug(), err)
	case commonErrors.ErrorTypeIncorrectInput:
		badRequestError(w, r, slugErr.Slug(), err)
	default:
		internalError(w, r, slugErr.Slug(), err)
	}

}

func internalError(w http.ResponseWriter, r *http.Request, slug string, err error) {
	handleError(w, r, "Internal error", slug, err, http.StatusInternalServerError)
}

func unauthorizedError(w http.ResponseWriter, r *http.Request, slug string, err error) {
	handleError(w, r, "Unauthorized", slug, err, http.StatusUnauthorized)
}

func badRequestError(w http.ResponseWriter, r *http.Request, slug string, err error) {
	handleError(w, r, "Bad request", slug, err, http.StatusBadRequest)
}

func handleError(w http.ResponseWriter, r *http.Request, msg string, slug string, err error, status int) {
	log.Ctx(r.Context()).Error().Err(err).Msg(msg)
	render.Status(r, status)
	render.JSON(w, r, ErrorResponse{
		Code:    slug,
		Message: msg,
	})
}
