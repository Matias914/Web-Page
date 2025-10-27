package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
)

func (service *RatingService) adaptQueryResult(result sqlc.Rating) Rating {
	return Rating{
		UserID:    int(result.UserID),
		MovieID:   int(result.MovieID),
		Rating:    int(result.Rating),
		CreatedAt: result.CreatedAt,
	}
}

func (service *RatingService) adaptMovieRatingsResults(results []sqlc.ListMovieRatingsRow) []Rating {
	var ratings = make([]Rating, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, rating := range results {
		ratings[i] = Rating{
			UserID:    int(rating.UserID.Int64),
			MovieID:   int(rating.MovieID.Int64),
			Rating:    int(rating.Rating.Int32),
			CreatedAt: rating.CreatedAt.Time,
		}
	}
	return ratings
}

func (service *RatingService) adaptUserRatingsResults(results []sqlc.ListUserRatingsRow) []Rating {
	var ratings = make([]Rating, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, rating := range results {
		ratings[i] = Rating{
			UserID:    int(rating.UserID.Int64),
			MovieID:   int(rating.MovieID.Int64),
			Rating:    int(rating.Rating.Int32),
			CreatedAt: rating.CreatedAt.Time,
		}
	}
	return ratings
}

// AddRating agrega un rating de un usuario a la base de datos. Para ello necesita del ID del
// usuario al cual agregarle su rating y los campos de datos. Si ocurre un error con la función
// es porque ya hay ratings del usuario para la película, algún dato de la misma no cumple con
// alguna de las restricciones de integridad u ocurrió un error inesperado.
func (service *RatingService) AddRating(ctx context.Context, userID int, rating RatingData) (Rating, error) {
	result, err := service.Queries.AddRating(ctx, sqlc.AddRatingParams{
		UserID:  int64(userID),
		MovieID: int64(rating.MovieID),
		Rating:  int32(rating.Rating),
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintPK:
			return Rating{}, ErrRatingDuplicated
		case ErrCodeConstraintCHK:
			return Rating{}, ErrInvalidRating
		case ErrCodeConstraintFK:
			if strings.Contains(pgErr.ConstraintName, "user") {
				return Rating{}, ErrUserNotFound
			}
			return Rating{}, ErrMovieNotFound
		}
	}
	if err != nil {
		return Rating{}, err
	}
	return service.adaptQueryResult(result), nil
}

// GetRating obtiene un rating dado el ID del usuario y la película. Si ocurre un error
// con la función es porque el usuario hizo el rating u ocurrió un error inesperado.
func (service *RatingService) GetRating(ctx context.Context, userID int, movieID int) (Rating, error) {
	result, err := service.Queries.GetRating(ctx, sqlc.GetRatingParams{
		UserID:  int64(userID),
		MovieID: int64(movieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Rating{}, ErrRatingNotFound
	}
	if err != nil {
		return Rating{}, err
	}
	return service.adaptQueryResult(result), nil
}

// UpdateRating actualiza todos los datos de un rating. Si ocurre un error con la función
// es porque algún dato de la misma no cumple con alguna de las restricciones de integridad
// u ocurrió un error inesperado.
func (service *RatingService) UpdateRating(ctx context.Context, userID int, movieID int, data UpdatableRatingData) (Rating, error) {
	result, err := service.Queries.UpdateRating(ctx, sqlc.UpdateRatingParams{
		UserID:  int64(userID),
		MovieID: int64(movieID),
		Rating:  int32(data.Rating),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Rating{}, ErrRatingNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintCHK:
			return Rating{}, ErrInvalidRating
		}
	}
	return service.adaptQueryResult(result), nil
}

// DeleteRating borra un rating dado el ID del usuario y de la película. Si ocurre un error
// es porque el rating no existía u ocurrió un error inesperado.
func (service *RatingService) DeleteRating(ctx context.Context, userID int, movieID int) error {
	_, err := service.Queries.DeleteRating(ctx, sqlc.DeleteRatingParams{
		UserID:  int64(userID),
		MovieID: int64(movieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return ErrRatingNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

// GetMovieRatingsList obtiene una lista con los ratings de una película dada por ID. Si ocurre un
// error es porque la película no existe u ocurrió un error inesperado. Los resultados se piden por
// páginas de cierto tamaño.
func (service *RatingService) GetMovieRatingsList(ctx context.Context, movieID int, page int, rows int) ([]Rating, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListMovieRatings(ctx, sqlc.ListMovieRatingsParams{
		ID:     int64(movieID),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return []Rating{}, ErrMovieNotFound
	}
	if len(results) == 1 && !results[0].MovieID.Valid {
		return []Rating{}, nil
	}
	return service.adaptMovieRatingsResults(results), nil
}

// GetUserRatingsList obtiene una lista con los ratings de un usuario dado por ID. Si ocurre un
// error es porque el usuario no existe u ocurrió un error inesperado. Los resultados se piden
// por páginas de cierto tamaño.
func (service *RatingService) GetUserRatingsList(ctx context.Context, userID int, page int, rows int) ([]Rating, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListUserRatings(ctx, sqlc.ListUserRatingsParams{
		ID:     int64(userID),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return []Rating{}, ErrUserNotFound
	}
	if len(results) == 1 && !results[0].UserID.Valid {
		return []Rating{}, nil
	}
	return service.adaptUserRatingsResults(results), nil
}
