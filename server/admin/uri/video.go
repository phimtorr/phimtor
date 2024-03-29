package uri

import "strconv"

func ViewVideo(videoID int64) string {
	return Prefix + "/videos/" + strconv.FormatInt(videoID, 10)
}

func CreateTorrent(videoID int64) string {
	return Prefix + "/videos/" + strconv.FormatInt(videoID, 10) + "/torrents/create"
}

func DeleteTorrent(videoID, id int64) string {
	return Prefix + "/videos/" + strconv.FormatInt(videoID, 10) + "/torrents/" + strconv.FormatInt(id, 10)
}
