package uri

import "strconv"

func ListTVSeries(page int) string {
	if page > 1 {
		return "/tv-series?page=" + strconv.Itoa(page)
	}
	return "/tv-series"
}

func ViewTVSeriesShow(showID int64) string {
	return "/tv-series/" + strconv.FormatInt(showID, 10)
}

func ViewTVSeason(showID int64, seasonNumber int) string {
	return "/tv-series/" + strconv.FormatInt(showID, 10) + "/seasons/" + strconv.Itoa(seasonNumber)
}

func ViewTVEpisode(showID int64, seasonNumber, episodeNumber int) string {
	return "/tv-series/" + strconv.FormatInt(showID, 10) + "/seasons/" + strconv.Itoa(seasonNumber) + "/episodes/" + strconv.Itoa(episodeNumber)
}

func CreateTVSeries() string {
	return "/tv-series/create"
}

func FetchTVSeriesFromTMDB(showID int64) string {
	return "/tv-series/" + strconv.FormatInt(showID, 10) + "/fetch-from-tmdb"
}

func CreateTVEpisodeVideo(showID int64, seasonNumber, episodeNumber int) string {
	return "/tv-series/" + strconv.FormatInt(showID, 10) + "/seasons/" + strconv.Itoa(seasonNumber) + "/episodes/" + strconv.Itoa(episodeNumber) + "/create-video"
}
