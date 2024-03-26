package uri

import (
	"fmt"
	"net/url"

	"github.com/phimtorr/phimtor/desktop/client/api"
)

func Home() string {
	return "/"
}

func ListShows(page, pageSize int, showType api.ShowType) string {
	return fmt.Sprintf("/shows?page=%d&pageSize=%d&type=%s", page, pageSize, showType)
}

func SearchShows(query string, page int) string {
	queryData := url.Values{}
	queryData.Set("q", query)
	queryData.Set("page", fmt.Sprintf("%d", page))
	uri := "/shows/search"
	if queryData.Encode() != "" {
		uri += "?" + queryData.Encode()
	}
	return uri
}

func GetMovie(id int64) string {
	return fmt.Sprintf("/movies/%d", id)
}

func GetSeries(id int64) string {
	return fmt.Sprintf("/series/%d", id)
}
