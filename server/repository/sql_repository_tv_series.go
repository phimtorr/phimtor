package repository

import (
	"context"
	"fmt"

	tmdb "github.com/cyruzin/golang-tmdb"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/phimtorr/phimtor/server/http"
	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func (r SQLRepository) GetTvSeries(ctx context.Context, id int64) (http.TvSeries, error) {
	tvSeries, err := dbmodels.TVSeriesShows(
		dbmodels.TVSeriesShowWhere.ID.EQ(id),
		qm.Load(dbmodels.TVSeriesShowRels.ShowTVSeasons),
	).One(ctx, r.db)
	if err != nil {
		return http.TvSeries{}, fmt.Errorf("get tv series: %w", err)
	}

	return toHTTP2TvSeries(tvSeries), nil
}

func toHTTP2TvSeries(dbTvSeries *dbmodels.TVSeriesShow) http.TvSeries {
	episodes := make([]struct {
		AirDate      *openapiTypes.Date `json:"airDate,omitempty"`
		Id           int64              `json:"id"`
		Name         string             `json:"name"`
		Overview     string             `json:"overview"`
		PosterLink   string             `json:"posterLink"`
		SeasonNumber int                `json:"seasonNumber"`
		VoteAverage  float32            `json:"voteAverage"`
	}, len(dbTvSeries.R.GetShowTVSeasons()))

	for i, s := range dbTvSeries.R.GetShowTVSeasons() {
		episodes[i].AirDate = toOpenapiTypeDate(s.AirDate)
		episodes[i].Id = s.ID
		episodes[i].Name = s.Name
		episodes[i].Overview = s.Overview
		episodes[i].PosterLink = tmdb.GetImageURL(s.PosterPath, tmdb.W300)
		episodes[i].SeasonNumber = s.SeasonNumber
		episodes[i].VoteAverage = s.VoteAverage
	}

	return http.TvSeries{
		BackdropLink:     tmdb.GetImageURL(dbTvSeries.BackdropPath, tmdb.Original),
		FirstAirDate:     toOpenapiTypeDate(dbTvSeries.FirstAirDate),
		Genres:           toHTTP2Genres(dbTvSeries.Genres),
		Id:               dbTvSeries.ID,
		LastAirDate:      toOpenapiTypeDate(dbTvSeries.LastAirDate),
		Name:             dbTvSeries.Name,
		NumberOfEpisodes: dbTvSeries.NumberOfEpisodes,
		NumberOfSeasons:  dbTvSeries.NumberOfSeasons,
		OriginalName:     dbTvSeries.OriginalName,
		Overview:         dbTvSeries.Overview,
		PosterLink:       tmdb.GetImageURL(dbTvSeries.PosterPath, tmdb.W300),
		Seasons:          episodes,
		Status:           dbTvSeries.Status,
		Tagline:          dbTvSeries.Tagline,
		VoteAverage:      dbTvSeries.VoteAverage,
	}
}

func (r SQLRepository) GetTvSeason(ctx context.Context, showID int64, seasonNumber int) (http.TVSeason, error) {
	season, err := dbmodels.TVSeasons(
		dbmodels.TVSeasonWhere.ShowID.EQ(showID),
		dbmodels.TVSeasonWhere.SeasonNumber.EQ(seasonNumber),
	).One(ctx, r.db)
	if err != nil {
		return http.TVSeason{}, fmt.Errorf("get tv season: %w", err)
	}

	episodes, err := dbmodels.TVEpisodes(
		dbmodels.TVEpisodeWhere.ShowID.EQ(showID),
		dbmodels.TVEpisodeWhere.SeasonNumber.EQ(seasonNumber),
	).All(ctx, r.db)
	if err != nil {
		return http.TVSeason{}, fmt.Errorf("get tv season episodes: %w", err)
	}

	return toHTTP2TVSeason(season, episodes), nil
}

func toHTTP2TVSeason(dbSeason *dbmodels.TVSeason, dbEpisodes []*dbmodels.TVEpisode) http.TVSeason {
	episodes := make([]struct {
		AirDate       *openapiTypes.Date `json:"airDate,omitempty"`
		EpisodeNumber int                `json:"episodeNumber"`
		Id            int64              `json:"id"`
		Name          string             `json:"name"`
		Overview      string             `json:"overview"`
		Runtime       int                `json:"runtime"`
		StillLink     string             `json:"stillLink"`
		VideoID       int64              `json:"videoID"`
		VoteAverage   float32            `json:"voteAverage"`
	}, len(dbEpisodes))

	for i, e := range dbEpisodes {
		episodes[i].AirDate = toOpenapiTypeDate(e.AirDate)
		episodes[i].EpisodeNumber = e.EpisodeNumber
		episodes[i].Id = e.ID
		episodes[i].Name = e.Name
		episodes[i].Overview = e.Overview
		episodes[i].Runtime = e.Runtime
		episodes[i].StillLink = tmdb.GetImageURL(e.StillPath, tmdb.W300)
		episodes[i].VideoID = e.VideoID
		episodes[i].VoteAverage = e.VoteAverage
	}

	return http.TVSeason{
		AirDate:      toOpenapiTypeDate(dbSeason.AirDate),
		Episodes:     episodes,
		Id:           dbSeason.ID,
		Name:         dbSeason.Name,
		Overview:     dbSeason.Overview,
		PosterLink:   tmdb.GetImageURL(dbSeason.PosterPath, tmdb.W300),
		SeasonNumber: dbSeason.SeasonNumber,
		VoteAverage:  dbSeason.VoteAverage,
	}
}

func toOpenapiTypeDate(date null.Time) *openapiTypes.Date {
	if !date.Valid {
		return nil
	}

	return &openapiTypes.Date{
		Time: date.Time,
	}
}
