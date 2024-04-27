package handler

import "github.com/phimtorr/phimtor/desktop/client/api"

const (
	stateKeyTorrent = "torrent"
)

type TorrentState struct {
	SelectedTorrent api.TorrentLink
}

func (h *Handler) SetTorrentState(s TorrentState) {
	h.state.Set(stateKeyTorrent, s)
}

func (h *Handler) GetTorrentState() (TorrentState, bool) {
	t, ok := h.state.Get(stateKeyTorrent)
	if !ok {
		return TorrentState{}, false
	}
	return t.(TorrentState), true
}
