package yts

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetMovieByIMDbID(t *testing.T) {
	baseURL := os.Getenv("YTS_BASE_URL")
	if baseURL == "" {
		t.Skip("empty YTS_BASE_URL")
	}

	c := NewClient(baseURL)

	movie, err := c.GetMovieByIMDbID(context.Background(), "tt0266543")
	require.NoError(t, err)

	assert.EqualValues(t, 1162, movie.ID)
	assert.Len(t, movie.Torrents, 3)
}

func TestClient_GetMovieByID(t *testing.T) {
	baseURL := os.Getenv("YTS_BASE_URL")
	if baseURL == "" {
		t.Skip("empty YTS_BASE_URL")
	}

	c := NewClient(baseURL)

	movie, err := c.GetMovieByID(context.Background(), 1162)
	require.NoError(t, err)

	assert.EqualValues(t, 1162, movie.ID)
	assert.Len(t, movie.Torrents, 3)
}
