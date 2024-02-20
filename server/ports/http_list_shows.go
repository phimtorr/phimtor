package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h HttpServer) ListShows(w http.ResponseWriter, r *http.Request, params ListShowsParams) {
	shows, pagination, err := h.repo.ListShows(r.Context(), params)
	if err != nil {
		handleError(w, r, "List shows", err, http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"shows":      shows,
		"pagination": pagination,
	})
}
