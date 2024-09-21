package uri

import "strconv"

func ViewVideo(videoID int64) string {
	return "/videos/" + strconv.FormatInt(videoID, 10)
}

func SyncVideo(videoID int64) string {
	return "/videos/" + strconv.FormatInt(videoID, 10) + "/sync"
}

func CreateTorrent(videoID int64) string {
	return "/videos/" + strconv.FormatInt(videoID, 10) + "/torrents/create"
}

func DeleteTorrent(videoID, id int64) string {
	return "/videos/" + strconv.FormatInt(videoID, 10) + "/torrents/" + strconv.FormatInt(id, 10)
}

func CreateSubtitle(videoID int64) string {
	return "/videos/" + strconv.FormatInt(videoID, 10) + "/subtitles/create"
}

func DeleteSubtitle(videoID, id int64) string {
	return "/videos/" + strconv.FormatInt(videoID, 10) + "/subtitles/" + strconv.FormatInt(id, 10)
}
