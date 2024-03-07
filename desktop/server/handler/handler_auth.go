package handler

import (
	"github.com/a-h/templ"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/server/uri"
	"net/http"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		handleError(w, r, "Parse form", err, http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if err := h.authService.SignIn(r.Context(), email, password); err != nil {
		handleError(w, r, "Sign in", err, http.StatusInternalServerError)
		return
	}

	redirect(w, r, uri.Home())
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	templ.Handler(ui.SignUp()).ServeHTTP(w, r)
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	h.authService.SignOut()
	redirect(w, r, uri.Home())
}
