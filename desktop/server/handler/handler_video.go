package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	commonErrors "github.com/phimtorr/phimtor/common/errors"

	"github.com/go-chi/chi/v5"

	"github.com/phimtorr/phimtor/desktop/server/uri"
	"github.com/phimtorr/phimtor/desktop/vlc"

	"github.com/gabriel-vasile/mimetype"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/torrent"
)

func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-id", fmt.Sprintf("invalid id=%s, err=%v", idStr, err))
	}

	torrentName, err := url.QueryUnescape(r.URL.Query().Get("torrent"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-torrent-name", fmt.Sprintf("invalid torrent_name=%s, err=%v", torrentName, err))
	}

	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get video")
	}
	selectedTorrent := getSelectedTorrentLink(video.TorrentLinks, torrentName)
	selectedSubtitle := getSelectedSubtitle(video.Subtitles)

	infoHash, err := h.torManager.AddFromLink(selectedTorrent.Link)
	if err != nil {
		return errors.Wrap(err, "add torrent from link")
	}

	// This step is for speed up the download!
	h.torManager.CancelOthers(infoHash)

	return ui.Video(video, infoHash, selectedTorrent, selectedSubtitle.Name).Render(r.Context(), w)
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

func (h *Handler) Stream(w http.ResponseWriter, r *http.Request) error {
	infoHashStr := chi.URLParam(r, "infoHash")
	infoHash, err := torrent.InfoHashFromString(infoHashStr)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-info-hash", fmt.Sprintf("invalid info_hash=%s, err=%v", infoHashStr, err))
	}
	fileIndexStr := chi.URLParam(r, "fileIndex")
	fileIndex, err := strconv.Atoi(fileIndexStr)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-file-index", fmt.Sprintf("invalid file_index=%s, err=%v", fileIndexStr, err))
	}
	file, err := h.torManager.GetFile(infoHash, fileIndex)
	if err != nil {
		return errors.Wrap(err, "get file")
	}

	file.Download()
	reader := file.NewReader()
	reader.SetResponsive()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return errors.Wrap(err, "detect mime type")
	} else {
		w.Header().Set("Content-Type", mime.String())
	}

	http.ServeContent(w, r, file.DisplayPath(), time.Time{}, reader)
	return nil
}

func (h *Handler) OpenInVLC(w http.ResponseWriter, r *http.Request) error {
	infoHashStr := chi.URLParam(r, "infoHash")
	infoHash, err := torrent.InfoHashFromString(infoHashStr)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-info-hash", fmt.Sprintf("invalid info_hash=%s, err=%v", infoHashStr, err))
	}
	fileIndexStr := chi.URLParam(r, "fileIndex")
	fileIndex, err := strconv.Atoi(fileIndexStr)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-file-index", fmt.Sprintf("invalid file_index=%s, err=%v", fileIndexStr, err))
	}
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}

	streamURL := protocol + "://" + r.Host + uri.GetStream(infoHash, fileIndex)

	if err := vlc.OpenURL(streamURL); err != nil {
		return errors.Wrap(err, "open url in vlc")
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) error {
	infoHashStr := chi.URLParam(r, "infoHash")
	infoHash, err := torrent.InfoHashFromString(infoHashStr)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-info-hash", fmt.Sprintf("invalid info_hash=%s, err=%v", infoHashStr, err))
	}
	fileIndexStr := chi.URLParam(r, "fileIndex")
	fileIndex, err := strconv.Atoi(fileIndexStr)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-file-index", fmt.Sprintf("invalid file_index=%s, err=%v", fileIndexStr, err))
	}

	stats := h.torManager.Stats(infoHash, fileIndex)

	return ui.VideoStatistics(stats).Render(r.Context(), w)
}
