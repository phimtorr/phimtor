package uri

import (
	"fmt"
)

func SelectSubtitle(videoID int64, subtitleID int64) string {
	if subtitleID == 0 {
		return UnselectSubtitle(videoID)
	}
	return fmt.Sprintf("/videos/%d/subtitles/%d", videoID, subtitleID)
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

func DownloadSubtitle(videoID int64, subtitleID int64) string {
	return fmt.Sprintf("/videos/%d/subtitles/%d/download", videoID, subtitleID)
}
