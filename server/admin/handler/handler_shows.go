package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/phimtorr/phimtor/server/admin/uri"

	"github.com/go-chi/chi/v5"
	"github.com/phimtorr/phimtor/server/admin/ui"

	"github.com/friendsofgo/errors"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
)

func (h *Handler) CreateShow(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.Wrap(err, "parsing form")
	}

	showType := r.Form.Get("showType")
	if showType == "" {
		return commonErrors.NewIncorrectInputError("empty-show-type", "empty show type")
	}

	title := r.Form.Get("title")
	if title == "" {
		return commonErrors.NewIncorrectInputError("empty-title", "empty title")

	}

	originalTitle := r.Form.Get("originalTitle")
	if originalTitle == "" {
		return commonErrors.NewIncorrectInputError("empty-original-title", "empty original title")
	}

	posterLink := r.Form.Get("posterLink")
	if posterLink == "" {
		return commonErrors.NewIncorrectInputError("empty-poster-link", "empty poster link")
	}

	description := r.Form.Get("description")
	if description == "" {
		return commonErrors.NewIncorrectInputError("empty-description", "empty description")
	}

	releaseYear, err := strconv.Atoi(r.Form.Get("releaseYear"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-release-year",
			fmt.Sprintf("invalid release_year=%s, err=%v", r.Form.Get("releaseYear"), err))
	}

	score, err := strconv.ParseFloat(r.Form.Get("score"), 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-score",
			fmt.Sprintf("invalid score=%s, err=%v", r.Form.Get("score"), err))
	}

	durationInMinutes, err := strconv.Atoi(r.Form.Get("durationInMinutes"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-duration-in-minutes",
			fmt.Sprintf("invalid duration_in_minutes=%s, err=%v", r.Form.Get("durationInMinutes"), err))

	}

	quality := r.Form.Get("quality")

	var totalEpisodes int
	if showType == "series" {
		totalEpisodes, err = strconv.Atoi(r.Form.Get("totalEpisodes"))
		if err != nil {
			return commonErrors.NewIncorrectInputError("invalid-total-episodes",
				fmt.Sprintf("invalid total_episodes=%s, err=%v", r.Form.Get("totalEpisodes"), err))
		}
	}

	id, err := h.repo.CreateShow(r.Context(), ShowToCreate{
		ShowType:          showType,
		Title:             title,
		OriginalTitle:     originalTitle,
		PosterLink:        posterLink,
		Description:       description,
		ReleaseYear:       releaseYear,
		Score:             score,
		DurationInMinutes: durationInMinutes,
		Quality:           quality,
		TotalEpisodes:     totalEpisodes,
	})
	if err != nil {
		return errors.Wrap(err, "creating show")
	}

	redirect(w, r, uri.ViewShow(id))

	return nil
}

func (h *Handler) ViewShow(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-id", err.Error())
	}

	show, err := h.repo.GetShow(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "getting show")
	}

	return ui.ViewShow(show).Render(r.Context(), w)
}

func (h *Handler) ViewUpdateShowForm(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-id", err.Error())
	}

	show, err := h.repo.GetShow(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "getting show")
	}

	return ui.UpdateShowForm(show).Render(r.Context(), w)
}

func (h *Handler) UpdateShow(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.Wrap(err, "parsing form")
	}

	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-id", err.Error())
	}

	title := r.Form.Get("title")
	if title == "" {
		return commonErrors.NewIncorrectInputError("empty-title", "empty title")
	}

	originalTitle := r.Form.Get("originalTitle")
	if originalTitle == "" {
		return commonErrors.NewIncorrectInputError("empty-original-title", "empty original title")
	}

	posterLink := r.Form.Get("posterLink")
	if posterLink == "" {
		return commonErrors.NewIncorrectInputError("empty-poster-link", "empty poster link")
	}

	description := r.Form.Get("description")
	if description == "" {
		return commonErrors.NewIncorrectInputError("empty-description", "empty description")
	}

	releaseYear, err := strconv.Atoi(r.Form.Get("releaseYear"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-release-year",
			fmt.Sprintf("invalid release_year=%s, err=%v", r.Form.Get("releaseYear"), err))
	}

	score, err := strconv.ParseFloat(r.Form.Get("score"), 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-score",
			fmt.Sprintf("invalid score=%s, err=%v", r.Form.Get("score"), err))
	}

	durationInMinutes, err := strconv.Atoi(r.Form.Get("durationInMinutes"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-duration-in-minutes",
			fmt.Sprintf("invalid duration_in_minutes=%s, err=%v", r.Form.Get("durationInMinutes"), err))
	}

	quality := r.Form.Get("quality")

	var totalEpisodes int
	showType := r.Form.Get("showType")
	if showType == "series" {
		totalEpisodes, err = strconv.Atoi(r.Form.Get("totalEpisodes"))
		if err != nil {
			return commonErrors.NewIncorrectInputError("invalid-total-episodes",
				fmt.Sprintf("invalid total_episodes=%s, err=%v", r.Form.Get("totalEpisodes"), err))
		}
	}

	err = h.repo.UpdateShow(r.Context(), ShowToUpdate{
		ID:                id,
		Title:             title,
		OriginalTitle:     originalTitle,
		PosterLink:        posterLink,
		Description:       description,
		ReleaseYear:       releaseYear,
		Score:             score,
		DurationInMinutes: durationInMinutes,
		Quality:           quality,
		TotalEpisodes:     totalEpisodes,
	})

	if err != nil {
		return errors.Wrap(err, "updating show")
	}

	redirect(w, r, uri.ViewShow(id))
	return nil
}

func parseID(idRaw string) (int64, error) {
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "parsing id")
	}
	return id, nil
}
