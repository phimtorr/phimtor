package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/common/strval"
	adminHttp "github.com/phimtorr/phimtor/server/admin/http"
	"github.com/phimtorr/phimtor/server/pkg/database"
	"github.com/rs/zerolog/log"
)

func main() {
	logs.Init(strval.MustBool(os.Getenv("LOCAL_ENV")))

	db := database.NewMySqlDB()
	httpServer := adminHttp.NewHTTPServer(db)

	r := chi.NewRouter()
	httpServer.Register(r)

	addr := ":" + os.Getenv("HTTP_PORT")
	log.Info().Str("address", addr).Msg("Starting HTTP server")
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("Stopped HTTP server")
	}
}
