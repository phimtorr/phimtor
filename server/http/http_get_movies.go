package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h HttpServer) GetMovie(w http.ResponseWriter, r *http.Request, id int64) {
	movie, err := h.repo.GetMovie(r.Context(), id)
	if err != nil {
		respondError(w, r, err)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"movie": movie,
	})
}
