package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

const syncToLatestShowsQuery = `
    INSERT INTO latest_shows(
		type, 
		show_id, 
		title, 
		original_title, 
		poster_path, 
		air_date, 
		runtime, 
		vote_average, 
		quality, 
		season_number, 
		episode_number
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)
	ON DUPLICATE KEY UPDATE
		type = VALUES(type),
		show_id = VALUES(show_id),
		title = VALUES(title),
		original_title = VALUES(original_title),
		poster_path = VALUES(poster_path),
		air_date = VALUES(air_date),
		runtime = VALUES(runtime),
		vote_average = VALUES(vote_average),
		quality = VALUES(quality),
		season_number = VALUES(season_number),
		episode_number = VALUES(episode_number);
`

func (r TMDBRepository) SyncMovie(ctx context.Context, movieID int64) error {
	movie, err := dbmodels.FindMovie(ctx, r.db, movieID)
	if err != nil {
		return fmt.Errorf("find movie: %w", err)
	}

	// not uploaded video yet, so ignore.
	if movie.VideoID == 0 {
		return nil
	}

	_, err = r.db.ExecContext(ctx, syncToLatestShowsQuery,
		dbmodels.LatestShowsTypeMovie,
		movie.ID,
		movie.Title,
		movie.OriginalTitle,
		movie.PosterPath,
		movie.ReleaseDate,
		movie.Runtime,
		movie.VoteAverage,
		"", // TODO: get quality from video
		nil,
		nil,
	)

	if err != nil {
		return fmt.Errorf("sync movie: %w", err)
	}

	return nil
}

func (r TMDBRepository) SyncTVSeries(ctx context.Context, showID int64) error {
	tvSeriesShow, err := dbmodels.FindTVSeriesShow(ctx, r.db, showID)
	if err != nil {
		return fmt.Errorf("find tv series show: %w", err)
	}

	lastedEpisode, err := dbmodels.TVEpisodes(
		dbmodels.TVEpisodeWhere.ShowID.EQ(showID),
		dbmodels.TVEpisodeWhere.VideoID.GT(0),
		qm.OrderBy(dbmodels.TVEpisodeColumns.AirDate+" DESC"),
		qm.OrderBy(dbmodels.TVEpisodeColumns.SeasonNumber+" DESC"),
		qm.OrderBy(dbmodels.TVEpisodeColumns.EpisodeNumber+" DESC"),
		qm.Limit(1),
	).One(ctx, r.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil // no episode uploaded yet
	}
	if err != nil {
		return fmt.Errorf("find lasted episode: %w", err)
	}

	lastedSeason, err := dbmodels.TVSeasons(
		dbmodels.TVSeasonWhere.ShowID.EQ(showID),
		dbmodels.TVSeasonWhere.SeasonNumber.EQ(lastedEpisode.SeasonNumber),
	).One(ctx, r.db)
	if err != nil {
		return fmt.Errorf("find season: %w", err)
	}

	airDate := lastedEpisode.AirDate
	if !lastedSeason.AirDate.Valid {
		airDate = lastedSeason.AirDate
	}
	if !airDate.Valid {
		airDate = tvSeriesShow.LastAirDate
	}

	_, err = r.db.ExecContext(ctx, syncToLatestShowsQuery,
		dbmodels.LatestShowsTypeTVSeries,
		tvSeriesShow.ID,
		tvSeriesShow.Name,
		tvSeriesShow.OriginalName,
		tvSeriesShow.PosterPath,
		airDate,
		nil,
		tvSeriesShow.VoteAverage,
		"",
		nil,
		nil,
	)

	if err != nil {
		return fmt.Errorf("sync tv series: %w", err)
	}

	_, err = r.db.ExecContext(ctx, syncToLatestShowsQuery,
		dbmodels.LatestShowsTypeEpisode,
		lastedEpisode.ShowID,
		fmt.Sprintf("%s S%02dE%02d", tvSeriesShow.Name, lastedEpisode.SeasonNumber, lastedEpisode.EpisodeNumber),
		fmt.Sprintf("%s S%02dE%02d", tvSeriesShow.OriginalName, lastedEpisode.SeasonNumber, lastedEpisode.EpisodeNumber),
		lastedSeason.PosterPath,
		lastedEpisode.AirDate,
		lastedEpisode.Runtime,
		lastedEpisode.VoteAverage,
		"",
		null.IntFrom(lastedEpisode.SeasonNumber),
		null.IntFrom(lastedEpisode.EpisodeNumber),
	)
	if err != nil {
		return fmt.Errorf("sync lasted episode: %w", err)
	}

	return nil
}
