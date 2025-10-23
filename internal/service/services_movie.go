package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

func (service *MovieService) adaptQueryResult(result sqlc.Movie) Movie {
	posterUrl := ""
	if result.PosterUrl.Valid {
		posterUrl = result.PosterUrl.String
	}
	return Movie{
		ID:              int(result.ID),
		Title:           result.Title,
		Synopsis:        result.Synopsis,
		DurationMinutes: int(result.DurationMinutes),
		ReleasedAt:      result.ReleasedAt,
		PosterUrl:       posterUrl,
	}
}

func (service *MovieService) adaptQueryResults(results []sqlc.Movie) []Movie {
	var movies = make([]Movie, len(results))
	for i, movie := range results {
		posterUrl := ""
		if movie.PosterUrl.Valid {
			posterUrl = movie.PosterUrl.String
		}
		movies[i] = Movie{
			ID:              int(movie.ID),
			Title:           movie.Title,
			Synopsis:        movie.Synopsis,
			DurationMinutes: int(movie.DurationMinutes),
			ReleasedAt:      movie.ReleasedAt,
			PosterUrl:       posterUrl,
		}
	}
	return movies
}

func (service *MovieService) adaptNullablePoster(posterUrl string) sql.NullString {
	var nullableposterUrl = sql.NullString{Valid: false}
	if posterUrl != "" {
		nullableposterUrl.String = posterUrl
		nullableposterUrl.Valid = true
	}
	return nullableposterUrl
}

func (service *MovieService) GetMoviesList(ctx context.Context, page int, rows int) ([]Movie, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListMovies(ctx, sqlc.ListMoviesParams{
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}

func (service *MovieService) AddMovie(ctx context.Context, movie AddMovieInput) (Movie, error) {
	result, err := service.Queries.AddMovie(ctx, sqlc.AddMovieParams{
		Title:           movie.Title,
		Synopsis:        movie.Synopsis,
		DurationMinutes: int32(movie.DurationMinutes),
		ReleasedAt:      movie.ReleasedAt,
		PosterUrl:       service.adaptNullablePoster(movie.PosterUrl),
	})
	if err != nil {
		return Movie{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *MovieService) GetMovie(ctx context.Context, id int) (Movie, error) {
	result, err := service.Queries.GetMovie(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return Movie{}, errors.New("invalid movie identifier")
	}
	if err != nil {
		return Movie{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *MovieService) UpdateMovie(ctx context.Context, id int, movie UpdateMovieInput) (Movie, error) {
	result, err := service.Queries.UpdateMovie(ctx, sqlc.UpdateMovieParams{
		ID:              int64(id),
		Title:           movie.Title,
		Synopsis:        movie.Synopsis,
		DurationMinutes: int32(movie.DurationMinutes),
		ReleasedAt:      movie.ReleasedAt,
		PosterUrl:       service.adaptNullablePoster(movie.PosterUrl),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Movie{}, errors.New("invalid movie identifier")
	}
	if err != nil {
		return Movie{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *MovieService) DeleteMovie(ctx context.Context, id int) error {
	if err := service.Queries.DeleteMovie(ctx, int64(id)); err != nil {
		return err
	}
	return nil
}

func (service *MovieService) GetGenreMoviesList(ctx context.Context, genre int, page int, rows int) ([]Movie, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListGenreMovies(ctx, sqlc.ListGenreMoviesParams{
		GenreID: int32(genre),
		Limit:   int32(rows),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}

func (service *MovieService) GetCelebrityMoviesList(ctx context.Context, celebrity int, page int, rows int) ([]Movie, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListCelebrityMovies(ctx, sqlc.ListCelebrityMoviesParams{
		CelebrityID: int64(celebrity),
		Limit:       int32(rows),
		Offset:      int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}
