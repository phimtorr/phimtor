package client

import (
	"context"
	"net/http"
	"strings"

	"github.com/friendsofgo/errors"

	"github.com/phimtorr/phimtor/desktop/build"
	"github.com/phimtorr/phimtor/desktop/client/api"
)

type Client struct {
	*api.ClientWithResponses
}

type AuthService interface {
	GetJWTToken(ctx context.Context) (string, error)
}

func tokenRequestEditor(authService AuthService) api.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		token, err := authService.GetJWTToken(ctx)
		if err != nil {
			return errors.Wrap(err, "get jwt token")
		}
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		return nil
	}

}

func NewClient(authService AuthService) *Client {
	if authService == nil {
		panic("authService is required")
	}

	apiBaseURL := strings.TrimRight(build.ServerAddr, "/") + "/api/v1"
	cl, err := api.NewClientWithResponses(apiBaseURL, api.WithRequestEditorFn(tokenRequestEditor(authService)))
	if err != nil {
		panic(err)
	}
	return &Client{cl}
}

func (c *Client) ListShows(ctx context.Context, page, pageSize int, showType api.ShowType) ([]api.Show, api.Pagination, error) {
	resp, err := c.ListShowsWithResponse(ctx, &api.ListShowsParams{
		Page:     &page,
		PageSize: &pageSize,
		Type:     &showType,
	})
	if err != nil {
		return nil, api.Pagination{}, errors.Wrap(err, "list shows")
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, api.Pagination{}, errors.Errorf("list shows with status code %d", resp.StatusCode())
	}

	return resp.JSON200.Shows, resp.JSON200.Pagination, nil
}

func (c *Client) GetVideo(ctx context.Context, id int64) (api.Video, error) {
	resp, err := c.GetVideoWithResponse(ctx, id)
	if err != nil {
		return api.Video{}, errors.Wrap(err, "get video")
	}
	if resp.StatusCode() != http.StatusOK {
		return api.Video{}, errors.Errorf("get video with status code %d", resp.StatusCode())
	}

	return resp.JSON200.Video, nil
}
