package torrent

import "github.com/anacrolix/torrent/metainfo"

// CancelOthers cancels all pieces of all torrents except the one with the given info hash.
func (m *Manager) CancelOthers(infoHash InfoHash) {
	for _, tor := range m.client.Torrents() {
		if tor.InfoHash() != metainfo.Hash(infoHash) {
			tor.Drop()
		}
	}

	return
}
