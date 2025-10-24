package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
)

// HasMovieWithGenre verifica si una película dada pertenece a un género dado. Si ocurre un error
// es porque no existe la relación entre ambos u ocurrió un error inesperado.
func (service *CategoryService) HasMovieWithGenre(ctx context.Context, movieID int, genreID int) error {
	_, err := service.Queries.GetCategory(ctx, sqlc.GetCategoryParams{
		GenreID: int32(genreID),
		MovieID: int64(movieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return ErrCategoryNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

// AddMovieGenre verifica si una película dada pertenece a un género dado. Si ocurre un error
// es porque no existe la relación entre ambos u ocurrió un error inesperado.
func (service *CategoryService) AddMovieGenre(ctx context.Context, movieID int, genreID int) error {
	err := service.Queries.AddCategory(ctx, sqlc.AddCategoryParams{
		GenreID: int32(genreID),
		MovieID: int64(movieID),
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23503":
			if strings.Contains(pgErr.ConstraintName, "MOVIE") {
				return ErrMovieNotFound
			}
			return ErrGenreNotFound
		case "23505":
			return ErrCategoryDuplicated
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteMovieGenre borra la el género de una pelicula. Si ocurre un error es porque no
// existe la relación u ocurrió un error inesperado.
func (service *CategoryService) DeleteMovieGenre(ctx context.Context, movieID int, genreID int) error {
	_, err := service.Queries.DeleteCategory(ctx, sqlc.DeleteCategoryParams{
		GenreID: int32(genreID),
		MovieID: int64(movieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return ErrCategoryNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
