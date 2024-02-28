package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/phimtorr/phimtor/desktop/server/ui"
)

func (h *Handler) GetSeries(w http.ResponseWriter, r *http.Request, id int64) {
	resp, err := h.apiClient.GetSeriesWithResponse(r.Context(), id)
	if err != nil {
		handleError(w, r, "Get series", err, http.StatusInternalServerError)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		handleError(w, r, "Get series", fmt.Errorf("http_status=%d", resp.StatusCode()), resp.StatusCode())
		return
	}

	templ.Handler(ui.Series(resp.JSON200.Series)).ServeHTTP(w, r)
}
