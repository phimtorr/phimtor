package upnp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_buildCurrentURIMetaData(t *testing.T) {
	content, err := buildCurrentURIMetaData(
		"test.mp4",
		"video/mp4",
		"http://localhost:123/video.mp4",
		"http://localhost:123/sub.srt",
	)
	require.NoError(t, err)

	t.Logf("%s", content)
}
func Test_buildCurrentURIMetaData_without_sub(t *testing.T) {
	content, err := buildCurrentURIMetaData(
		"test.mp4",
		"video/mp4",
		"http://localhost:123/video.mp4",
		"",
	)
	require.NoError(t, err)

	t.Logf("%s", content)
}
