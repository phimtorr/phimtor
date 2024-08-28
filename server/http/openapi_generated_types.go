// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package http

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for ShowType.
const (
	ShowTypeMovie  ShowType = "movie"
	ShowTypeSeries ShowType = "series"
)

// Episode defines model for Episode.
type Episode struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	VideoId int64  `json:"videoId"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Movie defines model for Movie.
type Movie struct {
	Description       string  `json:"description"`
	DurationInMinutes int     `json:"durationInMinutes"`
	Id                int64   `json:"id"`
	OriginalTitle     string  `json:"originalTitle"`
	PosterLink        string  `json:"posterLink"`
	Quantity          string  `json:"quantity"`
	ReleaseYear       int     `json:"releaseYear"`
	Score             float32 `json:"score"`
	Title             string  `json:"title"`
	VideoId           int64   `json:"videoId"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	Page         int   `json:"page"`
	TotalPages   int   `json:"totalPages"`
	TotalResults int64 `json:"totalResults"`
}

// PremiumTorrentLink defines model for PremiumTorrentLink.
type PremiumTorrentLink struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

// Series defines model for Series.
type Series struct {
	CurrentEpisode    int       `json:"currentEpisode"`
	Description       string    `json:"description"`
	DurationInMinutes int       `json:"durationInMinutes"`
	Episodes          []Episode `json:"episodes"`
	Id                int64     `json:"id"`
	OriginalTitle     string    `json:"originalTitle"`
	PosterLink        string    `json:"posterLink"`
	ReleaseYear       int       `json:"releaseYear"`
	Score             float32   `json:"score"`
	Title             string    `json:"title"`
	TotalEpisodes     int       `json:"totalEpisodes"`
}

// Show defines model for Show.
type Show struct {
	CurrentEpisode    int      `json:"currentEpisode"`
	DurationInMinutes int      `json:"durationInMinutes"`
	Id                int64    `json:"id"`
	OriginalTitle     string   `json:"originalTitle"`
	PosterLink        string   `json:"posterLink"`
	Quantity          string   `json:"quantity"`
	ReleaseYear       int      `json:"releaseYear"`
	Score             float32  `json:"score"`
	Title             string   `json:"title"`
	TotalEpisodes     int      `json:"totalEpisodes"`
	Type              ShowType `json:"type"`
}

// ShowType defines model for ShowType.
type ShowType string

// Subtitle defines model for Subtitle.
type Subtitle struct {
	Id       int64  `json:"id"`
	Language string `json:"language"`
	Link     string `json:"link"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Priority int    `json:"priority"`
}

// TorrentLink defines model for TorrentLink.
type TorrentLink struct {
	FileIndex      int    `json:"fileIndex"`
	Id             int64  `json:"id"`
	Link           string `json:"link"`
	Name           string `json:"name"`
	Priority       int    `json:"priority"`
	RequirePremium bool   `json:"requirePremium"`
}

// Video defines model for Video.
type Video struct {
	Id                  int64                `json:"id"`
	PremiumTorrentLinks []PremiumTorrentLink `json:"premiumTorrentLinks"`
	Subtitles           []Subtitle           `json:"subtitles"`
	Title               string               `json:"title"`
	TorrentLinks        []TorrentLink        `json:"torrentLinks"`
}

// BadRequest defines model for BadRequest.
type BadRequest = ErrorResponse

// InternalError defines model for InternalError.
type InternalError = ErrorResponse

// ListShowsParams defines parameters for ListShows.
type ListShowsParams struct {
	Page     *int      `form:"page,omitempty" json:"page,omitempty"`
	PageSize *int      `form:"pageSize,omitempty" json:"pageSize,omitempty"`
	Type     *ShowType `form:"type,omitempty" json:"type,omitempty"`
}

// SearchShowsParams defines parameters for SearchShows.
type SearchShowsParams struct {
	Query string `form:"query" json:"query"`
	Page  *int   `form:"page,omitempty" json:"page,omitempty"`
}
