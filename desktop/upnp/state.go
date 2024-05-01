package upnp

import (
	"path/filepath"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/huin/goupnp/dcps/av1"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
	"github.com/phimtorr/phimtor/desktop/subtitle"
	"github.com/phimtorr/phimtor/desktop/torrent"
)

type State struct {
	AvailableClients []*av1.AVTransport1
	SelectedUDN      string

	InfoHash         torrent.InfoHash
	FileIndex        int
	SubtitleFileName string
	SubtitleContent  []byte
}

func (s *State) Reset() {
	s.InfoHash = torrent.InfoHash{}
	s.FileIndex = 0
	s.SubtitleFileName = ""
	s.SubtitleContent = nil
}

func (s *State) GetSelectedClient() (*av1.AVTransport1, error) {
	if s.SelectedUDN == "" {
		return nil, commonErrors.NewIncorrectInputError("no-selected-device", "no selected device")
	}

	for _, client := range s.AvailableClients {
		if client.RootDevice.Device.UDN == s.SelectedUDN {
			return client, nil
		}
	}

	return nil, commonErrors.NewIncorrectInputError("device-not-found", "device not found")
}

func (s *State) SetSubtitle(fileName string, content []byte) error {
	s.SubtitleFileName = strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".srt" // force to srt

	srtContent, err := subtitle.NormalizeToSRT(fileName, content)
	if err != nil {
		return errors.Wrap(err, "normalize subtitle to srt")
	}
	s.SubtitleContent = srtContent
	return nil
}
