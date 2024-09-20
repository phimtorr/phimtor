package http

import (
	"database/sql"
	"errors"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/server/admin/http/handler"
	"github.com/phimtorr/phimtor/server/admin/repository"
	"github.com/phimtorr/phimtor/server/admin/s3"
	"github.com/phimtorr/phimtor/server/admin/tmdb"
)

type Server struct {
	videoHandler *handler.VideoHandler
	userHandler  *handler.UserHandler
	tmdbHandler  *handler.TMDBHandler
}

func NewHTTPServer(db *sql.DB, authClient *auth.Client) Server {
	return Server{
		videoHandler: handler.NewVideoHandler(
			repository.NewRepository(db),
			s3.NewService(),
		),
		userHandler: handler.NewUserHandler(authClient),
		tmdbHandler: handler.NewTMDBHandler(
			tmdb.NewClient(),
			repository.NewTMDBRepository(db),
		),
	}
}

func (s Server) Register(r chi.Router) {
	r.Get("/", errHandlerFunc(handler.Home))

	r.Get("/videos/{id}", errHandlerFunc(s.videoHandler.ViewVideo))
	r.Post("/videos/{id}/torrents/create", errHandlerFunc(s.videoHandler.CreateTorrent))
	r.Delete("/videos/{id}/torrents/{torrentID}", errHandlerFunc(s.videoHandler.DeleteTorrent))
	r.Post("/videos/{id}/subtitles/create", errHandlerFunc(s.videoHandler.CreateSubtitle))
	r.Delete("/videos/{id}/subtitles/{subtitleID}", errHandlerFunc(s.videoHandler.DeleteSubtitle))

	r.Get("/users", errHandlerFunc(s.userHandler.ListUsers))
	r.Get("/users/{uid}", errHandlerFunc(s.userHandler.ViewUser))
	r.Post("/users/{uid}/update-premium", errHandlerFunc(s.userHandler.UpdatePremium))

	r.Get("/latest-shows", errHandlerFunc(s.tmdbHandler.ListLatestShows))

	r.Get("/movies", errHandlerFunc(s.tmdbHandler.ViewMovies))
	r.Post("/movies/create", errHandlerFunc(s.tmdbHandler.CreateMovie))
	r.Get("/movies/{id}", errHandlerFunc(s.tmdbHandler.ViewMovie))
	r.Post("/movies/{id}/fetch-from-tmdb", errHandlerFunc(s.tmdbHandler.FetchMovieFromTMDB))
	r.Post("/movies/{id}/create-video", errHandlerFunc(s.tmdbHandler.CreateMovieVideo))
	r.Post("/movies/{id}/sync", errHandlerFunc(s.tmdbHandler.SyncMovie))

	r.Get("/tv-series", errHandlerFunc(s.tmdbHandler.ViewTVSeriesShows))
	r.Post("/tv-series/create", errHandlerFunc(s.tmdbHandler.CreateTVSeries))
	r.Get("/tv-series/{showID}", errHandlerFunc(s.tmdbHandler.ViewTVSeriesShow))
	r.Post("/tv-series/{showID}/fetch-from-tmdb", errHandlerFunc(s.tmdbHandler.FetchTVSeriesFromTMDB))
	r.Get("/tv-series/{showID}/seasons/{seasonNumber}", errHandlerFunc(s.tmdbHandler.ViewTVSeason))
	r.Get("/tv-series/{showID}/seasons/{seasonNumber}/episodes/{episodeNumber}", errHandlerFunc(s.tmdbHandler.ViewTVEpisode))
	r.Post("/tv-series/{showID}/seasons/{seasonNumber}/episodes/{episodeNumber}/create-video", errHandlerFunc(s.tmdbHandler.CreateTVEpisodeVideo))
	r.Post("/tv-series/{showID}/sync", errHandlerFunc(s.tmdbHandler.SyncTVSeries))
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
