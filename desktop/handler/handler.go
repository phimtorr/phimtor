package handler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/setting"
	"github.com/phimtorr/phimtor/desktop/torrent"
)

type Handler struct {
	torManager      *torrent.Manager
	settingsStorage *setting.Storage
	apiClient       api.ClientWithResponsesInterface
}

func New(
	torManager *torrent.Manager,
	settingsStorage *setting.Storage,
	apiClient api.ClientWithResponsesInterface,
) *Handler {
	if torManager == nil {
		panic("torrent manager is required")
	}
	if settingsStorage == nil {
		panic("settings storage is required")
	}
	if apiClient == nil {
		panic("apiClient is required")
	}
	return &Handler{
		torManager:      torManager,
		settingsStorage: settingsStorage,
		apiClient:       apiClient,
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

		torrentName, err := url.QueryUnescape(r.URL.Query().Get("torrent"))
		if err != nil {
			handleError(w, r, "Unescape torrent name", err, http.StatusBadRequest)
			return
		}

		h.GetVideo(w, r, id, torrentName)
		return
	})

	r.Get("/stream/{infoHash}/{fileIndex}", func(w http.ResponseWriter, r *http.Request) {
		infoHashStr := chi.URLParam(r, "infoHash")
		infoHash, err := torrent.InfoHashFromString(infoHashStr)
		if err != nil {
			handleError(w, r, "Parse info hash", err, http.StatusBadRequest)
			return
		}
		fileIndexStr := chi.URLParam(r, "fileIndex")
		fileIndex, err := strconv.Atoi(fileIndexStr)
		if err != nil {
			handleError(w, r, "Parse file index", err, http.StatusBadRequest)
			return
		}

		h.Stream(w, r, infoHash, fileIndex)
		return
	})

	// settings
	r.Get("/settings", h.GetSettings)
	r.Post("/settings", h.UpdateSetting)
	r.Post("/settings/change-data-dir", h.ChangeDataDir)
}

func handleError(w http.ResponseWriter, r *http.Request, msg string, err error, status int) {
	log.Ctx(r.Context()).Error().Err(err).Msg(msg)
	http.Error(w, msg+": "+err.Error(), status)
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("HX-Redirect", url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
