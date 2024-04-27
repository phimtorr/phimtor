package uri

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"

	"github.com/phimtorr/phimtor/desktop/torrent"
)

func GetVideo(id int64) string {
	return fmt.Sprintf("/videos/%d", id)
}

func GetVideoWithTorrentID(id int64, torrentID int64) string {
	return fmt.Sprintf("/videos/%d?torrentID=%d", id, torrentID)
}

func GetStream(infoHash torrent.InfoHash, fileIndex int, fileName string) string {
	return fmt.Sprintf("/stream/%s/%d/%s", infoHash, fileIndex, url.QueryEscape(fileName))
}

func OpenInVLC(infoHash torrent.InfoHash, fileIndex int) string {
	return fmt.Sprintf("/open-in-vlc/%s/%d", infoHash, fileIndex)
}

func GetStats(infoHash torrent.InfoHash, fileIndex int) string {
	return fmt.Sprintf("/stats/%s/%d", infoHash, fileIndex)
}

func GetMemFile(id uuid.UUID, fileName string) string {
	return fmt.Sprintf("/mem-files/%s/%s", id, url.QueryEscape(fileName))
}
