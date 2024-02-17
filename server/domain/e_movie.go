package domain

type Movie struct {
	ID int `json:"id"`

	DisplayInfo `json:"display_info"`

	Quantity string `json:"quantity"`

	TorrentLinks []TorrentLink `json:"torrent_links"`
	Subtitles    []Subtitle    `json:"subtitles"`
}
