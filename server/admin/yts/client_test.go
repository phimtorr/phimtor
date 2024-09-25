package yts

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_GetMovie(t *testing.T) {
	baseURL := os.Getenv("YTS_BASE_URL")
	if baseURL == "" {
		t.Skip("empty YTS_BASE_URL")
	}

	c := NewClient(baseURL)

	movie, err := c.GetMovie(context.Background(), "tt0266543")
	require.NoError(t, err)

	assert.Equal(t, 1162, movie.ID)
	assert.Len(t, movie.Torrents, 3)
}
