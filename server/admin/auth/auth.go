package auth

import (
	"firebase.google.com/go/v4/auth"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/phimtorr/phimtor/server/admin/auth/ui"
)

type Auth struct {
	authClient *auth.Client
}

func NewAuth(authClient *auth.Client) Auth {
	if authClient == nil {
		panic("authClient is required")
	}
	return Auth{authClient: authClient}
}

func (a Auth) Register(r chi.Router) {
	r.Get("/login", templ.Handler(ui.LoginPage(LoadFirebaseConfig())).ServeHTTP)
	r.Post("/sessionLogin", a.SessionLogin)
	r.Get("/logout", a.SessionLogout)
}
