package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s Server) ListRecentlyAddedMovies(w http.ResponseWriter, r *http.Request, params ListRecentlyAddedMoviesParams) {
	movies, pag, err := s.repo.ListRecentlyAddedMovies(r.Context(), params)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"movies":     movies,
		"pagination": pag,
	})
}

func (s Server) ListLatestMovies(w http.ResponseWriter, r *http.Request, params ListLatestMoviesParams) {
	movies, pag, err := s.repo.ListLatestMovies(r.Context(), params)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"movies":     movies,
		"pagination": pag,
	})
}

func (s Server) ListLatestEpisodes(w http.ResponseWriter, r *http.Request, params ListLatestEpisodesParams) {
	episodes, pag, err := s.repo.ListLatestEpisodes(r.Context(), params)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"episodes":   episodes,
		"pagination": pag,
	})
}

func (s Server) ListLatestTvSeries(w http.ResponseWriter, r *http.Request, params ListLatestTvSeriesParams) {
	tvSeries, pag, err := s.repo.ListLatestTvSeries(r.Context(), params)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"tvSeries":   tvSeries,
		"pagination": pag,
	})
}

func (s Server) SearchShows(w http.ResponseWriter, r *http.Request, params SearchShowsParams) {
	shows, pagination, err := s.repo.SearchShows(r.Context(), params)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"shows":      shows,
		"pagination": pagination,
	})
}
