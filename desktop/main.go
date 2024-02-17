package main

import (
	"github.com/labstack/echo/v4"

	"github.com/phimtorr/phimtor/common/logs"
	"github.com/phimtorr/phimtor/common/strval"
	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/client"
	"github.com/phimtorr/phimtor/desktop/handler"
)

func main() {
	logs.Init(strval.MustBool(build.IsLocal))

	cl := client.NewClient()
	hl := handler.New(cl)

	e := echo.New()

	e.Static("/static/style", "ui/style")

	hl.Register(e)

	e.Logger.Fatal(e.Start(":" + build.ServePort))

}
