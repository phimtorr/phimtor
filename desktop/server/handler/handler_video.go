package handler

import (
	"net/http"
	"time"

	"github.com/phimtorr/phimtor/desktop/server/uri"
	"github.com/phimtorr/phimtor/desktop/vlc"

	"github.com/a-h/templ"
	"github.com/gabriel-vasile/mimetype"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/torrent"
)

func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request, id int64, torrentName string) {
	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		handleError(w, r, "Failed to get video", err, http.StatusInternalServerError)
		return
	}
	selectedTorrent := getSelectedTorrentLink(video.TorrentLinks, torrentName)
	selectedSubtitle := getSelectedSubtitle(video.Subtitles)

	infoHash, err := h.torManager.AddFromLink(selectedTorrent.Link)
	if err != nil {
		handleError(w, r, "Failed to add torrent", err, http.StatusInternalServerError)
		return
	}

	// This step is for speed up the download!
	h.torManager.CancelOthers(infoHash)

	templ.Handler(ui.Video(video, infoHash, selectedTorrent, selectedSubtitle.Name)).ServeHTTP(w, r)
}

func getSelectedTorrentLink(torrentLinked []api.TorrentLink, torrentName string) api.TorrentLink {
	for _, t := range torrentLinked {
		if t.Name == torrentName {
			return t
		}
	}
	return torrentLinked[0]
}

func getSelectedSubtitle(subtitles []api.Subtitle) api.Subtitle {
	if len(subtitles) == 0 {
		return api.Subtitle{}
	}
	for _, s := range subtitles {
		if s.Language == "vi" {
			return s
		}
	}
	return subtitles[0]
}

func (h *Handler) Stream(w http.ResponseWriter, r *http.Request, infoHash torrent.InfoHash, fileIndex int) {
	file, err := h.torManager.GetFile(infoHash, fileIndex)
	if err != nil {
		handleError(w, r, "Failed to get file", err, http.StatusBadRequest)
		return
	}

	file.Download()
	reader := file.NewReader()
	reader.SetResponsive()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		handleError(w, r, "detect mime type", err, http.StatusBadRequest)
		return
	} else {
		w.Header().Set("Content-Type", mime.String())
	}

	http.ServeContent(w, r, file.DisplayPath(), time.Time{}, reader)
}

func (h *Handler) OpenInVLC(w http.ResponseWriter, r *http.Request, infoHash torrent.InfoHash, fileIndex int) {
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}

	streamURL := protocol + "://" + r.Host + uri.GetStream(infoHash, fileIndex)

	if err := vlc.OpenURL(streamURL); err != nil {
		handleError(w, r, "Failed to open in VLC", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request, infoHash torrent.InfoHash, fileIndex int) {
	stats := h.torManager.Stats(infoHash, fileIndex)
	templ.Handler(ui.VideoStatistics(stats)).ServeHTTP(w, r)
}
