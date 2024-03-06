package handler

import (
	"github.com/a-h/templ"
	"github.com/phimtorr/phimtor/desktop/server/handler/uri"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		templ.Handler(ui.Login()).ServeHTTP(w, r)
		return
	} else {
		if err := r.ParseForm(); err != nil {
			handleError(w, r, "Parse form", err, http.StatusBadRequest)
			return
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		log.Ctx(r.Context()).Debug().
			Str("email", email).
			Str("password", password).
			Msg("Login")

		redirect(w, r, uri.Home())
	}
}
