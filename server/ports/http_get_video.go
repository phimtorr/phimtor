package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h HttpServer) GetVideo(w http.ResponseWriter, r *http.Request, id int64) {
	video, err := h.repo.GetVideo(r.Context(), id)
	if err != nil {
		handleError(w, r, "Failed to get video", err, http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"video": video,
	})
}
