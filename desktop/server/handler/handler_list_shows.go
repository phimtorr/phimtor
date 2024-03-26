package handler

import (
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	commonErrors "github.com/phimtorr/phimtor/common/errors"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
)

func (h *Handler) ListShows(w http.ResponseWriter, r *http.Request) error {
	qType := r.URL.Query().Get("type")

	page, err := parsePage(r)
	if err != nil {
		return err
	}

	pageSize, err := parsePageSize(r)
	if err != nil {
		return err
	}

	showType := api.ShowType(qType)

	shows, pagination, err := h.apiClient.ListShows(r.Context(), page, pageSize, showType)
	if err != nil {
		return errors.Wrap(err, "list shows")
	}

	return ui.Shows(shows, pagination, showType).Render(r.Context(), w)
}

func (h *Handler) SearchShows(w http.ResponseWriter, r *http.Request) error {
	page, err := parsePage(r)
	if err != nil {
		return err
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		return commonErrors.NewIncorrectInputError("empty-query", "empty query")
	}

	shows, pagination, err := h.apiClient.SearchShows(r.Context(), query, page)
	if err != nil {
		return errors.Wrap(err, "search shows")
	}

	return ui.SearchPage(query, shows, pagination).Render(r.Context(), w)
}

var (
	ErrInvalidPage = commonErrors.NewIncorrectInputError("invalid-page", "invalid page")
)

func parsePage(r *http.Request) (int, error) {
	qPage := r.URL.Query().Get("page")
	if qPage == "" {
		return 1, nil
	}
	page, err := strconv.Atoi(qPage)
	if err != nil {
		return 0, errors.Wrapf(ErrInvalidPage, "parse page=%s, err=%v", qPage, err)
	}
	return page, nil
}

var (
	ErrInvalidPageSize = commonErrors.NewIncorrectInputError("invalid-page-size", "invalid page size")
)

func parsePageSize(r *http.Request) (int, error) {
	qPageSize := r.URL.Query().Get("pageSize")
	if qPageSize == "" {
		return 1, nil
	}
	pageSize, err := strconv.Atoi(qPageSize)
	if err != nil {
		return 0, errors.Wrapf(ErrInvalidPageSize, "parse page_size=%s, err=%v", qPageSize, err)
	}
	return pageSize, nil
}
