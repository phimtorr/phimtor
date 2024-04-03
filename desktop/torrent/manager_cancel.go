package torrent

func (m *Manager) DropAll() {
	for _, tor := range m.client.Torrents() {
		tor.Drop()
	}
	return
}
