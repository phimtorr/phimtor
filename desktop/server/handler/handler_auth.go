package handler

import (
	"github.com/friendsofgo/errors"
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
	if err := r.ParseForm(); err != nil {
		handleError(w, r, "Parse form", err, http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	if email == "" {
		handleError(w, r, "Email is required", errors.New("email is empty"), http.StatusBadRequest)
		return
	}

	displayName := r.Form.Get("displayName")
	if displayName == "" {
		handleError(w, r, "Display name is required", errors.New("display name is empty"), http.StatusBadRequest)
		return

	}

	password := r.Form.Get("password")
	if password == "" {
		handleError(w, r, "Password is required", errors.New("password is empty"), http.StatusBadRequest)
		return
	}

	confirmPassword := r.Form.Get("confirmPassword")
	if confirmPassword == "" {
		handleError(w, r, "Confirm password is required", errors.New("confirm password is empty"), http.StatusBadRequest)
		return
	}

	if password != confirmPassword {
		handleError(w, r, "Passwords do not match", errors.New("passwords not match"), http.StatusBadRequest)
		return
	}

	if err := h.authService.SignUp(r.Context(), email, password, displayName); err != nil {
		handleError(w, r, "Sign up", err, http.StatusInternalServerError)
		return
	}

	redirect(w, r, uri.Home())
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	h.authService.SignOut()
	redirect(w, r, uri.Home())
}
