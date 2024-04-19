package auth

import "time"

type User struct {
	Email       string `yaml:"email"`
	DisplayName string `yaml:"display_name"`
}

func (u User) IsZero() bool {
	return u == User{}
}

type Credentials struct {
	User         User      `yaml:"user"`
	IDToken      string    `yaml:"id_token"`
	ExpiresIn    int64     `yaml:"expires_in"` // seconds
	CreateAt     time.Time `yaml:"create_at"`
	RefreshToken string    `yaml:"refresh_token"`
}

func NewCredentials(user User, idToken string, expiresIn int64, refreshToken string) Credentials {
	return Credentials{
		User:         user,
		IDToken:      idToken,
		ExpiresIn:    expiresIn,
		CreateAt:     time.Now(),
		RefreshToken: refreshToken,
	}
}

func (c Credentials) IsExpired() bool {
	if c == (Credentials{}) {
		return false // no credentials mean it's not expired
	}

	return time.Now().After(c.CreateAt.Add(time.Duration(c.ExpiresIn-2) * time.Second)) // 2 seconds for safety
}
