package uri

import "strconv"

func ListShows(page int) string {
	if page > 1 {
		return "/shows?page=" + strconv.Itoa(page)
	}
	return "/shows"
}

func CreateShow() string {
	return "/shows/create"
}

func ViewShow(id int64) string {
	return "/shows/" + strconv.FormatInt(id, 10)
}

func UpdateShow(id int64) string {
	return "/shows/" + strconv.FormatInt(id, 10) + "/update"
}

func ListEpisodes(id int64) string {
	return "/shows/" + strconv.FormatInt(id, 10) + "/episodes"
}

func CreateEpisode(id int64) string {
	return "/shows/" + strconv.FormatInt(id, 10) + "/episodes/create"
}

func ViewEpisode(id, episodeID int64) string {
	return "/shows/" + strconv.FormatInt(id, 10) + "/episodes/" + strconv.FormatInt(episodeID, 10) + "/view"
}
