package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h HttpServer) GetMovie(w http.ResponseWriter, r *http.Request, id int64) {
	movie, err := h.repo.GetMovie(r.Context(), id)
	if err != nil {
		handleError(w, r, "Get movie", err, http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"movie": movie,
	})
}
