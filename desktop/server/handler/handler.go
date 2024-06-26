package handler

import (
	"net/http"

	"github.com/phimtorr/phimtor/desktop/upnp"

	"github.com/a-h/templ"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/desktop/auth"
	"github.com/phimtorr/phimtor/desktop/client"
	"github.com/phimtorr/phimtor/desktop/server/state"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/setting"
	"github.com/phimtorr/phimtor/desktop/torrent"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	torManager      *torrent.Manager
	settingsStorage *setting.Storage
	apiClient       *client.Client
	authService     *auth.FirebaseAuth
	upnpService     *upnp.UPnP

	state *state.State
}

func New(
	torManager *torrent.Manager,
	settingsStorage *setting.Storage,
	apiClient *client.Client,
	authService *auth.FirebaseAuth,
	upnpService *upnp.UPnP,
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
	if upnpService == nil {
		panic("upnpService is required")
	}
	return &Handler{
		torManager:      torManager,
		settingsStorage: settingsStorage,
		apiClient:       apiClient,
		authService:     authService,
		upnpService:     upnpService,
		state:           state.New(),
	}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/", errHandlerFunc(h.Home))

	r.Get("/shows", errHandlerFunc(h.ListShows))
	r.Get("/shows/search", errHandlerFunc(h.SearchShows))

	r.Get("/movies/{id}", errHandlerFunc(h.GetMovie))
	r.Get("/series/{id}", errHandlerFunc(h.GetSeries))

	r.Get("/videos/{id}", errHandlerFunc(h.GetVideo))
	r.Get("/stream/{infoHash}/{fileIndex}/{fileName}", errHandlerFunc(h.Stream))
	r.Post("/open-in-vlc/{infoHash}/{fileIndex}", errHandlerFunc(h.OpenInVLC))
	r.Get("/stats/{infoHash}/{fileIndex}", errHandlerFunc(h.Stats))

	// subtitles
	// select subtitle
	r.Post("/videos/{videoID}/subtitles/{subtitleID}", errHandlerFunc(h.SelectSubtitle))
	// unset subtitle
	r.Post("/videos/{videoID}/subtitles", errHandlerFunc(h.UnselectSubtitle))
	// upload file
	r.Post("/videos/{videoID}/subtitles/upload", errHandlerFunc(h.UploadSubtitle))
	r.Post("/videos/{videoID}/subtitles/adjust", errHandlerFunc(h.AdjustSubtitle))
	r.Post("/videos/{videoID}/subtitles/{subtitleID}/download", errHandlerFunc(h.DownloadSubtitle))
	// settings
	r.Get("/settings", errHandlerFunc(h.GetSettings))
	r.Post("/settings", errHandlerFunc(h.UpdateSetting))
	r.Post("/settings/change-data-dir", errHandlerFunc(h.ChangeDataDir))

	// auth
	r.Get("/sign-in", templ.Handler(ui.SignIn()).ServeHTTP)
	r.Post("/sign-in", errHandlerFunc(h.SignIn))

	r.Get("/sign-up", templ.Handler(ui.SignUp()).ServeHTTP)
	r.Post("/sign-up", errHandlerFunc(h.SignUp))

	r.HandleFunc("/sign-out", h.SignOut)

	// UPnP
	r.Route("/upnp/videos/{id}", func(r chi.Router) {
		r.Get("/", errHandlerFunc(h.ViewUPnP))

		r.Get("/torrents", errHandlerFunc(h.UPnPListTorrents))
		r.Post("/torrents/{torrentID}", errHandlerFunc(h.UPnSetTorrent))

		r.Get("/subtitles", errHandlerFunc(h.UPnPListSubtitles))
		r.Post("/subtitles/{subtitleID}", errHandlerFunc(h.UPnPSetSubtitle))
		r.Post("/subtitles/upload", errHandlerFunc(h.UPnPUploadSubtitle))
	})
	r.Route("/upnp/devices", func(r chi.Router) {
		r.Get("/", errHandlerFunc(h.UPnPListDevices))
		r.Post("/{udn}", errHandlerFunc(h.UPnPSelectDevice))
		r.Post("/scan", errHandlerFunc(h.ScanDevices))
	})
	r.Route("/upnp/actions", func(r chi.Router) {
		r.Post("/play", errHandlerFunc(h.UPnPPlay))
		r.Post("/pause", errHandlerFunc(h.UPnPPause))
		r.Post("/stop", errHandlerFunc(h.UPnPStop))
	})

}

func errHandlerFunc(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
			return
		}

		var slugErr commonErrors.SlugError
		if !errors.As(err, &slugErr) {
			handleError(w, r, "Internal error", "internal-error", err, http.StatusInternalServerError)
			return
		}

		switch slugErr.ErrorType() {
		case commonErrors.ErrorTypeIncorrectInput:
			handleError(w, r, "Incorrect input", slugErr.Slug(), err, http.StatusBadRequest)
		case commonErrors.ErrorTypeAuthorization:
			handleError(w, r, "Authorization error", slugErr.Slug(), err, http.StatusUnauthorized)
		default:
			handleError(w, r, "Internal error", slugErr.Slug(), err, http.StatusInternalServerError)
		}
	}
}

func handleError(w http.ResponseWriter, r *http.Request, msg string, slug string, err error, status int) {
	log.Ctx(r.Context()).Error().
		Str("slug", slug).
		Err(err).
		Msg(msg)

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Reswap", "innerHTML")
	}
	http.Error(w, msg+": "+err.Error(), status)
}

func fullyRedirect(w http.ResponseWriter, r *http.Request, url string) {
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
