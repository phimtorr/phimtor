package repository

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/ports"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r Repository) GetSeries(ctx context.Context, id int64) (ports.Series, error) {
	dbSeries, err := dbmodels.Shows(
		dbmodels.ShowWhere.ID.EQ(id),
		dbmodels.ShowWhere.Type.EQ(dbmodels.ShowsTypeSeries),
		qm.Load(dbmodels.ShowRels.Episodes),
	).One(ctx, r.db)
	if err != nil {
		return ports.Series{}, errors.Wrap(err, "get series")
	}

	return toHTTPSeries(dbSeries), nil
}

func toHTTPSeries(dbSeries *dbmodels.Show) ports.Series {
	return ports.Series{
		CurrentEpisode:    dbSeries.CurrentEpisode,
		Description:       dbSeries.Description,
		DurationInMinutes: dbSeries.DurationInMinutes,
		Episodes:          toHTTPEpisodes(dbSeries.R.Episodes),
		Id:                dbSeries.ID,
		OriginalTitle:     dbSeries.OriginalTitle,
		PosterLink:        dbSeries.PosterLink,
		ReleaseYear:       dbSeries.ReleaseYear,
		Score:             float32(dbSeries.Score),
		Title:             dbSeries.Title,
		TotalEpisodes:     dbSeries.TotalEpisodes,
	}
}

func toHTTPEpisodes(dbEpisodes dbmodels.EpisodeSlice) []ports.Episode {
	episodes := make([]ports.Episode, len(dbEpisodes))
	for i, dbEpisode := range dbEpisodes {
		episodes[i] = toHTTPEpisode(dbEpisode)
	}
	return episodes
}

func toHTTPEpisode(dbEpisode *dbmodels.Episode) ports.Episode {
	return ports.Episode{
		Id:      dbEpisode.ID,
		Name:    dbEpisode.Name,
		VideoId: dbEpisode.VideoID,
	}
}
