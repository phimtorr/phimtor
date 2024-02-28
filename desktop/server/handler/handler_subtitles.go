package handler

import (
	"io"
	"net/http"

	"github.com/a-h/templ"
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
		ui.SubtitleSection(videoID, video.Subtitles, selectedSubtitle.Name, fileName, originalContent, normalizedContent),
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
		ui.SubtitleSection(videoID, video.Subtitles, fileName, fileName, originalContent, content),
	).ServeHTTP(w, r)
}

func (h *Handler) AdjustSubtitle(w http.ResponseWriter, r *http.Request, videoID int64) {

}
