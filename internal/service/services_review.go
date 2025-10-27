package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
)

func (service *ReviewService) adaptQueryResult(result sqlc.Review) Review {
	return Review{
		UserID:    int(result.UserID),
		MovieID:   int(result.MovieID),
		Comment:   result.Comment,
		CreatedAt: result.CreatedAt,
	}
}

func (service *ReviewService) adaptMovieReviewResults(results []sqlc.ListMovieReviewsRow) []Review {
	var reviews = make([]Review, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, review := range results {
		reviews[i] = Review{
			UserID:    int(review.UserID.Int64),
			MovieID:   int(review.MovieID.Int64),
			Comment:   review.Comment.String,
			CreatedAt: review.CreatedAt.Time,
		}
	}
	return reviews
}

func (service *ReviewService) adaptUserReviewResults(results []sqlc.ListUserReviewsRow) []Review {
	var reviews = make([]Review, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, review := range results {
		reviews[i] = Review{
			UserID:    int(review.UserID.Int64),
			MovieID:   int(review.MovieID.Int64),
			Comment:   review.Comment.String,
			CreatedAt: review.CreatedAt.Time,
		}
	}
	return reviews
}

// AddReview agrega una review de un usuario a la base de datos. Para ello necesita del ID del
// usuario al cual agregarle su review y los campos de datos. Si ocurre un error con la función
// es porque ya hay reviews del usuario para la película, algún dato de la misma no cumple con
// alguna de las restricciones de integridad u ocurrió un error inesperado.
func (service *ReviewService) AddReview(ctx context.Context, userID int, review ReviewData) (Review, error) {
	result, err := service.Queries.AddReview(ctx, sqlc.AddReviewParams{
		UserID:  int64(userID),
		MovieID: int64(review.MovieID),
		Comment: review.Comment,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintPK:
			return Review{}, ErrReviewDuplicated
		case ErrCodeConstraintCHK:
			return Review{}, ErrInvalidReview
		case ErrCodeConstraintFK:
			if strings.Contains(pgErr.ConstraintName, "user") {
				return Review{}, ErrUserNotFound
			}
			return Review{}, ErrMovieNotFound
		}
	}
	if err != nil {
		return Review{}, err
	}
	return service.adaptQueryResult(result), nil
}

// GetReview obtiene una review dado el ID del usuario y la película. Si ocurre un error
// con la función es porque el usuario hizo la review u ocurrió un error inesperado.
func (service *ReviewService) GetReview(ctx context.Context, userID int, movieID int) (Review, error) {
	result, err := service.Queries.GetReview(ctx, sqlc.GetReviewParams{
		UserID:  int64(userID),
		MovieID: int64(movieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Review{}, ErrReviewNotFound
	}
	if err != nil {
		return Review{}, err
	}
	return service.adaptQueryResult(result), nil
}

// UpdateReview actualiza todos los datos de una review. Si ocurre un error con la función
// es porque algún dato de la misma no cumple con alguna de las restricciones de integridad
// u ocurrió un error inesperado.
func (service *ReviewService) UpdateReview(ctx context.Context, userID int, movieID int, data UpdatableReviewData) (Review, error) {
	result, err := service.Queries.UpdateReview(ctx, sqlc.UpdateReviewParams{
		UserID:  int64(userID),
		MovieID: int64(movieID),
		Comment: data.Comment,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Review{}, ErrReviewNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintCHK:
			return Review{}, ErrInvalidReview
		}
	}
	return service.adaptQueryResult(result), nil
}

// DeleteReview borra una review dado el ID del usuario y de la película. Si ocurre un error
// es porque la review no existía u ocurrió un error inesperado.
func (service *ReviewService) DeleteReview(ctx context.Context, userID int, movieID int) error {
	_, err := service.Queries.DeleteReview(ctx, sqlc.DeleteReviewParams{
		UserID:  int64(userID),
		MovieID: int64(movieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return ErrReviewNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

// GetMovieReviewsList obtiene una lista con los reviews de una película dada por ID. Si ocurre un
// error es porque la película no existe u ocurrió un error inesperado. Los resultados se piden por
// páginas de cierto tamaño.
func (service *ReviewService) GetMovieReviewsList(ctx context.Context, movie int, page int, rows int) ([]Review, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListMovieReviews(ctx, sqlc.ListMovieReviewsParams{
		ID:     int64(movie),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return []Review{}, ErrMovieNotFound
	}
	if len(results) == 1 && !results[0].MovieID.Valid {
		return []Review{}, nil
	}
	return service.adaptMovieReviewResults(results), nil
}

// GetUserReviewsList obtiene una lista con los reviews de un usuario dado por ID. Si ocurre un
// error es porque el usuario no existe u ocurrió un error inesperado. Los resultados se piden
// por páginas de cierto tamaño.
func (service *ReviewService) GetUserReviewsList(ctx context.Context, user int, page int, rows int) ([]Review, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListUserReviews(ctx, sqlc.ListUserReviewsParams{
		ID:     int64(user),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return []Review{}, ErrUserNotFound
	}
	if len(results) == 1 && !results[0].UserID.Valid {
		return []Review{}, nil
	}
	return service.adaptUserReviewResults(results), nil
}
