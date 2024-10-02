package tmdb

import (
	"context"
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
)

type Client struct {
	client         *tmdb.Client
	defaultOptions map[string]string
}

func NewClient() Client {
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		panic("TMDB_API_KEY is required")
	}

	tmdbClient, err := tmdb.Init(apiKey)
	if err != nil {
		panic(err)
	}

	tmdbClient.SetClientAutoRetry()

	return Client{
		client: tmdbClient,
		defaultOptions: map[string]string{
			"language": "vi-VN",
		},
	}
}

func (c Client) GetMovieDetails(_ context.Context, movieID int) (*tmdb.MovieDetails, error) {
	return c.client.GetMovieDetails(movieID, c.defaultOptions)
}

func (c Client) GetTVSeriesDetails(ctx context.Context, tvID int) (*tmdb.TVDetails, []*tmdb.TVSeasonDetails, error) {
	tvDetails, err := c.client.GetTVDetails(tvID, c.defaultOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("get tv details: %w", err)
	}

	tvSeasons := make([]*tmdb.TVSeasonDetails, 0, len(tvDetails.Seasons))

	for _, season := range tvDetails.Seasons {
		tvSeason, err := c.getTVSeasonDetails(ctx, tvID, season.SeasonNumber)
		if err != nil {
			return nil, nil, fmt.Errorf("get tv season details: %w", err)
		}

		tvSeasons = append(tvSeasons, tvSeason)
	}

	return tvDetails, tvSeasons, nil
}

func (c Client) getTVSeasonDetails(_ context.Context, tvID, seasonNumber int) (*tmdb.TVSeasonDetails, error) {
	return c.client.GetTVSeasonDetails(tvID, seasonNumber, c.defaultOptions)
}

func (c Client) ListTopRatedMovies(ctx context.Context, page int) (*tmdb.MovieTopRated, error) {
	options := map[string]string{
		"page": fmt.Sprintf("%d", page),
	}

	resp, err := c.client.GetMovieTopRated(options)
	if err != nil {
		return nil, fmt.Errorf("get top rated movies: %w", err)
	}

	return resp, nil
}
