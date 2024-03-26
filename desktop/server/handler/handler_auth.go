package handler

import (
	"net/http"

	"github.com/friendsofgo/errors"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/desktop/server/uri"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.Wrap(err, "parse form")
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if err := h.authService.SignIn(r.Context(), email, password); err != nil {
		return errors.Wrap(err, "sign in")
	}

	redirect(w, r, uri.Home())
	return nil
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.Wrap(err, "parse form")
	}

	email := r.Form.Get("email")
	if email == "" {
		return commonErrors.NewIncorrectInputError("empty-email", "Email is required")
	}

	displayName := r.Form.Get("displayName")
	if displayName == "" {
		return commonErrors.NewIncorrectInputError("empty-display-name", "Display name is required")
	}

	password := r.Form.Get("password")
	if password == "" {
		return commonErrors.NewIncorrectInputError("empty-password", "Password is required")
	}

	confirmPassword := r.Form.Get("confirmPassword")
	if confirmPassword == "" {
		return commonErrors.NewIncorrectInputError("empty-confirm-password", "Confirm password is required")
	}

	if password != confirmPassword {
		return commonErrors.NewIncorrectInputError("passwords-not-match", "Passwords do not match")
	}

	if err := h.authService.SignUp(r.Context(), email, password, displayName); err != nil {
		return errors.Wrap(err, "sign up")
	}

	redirect(w, r, uri.Home())
	return nil
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	h.authService.SignOut()
	redirect(w, r, uri.Home())
}
