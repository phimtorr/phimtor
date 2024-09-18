package handler2

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/phimtorr/phimtor/server/admin/http/ui"
)

func (h *Handler) ListLatestShows(w http.ResponseWriter, r *http.Request) error {
	var page int
	if p := r.URL.Query().Get("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if page < 1 {
		page = 1
	}

	shows, pagination, err := h.repo.ListLatestShows(r.Context(), page, pageSize)
	if err != nil {
		return fmt.Errorf("list latest shows: %w", err)
	}

	return ui.LatestShowsView(shows, pagination).Render(r.Context(), w)
}
