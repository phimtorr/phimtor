package repository

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/admin/http/handler"
	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r Repository) GetVideo(ctx context.Context, id int64) (ui.Video, error) {
	dbVideos, err := dbmodels.Videos(
		dbmodels.VideoWhere.ID.EQ(id),
		qm.Load(dbmodels.VideoRels.TorrentLinks),
		qm.Load(dbmodels.VideoRels.Subtitles),
	).One(ctx, r.db)
	if err != nil {
		return ui.Video{}, err
	}

	return toUIVideo(dbVideos), nil
}

func toUIVideo(vid *dbmodels.Video) ui.Video {
	torrents := make([]ui.Torrent, 0, len(vid.R.TorrentLinks))
	for _, t := range vid.R.TorrentLinks {
		torrents = append(torrents, ui.Torrent{
			ID:              t.ID,
			Name:            toTorrentLinkName(t),
			Link:            t.Link,
			FileIndex:       t.FileIndex,
			Priority:        t.Priority,
			RequiredPremium: t.RequiredPremium,
		})
	}

	subtitles := make([]ui.Subtitle, 0, len(vid.R.Subtitles))
	for _, s := range vid.R.Subtitles {
		subtitles = append(subtitles, toUISubtitle(s))
	}

	return ui.Video{
		ID:            vid.ID,
		MaxResolution: vid.MaxResolution,
		HasViSub:      vid.HasViSub,
		HasEnSub:      vid.HasEnSub,
		Torrents:      torrents,
		Subtitles:     subtitles,
	}
}

func toTorrentLinkName(t *dbmodels.TorrentLink) string {
	name := fmt.Sprintf("%dp.%s", t.Resolution, t.Type)
	if t.Codec != "" {
		name += "." + t.Codec
	}
	if t.Source != "" {
		name += "." + t.Source
	}
	return name
}

func (r Repository) CreateTorrent(ctx context.Context, torrent handler.TorrentToCreate) (int64, error) {
	dbTorrentLink := &dbmodels.TorrentLink{
		VideoID:         torrent.VideoID,
		Resolution:      torrent.Resolution,
		Type:            torrent.Type,
		Codec:           torrent.Codec,
		Source:          torrent.Source,
		Link:            torrent.Link,
		FileIndex:       torrent.FileIndex,
		Priority:        torrent.Priority,
		RequiredPremium: torrent.RequiredPremium,
	}
	if err := dbTorrentLink.Insert(ctx, r.db, boil.Infer()); err != nil {
		return 0, err
	}
	return dbTorrentLink.ID, nil
}

func (r Repository) DeleteTorrent(ctx context.Context, videoID, id int64) error {
	_, err := dbmodels.TorrentLinks(
		dbmodels.TorrentLinkWhere.ID.EQ(id),
		dbmodels.TorrentLinkWhere.VideoID.EQ(videoID),
	).DeleteAll(ctx, r.db)
	return err
}

func (r Repository) GetSubtitle(ctx context.Context, videoID, id int64) (ui.Subtitle, error) {
	dbSubtitle, err := dbmodels.Subtitles(
		dbmodels.SubtitleWhere.ID.EQ(id),
		dbmodels.SubtitleWhere.VideoID.EQ(videoID),
	).One(ctx, r.db)
	if err != nil {
		return ui.Subtitle{}, err
	}
	return toUISubtitle(dbSubtitle), nil
}

func (r Repository) CreateSubtitle(ctx context.Context, subtitle handler.SubtitleToCreate) (int64, error) {
	dbSubtitle := &dbmodels.Subtitle{
		VideoID:  subtitle.VideoID,
		Language: subtitle.Language,
		Name:     subtitle.Name,
		Owner:    subtitle.Owner,
		Link:     subtitle.Link,
		Priority: subtitle.Priority,
		FileKey:  subtitle.FileKey,
	}
	if err := dbSubtitle.Insert(ctx, r.db, boil.Infer()); err != nil {
		return 0, err
	}
	return dbSubtitle.ID, nil
}

func (r Repository) DeleteSubtitle(ctx context.Context, videoID, id int64) error {
	_, err := dbmodels.Subtitles(
		dbmodels.SubtitleWhere.ID.EQ(id),
		dbmodels.SubtitleWhere.VideoID.EQ(videoID),
	).DeleteAll(ctx, r.db)
	return err
}

func toUISubtitle(sub *dbmodels.Subtitle) ui.Subtitle {
	return ui.Subtitle{
		ID:       sub.ID,
		Language: sub.Language,
		Name:     sub.Name,
		Owner:    sub.Owner,
		Link:     sub.Link,
		FileKey:  sub.FileKey,
		Priority: sub.Priority,
	}
}
