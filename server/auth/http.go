package auth

import (
	"context"
	"net/http"
	"strings"

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

		r = setUserToRequestContext(r, User{
			UUID:        token.UID,
			Email:       token.Claims["email"].(string),
			DisplayName: token.Claims["name"].(string),
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

type User struct {
	UUID  string
	Email string

	DisplayName string
}

var (
	AnonymousUser = User{
		UUID:        "anonymous",
		Email:       "anonymous",
		DisplayName: "Anonymous",
	}
)

type ctxKey int

const (
	userContextKey ctxKey = iota
)

func UserFromCtx(ctx context.Context) User {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u
	}

	return AnonymousUser
}
