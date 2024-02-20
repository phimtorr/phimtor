package ports

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h HttpServer) GetSeries(w http.ResponseWriter, r *http.Request, id int64) {
	series, err := h.repo.GetSeries(r.Context(), id)
	if err != nil {
		handleError(w, r, "Get series", err, http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"series": series,
	})
}
