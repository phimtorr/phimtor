package jobs

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"

	"github.com/phimtorr/phimtor/server/repository/dbmodels"
)

func RunFetchNewMoviesFromYTS(
	ctx context.Context,
	tmdbClient TMDBClient,
	ytsClient YTSClient,
	repo Repository,
	db *sql.DB,
) {
	timer := time.NewTicker(20 * time.Minute)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			ctx := log.Ctx(ctx).With().
				Str("job", "fetch_new_movies_from_yts").
				Time("time", time.Now()).
				Logger().WithContext(ctx)
			if err := runFetchNewMoviesFromYTS(ctx, tmdbClient, ytsClient, repo, db); err != nil {
				log.Ctx(ctx).Error().Err(err).Msg("Failed to fetch new movies from YTS")
			}
		}
	}
}

func runFetchNewMoviesFromYTS(
	ctx context.Context,
	tmdbClient TMDBClient,
	ytsClient YTSClient,
	repo Repository,
	db *sql.DB,
) error {
	featuredMovies, err := getFeaturedMoviesFromYTS(ctx)
	if err != nil {
		return fmt.Errorf("get featured movies from YTS: %w", err)
	}

	// reverse the list so that we process the latest movies first
	for i := len(featuredMovies) - 1; i >= 0; i-- {
		movie := featuredMovies[i]
		if err := processFeaturedMovie(ctx, tmdbClient, ytsClient, repo, db, movie); err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Failed to process featured movie")
		}
	}

	return nil
}

type YTSFeaturedMovie struct {
	MovieID int64
	IMDBID  string
}

func getFeaturedMoviesFromYTS(ctx context.Context) ([]YTSFeaturedMovie, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.Ctx(ctx).Info().Str("url", r.URL.String()).Msg("Visiting")
	})

	var (
		mu     sync.Mutex
		movies []YTSFeaturedMovie
	)

	c.OnHTML(".browse-movie-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		log.Ctx(ctx).Info().Str("link", link).Msg("Link found")

		if err := e.Request.Visit(link); err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Error visiting")
		}
	})

	c.OnHTML("#movie-info", func(e *colly.HTMLElement) {
		movieID, err := strconv.ParseInt(e.Attr("data-movie-id"), 10, 64)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Error parsing movie ID")
			return
		}

		log.Ctx(ctx).Info().Int64("id", movieID).Msg("Movie ID")

		// here I want to get attribute href of the first a tag that has title = IMDb Rating
		imdbLink := e.ChildAttr("a[title='IMDb Rating']", "href")
		log.Ctx(ctx).Info().Str("link", imdbLink).Msg("IMDb link")
		imdbID := strings.Split(imdbLink, "/")[4]
		log.Ctx(ctx).Info().Str("id", imdbID).Msg("IMDb ID")

		mu.Lock()
		defer mu.Unlock()

		movies = append(movies, YTSFeaturedMovie{
			MovieID: movieID,
			IMDBID:  imdbID,
		})
	})

	err := c.Visit("https://yts.mx/browse-movies/0/all/all/5/featured/0/en")
	if err != nil {
		return nil, fmt.Errorf("visit: %w", err)
	}

	return movies, nil
}

func processFeaturedMovie(
	ctx context.Context,
	tmdbClient TMDBClient,
	ytsClient YTSClient,
	repo Repository,
	db *sql.DB,
	featuredMovie YTSFeaturedMovie,
) error {
	tmdbMovie, err := tmdbClient.GetMovieDetailsByIMDbID(ctx, featuredMovie.IMDBID)
	if err != nil {
		return fmt.Errorf("get movie details by IMDb ID: %w", err)
	}

	ytsMovie, err := ytsClient.GetMovieByID(ctx, featuredMovie.MovieID)
	if err != nil {
		return fmt.Errorf("get movie by ID: %w", err)
	}

	if err := repo.UpdateMovie(ctx, tmdbMovie); err != nil {
		return fmt.Errorf("update movie: %w", err)
	}

	if err := repo.CreateMovieVideo(ctx, tmdbMovie.ID); err != nil {
		return fmt.Errorf("create movie video: %w", err)
	}

	dbMovie, err := dbmodels.FindMovie(ctx, db, tmdbMovie.ID)
	if err != nil {
		return fmt.Errorf("find movie: %w", err)
	}

	if dbMovie.VideoID == 0 {
		return fmt.Errorf("video ID is 0")
	}

	if err := repo.UpdateYTSMovie(ctx, dbMovie.VideoID, ytsMovie); err != nil {
		return fmt.Errorf("update YTS movie: %w", err)
	}

	if err := repo.SyncMovie(ctx, tmdbMovie.ID); err != nil {
		return fmt.Errorf("sync movie: %w", err)
	}

	return nil
}
