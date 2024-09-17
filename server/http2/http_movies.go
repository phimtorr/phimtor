package http2

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s Server) GetMovie(w http.ResponseWriter, r *http.Request, movieId int64) {
	movie, err := s.repo.GetMovie(r.Context(), movieId)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"movie": movie,
	})
}
