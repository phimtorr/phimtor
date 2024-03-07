package auth

import "time"

type User struct {
	Email       string
	DisplayName string
}

func (u User) IsZero() bool {
	return u == User{}
}

type credentials struct {
	idToken      string
	expiresIn    int64 // seconds
	createAt     time.Time
	refreshToken string
}

func newCredentials(idToken string, expiresIn int64, refreshToken string) credentials {
	return credentials{
		idToken:      idToken,
		expiresIn:    expiresIn,
		createAt:     time.Now(),
		refreshToken: refreshToken,
	}
}

func (c credentials) IsExpired() bool {
	if c == (credentials{}) {
		return false // no credentials mean it's not expired
	}

	return time.Now().After(c.createAt.Add(time.Duration(c.expiresIn-2) * time.Second)) // 2 seconds for safety
}
