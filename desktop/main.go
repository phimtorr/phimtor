package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/common/strval"
	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/client"
	"github.com/phimtorr/phimtor/desktop/handler"
	"github.com/phimtorr/phimtor/desktop/i18n"
	"github.com/phimtorr/phimtor/desktop/setting"
	"github.com/phimtorr/phimtor/desktop/torrent"
	"github.com/phimtorr/phimtor/desktop/ui/style"
)

func main() {
	logs.Init(strval.MustBool(build.IsLocal))

	settingsStorage := setting.NewStorage("PhimTor")
	defer func() {
		if err := cleanUpStorage(settingsStorage.GetSettings()); err != nil {
			log.Error().Err(err).Msg("Failed to clean up storage")
		}
	}()

	settings := settingsStorage.GetSettings()

	torManager := torrent.NewManager(settings.GetCurrentDataDir())
	defer func() {
		if err := torManager.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close torrent manager")
		}
	}()

	apiClient := client.NewClient()
	httpHandler := handler.New(torManager, settingsStorage, apiClient)

	r := chi.NewRouter()
	setCommonMiddlewares(r)

	r.Use(setting.Middleware(settingsStorage))
	r.Use(i18n.Middleware(i18n.NewBundle(), settingsStorage))

	r.Handle("/static/style/*", http.StripPrefix("/static/style", http.FileServer(http.FS(style.FS))))

	httpHandler.Register(r)

	addr := ":" + build.ServePort
	log.Info().Str("address", addr).Msg("Starting HTTP server")
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("Stopped HTTP server")
	}
}

func setCommonMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(logs.NewHTTPStructuredLogger(log.Logger))
	router.Use(middleware.Recoverer)

	router.Use(cors.AllowAll().Handler)
	router.Use(middleware.NoCache)
}

func cleanUpStorage(setting setting.Settings) error {
	if setting.GetDeleteAfterClosed() {
		dif, err := os.ReadDir(setting.GetCurrentDataDir())
		if err != nil {
			return fmt.Errorf("read data directory: %w", err)
		}
		var errs []error
		for _, f := range dif {
			if err := os.RemoveAll(filepath.Join(setting.GetCurrentDataDir(), f.Name())); err != nil {
				errs = append(errs, fmt.Errorf("remove %s: %w", f.Name(), err))
			}
		}
		return errors.Join(errs...)
	}
	return nil
}
