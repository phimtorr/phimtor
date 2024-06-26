package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h HttpServer) GetSeries(w http.ResponseWriter, r *http.Request, id int64) {
	series, err := h.repo.GetSeries(r.Context(), id)
	if err != nil {
		respondError(w, r, err)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"series": series,
	})
}
