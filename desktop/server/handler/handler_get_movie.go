package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"

	"github.com/phimtorr/phimtor/desktop/server/ui"
)

func (h *Handler) GetMovie(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-id", fmt.Sprintf("invalid id=%s, err=%v", idStr, err))
	}

	resp, err := h.apiClient.GetMovieWithResponse(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get movie")
	}

	return ui.Movie(resp.JSON200.Movie).Render(r.Context(), w)
}
