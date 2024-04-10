package torrent

import (
	"fmt"

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
