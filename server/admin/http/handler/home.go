package handler

import (
	"net/http"

	"github.com/phimtorr/phimtor/server/admin/http/uri"
)

func Home(w http.ResponseWriter, r *http.Request) error {
	redirect(w, r, uri.ListLatestShows(1))
	return nil
}
