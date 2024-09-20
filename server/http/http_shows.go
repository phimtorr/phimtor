package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s Server) GetLatestEpisodes(w http.ResponseWriter, r *http.Request, params GetLatestEpisodesParams) {
	episodes, pag, err := s.repo.GetLatestEpisodes(r.Context(), params)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"episodes":   episodes,
		"pagination": pag,
	})
}

func (s Server) GetLatestMovies(w http.ResponseWriter, r *http.Request, params GetLatestMoviesParams) {
	movies, pag, err := s.repo.GetLatestMovies(r.Context(), params)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"movies":     movies,
		"pagination": pag,
	})
}

func (s Server) GetLatestTvSeries(w http.ResponseWriter, r *http.Request, params GetLatestTvSeriesParams) {
	tvSeries, pag, err := s.repo.GetLatestTvSeries(r.Context(), params)
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
