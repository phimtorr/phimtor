package uri

import (
	"fmt"

	"github.com/phimtorr/phimtor/desktop/torrent"
)

func GetVideo(id int64) string {
	return fmt.Sprintf("/videos/%d", id)
}

func GetVideoWithTorrentID(id int64, torrentID int64) string {
	return fmt.Sprintf("/videos/%d?torrentID=%d", id, torrentID)
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
