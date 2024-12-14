package yts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/context"
)

var (
	ErrNoMovieFound = fmt.Errorf("no movie found")
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClientFromEnv() *Client {
	baseURL := os.Getenv("YTS_BASE_URL")
	if baseURL == "" {
		panic("empty YTS_BASE_URL")
	}

	return NewClient(baseURL)
}

func NewClient(baseURL string) *Client {
	return NewClientWithHTTPClient(baseURL, http.DefaultClient)
}

func NewClientWithHTTPClient(baseURL string, httpClient *http.Client) *Client {
	if baseURL == "" {
		panic("empty baseURL")
	}

	if httpClient == nil {
		panic("nil httpClient")
	}
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

func (c *Client) GetMovieByIMDbID(ctx context.Context, imdbID string) (Movie, error) {
	listMoviesURL, err := url.JoinPath(c.baseURL, "list_movies.json")
	if err != nil {
		return Movie{}, fmt.Errorf("join path: %w", err)
	}

	listMoviesURL = listMoviesURL + "?query_term=" + imdbID

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, listMoviesURL, nil)
	if err != nil {
		return Movie{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Movie{}, err
	}
	defer resp.Body.Close()

	type response struct {
		Status        string `json:"status"`
		StatusMessage string `json:"status_message"`
		Data          struct {
			MovieCount int             `json:"movie_count"`
			Movies     []responseMovie `json:"movies"`
		}
	}

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return Movie{}, fmt.Errorf("decode: %w", err)
	}

	if res.Status != "ok" {
		return Movie{}, fmt.Errorf("status not ok: %s, message=%s", res.Status, res.StatusMessage)
	}

	if res.Data.MovieCount == 0 {
		return Movie{}, ErrNoMovieFound
	}

	movie, err := toMovie(res.Data.Movies[0])
	if err != nil {
		return Movie{}, fmt.Errorf("toMovie: %w", err)
	}

	return movie, nil
}

func (c *Client) GetMovieByID(ctx context.Context, id int64) (Movie, error) {
	listMoviesURL, err := url.JoinPath(c.baseURL, "movie_details.json")
	if err != nil {
		return Movie{}, fmt.Errorf("join path: %w", err)
	}

	listMoviesURL = listMoviesURL + "?movie_id=" + fmt.Sprintf("%d", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, listMoviesURL, nil)
	if err != nil {
		return Movie{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Movie{}, err
	}
	defer resp.Body.Close()

	type response struct {
		Status        string `json:"status"`
		StatusMessage string `json:"status_message"`
		Data          struct {
			Movie responseMovie `json:"movie"`
		}
	}

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return Movie{}, fmt.Errorf("decode: %w", err)
	}

	if res.Status != "ok" {
		return Movie{}, fmt.Errorf("status not ok: %s, message=%s", res.Status, res.StatusMessage)
	}

	return toMovie(res.Data.Movie)
}
