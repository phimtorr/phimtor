package handler

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gabriel-vasile/mimetype"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/torrent"
	"github.com/phimtorr/phimtor/desktop/ui"
)

func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request, id int64, torrentName string) {
	resp, err := h.client.GetVideoWithResponse(r.Context(), id)
	if err != nil {
		handleError(w, r, "Failed to get video", err, http.StatusInternalServerError)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		handleError(w, r, "Failed to get video", err, resp.StatusCode())
		return
	}

	video := resp.JSON200.Video
	selectedTorrent := getSelectedTorrentLink(video.TorrentLinks, torrentName)

	infoHash, err := h.torManager.AddFromLink(selectedTorrent.Link)
	if err != nil {
		handleError(w, r, "Failed to add torrent", err, http.StatusInternalServerError)
		return
	}

	// This step is for speed up the download!
	h.torManager.CancelOthers(infoHash)

	templ.Handler(ui.Video(resp.JSON200.Video, infoHash, selectedTorrent)).ServeHTTP(w, r)
}

func getSelectedTorrentLink(torrentLinked []api.TorrentLink, torrentName string) api.TorrentLink {
	for _, t := range torrentLinked {
		if t.Name == torrentName {
			return t
		}
	}
	return torrentLinked[0]
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
