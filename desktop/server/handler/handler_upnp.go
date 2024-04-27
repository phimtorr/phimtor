package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/phimtorr/phimtor/desktop/server/ui"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/subtitle"
	"github.com/rs/zerolog/log"
)

func (h *Handler) ViewUPnP(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	return ui.UPnP(video).Render(r.Context(), w)
}

func (h *Handler) UPnPListTorrents(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	selectedTorrent, _ := h.GetTorrentState()

	return ui.UPnPTorrents(video.Id, video.TorrentLinks, selectedTorrent.SelectedTorrent).Render(r.Context(), w)
}

func (h *Handler) UPnSetTorrent(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	torrentID, err := parseID(chi.URLParam(r, "torrentID"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	selectedTorrent := getSelectedTorrentLink(video.TorrentLinks, torrentID)

	h.SetTorrentState(TorrentState{SelectedTorrent: selectedTorrent})

	return ui.UPnPTorrents(video.Id, video.TorrentLinks, selectedTorrent).Render(r.Context(), w)
}

func (h *Handler) UPnPListSubtitles(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	selectedSubtitle, _ := h.GetSubtitleState()

	return ui.UPnPSubtitles(video.Id, video.Subtitles, selectedSubtitle.ID, selectedSubtitle.Name).Render(r.Context(), w)
}

func (h *Handler) UPnPSetSubtitle(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	subtitleID, err := parseID(chi.URLParam(r, "subtitleID"))
	if err != nil {
		return err
	}

	var selectedSubtitle api.Subtitle
	for _, s := range video.Subtitles {
		if s.Id == subtitleID {
			selectedSubtitle = s
		}
	}
	if selectedSubtitle.Id == 0 {
		return commonErrors.NewIncorrectInputError("subtitle-not-found", fmt.Sprintf("subtitle id not found: %d", subtitleID))
	}

	fileName, originalContent, err := subtitle.GetFileFromLink(selectedSubtitle.Link)
	if err != nil {
		return errors.Wrap(err, "get file from link")
	}

	h.SetSubtitleState(SubtitleState{
		ID:              selectedSubtitle.Id,
		Name:            selectedSubtitle.Name,
		FileName:        fileName,
		OriginalContent: originalContent,
	})

	return ui.UPnPSubtitles(video.Id, video.Subtitles, selectedSubtitle.Id, selectedSubtitle.Name).Render(r.Context(), w)
}

func (h *Handler) UPnPUploadSubtitle(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}

	video, err := h.apiClient.GetVideo(r.Context(), id)
	if err != nil {
		return errors.Wrap(err, "get video")
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return errors.Wrap(err, "parse multipart form")
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

	h.SetSubtitleState(SubtitleState{
		Name:            fileName,
		FileName:        fileName,
		OriginalContent: originalContent,
	})

	return ui.UPnPSubtitles(id, video.Subtitles, 0, fileName).Render(r.Context(), w)
}

func (h *Handler) UPnPPlay(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (h *Handler) UPnPStop(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (h *Handler) UPnPPause(w http.ResponseWriter, r *http.Request) error {

	return nil
}
