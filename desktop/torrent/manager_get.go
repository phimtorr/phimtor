package torrent

import (
	"fmt"
	"net/http"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/gabriel-vasile/mimetype"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

func (m *Manager) GetFile(infoHash InfoHash, fileIndex int) (*torrent.File, error) {
	tor, ok := m.GetTorrent(infoHash)
	if !ok {
		return nil, fmt.Errorf("torrent not found")
	}
	return tor.Files()[fileIndex], nil
}

func (m *Manager) GetTorrent(infoHash InfoHash) (*torrent.Torrent, bool) {
	tor, ok := m.client.Torrent(metainfo.Hash(infoHash))
	if !ok {
		return nil, false
	}
	<-tor.GotInfo()
	return tor, true
}

func (m *Manager) StreamFile(w http.ResponseWriter, r *http.Request, infoHash InfoHash, fileIndex int) error {
	file, err := m.GetFile(infoHash, fileIndex)
	if err != nil {
		return errors.Wrap(err, "get file")
	}

	file.Download()
	reader := file.NewReader()
	reader.SetResponsive()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return errors.Wrap(err, "detect mime type")
	} else {
		w.Header().Set("Content-Type", mime.String())
	}

	http.ServeContent(w, r, file.DisplayPath(), time.Time{}, reader)
	return nil
}
