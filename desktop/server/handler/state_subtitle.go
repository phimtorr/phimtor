package handler

const (
	stateKeySubtitle = "subtitle"
)

type SubtitleState struct {
	ID              int64
	Name            string
	FileName        string
	OriginalContent []byte
}

func (h *Handler) SetSubtitleState(s SubtitleState) {
	h.state.Set(stateKeySubtitle, s)
}

func (h *Handler) GetSubtitleState() (SubtitleState, bool) {
	t, ok := h.state.Get(stateKeySubtitle)
	if !ok {
		return SubtitleState{}, false
	}
	return t.(SubtitleState), true
}
