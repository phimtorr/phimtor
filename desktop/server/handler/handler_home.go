package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	resp, err := h.apiClient.ListShowsWithResponse(r.Context(), &api.ListShowsParams{})
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
