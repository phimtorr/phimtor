package main

import (
	"fmt"

	"github.com/YOMIkio/lorca"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/server"
)

func main() {
	logs.Init()

	aServer := server.NewServer(build.AppName)
	defer aServer.Close()

	listenPort := aServer.Start()

	runUI(listenPort)
}

func runUI(servePort int) {
	ui, err := lorca.New(fmt.Sprintf("http://localhost:%d", servePort), "", 1280, 800)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create UI")
	}
	defer func() {
		if err := ui.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close UI")
		}
	}()

	<-ui.Done()
}
