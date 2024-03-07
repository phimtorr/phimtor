package auth

import (
	"context"
	"net/http"
)

type ctxKey struct{}

func Middleware(authService *FirebaseAuth) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), ctxKey{}, authService))
			next.ServeHTTP(w, r)
		})
	}
}

func IsSignedIn(ctx context.Context) bool {
	authService, ok := ctx.Value(ctxKey{}).(*FirebaseAuth)
	if !ok {
		return false
	}
	return !authService.CurrentUser().IsZero()
}

func CurrentUser(ctx context.Context) User {
	authService, ok := ctx.Value(ctxKey{}).(*FirebaseAuth)
	if !ok {
		return User{}
	}
	return authService.CurrentUser()
}
