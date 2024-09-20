package main

import (
	"context"
	stdHttp "net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/server/auth"
	"github.com/phimtorr/phimtor/server/http2"
	"github.com/phimtorr/phimtor/server/migrations"
	"github.com/phimtorr/phimtor/server/pkg/database"
	"github.com/phimtorr/phimtor/server/repository"
)

func main() {
	logs.Init()

	firebaseApp := newFirebaseApp()

	db := database.NewMySqlDB()
	if err := migrations.Run(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to run migrations")
	}

	repo2 := repository.NewSQLRepo2(db)
	http2Server := http2.NewHttpServer(repo2)

	r := chi.NewRouter()
	setMiddlewares(r)
	setAuthMiddleware(r, firebaseApp)

	r.Get("/api/v1/*", http2Server.Unsupported)

	http2.HandlerFromMuxWithBaseURL(http2Server, r, "/api/v2")

	addr := ":" + os.Getenv("HTTP_PORT")

	log.Info().Str("address", addr).Msg("Starting HTTP server")

	if err := stdHttp.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("Stopped HTTP server")
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewHTTPStructuredLogger(log.Logger))
	router.Use(middleware.Recoverer)

	router.Use(cors.AllowAll().Handler)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func setAuthMiddleware(router *chi.Mux, firebaseApp *firebase.App) {
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create firebase auth client")
	}

	router.Use(auth.NewFirebaseHttpMiddleware(authClient).Middleware)
}

func newFirebaseApp() *firebase.App {
	// must set GOOGLE_APPLICATION_CREDENTIALS
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create firebase app")
	}
	return app
}
