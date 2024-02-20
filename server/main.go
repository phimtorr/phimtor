package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/common/strval"
	serverLogs "github.com/phimtorr/phimtor/server/pkg/server/logs"
	"github.com/phimtorr/phimtor/server/ports"
	"github.com/phimtorr/phimtor/server/repository"

	"github.com/phimtorr/phimtor/server/pkg/database"
)

func main() {
	logs.Init(strval.MustBool(os.Getenv("LOCAL_ENV")))

	db := database.NewMySqlDB()
	repo := repository.NewRepository(db)
	httpServer := ports.NewHttpServer(repo)

	r := chi.NewRouter()
	setMiddlewares(r)

	handler := ports.HandlerFromMuxWithBaseURL(httpServer, r, "/api/v1")

	addr := ":" + os.Getenv("HTTP_PORT")

	log.Info().Str("address", addr).Msg("Starting HTTP server")
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal().Err(err).Msg("Stopped HTTP server")
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(serverLogs.NewHTTPStructuredLogger(log.Logger))
	router.Use(middleware.Recoverer)

	router.Use(cors.AllowAll().Handler)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}
