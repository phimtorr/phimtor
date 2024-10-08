package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/admin/http/uri"
)

func (h *TMDBHandler) ViewTVSeriesShows(w http.ResponseWriter, r *http.Request) error {
	var page int
	if p := r.URL.Query().Get("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if page < 1 {
		page = 1
	}

	shows, pag, err := h.repo.ListTVSeriesShows(r.Context(), page, pageSize)
	if err != nil {
		return fmt.Errorf("list tv series shows: %w", err)
	}

	return ui.TVSeriesShowsView(shows, pag).Render(r.Context(), w)
}

func (h *TMDBHandler) ViewTVSeriesShow(w http.ResponseWriter, r *http.Request) error {
	showID, err := parseID(chi.URLParam(r, "showID"))
	if err != nil {
		return err
	}

	show, seasons, err := h.repo.GetTVSeriesShow(r.Context(), showID)
	if err != nil {
		return fmt.Errorf("get tv series show: %w", err)
	}

	return ui.TVSeriesShowView(show, seasons).Render(r.Context(), w)
}

func (h *TMDBHandler) ViewTVSeason(w http.ResponseWriter, r *http.Request) error {
	showID, err := parseID(chi.URLParam(r, "showID"))
	if err != nil {
		return err
	}

	seasonNumber, err := parseID(chi.URLParam(r, "seasonNumber"))
	if err != nil {
		return err
	}

	season, episodes, err := h.repo.GetTVSeason(r.Context(), showID, int(seasonNumber))
	if err != nil {
		return fmt.Errorf("get tv season: %w", err)
	}

	return ui.TVSeasonView(season, episodes).Render(r.Context(), w)
}

func (h *TMDBHandler) ViewTVEpisode(w http.ResponseWriter, r *http.Request) error {
	showID, err := parseID(chi.URLParam(r, "showID"))
	if err != nil {
		return err
	}

	seasonNumber, err := parseID(chi.URLParam(r, "seasonNumber"))
	if err != nil {
		return err
	}

	episodeNumber, err := parseID(chi.URLParam(r, "episodeNumber"))
	if err != nil {
		return err
	}

	episode, err := h.repo.GetTVEpisode(r.Context(), showID, int(seasonNumber), int(episodeNumber))
	if err != nil {
		return fmt.Errorf("get tv episode: %w", err)
	}

	return ui.TVEpisodeView(episode).Render(r.Context(), w)
}

func (h *TMDBHandler) CreateTVSeries(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("parsing form: %w", err)
	}

	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-tv-id",
			fmt.Sprintf("invalid tv_id=%s, err=%v", r.Form.Get("id"), err))
	}

	tv, seasons, err := h.tmdbClient.GetTVSeriesDetails(ctx, id)
	if err != nil {
		return fmt.Errorf("get tv series details: %w", err)
	}

	err = h.repo.UpdateTVSeries(ctx, tv, seasons)
	if err != nil {
		return fmt.Errorf("update tv series: %w", err)
	}

	redirect(w, r, uri.ViewTVSeriesShow(int64(id)))
	return nil
}

func (h *TMDBHandler) FetchTVSeriesFromTMDB(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "showID"))
	if err != nil {
		return fmt.Errorf("parsing id: %w", err)
	}

	tv, seasons, err := h.tmdbClient.GetTVSeriesDetails(ctx, int(id))
	if err != nil {
		return fmt.Errorf("get tv series details: %w", err)
	}

	err = h.repo.UpdateTVSeries(ctx, tv, seasons)
	if err != nil {
		return fmt.Errorf("update tv series: %w", err)
	}

	redirect(w, r, uri.ViewTVSeriesShow(id))
	return nil
}

func (h *TMDBHandler) CreateTVEpisodeVideo(w http.ResponseWriter, r *http.Request) error {
	showID, err := parseID(chi.URLParam(r, "showID"))
	if err != nil {
		return err
	}

	seasonNumber, err := parseID(chi.URLParam(r, "seasonNumber"))
	if err != nil {
		return err
	}

	episodeNumber, err := parseID(chi.URLParam(r, "episodeNumber"))
	if err != nil {
		return err
	}

	err = h.repo.CreateTVEpisodeVideo(r.Context(), showID, int(seasonNumber), int(episodeNumber))
	if err != nil {
		return fmt.Errorf("create tv episode video: %w", err)
	}

	redirect(w, r, uri.ViewTVEpisode(showID, int(seasonNumber), int(episodeNumber)))
	return nil
}

func (h *TMDBHandler) SyncTVSeries(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "showID"))
	if err != nil {
		return err
	}

	err = h.repo.SyncTVSeries(ctx, id)
	if err != nil {
		return fmt.Errorf("sync tv series: %w", err)
	}

	redirect(w, r, uri.ViewTVSeriesShow(id))
	return nil
}
