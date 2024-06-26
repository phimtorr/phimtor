package http

import (
	"database/sql"
	"net/http"

	"firebase.google.com/go/v4/auth"

	"github.com/phimtorr/phimtor/server/admin/http/handler"
	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/admin/repository"
	"github.com/phimtorr/phimtor/server/admin/s3"

	"github.com/a-h/templ"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	commonErrors "github.com/phimtorr/phimtor/common/errors"
)

type Server struct {
	handler *handler.Handler
}

func NewHTTPServer(db *sql.DB, authClient *auth.Client) Server {
	return Server{
		handler: handler.New(
			repository.NewRepository(db),
			s3.NewService(),
			authClient,
		),
	}
}

func (s Server) Register(r chi.Router) {
	r.Get("/", errHandlerFunc(s.handler.Home))
	r.Get("/shows", errHandlerFunc(s.handler.ListShows))

	r.Get("/shows/create", templ.Handler(ui.CreateShowForm()).ServeHTTP)
	r.Post("/shows/create", errHandlerFunc(s.handler.CreateShow))
	r.Get("/shows/{id}", errHandlerFunc(s.handler.ViewShow))
	r.Get("/shows/{id}/update", errHandlerFunc(s.handler.ViewUpdateShowForm))
	r.Post("/shows/{id}/update", errHandlerFunc(s.handler.UpdateShow))

	r.Get("/shows/{id}/episodes", errHandlerFunc(s.handler.ListEpisodes))
	r.Get("/shows/{id}/episodes/create", errHandlerFunc(s.handler.ViewCreateEpisodeForm))
	r.Post("/shows/{id}/episodes/create", errHandlerFunc(s.handler.CreateEpisode))
	r.Get("/shows/{id}/episodes/{episodeID}/view", errHandlerFunc(s.handler.ViewEpisode))

	r.Get("/videos/{id}", errHandlerFunc(s.handler.ViewVideo))
	r.Post("/videos/{id}/torrents/create", errHandlerFunc(s.handler.CreateTorrent))
	r.Delete("/videos/{id}/torrents/{torrentID}", errHandlerFunc(s.handler.DeleteTorrent))
	r.Post("/videos/{id}/subtitles/create", errHandlerFunc(s.handler.CreateSubtitle))
	r.Delete("/videos/{id}/subtitles/{subtitleID}", errHandlerFunc(s.handler.DeleteSubtitle))

	r.Get("/users", errHandlerFunc(s.handler.ListUsers))
	r.Get("/users/{uid}", errHandlerFunc(s.handler.ViewUser))
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
	http.Error(w, msg+": "+err.Error(), status)
}
