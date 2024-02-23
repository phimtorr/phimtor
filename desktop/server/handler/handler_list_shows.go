package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/ui"
)

func (h *Handler) ListShows(w http.ResponseWriter, r *http.Request, params api.ListShowsParams) {
	resp, err := h.apiClient.ListShowsWithResponse(r.Context(), &params)
	if err != nil {
		handleError(w, r, "List shows", err, http.StatusInternalServerError)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		handleError(w, r, "List shows", fmt.Errorf("http error=%d", resp.StatusCode()), resp.StatusCode())
		return
	}
	templ.Handler(ui.Shows(resp.JSON200.Shows, resp.JSON200.Pagination)).ServeHTTP(w, r)
}
