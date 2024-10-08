package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r TMDBRepository) SyncAll(ctx context.Context) error {
	var syncErr error
	if err := r.SyncAllMovies(ctx); err != nil {
		syncErr = errors.Join(syncErr, fmt.Errorf("sync all movies: %w", err))
	}
	if err := r.SyncAllTVSeries(ctx); err != nil {
		syncErr = errors.Join(syncErr, fmt.Errorf("sync all tv series: %w", err))
	}
	return syncErr
}

func (r TMDBRepository) SyncAllMovies(ctx context.Context) error {
	movies, err := dbmodels.Movies().All(ctx, r.db)
	if err != nil {
		return fmt.Errorf("find movies: %w", err)
	}

	var syncMoviesErr error
	for _, movie := range movies {
		if err := r.SyncMovie(ctx, movie.ID); err != nil {
			syncMoviesErr = errors.Join(syncMoviesErr, fmt.Errorf("sync movie %d: %w", movie.ID, err))
		}
	}

	return syncMoviesErr
}

func (r TMDBRepository) SyncAllTVSeries(ctx context.Context) error {
	tvSeriesShows, err := dbmodels.TVSeriesShows().All(ctx, r.db)
	if err != nil {
		return fmt.Errorf("find tv series shows: %w", err)
	}

	var syncTVSeriesErr error
	for _, show := range tvSeriesShows {
		if err := r.SyncTVSeries(ctx, show.ID); err != nil {
			syncTVSeriesErr = errors.Join(syncTVSeriesErr, fmt.Errorf("sync tv series %d: %w", show.ID, err))
		}
	}

	return syncTVSeriesErr
}

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
		has_vi_sub,
		season_number, 
		episode_number
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
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
		has_vi_sub = VALUES(has_vi_sub),
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

	video, err := dbmodels.FindVideo(ctx, r.db, movie.VideoID)
	if err != nil {
		return fmt.Errorf("find video: %w", err)
	}

	var quality string
	switch video.MaxResolution {
	case 2160:
		quality = "4K"
	case 1080:
		quality = "1080p"
	case 720:
		quality = "720p"
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
		quality,
		video.HasViSub,
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

	lastedEpisodeVideo, err := dbmodels.FindVideo(ctx, r.db, lastedEpisode.VideoID)
	if err != nil {
		return fmt.Errorf("find lasted episode video: %w", err)
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
		false,
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
		lastedEpisodeVideo.HasViSub,
		null.IntFrom(lastedEpisode.SeasonNumber),
		null.IntFrom(lastedEpisode.EpisodeNumber),
	)
	if err != nil {
		return fmt.Errorf("sync lasted episode: %w", err)
	}

	// sync count available episodes of the season
	episodesCount, err := r.countAvailableEpisodesBySeason(ctx, showID)
	if err != nil {
		return fmt.Errorf("count available episodes by season: %w", err)
	}

	for seasonNumber, count := range episodesCount {
		if err := r.updateCountAvailableEpisodesBySeason(ctx, showID, seasonNumber, count); err != nil {
			return fmt.Errorf("update count available episodes by season: %w", err)
		}
	}

	return nil
}

var countAvailableEpisodesBySeasonQuery = `
	SELECT season_number, COUNT(*) AS count
	FROM tv_episodes
	WHERE show_id = ?
	AND video_id > 0
	GROUP BY season_number;
`

func (r TMDBRepository) countAvailableEpisodesBySeason(ctx context.Context, showID int64) (map[int]int, error) {
	rows, err := r.db.QueryContext(ctx, countAvailableEpisodesBySeasonQuery, showID)
	if err != nil {
		return nil, fmt.Errorf("count available episodes by season: %w", err)
	}

	defer rows.Close()

	counts := make(map[int]int)

	for rows.Next() {
		var seasonNumber, count int
		if err := rows.Scan(&seasonNumber, &count); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		counts[seasonNumber] = count
	}

	return counts, nil
}

func (r TMDBRepository) updateCountAvailableEpisodesBySeason(ctx context.Context, showID int64, seasonNumber, count int) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE tv_seasons
		SET count_available_episodes = ?
		WHERE show_id = ?
		AND season_number = ?
	`, count, showID, seasonNumber)
	if err != nil {
		return fmt.Errorf("update count available episodes by season: %w", err)
	}
	return nil
}
