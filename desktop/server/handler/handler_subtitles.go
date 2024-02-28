package handler

import (
	"encoding/base64"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/pkg/browser"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/subtitle"
)

func (h *Handler) SelectSubtitle(w http.ResponseWriter, r *http.Request, videoID int64, subtitleName string) {
	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		handleError(w, r, "Get video", err, http.StatusInternalServerError)
		return
	}

	if subtitleName == "" {
		templ.Handler(
			ui.SubtitleSectionWithoutSubtitle(videoID, video.Subtitles),
		).ServeHTTP(w, r)
		return
	}

	selectedSubtitle, found := findSubtitle(video.Subtitles, subtitleName)
	if !found {
		handleError(w, r, "Subtitle not found", err, http.StatusBadRequest)
		return
	}

	fileName, originalContent, err := subtitle.GetFileFromLink(selectedSubtitle.Link)
	if err != nil {
		handleError(w, r, "Get subtitle file", err, http.StatusInternalServerError)
		return
	}

	normalizedContent, err := subtitle.Normalize(fileName, originalContent, 0)
	if err != nil {
		handleError(w, r, "Normalize subtitle", err, http.StatusInternalServerError)
		return
	}

	templ.Handler(
		ui.SubtitleSection(videoID, video.Subtitles, ui.SubtitleState{
			Name:                   selectedSubtitle.Name,
			FileName:               fileName,
			OriginalContent:        originalContent,
			Content:                normalizedContent,
			AdjustmentMilliseconds: 0,
		}),
	).ServeHTTP(w, r)
}

func findSubtitle(subtitles []api.Subtitle, subtitleName string) (api.Subtitle, bool) {
	for _, s := range subtitles {
		if s.Name == subtitleName {
			return s, true
		}
	}
	return api.Subtitle{}, false
}

func (h *Handler) DownloadSubtitle(w http.ResponseWriter, r *http.Request, videoID int64, subtitleName string) {
	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		handleError(w, r, "Get video", err, http.StatusInternalServerError)
		return
	}

	selectedSubtitle, found := findSubtitle(video.Subtitles, subtitleName)
	if !found {
		handleError(w, r, "Subtitle not found", err, http.StatusBadRequest)
		return
	}

	if err := browser.OpenURL(selectedSubtitle.Link); err != nil {
		handleError(w, r, "Open URL", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UploadSubtitle(w http.ResponseWriter, r *http.Request, videoID int64) {
	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		handleError(w, r, "Get video", err, http.StatusInternalServerError)
		return
	}
	file, header, err := r.FormFile("fileInput")
	if err != nil {
		handleError(w, r, "Get subtitle file", err, http.StatusBadRequest)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Ctx(r.Context()).Error().Err(err).Msg("Close file")
		}
	}()
	fileName := header.Filename

	originalContent, err := io.ReadAll(file)
	if err != nil {
		handleError(w, r, "Read file", err, http.StatusBadRequest)
		return
	}

	content, err := subtitle.Normalize(fileName, originalContent, 0)
	if err != nil {
		handleError(w, r, "Normalize subtitle", err, http.StatusBadRequest)
		return
	}

	templ.Handler(
		ui.SubtitleSection(videoID, video.Subtitles, ui.SubtitleState{
			Name:                   fileName,
			FileName:               fileName,
			OriginalContent:        originalContent,
			Content:                content,
			AdjustmentMilliseconds: 0,
		}),
	).ServeHTTP(w, r)
}

func (h *Handler) AdjustSubtitle(w http.ResponseWriter, r *http.Request, videoID int64) {
	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		handleError(w, r, "Get video", err, http.StatusInternalServerError)
		return
	}

	adjustMilliSeconds, err := strconv.ParseInt(r.URL.Query().Get("ms"), 10, 64)
	if err != nil {
		handleError(w, r, "Parse milliseconds", err, http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	fileName := r.FormValue("fileName")
	originalContent, err := base64.StdEncoding.DecodeString(r.FormValue("originalContent"))
	if err != nil {
		handleError(w, r, "Decode original content", err, http.StatusBadRequest)
		return
	}

	content, err := subtitle.Normalize(fileName, originalContent, time.Duration(adjustMilliSeconds)*time.Millisecond)
	if err != nil {
		handleError(w, r, "Normalize subtitle", err, http.StatusBadRequest)
		return
	}

	templ.Handler(
		ui.SubtitleSection(videoID, video.Subtitles, ui.SubtitleState{
			Name:                   name,
			FileName:               fileName,
			OriginalContent:        originalContent,
			Content:                content,
			AdjustmentMilliseconds: int(adjustMilliSeconds),
		}),
	).ServeHTTP(w, r)
}
