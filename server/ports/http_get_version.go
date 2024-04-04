package ports

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/phimtorr/phimtor/server/version"
)

func (h HttpServer) GetVersion(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"version": version.Version})
}
