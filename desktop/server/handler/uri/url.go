package uri

import (
	"fmt"
	"net/url"

	"github.com/phimtorr/phimtor/desktop/torrent"
)

func Home() string {
	return "/"
}

func ListShows() string {
	return "/shows"
}

func GetMovie(id int64) string {
	return fmt.Sprintf("/movies/%d", id)
}

func GetSeries(id int64) string {
	return fmt.Sprintf("/series/%d", id)
}

func GetVideo(id int64) string {
	return fmt.Sprintf("/videos/%d", id)
}

func GetVideoWithTorrentName(id int64, torrentName string) string {
	return fmt.Sprintf("/videos/%d?torrent=%s", id, url.QueryEscape(torrentName))
}

func GetStream(infoHash torrent.InfoHash, fileIndex int) string {
	return fmt.Sprintf("/stream/%s/%d", infoHash, fileIndex)
}

func SelectSubtitle(videoID int64, subtitleName string) string {
	return fmt.Sprintf("/videos/%d/subtitles/%s", videoID, url.QueryEscape(subtitleName))
}

func UnselectSubtitle(videoID int64) string {
	return fmt.Sprintf("/videos/%d/subtitles", videoID)
}

func UploadSubtitle(videoID int64) string {
	return fmt.Sprintf("/videos/%d/subtitles/upload", videoID)
}

func AdjustSubtitle(videoID int64, adjustMilliseconds int) string {
	return fmt.Sprintf("/videos/%d/subtitles/adjust?ms=%d", videoID, adjustMilliseconds)
}

func DownloadSubtitle(videoID int64, subtitleName string) string {
	return fmt.Sprintf("/videos/%d/subtitles/%s/download", videoID, url.QueryEscape(subtitleName))
}
