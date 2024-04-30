package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/desktop/client/api"
	"github.com/phimtorr/phimtor/desktop/server/ui"
	"github.com/phimtorr/phimtor/desktop/server/uri"
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

	// Clear the state
	h.upnpService.Reset()

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

	http.Redirect(w, r, uri.UPnPListTorrents(video.Id), http.StatusSeeOther)
	return nil
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

	http.Redirect(w, r, uri.UPnPListSubtitles(video.Id), http.StatusSeeOther)
	return nil
}

func (h *Handler) UPnPUploadSubtitle(w http.ResponseWriter, r *http.Request) error {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		return err
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

	http.Redirect(w, r, uri.UPnPListSubtitles(id), http.StatusSeeOther)
	return nil

}

func (h *Handler) UPnPListDevices(w http.ResponseWriter, r *http.Request) error {
	clients := h.upnpService.GetAvailableClients()
	if len(clients) == 0 {
		if err := h.upnpService.Scan(r.Context()); err != nil {
			return errors.Wrap(err, "scan devices")
		}
	}
	selectedUDN := h.upnpService.GetSelectedDeviceUDN()

	return ui.UPnPDevices(h.upnpService.GetAvailableClients(), selectedUDN).Render(r.Context(), w)
}

func (h *Handler) UPnPSelectDevice(w http.ResponseWriter, r *http.Request) error {
	udn := chi.URLParam(r, "udn")
	if udn == "" {
		return commonErrors.NewIncorrectInputError("empty-udn", "empty udn")
	}

	if err := h.upnpService.SelectDevice(r.Context(), udn); err != nil {
		return errors.Wrap(err, "select device")

	}

	http.Redirect(w, r, uri.UPnPListDevices(), http.StatusSeeOther)
	return nil
}

func (h *Handler) ScanDevices(w http.ResponseWriter, r *http.Request) error {
	if err := h.upnpService.Scan(r.Context()); err != nil {
		return errors.Wrap(err, "scan devices")
	}

	http.Redirect(w, r, uri.UPnPListDevices(), http.StatusSeeOther)
	return nil
}

func (h *Handler) UPnPPlay(w http.ResponseWriter, r *http.Request) error {
	selectedTorrent, _ := h.GetTorrentState()
	selectedSubtitle, _ := h.GetSubtitleState()

	if err := h.upnpService.Play(r.Context(), selectedTorrent.SelectedTorrent, selectedSubtitle.FileName, selectedSubtitle.OriginalContent); err != nil {
		return errors.Wrap(err, "play")
	}

	state := h.upnpService.GetState()

	return ui.UPnPController(state.InfoHash, state.FileIndex).Render(r.Context(), w)
}

func (h *Handler) UPnPStop(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (h *Handler) UPnPPause(w http.ResponseWriter, r *http.Request) error {

	return nil
}
