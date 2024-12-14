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

func (h *TMDBHandler) ViewMovies(w http.ResponseWriter, r *http.Request) error {
	var page int
	if p := r.URL.Query().Get("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if page < 1 {
		page = 1
	}

	movies, pag, err := h.repo.ListMovies(r.Context(), page, pageSize)
	if err != nil {
		return fmt.Errorf("list movies: %w", err)
	}

	return ui.MoviesView(movies, pag).Render(r.Context(), w)
}

func (h *TMDBHandler) ViewMovie(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	movie, err := h.repo.GetMovie(r.Context(), id)
	if err != nil {
		return fmt.Errorf("get movie: %w", err)
	}

	return ui.MovieView(movie).Render(r.Context(), w)
}

func (h *TMDBHandler) CreateMovie(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("parse form: %w", err)
	}

	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-movie-id",
			fmt.Sprintf("invalid movie_id=%s, err=%v", r.Form.Get("id"), err))
	}

	movie, err := h.tmdbClient.GetMovieDetails(ctx, id)
	if err != nil {
		return fmt.Errorf("get movie details: %w", err)
	}

	err = h.repo.UpdateMovie(ctx, movie)
	if err != nil {
		return fmt.Errorf("update movie: %w", err)
	}

	redirect(w, r, uri.ViewMovie(int64(id)))
	return nil
}

func (h *TMDBHandler) FetchMovieFromTMDB(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return fmt.Errorf("parsing id: %w", err)
	}

	movie, err := h.tmdbClient.GetMovieDetails(ctx, int(id))
	if err != nil {
		return fmt.Errorf("get movie details: %w", err)
	}

	err = h.repo.UpdateMovie(ctx, movie)
	if err != nil {
		return fmt.Errorf("update movie: %w", err)
	}

	redirect(w, r, uri.ViewMovie(id))

	return nil
}

func (h *TMDBHandler) CreateMovieVideo(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return fmt.Errorf("parsing id: %w", err)
	}

	err = h.repo.CreateMovieVideo(ctx, id)
	if err != nil {
		return fmt.Errorf("create movie video: %w", err)
	}

	redirect(w, r, uri.ViewMovie(id))
	return nil
}

func (h *TMDBHandler) SyncMovie(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return fmt.Errorf("parsing id: %w", err)
	}

	err = h.repo.SyncMovie(ctx, id)
	if err != nil {
		return fmt.Errorf("sync movie: %w", err)
	}

	redirect(w, r, uri.ViewMovie(id))
	return nil
}

func (h *TMDBHandler) SyncYTSMovie(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return fmt.Errorf("parsing id: %w", err)
	}

	movie, err := h.repo.GetMovie(ctx, id)
	if err != nil {
		return fmt.Errorf("get movie: %w", err)
	}

	if movie.VideoID == 0 {
		return commonErrors.NewIncorrectInputError("no-video-id", "no video id")
	}

	ytsMovie, err := h.ytsClient.GetMovieByIMDbID(ctx, movie.IMDBID)
	if err != nil {
		return fmt.Errorf("get yts movie: %w", err)
	}

	err = h.repo.UpdateYTSMovie(ctx, movie.VideoID, ytsMovie)
	if err != nil {
		return fmt.Errorf("update yts movie: %w", err)
	}

	err = h.repo.SyncMovie(ctx, id)
	if err != nil {
		return fmt.Errorf("sync movie: %w", err)
	}

	redirect(w, r, uri.ViewMovie(id))
	return nil
}
