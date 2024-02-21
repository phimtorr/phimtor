package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/common/strval"
	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/client"
	"github.com/phimtorr/phimtor/desktop/handler"
	"github.com/phimtorr/phimtor/desktop/torrent"
	"github.com/phimtorr/phimtor/desktop/ui/style"
)

func main() {
	logs.Init(strval.MustBool(build.IsLocal))

	// TODO: change the path from settings
	dataDir := "/Users/lap14897/Downloads/PhimTor"
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.Mkdir(dataDir, 0700); err != nil {
			panic(fmt.Errorf("create data directory: %w", err))
		}
	}

	torManager := torrent.NewManager(dataDir)
	defer func() {
		if err := torManager.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close torrent manager")
		}
	}()

	cl := client.NewClient()
	hl := handler.New(torManager, cl)

	r := chi.NewRouter()
	setMiddlewares(r)

	r.Handle("/static/style/*", http.StripPrefix("/static/style", http.FileServer(http.FS(style.FS))))

	hl.Register(r)

	addr := ":" + build.ServePort
	log.Info().Str("address", addr).Msg("Starting HTTP server")
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("Stopped HTTP server")
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(logs.NewHTTPStructuredLogger(log.Logger))
	router.Use(middleware.Recoverer)

	router.Use(cors.AllowAll().Handler)
	router.Use(middleware.NoCache)
}
