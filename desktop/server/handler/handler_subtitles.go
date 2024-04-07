package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/pkg/browser"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/subtitle"
)

func (h *Handler) SelectSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseVideoID(chi.URLParam(r, "videoID"))
	if err != nil {
		return err
	}

	subtitleID, err := parseSubtitleID(chi.URLParam(r, "subtitleID"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	selectedSubtitle, found := findSubtitle(video.Subtitles, subtitleID)
	if !found {
		return commonErrors.NewIncorrectInputError("subtitle-not-found", fmt.Sprintf("subtitle id not found: %d", subtitleID))
	}

	fileName, originalContent, err := subtitle.GetFileFromLink(selectedSubtitle.Link)
	if err != nil {
		return errors.Wrap(err, "get file from link")
	}

	normalizedContent, err := subtitle.Normalize(fileName, originalContent, 0)
	if err != nil {
		return errors.Wrap(err, "normalize subtitle")
	}

	return ui.SubtitleSection(videoID, video.Subtitles, ui.SubtitleState{
		Name:                   selectedSubtitle.Name,
		FileName:               fileName,
		OriginalContent:        originalContent,
		Content:                normalizedContent,
		AdjustmentMilliseconds: 0,
	}).Render(r.Context(), w)
}

func findSubtitle(subtitles []api.Subtitle, subtitleID int64) (api.Subtitle, bool) {
	for _, s := range subtitles {
		if s.Id == subtitleID {
			return s, true
		}
	}
	return api.Subtitle{}, false
}

func (h *Handler) UnselectSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseVideoID(chi.URLParam(r, "videoID"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}
	return ui.SubtitleSectionWithoutSubtitle(videoID, video.Subtitles).Render(r.Context(), w)
}

func (h *Handler) DownloadSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseVideoID(chi.URLParam(r, "videoID"))
	if err != nil {
		return err
	}

	subtitleID, err := parseSubtitleID(chi.URLParam(r, "subtitleID"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	selectedSubtitle, found := findSubtitle(video.Subtitles, subtitleID)
	if !found {
		return commonErrors.NewIncorrectInputError("subtitle-not-found", fmt.Sprintf("subtitle id not found: %d", subtitleID))
	}

	if err := browser.OpenURL(selectedSubtitle.Link); err != nil {
		return errors.Wrap(err, "open url")
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *Handler) UploadSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseVideoID(chi.URLParam(r, "videoID"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}
	file, header, err := r.FormFile("fileInput")
	if err != nil {
		return errors.Wrap(err, "get form file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Ctx(r.Context()).Error().Err(err).Msg("Close file")
		}
	}()
	fileName := header.Filename

	originalContent, err := io.ReadAll(file)
	if err != nil {
		return errors.Wrap(err, "read all")
	}

	content, err := subtitle.Normalize(fileName, originalContent, 0)
	if err != nil {
		return errors.Wrap(err, "normalize subtitle")
	}

	return ui.SubtitleSection(videoID, video.Subtitles, ui.SubtitleState{
		Name:                   fileName,
		FileName:               fileName,
		OriginalContent:        originalContent,
		Content:                content,
		AdjustmentMilliseconds: 0,
	}).Render(r.Context(), w)
}

func (h *Handler) AdjustSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := parseVideoID(chi.URLParam(r, "videoID"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	adjustMilliSeconds, err := strconv.ParseInt(r.URL.Query().Get("ms"), 10, 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-ms", fmt.Sprintf("invalid ms=%s, err=%v", r.URL.Query().Get("ms"), err))
	}

	name := r.FormValue("name")
	fileName := r.FormValue("fileName")
	originalContent, err := base64.StdEncoding.DecodeString(r.FormValue("originalContent"))
	if err != nil {
		return errors.Wrap(err, "decode original content")
	}

	content, err := subtitle.Normalize(fileName, originalContent, time.Duration(adjustMilliSeconds)*time.Millisecond)
	if err != nil {
		return errors.Wrap(err, "normalize subtitle")
	}

	return ui.SubtitleSection(videoID, video.Subtitles, ui.SubtitleState{
		Name:                   name,
		FileName:               fileName,
		OriginalContent:        originalContent,
		Content:                content,
		AdjustmentMilliseconds: int(adjustMilliSeconds),
	}).Render(r.Context(), w)
}

var (
	ErrInvalidVideoID = commonErrors.NewIncorrectInputError("invalid-video-id", "invalid video id")
)

func parseVideoID(videoIDRaw string) (int64, error) {
	videoID, err := strconv.ParseInt(videoIDRaw, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(ErrInvalidVideoID, "parse videoID=%s, err=%v", videoIDRaw, err)
	}
	return videoID, nil
}

var (
	ErrInvalidSubtitleID = commonErrors.NewIncorrectInputError("invalid-subtitle-id", "invalid subtitle id")
)

func parseSubtitleID(subtitleIDRaw string) (int64, error) {
	subtitleID, err := strconv.ParseInt(subtitleIDRaw, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(ErrInvalidSubtitleID, "parse subtitleID=%s, err=%v", subtitleIDRaw, err)
	}
	return subtitleID, nil
}
