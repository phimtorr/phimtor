package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/server/uri"
	"github.com/phimtorr/phimtor/desktop/torrent"
	"github.com/phimtorr/phimtor/desktop/vlc"
)

func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	torrentID, err := parseTorrentID(r.URL.Query().Get("torrentID"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	selectedTorrent := getSelectedTorrentLink(video.TorrentLinks, torrentID)
	selectedSubtitle := getSelectedSubtitle(video.Subtitles)

	infoHash, err := h.torManager.AddFromLink(selectedTorrent.Link)
	if err != nil {
		return errors.Wrap(err, "add torrent from link")
	}

	videoIndex, err := h.findVideoIndex(infoHash, selectedTorrent.FileIndex)
	if err != nil {
		return errors.Wrap(err, "find video index")
	}

	file, err := h.torManager.GetFile(infoHash, videoIndex)
	if err != nil {
		return errors.Wrap(err, "get file")
	}

	return ui.Video(video, infoHash, selectedTorrent, videoIndex, file.DisplayPath(), selectedSubtitle).Render(r.Context(), w)
}

func (h *Handler) findVideoIndex(infoHash torrent.InfoHash, configuredIndex int) (int, error) {
	tor, ok := h.torManager.GetTorrent(infoHash)
	if !ok {
		return 0, fmt.Errorf("torrent not found")
	}
	files := tor.Files()
	if isVideoFile(files[configuredIndex].DisplayPath()) {
		return configuredIndex, nil
	}

	for i, file := range files {
		if isVideoFile(file.DisplayPath()) {
			return i, nil
		}
	}

	return 0, fmt.Errorf("video file not found")
}

func isVideoFile(path string) bool {
	return strings.HasSuffix(path, ".mp4") || strings.HasSuffix(path, ".mkv") || strings.HasSuffix(path, ".avi")
}

func getSelectedTorrentLink(torrentLinked []api.TorrentLink, torrentID int64) api.TorrentLink {
	for _, t := range torrentLinked {
		if t.Id == torrentID {
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

func (h *Handler) OpenInVLC(w http.ResponseWriter, r *http.Request) error {
	infoHash, err := parseInfoHash(chi.URLParam(r, "infoHash"))
	if err != nil {
		return err
	}
	fileIndex, err := parseFileIndex(chi.URLParam(r, "fileIndex"))
	if err != nil {
		return err
	}

	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}

	file, err := h.torManager.GetFile(infoHash, fileIndex)
	if err != nil {
		return errors.Wrap(err, "get file")
	}

	streamURL := protocol + "://" + r.Host + uri.GetStream(infoHash, fileIndex, file.DisplayPath())

	if err := vlc.OpenURL(streamURL); err != nil {
		return errors.Wrap(err, "open url in vlc")
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) error {
	infoHash, err := parseInfoHash(chi.URLParam(r, "infoHash"))
	if err != nil {
		return err
	}
	fileIndex, err := parseFileIndex(chi.URLParam(r, "fileIndex"))
	if err != nil {
		return err
	}

	stats := h.torManager.Stats(infoHash, fileIndex)

	return ui.VideoStatistics(stats).Render(r.Context(), w)
}

var (
	ErrInvalidID = commonErrors.NewIncorrectInputError("invalid-id", "invalid id")
)

func parseID(idRaw string) (int64, error) {
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(ErrInvalidID, "parse id=%s, err=%v", idRaw, err)
	}
	return id, nil
}

var (
	ErrInvalidTorrentID = commonErrors.NewIncorrectInputError("invalid-torrent-id", "invalid torrent id")
)

func parseTorrentID(torrentIDRaw string) (int64, error) {
	if torrentIDRaw == "" {
		return 0, nil
	}
	id, err := strconv.ParseInt(torrentIDRaw, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(ErrInvalidTorrentID, "parse torrent_id=%s, err=%v", torrentIDRaw, err)
	}
	return id, nil
}

var (
	ErrInvalidInfoHash = commonErrors.NewIncorrectInputError("invalid-info-hash", "invalid info hash")
)

func parseInfoHash(infoHashRaw string) (torrent.InfoHash, error) {
	infoHash, err := torrent.InfoHashFromString(infoHashRaw)
	if err != nil {
		return torrent.InfoHash{}, errors.Wrap(ErrInvalidInfoHash, fmt.Sprintf("parse info_hash=%s, err=%v", infoHashRaw, err))
	}
	return infoHash, nil
}

var (
	ErrInvalidFileIndex = commonErrors.NewIncorrectInputError("invalid-file-index", "invalid file index")
)

func parseFileIndex(fileIndexRaw string) (int, error) {
	fileIndex, err := strconv.Atoi(fileIndexRaw)
	if err != nil {
		return 0, errors.Wrap(ErrInvalidFileIndex, fmt.Sprintf("parse file_index=%s, err=%v", fileIndexRaw, err))
	}
	return fileIndex, nil
}
