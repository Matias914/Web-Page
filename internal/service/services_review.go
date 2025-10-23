package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

func (service *ReviewService) adaptQueryResult(result sqlc.Review) Review {
	return Review{
		UserID:    int(result.UserID),
		MovieID:   int(result.MovieID),
		Comment:   result.Comment,
		CreatedAt: result.CreatedAt,
	}
}

func (service *ReviewService) adaptQueryResults(results []sqlc.Review) []Review {
	var reviews = make([]Review, len(results))
	for i, review := range results {
		reviews[i] = Review{
			UserID:    int(review.UserID),
			MovieID:   int(review.MovieID),
			Comment:   review.Comment,
			CreatedAt: review.CreatedAt,
		}
	}
	return reviews
}

func (service *ReviewService) AddReview(ctx context.Context, review Review) (Review, error) {
	result, err := service.Queries.AddReview(ctx, sqlc.AddReviewParams{
		UserID:  int64(review.UserID),
		MovieID: int64(review.MovieID),
		Comment: review.Comment,
	})
	if err != nil {
		return Review{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *ReviewService) GetReview(ctx context.Context, review Review) (Review, error) {
	result, err := service.Queries.GetReview(ctx, sqlc.GetReviewParams{
		UserID:  int64(review.UserID),
		MovieID: int64(review.MovieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Review{}, errors.New("invalid review identifier")
	}
	if err != nil {
		return Review{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *ReviewService) UpdateReview(ctx context.Context, review Review) (Review, error) {
	result, err := service.Queries.UpdateReview(ctx, sqlc.UpdateReviewParams{
		UserID:  int64(review.UserID),
		MovieID: int64(review.MovieID),
		Comment: review.Comment,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Review{}, errors.New("invalid review identifier")
	}
	if err != nil {
		return Review{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *ReviewService) DeleteReview(ctx context.Context, review Review) error {
	if err := service.Queries.DeleteReview(ctx, sqlc.DeleteReviewParams{
		UserID:  int64(review.UserID),
		MovieID: int64(review.MovieID),
	}); err != nil {
		return err
	}
	return nil
}

func (service *ReviewService) GetMovieReviewsList(ctx context.Context, movie int, page int, rows int) ([]Review, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListMovieReviews(ctx, sqlc.ListMovieReviewsParams{
		MovieID: int64(movie),
		Limit:   int32(rows),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}

func (service *ReviewService) GetUserReviewsList(ctx context.Context, user int, page int, rows int) ([]Review, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListUserReviews(ctx, sqlc.ListUserReviewsParams{
		UserID: int64(user),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}
