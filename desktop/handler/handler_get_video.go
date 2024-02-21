package handler

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/phimtorr/phimtor/desktop/ui"
)

func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request, id int64) {
	resp, err := h.client.GetVideoWithResponse(r.Context(), id)
	if err != nil {
		handleError(w, r, "Failed to get video", err, http.StatusInternalServerError)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		handleError(w, r, "Failed to get video", err, resp.StatusCode())
		return
	}

	templ.Handler(ui.Video(resp.JSON200.Video)).ServeHTTP(w, r)
}
