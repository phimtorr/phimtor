package uri

import "strconv"

const (
	Prefix = "/admin"
)

func ListShows(page int) string {
	if page > 1 {
		return Prefix + "/shows?page=" + strconv.Itoa(page)
	}
	return Prefix + "/shows"
}

func CreateShow() string {
	return Prefix + "/shows/create"
}

func ViewShow(id int64) string {
	return Prefix + "/shows/" + strconv.FormatInt(id, 10)
}

func UpdateShow(id int64) string {
	return Prefix + "/shows/" + strconv.FormatInt(id, 10) + "/update"
}
