package auth

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/friendsofgo/errors"
	"google.golang.org/api/identitytoolkit/v1"
	"google.golang.org/api/option"
)

type Storage interface {
	GetCredentials() Credentials
	SetCredentials(creds Credentials) error
}

type FirebaseAuth struct {
	apiKey  string
	svc     *identitytoolkit.Service
	storage Storage
}

func NewFirebaseAuth(apiKey string, storage Storage) *FirebaseAuth {
	if apiKey == "" {
		panic("apiKey is required")
	}
	if storage == nil {
		panic("storage is required")
	}
	svc, err := identitytoolkit.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}

	return &FirebaseAuth{
		apiKey:  apiKey,
		svc:     svc,
		storage: storage,
	}
}

func (a *FirebaseAuth) SignIn(ctx context.Context, email, password string) error {
	resp, err := a.svc.Accounts.SignInWithPassword(&identitytoolkit.GoogleCloudIdentitytoolkitV1SignInWithPasswordRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}).Context(ctx).Do()
	if err != nil {
		return errors.Wrap(err, "sign in")
	}

	user := User{
		Email:       resp.Email,
		DisplayName: resp.DisplayName,
	}
	a.setCredentials(NewCredentials(user, resp.IdToken, resp.ExpiresIn, resp.RefreshToken))
	return nil
}

func (a *FirebaseAuth) SignUp(ctx context.Context, email, password, displayName string) error {
	resp, err := a.svc.Accounts.SignUp(&identitytoolkit.GoogleCloudIdentitytoolkitV1SignUpRequest{
		DisplayName: displayName,
		Email:       email,
		Password:    password,
	}).Context(ctx).Do()
	if err != nil {
		return errors.Wrap(err, "sign up")
	}

	user := User{
		Email:       resp.Email,
		DisplayName: resp.DisplayName,
	}
	a.setCredentials(NewCredentials(user, resp.IdToken, resp.ExpiresIn, resp.RefreshToken))
	return nil
}

func (a *FirebaseAuth) SignOut() {
	a.setCredentials(Credentials{})
}

func (a *FirebaseAuth) setCredentials(cred Credentials) {
	if err := a.storage.SetCredentials(cred); err != nil {
		log.Warn().Err(err).Msg("Save credentials failed")
	}
}

func (a *FirebaseAuth) currentCredentials() Credentials {
	return a.storage.GetCredentials()
}

func (a *FirebaseAuth) CurrentUser() User {
	return a.currentCredentials().User
}

func (a *FirebaseAuth) GetJWTToken(ctx context.Context) (string, error) {
	if err := a.refreshIDTokenIfNeed(ctx); err != nil {
		return "", errors.Wrap(err, "refresh ID token")
	}

	return a.currentCredentials().IDToken, nil
}

func (a *FirebaseAuth) refreshIDTokenIfNeed(ctx context.Context) error {
	if !a.currentCredentials().IsExpired() {
		return nil
	}

	resp, err := refreshToken(ctx, a.apiKey, a.currentCredentials().RefreshToken)
	if err != nil {
		return err
	}

	user := a.currentCredentials().User
	a.setCredentials(NewCredentials(user, resp.IDToken, resp.ExpiresIn, resp.RefreshToken))
	return nil
}
