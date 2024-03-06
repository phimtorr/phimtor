package handler

import (
	"github.com/a-h/templ"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/server/uri"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		templ.Handler(ui.SignIn()).ServeHTTP(w, r)
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

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	templ.Handler(ui.SignUp()).ServeHTTP(w, r)
}
