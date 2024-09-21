package repository

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r Repository) SyncVideo(ctx context.Context, videoID int64) error {
	video, err := dbmodels.Videos(
		dbmodels.VideoWhere.ID.EQ(videoID),
		qm.Load(dbmodels.VideoRels.TorrentLinks),
		qm.Load(dbmodels.VideoRels.Subtitles),
	).One(ctx, r.db)
	if err != nil {
		return fmt.Errorf("find video: %w", err)
	}

	var maxResolution int
	for _, link := range video.R.TorrentLinks {
		if link.Resolution > maxResolution {
			maxResolution = link.Resolution
		}
	}

	video.MaxResolution = maxResolution

	var hasViSub, hasEnSub bool
	for _, sub := range video.R.Subtitles {
		if sub.Language == "vi" {
			hasViSub = true
		}
		if sub.Language == "en" {
			hasEnSub = true
		}
	}

	video.HasViSub = hasViSub
	video.HasEnSub = hasEnSub

	if _, err := video.Update(ctx, r.db, boil.Infer()); err != nil {
		return fmt.Errorf("update video: %w", err)
	}

	return nil
}
