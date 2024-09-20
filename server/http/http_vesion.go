package http

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/phimtorr/phimtor/server/version"
)

func (s Server) GetVersion(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"version": version.Version})
}
