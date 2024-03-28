package handler

import (
	"context"

	"github.com/phimtorr/phimtor/server/admin/ui"
)

type Repository interface {
	ListShows(ctx context.Context, page int, pageSize int) ([]ui.Show, ui.Pagination, error)
}

type Handler struct {
	repo Repository
}

func New(repo Repository) *Handler {
	if repo == nil {
		panic("nil repository")

	}
	return &Handler{
		repo: repo,
	}
}
