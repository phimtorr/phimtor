package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/server/admin/auth"
	adminHttp "github.com/phimtorr/phimtor/server/admin/http"
	"github.com/phimtorr/phimtor/server/admin/repository"
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

	//setAdmin(authClient)

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

func setAdmin(authClient *firebaseAuth.Client) {
	uid := os.Getenv("ADMIN_UID")
	user, err := authClient.GetUser(context.Background(), uid)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get user")
	}

	var claims map[string]interface{}
	if user.CustomClaims != nil {
		claims = user.CustomClaims
	} else {
		claims = make(map[string]interface{})
	}

	claims["is_admin"] = true

	if err := authClient.SetCustomUserClaims(context.Background(), uid, claims); err != nil {
		log.Fatal().Err(err).Msg("Failed to set admin")
	}
}

func syncAll(db *sql.DB) {
	r := repository.NewTMDBRepository(db)
	if err := r.SyncAll(context.Background()); err != nil {
		log.Error().Err(err).Msg("Failed to sync all")
	}
}
