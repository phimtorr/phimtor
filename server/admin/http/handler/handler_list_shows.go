package handler

import (
	"net/http"
	"strconv"

	"github.com/phimtorr/phimtor/server/admin/http/ui"

	"github.com/friendsofgo/errors"
)

const (
	pageSize = 10
)

func (h *Handler) ListShows(w http.ResponseWriter, r *http.Request) error {
	var page int
	if p := r.URL.Query().Get("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if page < 1 {
		page = 1
	}

	shows, pag, err := h.repo.ListShowDisplays(r.Context(), page, pageSize)
	if err != nil {
		return errors.Wrap(err, "list shows")
	}

	return ui.Shows(shows, pag).Render(r.Context(), w)
}
