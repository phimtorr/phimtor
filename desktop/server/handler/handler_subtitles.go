package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	videoID, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-video-id", fmt.Sprintf("invalid videoID err=%v", err))
	}
	subtitleName, err := url.QueryUnescape(chi.URLParam(r, "subtitleName"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-subtitle-name", fmt.Sprintf("invalid subtitle name err=%v", err))
	}
	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	selectedSubtitle, found := findSubtitle(video.Subtitles, subtitleName)
	if !found {
		return commonErrors.NewIncorrectInputError("subtitle-not-found", fmt.Sprintf("subtitle not found: %s", subtitleName))
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

func findSubtitle(subtitles []api.Subtitle, subtitleName string) (api.Subtitle, bool) {
	for _, s := range subtitles {
		if s.Name == subtitleName {
			return s, true
		}
	}
	return api.Subtitle{}, false
}

func (h *Handler) UnselectSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-video-id", fmt.Sprintf("invalid videoID err=%v", err))
	}

	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}
	return ui.SubtitleSectionWithoutSubtitle(videoID, video.Subtitles).Render(r.Context(), w)
}

func (h *Handler) DownloadSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-video-id", fmt.Sprintf("invalid videoID err=%v", err))
	}
	subtitleName, err := url.QueryUnescape(chi.URLParam(r, "subtitleName"))
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-subtitle-name", fmt.Sprintf("invalid subtitle name err=%v", err))
	}
	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	selectedSubtitle, found := findSubtitle(video.Subtitles, subtitleName)
	if !found {
		return commonErrors.NewIncorrectInputError("subtitle-not-found", fmt.Sprintf("subtitle not found: %s", subtitleName))
	}

	if err := browser.OpenURL(selectedSubtitle.Link); err != nil {
		return errors.Wrap(err, "open url")
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *Handler) UploadSubtitle(w http.ResponseWriter, r *http.Request) error {
	videoID, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-video-id", fmt.Sprintf("invalid videoID err=%v", err))
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
	videoID, err := strconv.ParseInt(chi.URLParam(r, "videoID"), 10, 64)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-video-id", fmt.Sprintf("invalid videoID err=%v", err))
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
