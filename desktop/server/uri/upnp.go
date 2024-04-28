package uri

import "fmt"

func UPnPView(videoID int64) string {
	return fmt.Sprintf("/upnp/videos/%d", videoID)
}

func UPnPListTorrents(videoID int64) string {
	return fmt.Sprintf("/upnp/videos/%d/torrents", videoID)
}

func UPnPSetTorrent(videoID int64, torrentID int64) string {
	return fmt.Sprintf("/upnp/videos/%d/torrents/%d", videoID, torrentID)
}

func UPnPListSubtitles(videoID int64) string {
	return fmt.Sprintf("/upnp/videos/%d/subtitles", videoID)
}

func UPnPSetSubtitle(videoID int64, subtitleID int64) string {
	return fmt.Sprintf("/upnp/videos/%d/subtitles/%d", videoID, subtitleID)
}

func UPnPUploadSubtitle(videoID int64) string {
	return fmt.Sprintf("/upnp/videos/%d/subtitles/upload", videoID)
}

func UPnPDevices() string {
	return "/upnp/devices"
}

func UPnPSelectDevice(udn string) string {
	return fmt.Sprintf("/upnp/devices/%s", udn)
}

func UPnPScanDevices() string {
	return "/upnp/devices/scan"
}
