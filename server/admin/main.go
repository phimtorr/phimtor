package main

import (
	"context"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/server/admin/auth"
	adminHttp "github.com/phimtorr/phimtor/server/admin/http"
	"github.com/phimtorr/phimtor/server/pkg/database"
)

func init() {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	time.Local = loc
}

func main() {
	logs.Init()

	firebaseApp := newFirebaseApp()
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create auth client")
	}

	authService := auth.NewAuth(authClient)

	db := database.NewMySqlDB()
	httpServer := adminHttp.NewHTTPServer(db, authClient)

	r := chi.NewRouter()

	r.Use(authService.Middleware)

	authService.Register(r)
	httpServer.Register(r)

	addr := ":" + os.Getenv("HTTP_PORT")
	log.Info().Str("address", addr).Msg("Starting HTTP server")
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("Stopped HTTP server")
	}
}

func newFirebaseApp() *firebase.App {
	// must set GOOGLE_APPLICATION_CREDENTIALS
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create firebase app")
	}
	return app
}
