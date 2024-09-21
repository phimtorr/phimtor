package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/server/admin/http/ui"
)

type FileService interface {
	UploadFile(ctx context.Context, key string, body io.Reader) (string, error)
	DeleteFile(ctx context.Context, key string) error
}

type VideoHandler struct {
	repo        VideoRepository
	fileService FileService
}

func NewVideoHandler(repo VideoRepository, fileService FileService) *VideoHandler {
	if repo == nil {
		panic("nil repository")

	}
	if fileService == nil {
		panic("nil file service")
	}

	return &VideoHandler{
		repo:        repo,
		fileService: fileService,
	}
}

func (h *VideoHandler) ViewVideo(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	video, err := h.repo.GetVideo(r.Context(), videoID)
	if err != nil {
		return fmt.Errorf("get video: %w", err)
	}

	return ui.ViewVideo(video).Render(r.Context(), w)
}

func (h *VideoHandler) CreateTorrent(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return fmt.Errorf("parsing form: %w", err)
	}

	resolution, err := strconv.Atoi(r.Form.Get("resolution"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-resolution",
			fmt.Sprintf("invalid resolution: %v", err))
	}

	videoType := r.Form.Get("type")
	if videoType == "" {
		return commonErrors.NewIncorrectInputError("empty-type", "empty type")
	}

	codec := r.Form.Get("codec")
	source := r.Form.Get("source")

	link := r.Form.Get("link")
	if strings.TrimSpace(link) == "" {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return fmt.Errorf("get file: %w", err)
		}

		fileKey := strconv.FormatInt(videoID, 10) + "/torrents/" + fileHeader.Filename
		link, err = h.fileService.UploadFile(r.Context(), fileKey, file)
		if err != nil {
			return fmt.Errorf("upload file: %w", err)
		}
	}

	if link == "" {
		return commonErrors.NewIncorrectInputError("empty-link", "empty link")
	}

	fileIndex, err := strconv.Atoi(r.Form.Get("fileIndex"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-file-index",
			fmt.Sprintf("invalid file index: %v", err))
	}
	priority, err := strconv.Atoi(r.Form.Get("priority"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-priority",
			fmt.Sprintf("invalid priority: %v", err))
	}

	requiredPremium := r.Form.Get("requiredPremium") == "on"

	if _, err := h.repo.CreateTorrent(r.Context(), TorrentToCreate{
		VideoID:         videoID,
		Resolution:      resolution,
		Type:            videoType,
		Codec:           codec,
		Source:          source,
		Link:            link,
		FileIndex:       fileIndex,
		Priority:        priority,
		RequiredPremium: requiredPremium,
	}); err != nil {
		return fmt.Errorf("create torrent: %w", err)
	}

	video, err := h.repo.GetVideo(r.Context(), videoID)
	if err != nil {
		return fmt.Errorf("get video: %w", err)
	}

	return ui.ViewTorrents(video.ID, video.Torrents).Render(r.Context(), w)
}

func (h *VideoHandler) DeleteTorrent(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	torrentID, err := parseID(chi.URLParam(r, "torrentID"))
	if err != nil {
		return err
	}

	if err := h.repo.DeleteTorrent(r.Context(), videoID, torrentID); err != nil {
		return fmt.Errorf("delete torrent: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *VideoHandler) CreateSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return fmt.Errorf("parsing form: %w", err)
	}

	language := r.Form.Get("language")
	if language == "" {
		return commonErrors.NewIncorrectInputError("empty-language", "empty language")
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("get file: %w", err)
	}
	fileKey := strconv.FormatInt(videoID, 10) + "/" + language + "/" + fileHeader.Filename

	objectURL, err := h.fileService.UploadFile(r.Context(), fileKey, file)
	if err != nil {
		return fmt.Errorf("upload file: %w", err)
	}

	name := strings.TrimSpace(r.Form.Get("name"))
	owner := strings.TrimSpace(r.Form.Get("owner"))

	if name == "" {
		name = fileHeader.Filename
	}

	priority, err := strconv.Atoi(r.Form.Get("priority"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-priority",
			fmt.Sprintf("invalid priority: %v", err))
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
		return fmt.Errorf("create subtitle: %w", err)
	}

	video, err := h.repo.GetVideo(r.Context(), videoID)
	if err != nil {
		return fmt.Errorf("get video: %w", err)
	}

	return ui.ViewSubtitles(video.ID, video.Subtitles).Render(r.Context(), w)
}

func (h *VideoHandler) DeleteSubtitle(w http.ResponseWriter, r *http.Request) error {
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
		return fmt.Errorf("get subtitle: %w", err)
	}

	if err := h.repo.DeleteSubtitle(r.Context(), videoID, subtitleID); err != nil {
		return fmt.Errorf("delete subtitle: %w", err)
	}

	if err := h.fileService.DeleteFile(r.Context(), sub.FileKey); err != nil {
		return fmt.Errorf("delete file: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
