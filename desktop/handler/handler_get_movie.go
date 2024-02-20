package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"

	"github.com/phimtorr/phimtor/desktop/ui"
)

func (h *Handler) GetMovie(w http.ResponseWriter, r *http.Request, id int64) {
	resp, err := h.client.GetMovieWithResponse(r.Context(), id)
	if err != nil {
		handleError(w, r, "Get movie", err, http.StatusInternalServerError)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		handleError(w, r, "Get movie", fmt.Errorf("http error=%d", resp.StatusCode()), resp.StatusCode())
		return
	}

	templ.Handler(ui.Movie(resp.JSON200.Movie)).ServeHTTP(w, r)
}
