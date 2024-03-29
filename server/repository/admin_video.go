package repository

import (
	"context"

	"github.com/phimtorr/phimtor/server/admin/handler"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/phimtorr/phimtor/server/admin/ui"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r AdminRepository) GetVideo(ctx context.Context, id int64) (ui.Video, error) {
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
			ID:        t.ID,
			Name:      t.Name,
			Link:      t.Link,
			FileIndex: t.FileIndex,
			Priority:  t.Priority,
		})
	}

	subtitles := make([]ui.Subtitle, 0, len(vid.R.Subtitles))
	for _, s := range vid.R.Subtitles {
		subtitles = append(subtitles, ui.Subtitle{
			ID:       s.ID,
			Language: s.Language,
			Name:     s.Name,
			Owner:    s.Owner,
			Link:     s.Link,
		})
	}

	return ui.Video{
		ID:        vid.ID,
		Torrents:  torrents,
		Subtitles: subtitles,
	}
}

func (r AdminRepository) CreateTorrent(ctx context.Context, torrent handler.TorrentToCreate) (int64, error) {
	dbTorrentLink := &dbmodels.TorrentLink{
		VideoID:   torrent.VideoID,
		Name:      torrent.Name,
		Link:      torrent.Link,
		FileIndex: torrent.FileIndex,
		Priority:  torrent.Priority,
	}
	if err := dbTorrentLink.Insert(ctx, r.db, boil.Infer()); err != nil {
		return 0, err
	}
	return dbTorrentLink.ID, nil
}

func (r AdminRepository) DeleteTorrent(ctx context.Context, videoID, id int64) error {
	_, err := dbmodels.TorrentLinks(
		dbmodels.TorrentLinkWhere.ID.EQ(id),
		dbmodels.TorrentLinkWhere.VideoID.EQ(videoID),
	).DeleteAll(ctx, r.db)
	return err
}

func (r AdminRepository) CreateSubtitle(ctx context.Context, subtitle handler.SubtitleToCreate) (int64, error) {
	dbSubtitle := &dbmodels.Subtitle{
		VideoID:  subtitle.VideoID,
		Language: subtitle.Language,
		Name:     subtitle.Name,
		Owner:    subtitle.Owner,
		Link:     subtitle.Link,
		FileKey:  subtitle.FileKey,
	}
	if err := dbSubtitle.Insert(ctx, r.db, boil.Infer()); err != nil {
		return 0, err
	}
	return dbSubtitle.ID, nil
}
