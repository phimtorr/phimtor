package handler

import (
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"

	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/admin/http/uri"
)

func (h *Handler) ListEpisodes(w http.ResponseWriter, r *http.Request) error {
	showID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	episodes, err := h.repo.ListEpisodes(r.Context(), showID)
	if err != nil {
		return errors.Wrap(err, "list episodes")
	}

	return ui.Episodes(showID, episodes).Render(r.Context(), w)
}

func (h *Handler) ViewCreateEpisodeForm(w http.ResponseWriter, r *http.Request) error {
	showID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	return ui.CreateEpisodeForm(showID).Render(r.Context(), w)
}

func (h *Handler) CreateEpisode(w http.ResponseWriter, r *http.Request) error {
	showID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	if err := r.ParseForm(); err != nil {
		return errors.Wrap(err, "parsing form")
	}

	name := r.Form.Get("name")
	if name == "" {
		return commonErrors.NewIncorrectInputError("empty-name", "empty name")
	}

	_, err = h.repo.CreateEpisode(r.Context(), EpisodeToCreate{
		ShowID: showID,
		Name:   name,
	})
	if err != nil {
		return errors.Wrap(err, "create episode")
	}

	redirect(w, r, uri.ViewShow(showID))
	return nil
}
