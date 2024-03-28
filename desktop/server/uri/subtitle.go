package uri

import (
	"fmt"
	"net/url"
)

func SelectSubtitle(videoID int64, subtitleName string) string {
	if subtitleName == "" {
		return UnselectSubtitle(videoID)
	}
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
