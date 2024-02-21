package torrent

// CancelOthers cancels all pieces of all torrents except the one with the given info hash.
func (m *Manager) CancelOthers(infoHash InfoHash) {
	for _, tor := range m.client.Torrents() {
		if tor.InfoHash().String() != infoHash.String() {
			tor.CancelPieces(0, tor.NumPieces())
		}
	}

	return
}
