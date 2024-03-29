package admin

import (
	"database/sql"
	"net/http"

	"github.com/a-h/templ"
	"github.com/phimtorr/phimtor/server/admin/ui"

	"github.com/phimtorr/phimtor/server/repository"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/server/admin/handler"
	"github.com/rs/zerolog/log"
)

type HTTPServer struct {
	handler *handler.Handler
}

func NewHTTPServer(db *sql.DB) HTTPServer {
	return HTTPServer{
		handler: handler.New(repository.NewAdminRepository(db)),
	}
}

func (s HTTPServer) Register(r chi.Router) {
	r.Get("/shows", errHandlerFunc(s.handler.ListShows))

	r.Get("/movies/create", templ.Handler(ui.MovieForm()).ServeHTTP)
}

func errHandlerFunc(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			respondError(w, r, err)
		}
	}
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
	log.Ctx(r.Context()).Error().
		Err(err).
		Str("slug", slug).
		Msg(msg)
	http.Error(w, msg, status)
}
