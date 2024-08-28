package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"firebase.google.com/go/v4/auth"
)

type FirebaseHttpMiddleware struct {
	AuthClient *auth.Client
}

func NewFirebaseHttpMiddleware(authClient *auth.Client) FirebaseHttpMiddleware {
	if authClient == nil {
		panic("authClient is required")
	}
	return FirebaseHttpMiddleware{AuthClient: authClient}
}

func (a FirebaseHttpMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bearerToken := a.tokenFromHeader(r)
		if bearerToken == "" {
			r = setUserToRequestContext(r, AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		token, err := a.AuthClient.VerifyIDToken(ctx, bearerToken)
		if err != nil {
			http.Error(w, "unable to verify jwt", http.StatusUnauthorized)
			return
		}

		email := token.Claims["email"].(string) // always present
		var premiumUntil time.Time
		if v, ok := token.Claims["premium_until"].(float64); ok {
			premiumUntil = time.Unix(int64(v), 0)
		}

		emailVerified := false
		if v, ok := token.Claims["email_verified"].(bool); ok {
			emailVerified = v
		}

		name := email
		if token.Claims["name"] != nil {
			name = token.Claims["name"].(string)
		}

		r = setUserToRequestContext(r, User{
			UUID:          token.UID,
			Email:         email,
			DisplayName:   name,
			EmailVerified: emailVerified,
			PremiumUntil:  premiumUntil,
		})

		next.ServeHTTP(w, r)
	})
}

func (a FirebaseHttpMiddleware) tokenFromHeader(r *http.Request) string {
	headerValue := r.Header.Get("Authorization")

	if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
		return headerValue[7:]
	}

	return ""
}

func setUserToRequestContext(r *http.Request, u User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, u)
	return r.WithContext(ctx)
}
