package handler

import (
	"net/http"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/gabriel-vasile/mimetype"
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

	file, err := h.torManager.GetFile(infoHash, fileIndex)
	if err != nil {
		return errors.Wrap(err, "get file")
	}

	file.Download()
	reader := file.NewReader()
	reader.SetResponsive()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return errors.Wrap(err, "detect mime type")
	} else {
		w.Header().Set("Content-Type", mime.String())
	}

	http.ServeContent(w, r, file.DisplayPath(), time.Time{}, reader)
	return nil
}
