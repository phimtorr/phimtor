package handler

import (
	"context"
	"net/http"

	"github.com/friendsofgo/errors"
	"golang.org/x/sync/errgroup"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) error {
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
		return errors.Wrap(err, "fetch data")
	}

	return ui.Home(movies, series).Render(r.Context(), w)
}
