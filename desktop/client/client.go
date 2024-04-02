package client

import (
	"context"
	"net/http"
	"strings"

	commonErrors "github.com/phimtorr/phimtor/common/errors"

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
	if isNotResponseOk(resp, err) {
		return nil, api.Pagination{}, handleError(resp.JSON400, resp.JSON500, err)
	}

	return resp.JSON200.Shows, resp.JSON200.Pagination, nil
}

func (c *Client) SearchShows(ctx context.Context, query string, page int) ([]api.Show, api.Pagination, error) {
	resp, err := c.SearchShowsWithResponse(ctx, &api.SearchShowsParams{
		Query: query,
		Page:  &page,
	})
	if isNotResponseOk(resp, err) {
		return nil, api.Pagination{}, handleError(resp.JSON400, resp.JSON500, err)
	}

	return resp.JSON200.Shows, resp.JSON200.Pagination, nil
}

func (c *Client) GetVideo(ctx context.Context, id int64) (api.Video, error) {
	resp, err := c.GetVideoWithResponse(ctx, id)
	if isNotResponseOk(resp, err) {
		return api.Video{}, handleError(resp.JSON400, resp.JSON500, err)
	}

	return resp.JSON200.Video, nil
}

type response interface {
	StatusCode() int
}

func isNotResponseOk(resp response, err error) bool {
	if err != nil {
		return true
	}
	statusCode := resp.StatusCode()
	return statusCode < 200 || statusCode >= 300
}

func handleError(json400 *api.BadRequest, json500 *api.InternalError, err error) error {
	if json400 != nil {
		return commonErrors.NewIncorrectInputError(json400.Code, json400.Message)
	}
	if json500 != nil {
		return commonErrors.NewUnknownError(json500.Code, json500.Message)
	}
	return err
}
