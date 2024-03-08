package ports

import "context"

type Repository interface {
	ListShows(ctx context.Context, params ListShowsParams) ([]Show, Pagination, error)
	SearchShow(ctx context.Context, params SearchShowsParams) ([]Show, Pagination, error)
	GetMovie(ctx context.Context, id int64) (Movie, error)
	GetSeries(ctx context.Context, id int64) (Series, error)
	GetVideo(ctx context.Context, id int64) (Video, error)
}
