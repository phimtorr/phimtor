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

func NewClient() *Client {
	apiBaseURL := strings.TrimRight(build.ServerAddr, "/") + "/api/v1"
	cl, err := api.NewClientWithResponses(apiBaseURL)
	if err != nil {
		panic(err)
	}
	return &Client{cl}
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
