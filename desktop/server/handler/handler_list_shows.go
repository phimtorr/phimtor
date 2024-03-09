package handler

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
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

func (h *Handler) SearchShows(w http.ResponseWriter, r *http.Request) {
	qPage := r.URL.Query().Get("page")
	query := r.URL.Query().Get("q")

	page := 1

	if qPage != "" {
		_page, err := strconv.Atoi(qPage)
		if err != nil {
			handleError(w, r, "Parse page", err, http.StatusBadRequest)
			return
		}
		page = _page
	}
	
	if query == "" {
		handleError(w, r, "Empty query", nil, http.StatusBadRequest)
		return
	}

	shows, pagination, err := h.apiClient.SearchShows(r.Context(), query, page)
	if err != nil {
		handleError(w, r, "Search shows", err, http.StatusInternalServerError)
		return
	}

	templ.Handler(ui.SearchPage(query, shows, pagination)).ServeHTTP(w, r)
}
