package handler

import (
	"context"

	"github.com/phimtorr/phimtor/server/admin/http/ui"
)

type Repository interface {
	GetVideo(ctx context.Context, id int64) (ui.Video, error)

	CreateTorrent(ctx context.Context, torrent TorrentToCreate) (int64, error)
	DeleteTorrent(ctx context.Context, videoID, id int64) error

	GetSubtitle(ctx context.Context, videoID, id int64) (ui.Subtitle, error)
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
	VideoID         int64
	Name            string
	Link            string
	FileIndex       int
	Priority        int
	RequiredPremium bool
}

type SubtitleToCreate struct {
	VideoID  int64
	Language string
	Name     string
	Owner    string
	Link     string
	FileKey  string
	Priority int
}

type EpisodeToCreate struct {
	ShowID int64
	Name   string
}
