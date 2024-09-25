package uri

import "strconv"

func ListMovies(page int) string {
	if page > 1 {
		return "/movies?page=" + strconv.Itoa(page)
	}
	return "/movies"
}

func ViewMovie(id int64) string {
	return "/movies/" + strconv.FormatInt(id, 10)
}

func CreateMovie() string {
	return "/movies/create"
}

func FetchMovieFromTMDB(id int64) string {
	return "/movies/" + strconv.FormatInt(id, 10) + "/fetch-from-tmdb"
}

func CreateMovieVideo(id int64) string {
	return "/movies/" + strconv.FormatInt(id, 10) + "/create-video"
}

func SyncMovie(id int64) string {
	return "/movies/" + strconv.FormatInt(id, 10) + "/sync"
}

func SyncYTSMovie(id int64) string {
	return "/movies/" + strconv.FormatInt(id, 10) + "/sync-yts"
}
