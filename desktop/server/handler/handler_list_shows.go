package handler

import (
	"github.com/a-h/templ"
	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"net/http"
	"strconv"
)

func (h *Handler) ListShows(w http.ResponseWriter, r *http.Request) {
	qPage := r.URL.Query().Get("page")
	qPageSize := r.URL.Query().Get("pageSize")
	qType := r.URL.Query().Get("type")

	page, err := strconv.Atoi(qPage)
	if err != nil {
		handleError(w, r, "Parse page", err, http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(qPageSize)
	if err != nil {
		handleError(w, r, "Parse pageSize", err, http.StatusBadRequest)
		return
	}

	showType := api.ShowType(qType)

	shows, pagination, err := h.apiClient.ListShows(r.Context(), page, pageSize, showType)
	if err != nil {
		handleError(w, r, "List shows", err, http.StatusInternalServerError)
		return
	}

	templ.Handler(ui.Shows(shows, pagination, showType)).ServeHTTP(w, r)
}
