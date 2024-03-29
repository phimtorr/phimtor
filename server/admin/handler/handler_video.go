package handler

import (
	"net/http"
	"strconv"

	commonErrors "github.com/phimtorr/phimtor/common/errors"

	"github.com/phimtorr/phimtor/server/admin/ui"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) ViewVideo(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	video, err := h.repo.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	return ui.ViewVideo(video).Render(r.Context(), w)
}

func (h *Handler) CreateTorrent(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	if err := r.ParseForm(); err != nil {
		return errors.Wrap(err, "parsing form")
	}

	name := r.Form.Get("name")
	if name == "" {
		return commonErrors.NewIncorrectInputError("empty-name", "empty name")
	}
	link := r.Form.Get("link")
	if link == "" {
		return commonErrors.NewIncorrectInputError("empty-link", "empty link")
	}

	fileIndex, err := strconv.Atoi(r.Form.Get("fileIndex"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-file-index",
			errors.Wrap(err, "invalid file index").Error())
	}
	priority, err := strconv.Atoi(r.Form.Get("priority"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-priority",
			errors.Wrap(err, "invalid priority").Error())
	}

	if _, err := h.repo.CreateTorrent(r.Context(), TorrentToCreate{
		VideoID:   videoID,
		Name:      name,
		Link:      link,
		FileIndex: fileIndex,
		Priority:  priority,
	}); err != nil {
		return errors.Wrap(err, "create torrent")
	}

	video, err := h.repo.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	return ui.ViewTorrents(video.ID, video.Torrents).Render(r.Context(), w)
}
