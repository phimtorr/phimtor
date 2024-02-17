package domain

type DisplayInfo struct {
	Title         string `json:"title"`
	OriginalTitle string `json:"original_title"`
	PosterLink    string `json:"poster_link"`
	Description   string `json:"description"`

	ReleaseYear       int     `json:"release_year"`
	Score             float64 `json:"score"`
	DurationInMinutes int     `json:"duration_in_minutes"`
}
