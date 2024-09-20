package http

import (
	"database/sql"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/server/admin/http/handler"
	"github.com/phimtorr/phimtor/server/admin/http/handler2"
	"github.com/phimtorr/phimtor/server/admin/repository"
	"github.com/phimtorr/phimtor/server/admin/s3"
	"github.com/phimtorr/phimtor/server/admin/tmdb"
)

type Server struct {
	handler  *handler.Handler
	handler2 *handler2.Handler
}

func NewHTTPServer(db *sql.DB, authClient *auth.Client) Server {
	return Server{
		handler: handler.New(
			repository.NewRepository(db),
			s3.NewService(),
			authClient,
		),
		handler2: handler2.NewHandler(
			tmdb.NewClient(),
			repository.NewTMDBRepository(db),
		),
	}
}

func (s Server) Register(r chi.Router) {
	r.Get("/", errHandlerFunc(s.handler.Home))

	r.Get("/videos/{id}", errHandlerFunc(s.handler.ViewVideo))
	r.Post("/videos/{id}/torrents/create", errHandlerFunc(s.handler.CreateTorrent))
	r.Delete("/videos/{id}/torrents/{torrentID}", errHandlerFunc(s.handler.DeleteTorrent))
	r.Post("/videos/{id}/subtitles/create", errHandlerFunc(s.handler.CreateSubtitle))
	r.Delete("/videos/{id}/subtitles/{subtitleID}", errHandlerFunc(s.handler.DeleteSubtitle))

	r.Get("/users", errHandlerFunc(s.handler.ListUsers))
	r.Get("/users/{uid}", errHandlerFunc(s.handler.ViewUser))
	r.Post("/users/{uid}/update-premium", errHandlerFunc(s.handler.UpdatePremium))

	r.Get("/latest-shows", errHandlerFunc(s.handler2.ListLatestShows))

	r.Get("/movies", errHandlerFunc(s.handler2.ViewMovies))
	r.Post("/movies/create", errHandlerFunc(s.handler2.CreateMovie))
	r.Get("/movies/{id}", errHandlerFunc(s.handler2.ViewMovie))
	r.Post("/movies/{id}/fetch-from-tmdb", errHandlerFunc(s.handler2.FetchMovieFromTMDB))
	r.Post("/movies/{id}/create-video", errHandlerFunc(s.handler2.CreateMovieVideo))
	r.Post("/movies/{id}/sync", errHandlerFunc(s.handler2.SyncMovie))

	r.Get("/tv-series", errHandlerFunc(s.handler2.ViewTVSeriesShows))
	r.Post("/tv-series/create", errHandlerFunc(s.handler2.CreateTVSeries))
	r.Get("/tv-series/{showID}", errHandlerFunc(s.handler2.ViewTVSeriesShow))
	r.Post("/tv-series/{showID}/fetch-from-tmdb", errHandlerFunc(s.handler2.FetchTVSeriesFromTMDB))
	r.Get("/tv-series/{showID}/seasons/{seasonNumber}", errHandlerFunc(s.handler2.ViewTVSeason))
	r.Get("/tv-series/{showID}/seasons/{seasonNumber}/episodes/{episodeNumber}", errHandlerFunc(s.handler2.ViewTVEpisode))
	r.Post("/tv-series/{showID}/seasons/{seasonNumber}/episodes/{episodeNumber}/create-video", errHandlerFunc(s.handler2.CreateTVEpisodeVideo))
	r.Post("/tv-series/{showID}/sync", errHandlerFunc(s.handler2.SyncTVSeries))
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
