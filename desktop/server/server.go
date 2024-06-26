package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/phimtorr/phimtor/desktop/upnp"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/desktop/auth"
	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/client"
	"github.com/phimtorr/phimtor/desktop/data"
	"github.com/phimtorr/phimtor/desktop/i18n"
	"github.com/phimtorr/phimtor/desktop/net"
	"github.com/phimtorr/phimtor/desktop/server/handler"
	"github.com/phimtorr/phimtor/desktop/server/ui/style"
	"github.com/phimtorr/phimtor/desktop/setting"
	"github.com/phimtorr/phimtor/desktop/torrent"
	"github.com/phimtorr/phimtor/desktop/updater"
)

type closeFn struct {
	resourceName string
	fn           func() error
}

func newCloseFn(resourceName string, fn func() error) closeFn {
	return closeFn{
		resourceName: resourceName,
		fn:           fn,
	}
}

type Server struct {
	appName  string
	closeFns []closeFn
}

func NewServer(appName string) *Server {
	if appName == "" {
		panic("app name is empty")
	}
	return &Server{
		appName: appName,
	}
}

func (s *Server) Start() int {
	settingsStorage := setting.NewStorage(s.appName)
	s.closeFns = append(s.closeFns, newCloseFn("settingsStorage", func() error {
		return setting.CleanUpStorage(settingsStorage.GetSettings())
	}))

	settings := settingsStorage.GetSettings()
	torManager := torrent.NewManager(settings.GetCurrentDataDir())
	s.closeFns = append(s.closeFns, newCloseFn("torManager", torManager.Close))

	authService := auth.NewFirebaseAuth(build.FirebaseAPIKey, auth.NewFileStorage(s.appName))
	apiClient := client.NewClient(authService)

	updaterSvc := updater.NewUpdater(build.Version, 30*time.Minute, apiClient)
	go updaterSvc.Start()
	s.closeFns = append(s.closeFns, newCloseFn("updater", updaterSvc.Stop))

	upnpSvc := upnp.NewUPnP(torManager)
	if err := upnpSvc.Run(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start UPnP server")
	}
	s.closeFns = append(s.closeFns, newCloseFn("upnp", upnpSvc.Close))

	httpHandler := handler.New(torManager, settingsStorage, apiClient, authService, upnpSvc)

	router := newChiRouter(settingsStorage, authService)
	httpHandler.Register(router)

	ln, listenPort, cleanUp, err := net.CreateListener()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create listener")
	}
	s.closeFns = append(s.closeFns, newCloseFn("listener", cleanUp))

	httpServer := &http.Server{
		Handler: router,
	}
	s.closeFns = append(s.closeFns, newCloseFn("httpServer", httpServer.Close))

	go func() {
		log.Info().Int("port", listenPort).Msg("Starting HTTP server")
		if err := httpServer.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}()

	return listenPort
}

func (s *Server) Close() {
	// Close resources in reverse order
	for i := len(s.closeFns) - 1; i >= 0; i-- {
		if err := s.closeFns[i].fn(); err != nil {
			log.Error().Err(err).Str("resource", s.closeFns[i].resourceName).Msg("Failed to close resource")
		}
	}
}

func newChiRouter(settingsStorage *setting.Storage, authService *auth.FirebaseAuth) *chi.Mux {
	r := chi.NewRouter()
	setCommonMiddlewares(r)

	r.Use(setting.Middleware(settingsStorage))
	r.Use(auth.Middleware(authService))
	r.Use(i18n.Middleware(i18n.NewBundle(), settingsStorage))

	r.Handle("/static/style/*", http.StripPrefix("/static/style", http.FileServer(http.FS(style.FS))))
	r.Handle("/static/assets/*", http.StripPrefix("/static/assets", http.FileServer(http.FS(data.Assets))))

	return r
}

func setCommonMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(logs.NewHTTPStructuredLogger(log.Logger))
	router.Use(middleware.Recoverer)

	router.Use(cors.AllowAll().Handler)
	router.Use(middleware.NoCache)
}
