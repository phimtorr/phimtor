package repository

import (
	"context"
	"fmt"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/phimtorr/phimtor/server/admin/yts"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r TMDBRepository) UpdateYTSMovie(ctx context.Context, videoID int64, movie yts.Movie) error {
	video, err := dbmodels.FindVideo(ctx, r.db, videoID)
	if err != nil {
		return fmt.Errorf("find video: %w", err)
	}

	video.YtsMovieID = null.NewInt64(movie.ID, true)

	if _, err := video.Update(ctx, r.db, boil.Whitelist(dbmodels.VideoColumns.YtsMovieID)); err != nil {
		return fmt.Errorf("update video: %w", err)
	}

	// Update torrents
	for _, t := range movie.Torrents {
		if err := r.updateYTSTorrent(ctx, movie.ID, t); err != nil {
			return fmt.Errorf("update yts torrent %s: %w", t.Hash, err)
		}
	}

	if err := r.getVideoRepo().SyncVideo(ctx, videoID); err != nil {
		return fmt.Errorf("sync video: %w", err)
	}

	return nil
}

var updateYTSTorrentQuery = `
INSERT INTO yts_torrents (
				hash, 
				movie_id, 
				quality, 
				resolution, 
				type, 
				is_repack, 
				video_codec, 
				bit_depth, 
				audio_channels, 
				seeds, 
				peers, 
				size_bytes, 
				date_uploaded
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
ON DUPLICATE KEY UPDATE
				quality = VALUES(quality),
				resolution = VALUES(resolution),
				type = VALUES(type),
				is_repack = VALUES(is_repack),
				video_codec = VALUES(video_codec),
				bit_depth = VALUES(bit_depth),	
				audio_channels = VALUES(audio_channels),
				seeds = VALUES(seeds),
				peers = VALUES(peers),
				size_bytes = VALUES(size_bytes),
				date_uploaded = VALUES(date_uploaded)
`

func (r TMDBRepository) updateYTSTorrent(ctx context.Context, movieID int64, torrent yts.Torrent) error {
	_, err := r.db.ExecContext(ctx, updateYTSTorrentQuery,
		torrent.Hash,
		movieID,
		torrent.Quality,
		torrent.Resolution,
		torrent.Type,
		torrent.IsRepack,
		torrent.VideoCodec,
		torrent.BitDepth,
		torrent.AudioChannels,
		torrent.Seeds,
		torrent.Peers,
		torrent.SizeBytes,
		torrent.DateUploadedUnix,
	)
	if err != nil {
		return fmt.Errorf("update yts torrent: %w", err)
	}

	return nil
}
