package repository

import (
	"context"
	"slices"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/ports"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r Repository) GetVideo(ctx context.Context, id int64) (ports.Video, error) {
	dbVideo, err := dbmodels.Videos(
		dbmodels.VideoWhere.ID.EQ(id),
		qm.Load(dbmodels.VideoRels.TorrentLinks),
		qm.Load(dbmodels.VideoRels.Subtitles),
		qm.Load(dbmodels.VideoRels.Show),
		qm.Load(dbmodels.VideoRels.Episode),
		qm.Load(qm.Rels(dbmodels.VideoRels.Episode, dbmodels.EpisodeRels.Show)),
	).One(ctx, r.db)
	if err != nil {
		return ports.Video{}, err
	}
	return toHTTPVideo(dbVideo), nil
}

func toHTTPVideo(dbVid *dbmodels.Video) ports.Video {
	return ports.Video{
		Id:           dbVid.ID,
		Subtitles:    toHTTPSubtitles(dbVid.R.Subtitles),
		Title:        toVideoTitle(dbVid),
		TorrentLinks: toHTTPTorrentLinks(dbVid.R.TorrentLinks),
	}
}

func toHTTPTorrentLinks(dbLinks dbmodels.TorrentLinkSlice) []ports.TorrentLink {
	var links []ports.TorrentLink
	for _, link := range dbLinks {
		links = append(links, ports.TorrentLink{
			Link:      link.Link,
			Name:      link.Name,
			FileIndex: link.FileIndex,
			Priority:  link.Priority,
		})
	}
	slices.SortStableFunc(links, func(a, b ports.TorrentLink) int {
		// sort by priority, descending
		return b.Priority - a.Priority
	})
	return links
}

func toHTTPSubtitles(dbSubs dbmodels.SubtitleSlice) []ports.Subtitle {
	var subs []ports.Subtitle
	for _, sub := range dbSubs {
		subs = append(subs, ports.Subtitle{
			Language: sub.Language,
			Link:     sub.Link,
			Name:     sub.Name,
			Owner:    sub.Owner,
		})
	}
	return subs
}

func toVideoTitle(dbVideo *dbmodels.Video) string {
	if dbVideo.R.Show != nil {
		return dbVideo.R.Show.Title
	}
	if dbVideo.R.Episode != nil {
		return dbVideo.R.Episode.R.Show.Title + " - " + dbVideo.R.Episode.Name
	}
	return ""
}
