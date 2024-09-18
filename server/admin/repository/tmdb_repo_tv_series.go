package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/admin/http/ui"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r TMDBRepository) ListTVSeriesShows(ctx context.Context, page, pageSize int) ([]ui.TVSeriesShow, ui.Pagination, error) {
	shows, err := dbmodels.TVSeriesShows(
		qm.OrderBy(dbmodels.TVSeriesShowColumns.UpdatedAt+" DESC"),
		qm.Limit(pageSize),
		qm.Offset((page-1)*pageSize),
	).All(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, fmt.Errorf("list tv series shows: %w", err)
	}

	count, err := dbmodels.TVSeriesShows().Count(ctx, r.db)
	if err != nil {
		return nil, ui.Pagination{}, fmt.Errorf("count tv series shows: %w", err)
	}

	totalPages := int(math.Ceil(float64(count) / float64(pageSize)))

	return toUITVSeriesShows(shows), ui.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}, nil

}
func (r TMDBRepository) GetTVSeriesShow(ctx context.Context, showID int64) (ui.TVSeriesShow, []ui.TVSeason, error) {
	show, err := dbmodels.FindTVSeriesShow(ctx, r.db, showID)
	if err != nil {
		return ui.TVSeriesShow{}, nil, fmt.Errorf("get tv series show: %w", err)
	}

	seasons, err := dbmodels.TVSeasons(
		dbmodels.TVSeasonWhere.ShowID.EQ(showID),
		qm.OrderBy(dbmodels.TVSeasonColumns.ID+" ASC"),
	).All(ctx, r.db)
	if err != nil {
		return ui.TVSeriesShow{}, nil, fmt.Errorf("get tv seasons: %w", err)
	}

	return toUITVSeriesShow(show), toUITVSeasons(seasons), nil
}

func (r TMDBRepository) GetTVSeason(ctx context.Context, showID int64, seasonNumber int) (ui.TVSeason, []ui.TVEpisode, error) {
	season, err := dbmodels.TVSeasons(
		dbmodels.TVSeasonWhere.ShowID.EQ(showID),
		dbmodels.TVSeasonWhere.SeasonNumber.EQ(seasonNumber),
	).One(ctx, r.db)
	if err != nil {
		return ui.TVSeason{}, nil, fmt.Errorf("get tv season: %w", err)
	}

	episodes, err := dbmodels.TVEpisodes(
		dbmodels.TVEpisodeWhere.ShowID.EQ(showID),
		dbmodels.TVEpisodeWhere.SeasonNumber.EQ(seasonNumber),
		qm.OrderBy(dbmodels.TVEpisodeColumns.EpisodeNumber+" ASC"),
	).All(ctx, r.db)
	if err != nil {
		return ui.TVSeason{}, nil, fmt.Errorf("get tv episodes: %w", err)
	}

	return toUITVSeason(season), toUITVEpisodes(episodes), nil
}
func (r TMDBRepository) GetTVEpisode(ctx context.Context, showID int64, seasonNumber, episodeNumber int) (ui.TVEpisode, error) {
	episode, err := dbmodels.TVEpisodes(
		dbmodels.TVEpisodeWhere.ShowID.EQ(showID),
		dbmodels.TVEpisodeWhere.ShowID.EQ(showID),
		dbmodels.TVEpisodeWhere.SeasonNumber.EQ(seasonNumber),
		dbmodels.TVEpisodeWhere.EpisodeNumber.EQ(episodeNumber),
	).One(ctx, r.db)
	if err != nil {
		return ui.TVEpisode{}, fmt.Errorf("get tv episode: %w", err)
	}

	return toUITVEpisode(episode), nil
}

func (r TMDBRepository) CreateTVEpisodeVideo(ctx context.Context, showID int64, seasonNumber, episodeNumber int) error {
	return withTx(ctx, r.db, func(ctx context.Context, tx *sql.Tx) error {
		episode, err := dbmodels.TVEpisodes(
			dbmodels.TVEpisodeWhere.ShowID.EQ(showID),
			dbmodels.TVEpisodeWhere.SeasonNumber.EQ(seasonNumber),
			dbmodels.TVEpisodeWhere.EpisodeNumber.EQ(episodeNumber),
		).One(ctx, tx)
		if err != nil {
			return fmt.Errorf("find episode: %w", err)
		}

		if episode.VideoID != 0 {
			return nil
		}

		video := &dbmodels.Video{}

		if err := video.Insert(ctx, tx, boil.Infer()); err != nil {
			return fmt.Errorf("insert video: %w", err)
		}

		episode.VideoID = video.ID

		if _, err := episode.Update(ctx, tx, boil.Whitelist(dbmodels.TVEpisodeColumns.VideoID)); err != nil {
			return fmt.Errorf("update episode: %w", err)
		}

		return nil
	})
}

func toUITVSeriesShows(dbShows dbmodels.TVSeriesShowSlice) []ui.TVSeriesShow {
	shows := make([]ui.TVSeriesShow, len(dbShows))
	for i, dbShow := range dbShows {
		shows[i] = toUITVSeriesShow(dbShow)
	}
	return shows
}

func toUITVSeriesShow(dbShow *dbmodels.TVSeriesShow) ui.TVSeriesShow {
	return ui.TVSeriesShow{
		ID:               dbShow.ID,
		Name:             dbShow.Name,
		OriginalName:     dbShow.OriginalName,
		Status:           dbShow.Status,
		Tagline:          dbShow.Tagline,
		Genres:           string(dbShow.Genres),
		Overview:         dbShow.Overview,
		PosterLink:       tmdb.GetImageURL(dbShow.PosterPath, tmdb.W300),
		BackdropLink:     tmdb.GetImageURL(dbShow.BackdropPath, tmdb.W780),
		FirstAirDate:     dbShow.FirstAirDate.Time.Format("2006-01-02"),
		LastAirDate:      dbShow.LastAirDate.Time.Format("2006-01-02"),
		VoteAverage:      float64(dbShow.VoteAverage),
		VoteCount:        dbShow.VoteCount,
		NumberOfEpisodes: dbShow.NumberOfEpisodes,
		NumberOfSeasons:  dbShow.NumberOfSeasons,
		CreatedAt:        dbShow.CreatedAt.String(),
		UpdatedAt:        dbShow.UpdatedAt.String(),
	}
}

func toUITVSeasons(dbSeasons dbmodels.TVSeasonSlice) []ui.TVSeason {
	seasons := make([]ui.TVSeason, len(dbSeasons))
	for i, dbSeason := range dbSeasons {
		seasons[i] = toUITVSeason(dbSeason)
	}
	return seasons
}

func toUITVSeason(dbSeason *dbmodels.TVSeason) ui.TVSeason {
	return ui.TVSeason{
		ID:            dbSeason.ID,
		ShowID:        dbSeason.ShowID,
		SeasonNumber:  dbSeason.SeasonNumber,
		Name:          dbSeason.Name,
		PosterLink:    tmdb.GetImageURL(dbSeason.PosterPath, tmdb.W300),
		Overview:      dbSeason.Overview,
		AirDate:       dbSeason.AirDate.Time.Format("2006-01-02"),
		VoteAverage:   float64(dbSeason.VoteAverage),
		TotalEpisodes: dbSeason.TotalEpisodes,
		CreatedAt:     dbSeason.CreatedAt.String(),
		UpdatedAt:     dbSeason.UpdatedAt.String(),
	}
}

func toUITVEpisodes(dbEpisodes dbmodels.TVEpisodeSlice) []ui.TVEpisode {
	episodes := make([]ui.TVEpisode, len(dbEpisodes))
	for i, dbEpisode := range dbEpisodes {
		episodes[i] = toUITVEpisode(dbEpisode)
	}
	return episodes
}

func toUITVEpisode(dbEpisode *dbmodels.TVEpisode) ui.TVEpisode {
	return ui.TVEpisode{
		ID:            dbEpisode.ID,
		ShowID:        dbEpisode.ShowID,
		SeasonNumber:  dbEpisode.SeasonNumber,
		EpisodeNumber: dbEpisode.EpisodeNumber,
		Name:          dbEpisode.Name,
		Overview:      dbEpisode.Overview,
		AirDate:       dbEpisode.AirDate.Time.Format("2006-01-02"),
		Runtime:       dbEpisode.Runtime,
		StillLink:     tmdb.GetImageURL(dbEpisode.StillPath, tmdb.W300),
		VoteAverage:   float64(dbEpisode.VoteAverage),
		VoteCount:     dbEpisode.VoteCount,
		VideoID:       dbEpisode.VideoID,
		CreatedAt:     dbEpisode.CreatedAt.String(),
		UpdatedAt:     dbEpisode.UpdatedAt.String(),
	}
}
