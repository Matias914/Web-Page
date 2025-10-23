package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

func (service *CelebrityService) adaptQueryResult(result sqlc.Celebrity) Celebrity {
	return Celebrity{
		ID:        int(result.ID),
		Name:      result.Name,
		BirthDate: result.BirthDate,
	}
}

func (service *CelebrityService) adaptQueryResults(results []sqlc.Celebrity) []Celebrity {
	var celebrities = make([]Celebrity, len(results))
	for i, celebrity := range results {
		celebrities[i] = Celebrity{
			ID:        int(celebrity.ID),
			Name:      celebrity.Name,
			BirthDate: celebrity.BirthDate,
		}
	}
	return celebrities
}

func (service *CelebrityService) ValidateCelebrity(celebrity Celebrity) error {
	if celebrity.Name == "" {
		return errors.New("name is required")
	}
	if celebrity.BirthDate.IsZero() || celebrity.BirthDate.After(time.Now()) {
		return errors.New("celebrity birth date does not exists")
	}
	return nil
}

func (service *CelebrityService) GetCelebritiesList(ctx context.Context, page int, rows int) ([]Celebrity, error) {
	offset := (page - 1) * rows
	celebrities, err := service.Queries.ListCelebrities(ctx, sqlc.ListCelebritiesParams{
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(celebrities), nil
}

func (service *CelebrityService) GetMovieCelebritiesList(ctx context.Context, movie int, page int, rows int) ([]Celebrity, error) {
	offset := (page - 1) * rows
	celebrities, err := service.Queries.ListMovieCelebrities(ctx, sqlc.ListMovieCelebritiesParams{
		MovieID: int64(movie),
		Limit:   int32(rows),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(celebrities), nil
}

func (service *CelebrityService) AddCelebrity(ctx context.Context, celebrity Celebrity) (Celebrity, error) {
	added, err := service.Queries.AddCelebrity(ctx, sqlc.AddCelebrityParams{
		Name:      celebrity.Name,
		BirthDate: celebrity.BirthDate,
	})
	if err != nil {
		return Celebrity{}, err
	}
	return service.adaptQueryResult(added), nil
}

func (service *CelebrityService) GetCelebrity(ctx context.Context, id int) (Celebrity, error) {
	celebrity, err := service.Queries.GetCelebrity(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return Celebrity{}, errors.New("invalid celebrity identifier")
	}
	if err != nil {
		return Celebrity{}, err
	}
	return service.adaptQueryResult(celebrity), nil
}

func (service *CelebrityService) UpdateCelebrity(ctx context.Context, celebrity Celebrity) (Celebrity, error) {
	updated, err := service.Queries.UpdateCelebrity(ctx, sqlc.UpdateCelebrityParams{
		ID:        int64(celebrity.ID),
		Name:      celebrity.Name,
		BirthDate: celebrity.BirthDate,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Celebrity{}, errors.New("invalid celebrity identifier")
	}
	if err != nil {
		return Celebrity{}, err
	}
	return service.adaptQueryResult(updated), nil
}

func (service *CelebrityService) DeleteCelebrity(ctx context.Context, id int) error {
	if err := service.Queries.DeleteCelebrity(ctx, int64(id)); err != nil {
		return err
	}
	return nil
}
