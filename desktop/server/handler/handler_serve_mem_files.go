package handler

import (
	"bytes"
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/gabriel-vasile/mimetype"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	commonErrors "github.com/phimtorr/phimtor/common/errors"
)

func (h *Handler) ServeMemoryFile(w http.ResponseWriter, r *http.Request) error {
	rawID := chi.URLParam(r, "id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		return commonErrors.NewIncorrectInputError("invalid-uuid", "invalid uuid")
	}

	file, ok := h.memStorage.Get(id)
	if !ok {
		return commonErrors.NewIncorrectInputError("file-not-found", "file not found")
	}

	mine := mimetype.Detect(file.Content)
	w.Header().Set("Content-Type", mine.String())

	_, err = bytes.NewBuffer(file.Content).WriteTo(w)
	if err != nil {
		return errors.Wrap(err, "write to response")
	}

	return nil
}
