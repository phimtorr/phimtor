package repository

import (
	"context"
	"database/sql"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/phimtorr/phimtor/server/admin/http/handler"
	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/repository"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r Repository) ListEpisodes(ctx context.Context, showID int64) ([]ui.Episode, error) {
	dbEpisodes, err := dbmodels.Episodes(
		dbmodels.EpisodeWhere.ShowID.EQ(showID),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	episodes := make([]ui.Episode, len(dbEpisodes))
	for i, dbEpisode := range dbEpisodes {
		episodes[i] = ui.Episode{
			ID:      dbEpisode.ID,
			Name:    dbEpisode.Name,
			VideoID: dbEpisode.VideoID,
		}
	}

	return episodes, nil
}

func (r Repository) CreateEpisode(ctx context.Context, episode handler.EpisodeToCreate) (int64, error) {
	var id int64
	err := repository.WithTx(ctx, r.db, func(tx *sql.Tx) error {
		vd := &dbmodels.Video{}
		if err := vd.Insert(ctx, tx, boil.Infer()); err != nil {
			return errors.Wrap(err, "inserting video")
		}

		dbEpisode := dbmodels.Episode{
			ShowID:  episode.ShowID,
			Name:    episode.Name,
			VideoID: vd.ID,
		}

		if err := dbEpisode.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}

		id = dbEpisode.ID
		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r Repository) GetEpisode(ctx context.Context, showID, id int64) (ui.Episode, error) {
	dbEpisode, err := dbmodels.Episodes(
		dbmodels.EpisodeWhere.ID.EQ(id),
		dbmodels.EpisodeWhere.ShowID.EQ(showID),
	).One(ctx, r.db)
	if err != nil {
		return ui.Episode{}, err
	}

	return ui.Episode{
		ID:      dbEpisode.ID,
		Name:    dbEpisode.Name,
		VideoID: dbEpisode.VideoID,
	}, nil
}
