package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/phimtorr/phimtor/desktop/ui"
)

func (h *Handler) ListShows(c echo.Context) error {
	resp, err := h.client.ListShowsWithResponse(c.Request().Context(), nil)
	if err != nil {
		return err
	}
	return render(c, http.StatusOK, ui.Shows(resp.JSON200.Shows, resp.JSON200.Pagination))
}
