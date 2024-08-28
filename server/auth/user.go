package auth

import (
	"context"
	"time"
)

const (
	anonymousUUID = "anonymous"
)

type User struct {
	UUID  string
	Email string

	DisplayName   string
	EmailVerified bool
	PremiumUntil  time.Time
}

func (u User) IsAnonymous() bool {
	return u.UUID == anonymousUUID
}

func (u User) IsVerified() bool {
	return u.EmailVerified
}

func (u User) IsPremium() bool {
	return u.PremiumUntil.After(time.Now())
}

var (
	AnonymousUser = User{
		UUID:          anonymousUUID,
		Email:         "anonymous",
		DisplayName:   "Anonymous",
		EmailVerified: false,
		PremiumUntil:  time.Time{},
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
