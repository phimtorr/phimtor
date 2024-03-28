package uri

import "strconv"

func ListShows(page int) string {
	if page > 1 {
		return "shows?page=" + strconv.Itoa(page)
	}
	return "shows"
}
