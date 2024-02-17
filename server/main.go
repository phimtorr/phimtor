package main

import (
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/common/strval"

	"github.com/phimtorr/phimtor/server/pkg/database"
	"github.com/phimtorr/phimtor/server/ports"
)

func main() {
	logs.Init(strval.MustBool(os.Getenv("LOCAL_ENV")))

	db := database.NewSqlDB()

	httpServer := ports.NewHttpServer(db)

	e := echo.New()

	api := e.Group("/api/v1")

	api.Use(middleware.CORS())
	api.Use(middleware.RequestID())
	api.Use(lecho.Middleware(lecho.Config{
		Logger:      lecho.From(log.Logger),
		HandleError: true,
	}))

	ports.RegisterHandlers(api, httpServer)

	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot parse HTTP_PORT")
	}
	addr := ":" + strconv.Itoa(port)

	log.Info().Str("addr", addr).Msg("Starting server")
	e.Logger.Fatal(e.Start(addr))
}
