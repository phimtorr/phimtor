package uri

import (
	"fmt"
	"net/url"

	"github.com/phimtorr/phimtor/desktop/torrent"
)

func GetVideo(id int64) string {
	return fmt.Sprintf("/videos/%d", id)
}

func GetVideoWithTorrentName(id int64, torrentName string) string {
	return fmt.Sprintf("/videos/%d?torrent=%s", id, url.QueryEscape(torrentName))
}

func GetStream(infoHash torrent.InfoHash, fileIndex int) string {
	return fmt.Sprintf("/stream/%s/%d", infoHash, fileIndex)
}

func OpenInVLC(infoHash torrent.InfoHash, fileIndex int) string {
	return fmt.Sprintf("/open-in-vlc/%s/%d", infoHash, fileIndex)
}

func GetStats(infoHash torrent.InfoHash, fileIndex int) string {
	return fmt.Sprintf("/stats/%s/%d", infoHash, fileIndex)
}
