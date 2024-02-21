package uri

import (
	"fmt"

	"github.com/phimtorr/phimtor/desktop/torrent"
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

func GetVideo(id int64) string {
	return fmt.Sprintf("/videos/%d", id)
}

func GetStream(infoHash torrent.InfoHash, fileIndex int) string {
	return fmt.Sprintf("/stream/%s/%d", infoHash, fileIndex)
}
