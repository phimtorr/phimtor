package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/desktop/client/api"
)

type Handler struct {
	client api.ClientWithResponsesInterface
}

func New(client api.ClientWithResponsesInterface) *Handler {
	if client == nil {
		panic("client is required")
	}
	return &Handler{
		client: client,
	}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/", h.Home)

	r.Get("/shows", func(w http.ResponseWriter, r *http.Request) {
		h.ListShows(w, r, api.ListShowsParams{})
		return
	})

	r.Get("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			handleError(w, r, "Parse movie id", err, http.StatusBadRequest)
			return
		}

		h.GetMovie(w, r, id)
		return
	})

	r.Get("/series/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			handleError(w, r, "Parse series id", err, http.StatusBadRequest)
			return
		}

		h.GetSeries(w, r, id)
		return
	})

	r.Get("/videos/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			handleError(w, r, "Parse video id", err, http.StatusBadRequest)
			return
		}

		h.GetVideo(w, r, id)
		return
	})
}

func handleError(w http.ResponseWriter, r *http.Request, msg string, err error, status int) {
	log.Ctx(r.Context()).Error().Err(err).Msg(msg)
	http.Error(w, msg+": "+err.Error(), status)
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("HX-Redirect", url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
