package auth

import (
	"context"
	"github.com/friendsofgo/errors"
	"google.golang.org/api/identitytoolkit/v1"
	"google.golang.org/api/option"
	"sync"
)

type FirebaseAuth struct {
	apiKey string
	svc    *identitytoolkit.Service

	mu          sync.RWMutex
	currentUser User
	cred        credentials
}

func NewFirebaseAuth(apiKey string) *FirebaseAuth {
	if apiKey == "" {
		panic("apiKey is required")
	}
	svc, err := identitytoolkit.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}

	return &FirebaseAuth{
		apiKey: apiKey,
		svc:    svc,
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

	a.setUser(User{
		Email:       resp.Email,
		DisplayName: resp.DisplayName,
	})
	a.setCredentials(newCredentials(resp.IdToken, resp.ExpiresIn, resp.RefreshToken))
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

	a.setUser(User{
		Email:       resp.Email,
		DisplayName: resp.DisplayName,
	})
	a.setCredentials(newCredentials(resp.IdToken, resp.ExpiresIn, resp.RefreshToken))
	return nil
}

func (a *FirebaseAuth) SignOut() {
	a.setUser(User{})
	a.setCredentials(credentials{})
}

func (a *FirebaseAuth) setUser(user User) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.currentUser = user
}

func (a *FirebaseAuth) setCredentials(cred credentials) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.cred = cred
}

func (a *FirebaseAuth) currentCredentials() credentials {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.cred
}

func (a *FirebaseAuth) CurrentUser() User {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.currentUser
}

func (a *FirebaseAuth) GetJWTToken(ctx context.Context) (string, error) {
	if err := a.refreshIDTokenIfNeed(ctx); err != nil {
		return "", errors.Wrap(err, "refresh ID token")
	}

	return a.currentCredentials().idToken, nil
}

func (a *FirebaseAuth) refreshIDTokenIfNeed(ctx context.Context) error {
	if !a.currentCredentials().IsExpired() {
		return nil
	}

	resp, err := refreshToken(ctx, a.apiKey, a.currentCredentials().refreshToken)
	if err != nil {
		return err
	}

	a.setCredentials(newCredentials(resp.IDToken, resp.ExpiresIn, resp.RefreshToken))
	return nil
}
