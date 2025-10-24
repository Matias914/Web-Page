package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
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

func (service *GenreService) adaptNullableMovieGenres(results []sqlc.ListMovieGenresRow) []Genre {
	var genres = make([]Genre, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, genre := range results {
		genres[i] = Genre{
			ID:   genre.ID.Int32,
			Name: genre.Name.String,
		}
	}
	return genres
}

// GetGenresList retorna una lista de géneros limitada. Se pueden pedir por páginas de cierto tamaño.
// Si ocurre un error con la función es porque hubo un error inesperado.
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

// AddGenre agrega un género a la base de datos. Por una cuestión de consistencia, se agregan en
// minúsculas. Si ocurre un error con la función es porque hay géneros duplicados, algún dato
// del mismo no cumple con alguna de las restricciones de integridad u ocurrió un error inesperado.
func (service *GenreService) AddGenre(ctx context.Context, genre AddGenreInput) (Genre, error) {
	result, err := service.Queries.AddGenre(ctx, strings.ToLower(genre.Name))
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return Genre{}, ErrGenreDuplicated
		case "23514":
			return Genre{}, ErrInvalidGenre
		default:
			return Genre{}, err
		}
	}
	if err != nil {
		return Genre{}, err
	}
	return service.adaptQueryResult(result), nil
}

// GetGenre obtiene un género dado su ID. Si ocurre un error con la función es porque
// el género no existe u ocurrió un error inesperado.
func (service *GenreService) GetGenre(ctx context.Context, id int) (Genre, error) {
	result, err := service.Queries.GetGenre(ctx, int32(id))
	if errors.Is(err, sql.ErrNoRows) {
		return Genre{}, ErrGenreNotFound
	}
	if err != nil {
		return Genre{}, err
	}
	return service.adaptQueryResult(result), nil
}

// UpdateGenre actualiza todos los datos de un género. Si ocurre un error con la función
// es porque hay géneros duplicados, algún dato del mismo no cumple con alguna de las
// restricciones de integridad u ocurrió un error inesperado.
func (service *GenreService) UpdateGenre(ctx context.Context, id int, genre UpdateGenreInput) (Genre, error) {
	result, err := service.Queries.UpdateGenre(ctx, sqlc.UpdateGenreParams{
		ID:   int32(id),
		Name: genre.Name,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return Genre{}, ErrGenreDuplicated
		case "23514":
			return Genre{}, ErrInvalidGenre
		default:
			return Genre{}, err
		}
	}
	if errors.Is(err, sql.ErrNoRows) {
		return Genre{}, ErrGenreNotFound
	}
	if err != nil {
		return Genre{}, err
	}
	return service.adaptQueryResult(result), nil
}

// DeleteGenre elimina un género de la base de datos dado su ID. Si ocurre un error con
// la función es porque el género no existe u ocurrió un error inesperado.
func (service *GenreService) DeleteGenre(ctx context.Context, id int) error {
	if err := service.Queries.DeleteGenre(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

// GetMovieGenresList retorna la lista de géneros de una película limitada. Se pueden pedir por páginas
// de cierto tamaño. Si ocurre un error con la función es porque la película no existe u ocurrió un error
// inesperado.
func (service *GenreService) GetMovieGenresList(ctx context.Context, movie int, page int, rows int) ([]Genre, error) {
	offset := (page - 1) * rows
	result, err := service.Queries.ListMovieGenres(ctx, sqlc.ListMovieGenresParams{
		ID:     int64(movie),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if len(result) == 0 {
		return nil, ErrMovieNotFound
	}
	if len(result) == 1 && !result[0].ID.Valid {
		return []Genre{}, nil
	}
	if err != nil {
		return nil, err
	}
	return service.adaptNullableMovieGenres(result), nil
}
