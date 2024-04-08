package auth

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func (a Auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" || r.URL.Path == "/sessionLogin" || r.URL.Path == "/logout" {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		// Get the ID token sent by the client
		cookie, err := r.Cookie("session")
		if err != nil {
			// Session cookie is unavailable. Force user to login.
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Verify the session cookie. In this case an additional check is added to detect
		// if the user's Firebase session was revoked, user deleted/disabled, etc.
		decoded, err := a.authClient.VerifySessionCookieAndCheckRevoked(ctx, cookie.Value)
		if err != nil {
			// Session cookie is invalid. Force user to login.
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if decoded.Claims["is_admin"] != true {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		log.Ctx(ctx).Info().
			Str("uid", decoded.UID).
			Str("email", decoded.Claims["email"].(string)).
			Msg("User is authenticated")

		next.ServeHTTP(w, r)
	})
}
