package handler2

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"

	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/admin/http/uri"
)

func (h *Handler) ViewMovies(w http.ResponseWriter, r *http.Request) error {
	var page int
	if p := r.URL.Query().Get("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if page < 1 {
		page = 1
	}

	movies, pag, err := h.repo.ListMovies(r.Context(), page, pageSize)
	if err != nil {
		return errors.Wrap(err, "list movies")
	}

	return ui.MoviesView(movies, pag).Render(r.Context(), w)
}

func (h *Handler) ViewMovie(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	movie, err := h.repo.GetMovie(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get movie")
	}

	return ui.MovieView(movie).Render(r.Context(), w)
}

func (h *Handler) CreateMovie(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		return errors.Wrap(err, "parsing form")
	}

	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-movie-id",
			fmt.Sprintf("invalid movie_id=%s, err=%v", r.Form.Get("id"), err))
	}

	movie, err := h.tmdbClient.GetMovieDetails(ctx, id)
	if err != nil {
		return errors.Wrap(err, "get movie details")
	}

	err = h.repo.UpdateMovie(ctx, movie)
	if err != nil {
		return errors.Wrap(err, "update movie")
	}

	redirect(w, r, uri.ViewMovie(int64(id)))
	return nil
}

func (h *Handler) FetchMovieFromTMDB(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return errors.Wrap(err, "parsing id")
	}

	movie, err := h.tmdbClient.GetMovieDetails(ctx, int(id))
	if err != nil {
		return errors.Wrap(err, "get movie details")
	}

	err = h.repo.UpdateMovie(ctx, movie)
	if err != nil {
		return errors.Wrap(err, "update movie")
	}

	redirect(w, r, uri.ViewMovie(id))

	return nil
}

func (h *Handler) CreateMovieVideo(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return errors.Wrap(err, "parsing id")
	}

	err = h.repo.CreateMovieVideo(ctx, id)
	if err != nil {
		return errors.Wrap(err, "create movie video")
	}

	redirect(w, r, uri.ViewMovie(id))
	return nil
}

func (h *Handler) SyncMovie(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return errors.Wrap(err, "parsing id")
	}

	err = h.repo.SyncMovie(ctx, id)
	if err != nil {
		return errors.Wrap(err, "sync movie")
	}

	redirect(w, r, uri.ViewMovie(id))
	return nil
}
