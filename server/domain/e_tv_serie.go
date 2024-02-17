package domain

type TVSeries struct {
	ID int `json:"id"`

	DisplayInfo DisplayInfo `json:"display_info"`

	TotalEpisodes   int `json:"total_episodes"`
	CurrentEpisodes int `json:"current_episodes"`

	Episodes []Episode `json:"episodes"`
}

type Episode struct {
	Name         string        `json:"name"`
	TorrentLinks []TorrentLink `json:"torrent_links"`
	Subtitles    []Subtitle    `json:"subtitles"`
}
