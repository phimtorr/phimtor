package uri

import "strconv"

func ListLatestShows(page int) string {
	if page > 0 {
		return "/latest-shows?page=" + strconv.Itoa(page)
	}
	return "/latest-shows"
}
