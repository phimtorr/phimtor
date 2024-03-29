package handler

import (
	"context"

	"github.com/phimtorr/phimtor/server/admin/ui"
)

type Repository interface {
	ListShowDisplays(ctx context.Context, page int, pageSize int) ([]ui.ShowDisplay, ui.Pagination, error)

	CreateShow(ctx context.Context, show ShowToCreate) (int64, error)
	UpdateShow(ctx context.Context, show ShowToUpdate) error
	GetShow(ctx context.Context, id int64) (ui.Show, error)
}

type ShowToCreate struct {
	ShowType          string
	Title             string
	OriginalTitle     string
	PosterLink        string
	Description       string
	ReleaseYear       int
	Score             float64
	DurationInMinutes int
	Quality           string
	TotalEpisodes     int
}

type ShowToUpdate struct {
	ID                int64
	Title             string
	OriginalTitle     string
	PosterLink        string
	Description       string
	ReleaseYear       int
	Score             float64
	DurationInMinutes int
	Quality           string
	TotalEpisodes     int
}
