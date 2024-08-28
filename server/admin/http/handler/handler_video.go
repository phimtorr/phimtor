package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"

	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/server/admin/http/ui"
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

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return errors.Wrap(err, "parsing form")
	}

	name := r.Form.Get("name")
	if name == "" {
		return commonErrors.NewIncorrectInputError("empty-name", "empty name")
	}

	link := r.Form.Get("link")
	if strings.TrimSpace(link) == "" {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return errors.Wrap(err, "get file")
		}

		fileKey := strconv.FormatInt(videoID, 10) + "/torrents/" + fileHeader.Filename
		link, err = h.fileService.UploadFile(r.Context(), fileKey, file)
		if err != nil {
			return errors.Wrap(err, "upload file")
		}
	}

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

	requiredPremium := r.Form.Get("requiredPremium") == "on"

	if _, err := h.repo.CreateTorrent(r.Context(), TorrentToCreate{
		VideoID:         videoID,
		Name:            name,
		Link:            link,
		FileIndex:       fileIndex,
		Priority:        priority,
		RequiredPremium: requiredPremium,
	}); err != nil {
		return errors.Wrap(err, "create torrent")
	}

	video, err := h.repo.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	return ui.ViewTorrents(video.ID, video.Torrents).Render(r.Context(), w)
}

func (h *Handler) DeleteTorrent(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	torrentID, err := parseID(chi.URLParam(r, "torrentID"))
	if err != nil {
		return err
	}

	if err := h.repo.DeleteTorrent(r.Context(), videoID, torrentID); err != nil {
		return errors.Wrap(err, "delete torrent")
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *Handler) CreateSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return errors.Wrap(err, "parsing form")
	}

	language := r.Form.Get("language")
	if language == "" {
		return commonErrors.NewIncorrectInputError("empty-language", "empty language")
	}
	name := r.Form.Get("name")
	if name == "" {
		return commonErrors.NewIncorrectInputError("empty-name", "empty name")
	}
	owner := r.Form.Get("owner")

	file, fileHeader, err := r.FormFile("file")
	fileKey := strconv.FormatInt(videoID, 10) + "/" + language + "/" + fileHeader.Filename

	objectURL, err := h.fileService.UploadFile(r.Context(), fileKey, file)
	if err != nil {
		return errors.Wrap(err, "upload file")
	}

	priority, err := strconv.Atoi(r.Form.Get("priority"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-priority",
			errors.Wrap(err, "invalid priority").Error())
	}

	if _, err := h.repo.CreateSubtitle(r.Context(), SubtitleToCreate{
		VideoID:  videoID,
		Language: language,
		Name:     name,
		Owner:    owner,
		Link:     objectURL,
		FileKey:  fileKey,
		Priority: priority,
	}); err != nil {
		return errors.Wrap(err, "create subtitle")
	}

	video, err := h.repo.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	return ui.ViewSubtitles(video.ID, video.Subtitles).Render(r.Context(), w)
}

func (h *Handler) DeleteSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	subtitleID, err := parseID(chi.URLParam(r, "subtitleID"))
	if err != nil {
		return err
	}

	sub, err := h.repo.GetSubtitle(r.Context(), videoID, subtitleID)
	if err != nil {
		return errors.Wrap(err, "get subtitle")
	}

	if err := h.repo.DeleteSubtitle(r.Context(), videoID, subtitleID); err != nil {
		return errors.Wrap(err, "delete subtitle")
	}

	if err := h.fileService.DeleteFile(r.Context(), sub.FileKey); err != nil {
		return errors.Wrap(err, "delete file")
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
