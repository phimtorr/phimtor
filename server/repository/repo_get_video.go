package repository

import (
	"context"
	"slices"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/auth"
	"github.com/phimtorr/phimtor/server/http"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r Repository) GetVideo(ctx context.Context, user auth.User, id int64) (http.Video, error) {
	dbVideo, err := dbmodels.Videos(
		dbmodels.VideoWhere.ID.EQ(id),
		qm.Load(dbmodels.VideoRels.TorrentLinks),
		qm.Load(dbmodels.VideoRels.Subtitles),
		qm.Load(dbmodels.VideoRels.Show),
		qm.Load(dbmodels.VideoRels.Episode),
		qm.Load(qm.Rels(dbmodels.VideoRels.Episode, dbmodels.EpisodeRels.Show)),
	).One(ctx, r.db)
	if err != nil {
		return http.Video{}, err
	}
	return toHTTPVideo(user, dbVideo), nil
}

func toHTTPVideo(user auth.User, dbVid *dbmodels.Video) http.Video {
	return http.Video{
		Id:                  dbVid.ID,
		Subtitles:           toHTTPSubtitles(dbVid.R.Subtitles),
		Title:               toVideoTitle(dbVid),
		TorrentLinks:        toHTTPTorrentLinks(user, dbVid.R.TorrentLinks),
		PremiumTorrentLinks: toHTTPPremiumTorrentLinks(user, dbVid.R.TorrentLinks),
	}
}

func toHTTPTorrentLinks(user auth.User, dbLinks dbmodels.TorrentLinkSlice) []http.TorrentLink {
	links := make([]http.TorrentLink, 0, len(dbLinks))
	for _, link := range dbLinks {
		if link.RequiredPremium && !user.IsPremium() {
			continue
		}
		links = append(links, http.TorrentLink{
			Id:             link.ID,
			Link:           link.Link,
			Name:           link.Name,
			FileIndex:      link.FileIndex,
			Priority:       link.Priority,
			RequirePremium: link.RequiredPremium,
		})
	}
	slices.SortStableFunc(links, func(a, b http.TorrentLink) int {
		// sort by priority, descending
		return b.Priority - a.Priority
	})
	return links
}

func toHTTPPremiumTorrentLinks(user auth.User, dbLinks dbmodels.TorrentLinkSlice) []http.PremiumTorrentLink {
	links := make([]http.PremiumTorrentLink, 0, len(dbLinks))
	if user.IsPremium() {
		return links // no need to filter, because user is premium, all links are available
	}
	for _, link := range dbLinks {
		if !link.RequiredPremium {
			continue
		}
		links = append(links, http.PremiumTorrentLink{
			Id:       link.ID,
			Name:     link.Name,
			Priority: link.Priority,
		})
	}
	slices.SortStableFunc(links, func(a, b http.PremiumTorrentLink) int {
		// sort by priority, descending
		return b.Priority - a.Priority
	})
	return links

}

func toHTTPSubtitles(dbSubs dbmodels.SubtitleSlice) []http.Subtitle {
	subs := make([]http.Subtitle, 0, len(dbSubs))
	for _, sub := range dbSubs {
		subs = append(subs, http.Subtitle{
			Id:       sub.ID,
			Language: sub.Language,
			Link:     sub.Link,
			Name:     sub.Name,
			Owner:    sub.Owner,
			Priority: sub.Priority,
		})
	}
	slices.SortStableFunc(subs, func(a, b http.Subtitle) int {
		// sort by priority, descending
		return b.Priority - a.Priority
	})
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
