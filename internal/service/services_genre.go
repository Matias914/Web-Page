package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

func (service *GenreService) adaptQueryResult(result sqlc.Genre) Genre {
	return Genre{
		ID:   result.ID,
		Name: result.Name,
	}
}

func (service *GenreService) adaptQueryResults(results []sqlc.Genre) []Genre {
	var genres = make([]Genre, len(results))
	for i, genre := range results {
		genres[i] = Genre{
			ID:   genre.ID,
			Name: genre.Name,
		}
	}
	return genres
}

func (service *GenreService) GetGenresList(ctx context.Context, page int, rows int) ([]Genre, error) {
	offset := (page - 1) * rows
	result, err := service.Queries.ListGenres(ctx, sqlc.ListGenresParams{
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(result), nil
}

func (service *GenreService) AddGenre(ctx context.Context, genre string) (Genre, error) {
	result, err := service.Queries.AddGenre(ctx, strings.ToLower(genre))
	if err != nil {
		return Genre{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *GenreService) GetGenre(ctx context.Context, id int) (Genre, error) {
	result, err := service.Queries.GetGenre(ctx, int32(id))
	if errors.Is(err, sql.ErrNoRows) {
		return Genre{}, errors.New("invalid genre identifier")
	}
	if err != nil {
		return Genre{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *GenreService) UpdateGenre(ctx context.Context, genre Genre) (Genre, error) {
	result, err := service.Queries.UpdateGenre(ctx, sqlc.UpdateGenreParams{
		ID:   genre.ID,
		Name: genre.Name,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Genre{}, errors.New("invalid genre identifier")
	}
	if err != nil {
		return Genre{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *GenreService) DeleteGenre(ctx context.Context, id int) error {
	if err := service.Queries.DeleteGenre(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (service *GenreService) GetMovieGenresList(ctx context.Context, movie int, page int, rows int) ([]Genre, error) {
	offset := (page - 1) * rows
	result, err := service.Queries.ListMovieGenres(ctx, sqlc.ListMovieGenresParams{
		MovieID: int64(movie),
		Limit:   int32(rows),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(result), nil
}
