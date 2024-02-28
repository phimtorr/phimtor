package handler

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/subtitle"
	"github.com/phimtorr/phimtor/desktop/ui"
)

func (h *Handler) SelectSubtitle(w http.ResponseWriter, r *http.Request, videoID int64, subtitleName string) {
	video, err := h.apiClient.GetVideo(r.Context(), videoID)
	if err != nil {
		handleError(w, r, "Get video", err, http.StatusInternalServerError)
		return
	}

	if subtitleName == "" {
		templ.Handler(
			ui.SubtitleSectionWithoutSubtitle(video.Subtitles),
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
		ui.SubtitleSection(video.Subtitles, selectedSubtitle.Name, fileName, originalContent, normalizedContent),
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

}

func (h *Handler) AdjustSubtitle(w http.ResponseWriter, r *http.Request, videoID int64) {

}
