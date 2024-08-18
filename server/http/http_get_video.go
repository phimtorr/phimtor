package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h Server) GetVideo(w http.ResponseWriter, r *http.Request, id int64) {
	video, err := h.repo.GetVideo(r.Context(), id)
	if err != nil {
		respondError(w, r, err)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"video": video,
	})
}
