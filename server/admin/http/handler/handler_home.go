package handler

import (
	"net/http"

	"github.com/phimtorr/phimtor/server/admin/http/uri"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) error {
	redirect(w, r, uri.ListShows(1))
	return nil
}
