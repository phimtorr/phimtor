package upnp

import (
	"github.com/huin/goupnp/dcps/av1"
	"github.com/phimtorr/phimtor/desktop/torrent"
)

type State struct {
	AvailableClients []*av1.AVTransport1

	InfoHash         torrent.InfoHash
	FileIndex        int
	SubtitleFileName string
	SubtitleContent  []byte
	SelectedUDN      string
}

func (s *State) Reset() {
	s.InfoHash = torrent.InfoHash{}
	s.FileIndex = 0
	s.SubtitleFileName = ""
	s.SubtitleContent = nil
	s.SelectedUDN = ""
}
