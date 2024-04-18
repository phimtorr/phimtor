package repository

import (
	"context"
	"database/sql"
	"math"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/admin/http/handler"
	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/repository"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r Repository) ListShowDisplays(ctx context.Context, page, pageSize int) ([]ui.ShowDisplay, ui.Pagination, error) {
	dbShows, err := dbmodels.Shows(
		qm.OrderBy(dbmodels.ShowColumns.UpdatedAt+" DESC"),
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
	).All(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, err
	}

	count, err := dbmodels.Shows().Count(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, err
	}

	shows := make([]ui.ShowDisplay, len(dbShows))
	for i, dbShow := range dbShows {
		shows[i] = ui.ShowDisplay{
			ID:            dbShow.ID,
			Title:         dbShow.Title,
			OriginalTitle: dbShow.OriginalTitle,
			Poster:        dbShow.PosterLink,
		}
	}

	totalPages := int(math.Ceil(float64(count) / float64(pageSize)))

	return shows, ui.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}, nil
}

func (r Repository) CreateShow(ctx context.Context, show handler.ShowToCreate) (int64, error) {
	var id int64
	err := repository.WithTx(ctx, r.db, func(tx *sql.Tx) error {
		var videoID int64
		if show.ShowType == "movie" {
			vd := &dbmodels.Video{}
			if err := vd.Insert(ctx, tx, boil.Infer()); err != nil {
				return errors.Wrap(err, "inserting video")
			}
			videoID = vd.ID
		}

		dbShow := &dbmodels.Show{
			Type:              dbmodels.ShowsType(show.ShowType),
			Title:             show.Title,
			OriginalTitle:     show.OriginalTitle,
			PosterLink:        show.PosterLink,
			Description:       show.Description,
			ReleaseYear:       show.ReleaseYear,
			Score:             show.Score,
			DurationInMinutes: show.DurationInMinutes,
			Quantity:          show.Quality,
			TotalEpisodes:     show.TotalEpisodes,
			VideoID:           null.NewInt64(videoID, videoID != 0),
		}

		if err := dbShow.Insert(ctx, tx, boil.Infer()); err != nil {
			return errors.Wrap(err, "inserting show")
		}

		id = dbShow.ID
		return nil

	})
	return id, err
}

func (r Repository) UpdateShow(ctx context.Context, show handler.ShowToUpdate) error {
	dbShow, err := dbmodels.FindShow(ctx, r.db, show.ID)
	if err != nil {
		return errors.Wrap(err, "finding show")
	}

	dbShow.Title = show.Title
	dbShow.OriginalTitle = show.OriginalTitle
	dbShow.PosterLink = show.PosterLink
	dbShow.Description = show.Description
	dbShow.ReleaseYear = show.ReleaseYear
	dbShow.Score = show.Score
	dbShow.DurationInMinutes = show.DurationInMinutes
	dbShow.Quantity = show.Quality
	if dbShow.Type == dbmodels.ShowsTypeSeries {
		dbShow.TotalEpisodes = show.TotalEpisodes
	}

	if _, err = dbShow.Update(ctx, r.db, boil.Infer()); err != nil {
		return errors.Wrap(err, "updating show")
	}

	return nil
}

func (r Repository) GetShow(ctx context.Context, id int64) (ui.Show, error) {
	dbShow, err := dbmodels.FindShow(ctx, r.db, id)
	if err != nil {
		return ui.Show{}, errors.Wrap(err, "finding show")
	}

	show := ui.Show{
		ID:                dbShow.ID,
		Title:             dbShow.Title,
		OriginalTitle:     dbShow.OriginalTitle,
		PosterLink:        dbShow.PosterLink,
		Description:       dbShow.Description,
		ReleaseYear:       dbShow.ReleaseYear,
		Score:             dbShow.Score,
		DurationInMinutes: dbShow.DurationInMinutes,
		Quality:           dbShow.Quantity,
		VideoID:           dbShow.VideoID.Int64,
	}

	if dbShow.Type == dbmodels.ShowsTypeSeries {
		show.TotalEpisodes = dbShow.TotalEpisodes
	}

	return show, nil
}
