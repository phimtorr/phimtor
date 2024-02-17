package domain

type TorrentLink struct {
	Quantity string `json:"quantity"`

	Link       string `json:"link"`
	MagnetLink string `json:"magnet_link"`
}
