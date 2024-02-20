package uri

import (
	"fmt"
)

func Home() string {
	return "/"
}

func ListShows() string {
	return "/shows"
}

func GetMovie(id int64) string {
	return fmt.Sprintf("/movies/%d", id)
}

func GetSeries(id int64) string {
	return fmt.Sprintf("/series/%d", id)
}
