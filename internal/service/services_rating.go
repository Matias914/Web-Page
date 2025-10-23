package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

func (service *RatingService) adaptQueryResult(result sqlc.Rating) Rating {
	return Rating{
		UserID:    int(result.UserID),
		MovieID:   int(result.MovieID),
		Rating:    int(result.Rating),
		CreatedAt: result.CreatedAt,
	}
}

func (service *RatingService) adaptQueryResults(results []sqlc.Rating) []Rating {
	var ratings = make([]Rating, len(results))
	for i, rating := range results {
		ratings[i] = Rating{
			UserID:    int(rating.UserID),
			MovieID:   int(rating.MovieID),
			Rating:    int(rating.Rating),
			CreatedAt: rating.CreatedAt,
		}
	}
	return ratings
}

func (service *RatingService) ValidateRating(rating Rating) error {
	if rating.UserID <= 0 {
		return errors.New("invalid user identifier")
	}
	if rating.MovieID <= 0 {
		return errors.New("invalid movie identifier")
	}
	if rating.Rating <= 0 || rating.Rating > 10 {
		return errors.New("rating is not between 0 and 10")
	}
	if rating.CreatedAt.IsZero() {
		return errors.New("created date does not exists")
	}
	return nil
}

func (service *RatingService) AddRating(ctx context.Context, rating Rating) (Rating, error) {
	result, err := service.Queries.AddRating(ctx, sqlc.AddRatingParams{
		UserID:  int64(rating.UserID),
		MovieID: int64(rating.MovieID),
		Rating:  int32(rating.Rating),
	})
	if err != nil {
		return Rating{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *RatingService) GetRating(ctx context.Context, rating Rating) (Rating, error) {
	result, err := service.Queries.GetRating(ctx, sqlc.GetRatingParams{
		UserID:  int64(rating.UserID),
		MovieID: int64(rating.MovieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Rating{}, errors.New("invalid rating identifier")
	}
	if err != nil {
		return Rating{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *RatingService) UpdateRating(ctx context.Context, rating Rating) (Rating, error) {
	result, err := service.Queries.UpdateRating(ctx, sqlc.UpdateRatingParams{
		UserID:  int64(rating.UserID),
		MovieID: int64(rating.MovieID),
		Rating:  int32(rating.Rating),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Rating{}, errors.New("invalid rating identifier")
	}
	if err != nil {
		return Rating{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *RatingService) DeleteRating(ctx context.Context, rating Rating) error {
	if err := service.Queries.DeleteRating(ctx, sqlc.DeleteRatingParams{
		UserID:  int64(rating.UserID),
		MovieID: int64(rating.MovieID),
	}); err != nil {
		return err
	}
	return nil
}

func (service *RatingService) GetMovieRatingsList(ctx context.Context, movie int, page int, rows int) ([]Rating, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListMovieRatings(ctx, sqlc.ListMovieRatingsParams{
		MovieID: int64(movie),
		Limit:   int32(rows),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}

func (service *RatingService) GetUserRatingsList(ctx context.Context, user int, page int, rows int) ([]Rating, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListUserRatings(ctx, sqlc.ListUserRatingsParams{
		UserID: int64(user),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}
