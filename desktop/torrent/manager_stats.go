package torrent

type Stats struct {
	TotalPeers       int
	PendingPeers     int
	ActivePeers      int
	ConnectedSeeders int
	HalfOpenPeers    int
	Length           int64
	BytesCompleted   int64
}

func (m *Manager) Stats(infoHash InfoHash, fileIndex int) Stats {
	tor, ok := m.GetTorrent(infoHash)
	if !ok {
		return Stats{}
	}
	stats := tor.Stats()
	file := tor.Files()[fileIndex]
	return Stats{
		TotalPeers:       stats.TotalPeers,
		PendingPeers:     stats.PendingPeers,
		ActivePeers:      stats.ActivePeers,
		ConnectedSeeders: stats.ConnectedSeeders,
		HalfOpenPeers:    stats.HalfOpenPeers,
		Length:           file.Length(),
		BytesCompleted:   file.BytesCompleted(),
	}
}
