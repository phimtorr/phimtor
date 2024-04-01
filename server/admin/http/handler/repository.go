package handler

import (
	"context"

	ui2 "github.com/phimtorr/phimtor/server/admin/http/ui"
)

type Repository interface {
	ListShowDisplays(ctx context.Context, page int, pageSize int) ([]ui2.ShowDisplay, ui2.Pagination, error)

	CreateShow(ctx context.Context, show ShowToCreate) (int64, error)
	UpdateShow(ctx context.Context, show ShowToUpdate) error
	GetShow(ctx context.Context, id int64) (ui2.Show, error)

	GetVideo(ctx context.Context, id int64) (ui2.Video, error)
	CreateTorrent(ctx context.Context, torrent TorrentToCreate) (int64, error)
	DeleteTorrent(ctx context.Context, videoID, id int64) error

	GetSubtitle(ctx context.Context, videoID, id int64) (ui2.Subtitle, error)
	CreateSubtitle(ctx context.Context, subtitle SubtitleToCreate) (int64, error)
	DeleteSubtitle(ctx context.Context, videoID, id int64) error
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

type TorrentToCreate struct {
	VideoID   int64
	Name      string
	Link      string
	FileIndex int
	Priority  int
}

type SubtitleToCreate struct {
	VideoID  int64
	Language string
	Name     string
	Owner    string
	Link     string
	FileKey  string
}
