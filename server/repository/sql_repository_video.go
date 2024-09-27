package repository

import (
	"context"
	"fmt"
	"net/url"
	"slices"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/auth"
	"github.com/phimtorr/phimtor/server/http"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

const (
	Resolution4K = 2160
)

func (r SQLRepository) GetVideo(ctx context.Context, user auth.User, id int64) (http.Video, error) {
	dbVideo, err := dbmodels.Videos(
		dbmodels.VideoWhere.ID.EQ(id),
		qm.Load(dbmodels.VideoRels.MovieYtsTorrents),
		qm.Load(dbmodels.VideoRels.TorrentLinks),
		qm.Load(dbmodels.VideoRels.Subtitles),
	).One(ctx, r.db)
	if err != nil {
		return http.Video{}, err
	}
	return r.toHTTP2Video(user, dbVideo), nil
}

func (r SQLRepository) toHTTP2Video(user auth.User, dbVid *dbmodels.Video) http.Video {
	torrentLinks, premiumTorrentLinks := r.toHTTP2TorrentLinks2(user, dbVid.R.MovieYtsTorrents, dbVid.R.TorrentLinks)
	return http.Video{
		Id:                  dbVid.ID,
		Subtitles:           toHTTP2Subtitles(dbVid.R.Subtitles),
		TorrentLinks:        torrentLinks,
		PremiumTorrentLinks: premiumTorrentLinks,
	}
}

func (r SQLRepository) toHTTP2TorrentLinks2(
	user auth.User, ytsMovieTorrents []*dbmodels.YtsTorrent, torrents []*dbmodels.TorrentLink,
) ([]http.TorrentLink, []http.PremiumTorrentLink) {
	slices.SortStableFunc(ytsMovieTorrents, func(a, b *dbmodels.YtsTorrent) int {
		return b.Seeds - a.Seeds // sort by seeds, descending
	})
	slices.SortStableFunc(torrents, func(a, b *dbmodels.TorrentLink) int {
		return b.Priority - a.Priority // sort by priority, descending
	})

	links := make([]http.TorrentLink, 0, len(torrents))
	premiumLinks := make([]http.PremiumTorrentLink, 0, len(torrents))

	for i, link := range ytsMovieTorrents {
		id := (link.MovieID << 32) + int64(i)
		if isExclude(user, link.Resolution) {
			premiumLinks = append(premiumLinks, http.PremiumTorrentLink{
				Id:       id,
				Name:     toYTSTorrentName(link),
				Priority: link.Seeds,
			})
		}
		links = append(links, http.TorrentLink{
			Id:             id,
			Link:           r.ToMagnetLink(link.Hash),
			Name:           toYTSTorrentName(link),
			FileIndex:      0,
			Priority:       link.Seeds,
			RequirePremium: false,
		})
	}

	for _, link := range torrents {
		if isExclude(user, link.Resolution) {
			premiumLinks = append(premiumLinks, http.PremiumTorrentLink{
				Id:       link.ID,
				Name:     toTorrentLinkName(link),
				Priority: link.Priority,
			})
		} else {
			links = append(links, http.TorrentLink{
				Id:             link.ID,
				Link:           link.Link,
				Name:           toTorrentLinkName(link),
				FileIndex:      link.FileIndex,
				Priority:       link.Priority,
				RequirePremium: link.RequiredPremium,
			})
		}
	}

	return links, premiumLinks
}

func isExclude(user auth.User, resolution int) bool {
	if user.IsPremium() {
		return false
	}
	if resolution >= Resolution4K {
		return true
	}
	return false
}

func toYTSTorrentName(torrent *dbmodels.YtsTorrent) string {
	name := fmt.Sprintf("%s.%s", torrent.Quality, torrent.Type)
	if torrent.VideoCodec != "x264" {
		name += "." + torrent.VideoCodec
	}
	if torrent.IsRepack {
		name += ".REPACK"
	}

	return name
}

func toTorrentLinkName(link *dbmodels.TorrentLink) string {
	name := fmt.Sprintf("%dp.%s", link.Resolution, link.Type)
	if link.Codec != "" {
		name += "." + link.Codec
	}
	if link.Source != "" {
		name += "." + link.Source
	}
	return name
}

func toHTTP2Subtitles(dbSubs dbmodels.SubtitleSlice) []http.Subtitle {
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

func (r SQLRepository) ToMagnetLink(torrentHash string) string {
	v := url.Values{}
	v.Set("dn", torrentHash)
	for _, tracker := range r.ytsTrackers {
		v.Add("tr", tracker)
	}

	return fmt.Sprintf("magnet:?xt=urn:btih:%s&%s", torrentHash, v.Encode())
}
