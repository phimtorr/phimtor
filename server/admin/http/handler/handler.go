package handler

import (
	"context"
	"io"
	"net/http"
)

type FileService interface {
	UploadFile(ctx context.Context, key string, body io.Reader) (string, error)
	DeleteFile(ctx context.Context, key string) error
}

type Handler struct {
	repo        Repository
	fileService FileService
}

func New(repo Repository, fileService FileService) *Handler {
	if repo == nil {
		panic("nil repository")

	}
	if fileService == nil {
		panic("nil file service")
	}

	return &Handler{
		repo:        repo,
		fileService: fileService,
	}
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
