package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/common/strval"
	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/client"
	"github.com/phimtorr/phimtor/desktop/handler"
	"github.com/phimtorr/phimtor/desktop/ui/style"
)

func main() {
	logs.Init(strval.MustBool(build.IsLocal))

	cl := client.NewClient()
	hl := handler.New(cl)

	r := chi.NewRouter()
	r.Use(logs.NewHTTPStructuredLogger(log.Logger))

	r.Handle("/static/style/*", http.StripPrefix("/static/style", http.FileServer(http.FS(style.FS))))

	hl.Register(r)

	addr := ":" + build.ServePort
	log.Info().Str("address", addr).Msg("Starting HTTP server")
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("Stopped HTTP server")
	}
}
