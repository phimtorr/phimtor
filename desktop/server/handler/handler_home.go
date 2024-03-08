package handler

import (
	"context"
	"github.com/a-h/templ"
	"golang.org/x/sync/errgroup"
	"net/http"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	errGrp, ctx := errgroup.WithContext(ctx)

	var movies []api.Show
	var series []api.Show

	errGrp.Go(func() error {
		shows, _, err := h.apiClient.ListShows(ctx, 1, 6, api.ShowTypeMovie)
		if err != nil {
			return err
		}

		movies = shows
		return nil
	})

	errGrp.Go(func() error {
		shows, _, err := h.apiClient.ListShows(ctx, 1, 6, api.ShowTypeSeries)
		if err != nil {
			return err
		}

		series = shows
		return nil
	})

	if err := errGrp.Wait(); err != nil {
		handleError(w, r, "Error fetch data", err, http.StatusInternalServerError)
		return
	}

	templ.Handler(ui.Home(movies, series)).ServeHTTP(w, r)
}
