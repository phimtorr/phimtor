package yts

import (
	"strconv"
	"time"
)

type Movie struct {
	ID       int64
	Torrents []Torrent
}

type Torrent struct {
	Hash             string
	Quality          string
	Resolution       int
	Type             string
	IsRepack         bool
	VideoCodec       string
	BitDepth         int
	AudioChannels    string
	Seeds            int
	Peers            int
	SizeBytes        int64
	DateUploadedUnix time.Time
}

type responseMovie struct {
	ID       int64             `json:"id"`
	Torrents []responseTorrent `json:"torrents"`
}

type responseTorrent struct {
	Hash             string `json:"hash"`
	Quality          string `json:"quality"`
	Type             string `json:"type"`
	IsRepack         string `json:"is_repack"`
	VideoCodec       string `json:"video_codec"`
	BitDepth         string `json:"bit_depth"`
	AudioChannels    string `json:"audio_channels"`
	Seeds            int    `json:"seeds"`
	Peers            int    `json:"peers"`
	SizeBytes        int64  `json:"size_bytes"`
	DateUploadedUnix int64  `json:"date_uploaded_unix"`
}

func toMovie(movie responseMovie) (Movie, error) {
	return Movie{
		ID:       movie.ID,
		Torrents: toTorrents(movie.Torrents),
	}, nil
}

func toTorrents(torrents []responseTorrent) []Torrent {
	ts := make([]Torrent, 0, len(torrents))
	for _, t := range torrents {
		tt := toTorrent(t)
		// ignore torrents with 0 resolution (3D, etc.)
		if tt.Resolution == 0 {
			continue
		}
		ts = append(ts, tt)
	}
	return ts
}

func toTorrent(t responseTorrent) Torrent {
	var resolution int
	switch t.Quality {
	case "480p":
		resolution = 480
	case "720p":
		resolution = 720
	case "1080p":
		resolution = 1080
	case "2160p":
		resolution = 2160
	}

	bitDepth, _ := strconv.Atoi(t.BitDepth)
	return Torrent{
		Hash:             t.Hash,
		Quality:          t.Quality,
		Resolution:       resolution,
		Type:             t.Type,
		IsRepack:         t.IsRepack == "1",
		VideoCodec:       t.VideoCodec,
		BitDepth:         bitDepth,
		AudioChannels:    t.AudioChannels,
		Seeds:            t.Seeds,
		Peers:            t.Peers,
		SizeBytes:        t.SizeBytes,
		DateUploadedUnix: time.Unix(t.DateUploadedUnix, 0),
	}
}
