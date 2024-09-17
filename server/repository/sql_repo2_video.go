package repository

import (
	"context"
	"slices"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/auth"
	"github.com/phimtorr/phimtor/server/http2"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r SQLRepo2) GetVideo(ctx context.Context, user auth.User, id int64) (http2.Video, error) {
	dbVideo, err := dbmodels.Videos(
		dbmodels.VideoWhere.ID.EQ(id),
		qm.Load(dbmodels.VideoRels.TorrentLinks),
		qm.Load(dbmodels.VideoRels.Subtitles),
	).One(ctx, r.db)
	if err != nil {
		return http2.Video{}, err
	}
	return toHTTP2Video(user, dbVideo), nil
}

func toHTTP2Video(user auth.User, dbVid *dbmodels.Video) http2.Video {
	return http2.Video{
		Id:                  dbVid.ID,
		Subtitles:           toHTTP2Subtitles(dbVid.R.Subtitles),
		TorrentLinks:        toHTTP2TorrentLinks(user, dbVid.R.TorrentLinks),
		PremiumTorrentLinks: toHTTP2PremiumTorrentLinks(user, dbVid.R.TorrentLinks),
	}
}

func toHTTP2TorrentLinks(user auth.User, dbLinks dbmodels.TorrentLinkSlice) []http2.TorrentLink {
	links := make([]http2.TorrentLink, 0, len(dbLinks))
	for _, link := range dbLinks {
		if link.RequiredPremium && !user.IsPremium() {
			continue
		}
		links = append(links, http2.TorrentLink{
			Id:             link.ID,
			Link:           link.Link,
			Name:           link.Name,
			FileIndex:      link.FileIndex,
			Priority:       link.Priority,
			RequirePremium: link.RequiredPremium,
		})
	}
	slices.SortStableFunc(links, func(a, b http2.TorrentLink) int {
		// sort by priority, descending
		return b.Priority - a.Priority
	})
	return links
}

func toHTTP2PremiumTorrentLinks(user auth.User, dbLinks dbmodels.TorrentLinkSlice) []http2.PremiumTorrentLink {
	links := make([]http2.PremiumTorrentLink, 0, len(dbLinks))
	if user.IsPremium() {
		return links // no need to filter, because user is premium, all links are available
	}
	for _, link := range dbLinks {
		if !link.RequiredPremium {
			continue
		}
		links = append(links, http2.PremiumTorrentLink{
			Id:       link.ID,
			Name:     link.Name,
			Priority: link.Priority,
		})
	}
	slices.SortStableFunc(links, func(a, b http2.PremiumTorrentLink) int {
		// sort by priority, descending
		return b.Priority - a.Priority
	})
	return links

}

func toHTTP2Subtitles(dbSubs dbmodels.SubtitleSlice) []http2.Subtitle {
	subs := make([]http2.Subtitle, 0, len(dbSubs))
	for _, sub := range dbSubs {
		subs = append(subs, http2.Subtitle{
			Id:       sub.ID,
			Language: sub.Language,
			Link:     sub.Link,
			Name:     sub.Name,
			Owner:    sub.Owner,
			Priority: sub.Priority,
		})
	}
	slices.SortStableFunc(subs, func(a, b http2.Subtitle) int {
		// sort by priority, descending
		return b.Priority - a.Priority
	})
	return subs
}
