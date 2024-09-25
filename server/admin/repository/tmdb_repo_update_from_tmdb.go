package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	tmdb "github.com/cyruzin/golang-tmdb"
)

type TMDBRepository struct {
	db *sql.DB
}

func NewTMDBRepository(db *sql.DB) TMDBRepository {
	return TMDBRepository{db}
}

func (r TMDBRepository) getVideoRepo() Repository {
	return NewRepository(r.db)
}

var insertMovieQuery = `
INSERT INTO movies (
	id,
	imdb_id,
	title,
	original_title,
	status,
	tagline,
	genres,
	overview,
	poster_path,
	backdrop_path,
	release_date,
	runtime,
	vote_average,
	vote_count
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
) ON DUPLICATE KEY UPDATE 
	id = VALUES(id),
	imdb_id = VALUES(imdb_id),
	title = VALUES(title),
	original_title = VALUES(original_title),
	status = VALUES(status),
	tagline = VALUES(tagline),
	genres = VALUES(genres),
	overview = VALUES(overview),
	poster_path = VALUES(poster_path),
	backdrop_path = VALUES(backdrop_path),
	release_date = VALUES(release_date),
	runtime = VALUES(runtime),
	vote_average = VALUES(vote_average),
	vote_count = VALUES(vote_count);
`

func (r TMDBRepository) UpdateMovie(ctx context.Context, movie *tmdb.MovieDetails) error {
	genres, err := json.Marshal(movie.Genres)
	if err != nil {
		return fmt.Errorf("marshal genres: %w", err)
	}
	_, err = r.db.ExecContext(ctx, insertMovieQuery,
		movie.ID,
		movie.IMDbID,
		movie.Title,
		movie.OriginalTitle,
		movie.Status,
		movie.Tagline,
		genres,
		movie.Overview,
		movie.PosterPath,
		movie.BackdropPath,
		movie.ReleaseDate,
		movie.Runtime,
		movie.VoteAverage,
		movie.VoteCount,
	)
	if err != nil {
		return fmt.Errorf("insert movie: %w", err)
	}
	return nil
}

var insertTVSeriesQuery = `
INSERT INTO tv_series_shows (
	id,
	name,
	original_name,
	status,
	tagline,
	genres,
	overview,
	poster_path,
	backdrop_path,
	first_air_date,
	last_air_date,
	vote_average,
	vote_count,
	number_of_episodes,
	number_of_seasons
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
) ON DUPLICATE KEY UPDATE
	id = VALUES(id),
	name = VALUES(name),
	original_name = VALUES(original_name),
	status = VALUES(status),
	tagline = VALUES(tagline),
	genres = VALUES(genres),
	overview = VALUES(overview),
	poster_path = VALUES(poster_path),
	backdrop_path = VALUES(backdrop_path),
	first_air_date = VALUES(first_air_date),
	last_air_date = VALUES(last_air_date),
	vote_average = VALUES(vote_average),
	vote_count = VALUES(vote_count),
	number_of_episodes = VALUES(number_of_episodes),
	number_of_seasons = VALUES(number_of_seasons);
`

var insertTVSeasonQuery = `
INSERT INTO tv_seasons (
	id,
	show_id,
	season_number,
	name,
	poster_path,
	overview,
	air_date,
    vote_average,
	total_episodes
	
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?
) ON DUPLICATE KEY UPDATE
	id = VALUES(id),
	show_id = VALUES(show_id),
	season_number = VALUES(season_number),
	name = VALUES(name),
	poster_path = VALUES(poster_path),
	overview = VALUES(overview),
	air_date = VALUES(air_date),
	vote_average = VALUES(vote_average),
	total_episodes = VALUES(total_episodes);
`

var insertTVEpisodeQuery = `
INSERT INTO tv_episodes (
	id,
	show_id,
	season_number,
	episode_number,
	name,
	overview,
	air_date,
	runtime,
	still_path,
	vote_average,
	vote_count
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
) ON DUPLICATE KEY UPDATE
	id = VALUES(id),
	show_id = VALUES(show_id),
	season_number = VALUES(season_number),
	episode_number = VALUES(episode_number),
	name = VALUES(name),
	overview = VALUES(overview),
	air_date = VALUES(air_date),
	runtime = VALUES(runtime),
	still_path = VALUES(still_path),
	vote_average = VALUES(vote_average),
	vote_count = VALUES(vote_count);
`

func (r TMDBRepository) UpdateTVSeries(ctx context.Context, tv *tmdb.TVDetails, seasons []*tmdb.TVSeasonDetails) error {
	genres, err := json.Marshal(tv.Genres)
	if err != nil {
		return fmt.Errorf("marshal genres: %w", err)
	}
	return withTx(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, insertTVSeriesQuery,
			tv.ID,
			tv.Name,
			tv.OriginalName,
			tv.Status,
			tv.Tagline,
			genres,
			tv.Overview,
			tv.PosterPath,
			tv.BackdropPath,
			toDBAirDate(tv.FirstAirDate),
			toDBAirDate(tv.LastAirDate),
			tv.VoteAverage,
			tv.VoteCount,
			tv.NumberOfEpisodes,
			tv.NumberOfSeasons,
		)
		if err != nil {
			return fmt.Errorf("insert tv series: %w", err)
		}

		for _, season := range seasons {
			_, err := tx.ExecContext(ctx, insertTVSeasonQuery,
				season.ID,
				tv.ID,
				season.SeasonNumber,
				season.Name,
				season.PosterPath,
				season.Overview,
				toDBAirDate(season.AirDate),
				season.VoteAverage,
				len(season.Episodes),
			)
			if err != nil {
				return fmt.Errorf("insert tv season: %w", err)
			}

			for _, episode := range season.Episodes {
				_, err := tx.ExecContext(ctx, insertTVEpisodeQuery,
					episode.ID,
					episode.ShowID,
					episode.SeasonNumber,
					episode.EpisodeNumber,
					episode.Name,
					episode.Overview,
					toDBAirDate(episode.AirDate),
					episode.Runtime,
					episode.StillPath,
					episode.VoteAverage,
					episode.VoteCount,
				)
				if err != nil {
					return fmt.Errorf("insert tv episode %+v: %w", episode, err)
				}
			}
		}

		return nil
	})
}

func toDBAirDate(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
