package handler

import (
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	"github.com/phimtorr/phimtor/desktop/server/ui"
)

func (h *Handler) GetMovie(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	resp, err := h.apiClient.GetMovieWithResponse(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get movie")
	}

	return ui.Movie(resp.JSON200.Movie).Render(r.Context(), w)
}
