package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/friendsofgo/errors"
)

const (
	refreshTokenEndpoint = "https://securetoken.googleapis.com/v1/token"
)

// tokenResponse is a type that represents the response body received from Google Token Service
type tokenResponse struct {
	// Probably 3600 (seconds)
	ExpiresIn    int64  `json:"expires_in,omitempty,string" description:"The number of second in which the ID token expires"`
	TokenType    string `json:"token_type" description:"The type of the access token; always Bearer"`
	RefreshToken string `json:"refresh_token" description:"The refresh token provided in the request or a new refresh token"`
	IDToken      string `json:"id_token" description:"The ID token"`
	UserID       string `json:"user_id" description:"A unique identifier of the User"`
	ProjectID    string `json:"project_id" description:"The Google Project ID"`
}

func refreshToken(ctx context.Context, apiKey string, refreshToken string) (tokenResponse, error) {
	refreshTokenURL := refreshTokenEndpoint + fmt.Sprintf("?key=%s", apiKey)
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, refreshTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return tokenResponse{}, errors.Wrap(err, "create refresh token request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return tokenResponse{}, errors.Wrap(err, "call refresh token endpoint")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return tokenResponse{}, errors.Wrap(err, "read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return tokenResponse{}, errors.Errorf("refresh token request failed with status code %d: %s", resp.StatusCode, body)
	}

	var tokenResp tokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return tokenResponse{}, errors.Wrap(err, "unmarshal response body")
	}

	return tokenResp, nil
}
