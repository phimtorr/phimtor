package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	commonErrors "github.com/phimtorr/phimtor/common/errors"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
)

func (h *Handler) ListShows(w http.ResponseWriter, r *http.Request) error {
	qPage := r.URL.Query().Get("page")
	qPageSize := r.URL.Query().Get("pageSize")
	qType := r.URL.Query().Get("type")

	page, err := strconv.Atoi(qPage)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-page", fmt.Sprintf("invalid page=%s, err=%v", qPage, err))
	}
	pageSize, err := strconv.Atoi(qPageSize)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-page-size", fmt.Sprintf("invalid page_size=%s, err=%v", qPageSize, err))
	}

	showType := api.ShowType(qType)

	shows, pagination, err := h.apiClient.ListShows(r.Context(), page, pageSize, showType)
	if err != nil {
		return errors.Wrap(err, "list shows")
	}

	return ui.Shows(shows, pagination, showType).Render(r.Context(), w)
}

func (h *Handler) SearchShows(w http.ResponseWriter, r *http.Request) error {
	qPage := r.URL.Query().Get("page")
	query := r.URL.Query().Get("q")

	page := 1

	if qPage != "" {
		_page, err := strconv.Atoi(qPage)
		if err != nil {
			return commonErrors.NewIncorrectInputError("invalid-page", fmt.Sprintf("invalid page=%s, err=%v", qPage, err))
		}
		page = _page
	}

	if query == "" {
		return commonErrors.NewIncorrectInputError("empty-query", "empty query")
	}

	shows, pagination, err := h.apiClient.SearchShows(r.Context(), query, page)
	if err != nil {
		return errors.Wrap(err, "search shows")
	}

	return ui.SearchPage(query, shows, pagination).Render(r.Context(), w)
}
