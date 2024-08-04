package http

import (
	"errors"
	"net/http"

	gotorrent "github.com/anacrolix/torrent"
	"github.com/go-chi/render"
)

func (s *Server) ListTorrents(w http.ResponseWriter, r *http.Request) {
	torrents := s.torManager.ListTorrents()

	render.Respond(w, r, map[string]any{
		"torrents": toHTTPTorrents(torrents),
	})
}

func (s *Server) AddTorrent(w http.ResponseWriter, r *http.Request, params AddTorrentParams) {
	dropOthers, deleteOthers := false, false
	if params.DropOthers != nil {
		dropOthers = *params.DropOthers
	}
	if params.DeleteOthers != nil {
		deleteOthers = *params.DeleteOthers
	}

	var requestBody AddTorrentJSONRequestBody
	if err := render.DecodeJSON(r.Body, &requestBody); err != nil {
		respondError(w, r, err, http.StatusBadRequest)
		return
	}

	var torrentLink string
	if requestBody.Url != nil {
		torrentLink = *requestBody.Url
	}
	if requestBody.MagnetUri != nil {
		torrentLink = *requestBody.MagnetUri
	}
	if torrentLink == "" {
		respondError(w, r, errors.New("url or magnetUri is required"), http.StatusBadRequest)
		return
	}

	infoHash, err := s.torManager.AddFromLink(torrentLink, dropOthers, deleteOthers)
	if err != nil {
		respondError(w, r, err, http.StatusInternalServerError)
		return
	}

	torr, ok := s.torManager.GetTorrent(infoHash)
	if !ok {
		respondError(w, r, errors.New("failed to get torrent"), http.StatusInternalServerError)
		return
	}

	render.Respond(w, r, map[string]any{
		"torrent": toHTTPTorrent(torr),
	})

}

func (s *Server) DeleteTorrent(w http.ResponseWriter, r *http.Request, infoHash InfoHash) {
	//TODO implement me
	panic("implement me")
}

func toHTTPTorrents(torrents []*gotorrent.Torrent) []Torrent {
	httpTorrents := make([]Torrent, len(torrents))
	for i, t := range torrents {
		httpTorrents[i] = toHTTPTorrent(t)
	}
	return httpTorrents
}

func toHTTPTorrent(t *gotorrent.Torrent) Torrent {
	return Torrent{
		InfoHash: t.InfoHash().String(),
		Name:     t.Name(),
		Size:     t.Length(),
		Files:    toHTTPFiles(t.Files()),
	}
}

func toHTTPFiles(files []*gotorrent.File) []File {
	httpFiles := make([]File, len(files))
	for i, f := range files {
		httpFiles[i] = toHTTPFile(f)
	}
	return httpFiles
}

func toHTTPFile(f *gotorrent.File) File {
	return File{
		Name: f.Path(),
		Size: f.Length(),
	}
}
