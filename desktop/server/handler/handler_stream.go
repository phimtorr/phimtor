package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) Stream(w http.ResponseWriter, r *http.Request) error {
	infoHash, err := parseInfoHash(chi.URLParam(r, "infoHash"))
	if err != nil {
		return err
	}
	fileIndex, err := parseFileIndex(chi.URLParam(r, "fileIndex"))
	if err != nil {
		return err
	}

	return h.torManager.StreamVideoFile(w, r, infoHash, fileIndex)
}
