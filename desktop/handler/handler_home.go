package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/phimtorr/phimtor/desktop/handler/uri"
)

func (h *Handler) Home(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, uri.ListShows())
}
