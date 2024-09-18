package http2

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s Server) GetTvSeries(w http.ResponseWriter, r *http.Request, tvSeriesId int64) {
	tvSeries, err := s.repo.GetTvSeries(r.Context(), tvSeriesId)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"tvSeries": tvSeries,
	})
}

func (s Server) GetTvSeason(w http.ResponseWriter, r *http.Request, tvSeriesId int64, seasonNumber int) {
	season, err := s.repo.GetTvSeason(r.Context(), tvSeriesId, seasonNumber)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"tvSeason": season,
	})
}
