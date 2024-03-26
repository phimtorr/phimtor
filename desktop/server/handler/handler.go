package handler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/phimtorr/phimtor/desktop/auth"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/desktop/client"
	"github.com/phimtorr/phimtor/desktop/setting"
	"github.com/phimtorr/phimtor/desktop/torrent"
)

type Handler struct {
	torManager      *torrent.Manager
	settingsStorage *setting.Storage
	apiClient       *client.Client
	authService     *auth.FirebaseAuth
}

func New(
	torManager *torrent.Manager,
	settingsStorage *setting.Storage,
	apiClient *client.Client,
	authService *auth.FirebaseAuth,
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
	if authService == nil {
		panic("authService is required")
	}
	return &Handler{
		torManager:      torManager,
		settingsStorage: settingsStorage,
		apiClient:       apiClient,
		authService:     authService,
	}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/", h.Home)

	r.Get("/shows", h.ListShows)
	r.Get("/shows/search", h.SearchShows)

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

	r.Post("/open-in-vlc/{infoHash}/{fileIndex}", func(w http.ResponseWriter, r *http.Request) {
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

		h.OpenInVLC(w, r, infoHash, fileIndex)
		return
	})

	r.Get("/stats/{infoHash}/{fileIndex}", func(w http.ResponseWriter, r *http.Request) {
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

		h.Stats(w, r, infoHash, fileIndex)
		return
	})

	// subtitles
	// select subtitle
	r.Post("/videos/{videoID}/subtitles/{subtitleName}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
		if err != nil {
			handleError(w, r, "Parse video id", err, http.StatusBadRequest)
			return
		}
		subtitleName, err := url.QueryUnescape(chi.URLParam(r, "subtitleName"))
		if err != nil {
			handleError(w, r, "Unescape subtitle name", err, http.StatusBadRequest)
			return
		}

		h.SelectSubtitle(w, r, id, subtitleName)
	})
	// unset subtitle
	r.Post("/videos/{videoID}/subtitles", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
		if err != nil {
			handleError(w, r, "Parse video id", err, http.StatusBadRequest)
			return
		}

		h.SelectSubtitle(w, r, id, "")
	})
	// upload file
	r.Post("/videos/{videoID}/subtitles/upload", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
		if err != nil {
			handleError(w, r, "Parse video id", err, http.StatusBadRequest)
			return
		}

		h.UploadSubtitle(w, r, id)
	})
	r.Post("/videos/{videoID}/subtitles/adjust", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
		if err != nil {
			handleError(w, r, "Parse video id", err, http.StatusBadRequest)
			return
		}

		h.AdjustSubtitle(w, r, id)
	})
	r.Post("/videos/{videoID}/subtitles/{subtitleName}/download", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
		if err != nil {
			handleError(w, r, "Parse video id", err, http.StatusBadRequest)
			return
		}
		subtitleName, err := url.QueryUnescape(chi.URLParam(r, "subtitleName"))
		if err != nil {
			handleError(w, r, "Unescape subtitle name", err, http.StatusBadRequest)
			return
		}

		h.DownloadSubtitle(w, r, id, subtitleName)
	})
	// settings
	r.Get("/settings", h.GetSettings)
	r.Post("/settings", h.UpdateSetting)
	r.Post("/settings/change-data-dir", h.ChangeDataDir)

	// auth
	r.Get("/sign-in", templ.Handler(ui.SignIn()).ServeHTTP)
	r.Post("/sign-in", h.SignIn)

	r.Get("/sign-up", templ.Handler(ui.SignUp()).ServeHTTP)
	r.Post("/sign-up", h.SignUp)

	r.HandleFunc("/sign-out", h.SignOut)
}

func handleError(w http.ResponseWriter, r *http.Request, msg string, err error, status int) {
	log.Ctx(r.Context()).Error().Err(err).Msg(msg)
	http.Error(w, msg+": "+err.Error(), status)
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
