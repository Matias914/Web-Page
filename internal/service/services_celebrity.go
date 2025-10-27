package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
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

func (service *CelebrityService) adaptNullableGenreCelebrities(results []sqlc.ListMovieCelebritiesRow) []Celebrity {
	var celebrities = make([]Celebrity, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, celebrity := range results {
		celebrities[i] = Celebrity{
			ID:        int(celebrity.ID.Int64),
			Name:      celebrity.Name.String,
			BirthDate: celebrity.BirthDate.Time,
		}
	}
	return celebrities
}

// GetCelebritiesList retorna una lista de celebridades limitada. Los resultados se piden por páginas
// de cierto tamaño. Si ocurre un error con la función es porque hubo un error inesperado.
func (service *CelebrityService) GetCelebritiesList(ctx context.Context, page int, rows int) ([]Celebrity, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListCelebrities(ctx, sqlc.ListCelebritiesParams{
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}

// AddCelebrity agrega una celebridad a la base de datos. Si ocurre un error con la función es porque
// hay celebridades duplicadas, algún dato de la misma no cumple con alguna de las restricciones de
// integridad u ocurrió un error inesperado.
func (service *CelebrityService) AddCelebrity(ctx context.Context, celebrity CelebrityData) (Celebrity, error) {
	result, err := service.Queries.AddCelebrity(ctx, sqlc.AddCelebrityParams{
		Name:      celebrity.Name,
		BirthDate: celebrity.BirthDate,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintPK:
			return Celebrity{}, ErrCelebrityDuplicated
		case ErrCodeConstraintCHK:
			return Celebrity{}, ErrInvalidCelebrity
		}
	}
	if err != nil {
		return Celebrity{}, err
	}
	return service.adaptQueryResult(result), nil
}

// GetCelebrity obtiene una celebridad dad su ID. Si ocurre un error con la función es porque
// la película no existe u ocurrió un error inesperado.
func (service *CelebrityService) GetCelebrity(ctx context.Context, id int) (Celebrity, error) {
	result, err := service.Queries.GetCelebrity(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return Celebrity{}, ErrCelebrityNotFound
	}
	if err != nil {
		return Celebrity{}, err
	}
	return service.adaptQueryResult(result), nil
}

// UpdateCelebrity actualiza todos los datos de una celebridad. Si ocurre un error con la función
// es porque hay celebridades duplicadas, algún dato de la misma no cumple con alguna de las
// restricciones de integridad u ocurrió un error inesperado.
func (service *CelebrityService) UpdateCelebrity(ctx context.Context, id int, celebrity CelebrityData) (Celebrity, error) {
	result, err := service.Queries.UpdateCelebrity(ctx, sqlc.UpdateCelebrityParams{
		ID:        int64(id),
		Name:      celebrity.Name,
		BirthDate: celebrity.BirthDate,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintCHK:
			return Celebrity{}, ErrInvalidCelebrity
		default:
			return Celebrity{}, err
		}
	}
	if errors.Is(err, sql.ErrNoRows) {
		return Celebrity{}, ErrCelebrityNotFound
	}
	if err != nil {
		return Celebrity{}, err
	}
	return service.adaptQueryResult(result), nil
}

// DeleteCelebrity elimina una celebridad de la base de datos dado su ID. Si ocurre un error con
// la función es porque la celebridad no existe u ocurrió un error inesperado.
func (service *CelebrityService) DeleteCelebrity(ctx context.Context, id int) error {
	_, err := service.Queries.DeleteCelebrity(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return ErrCelebrityNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintFK:
			return ErrCelebrityReferenced
		default:
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// GetMovieCelebritiesList retorna la lista de celebridades de una película limitada. Se pueden pedir por páginas
// de cierto tamaño. Si ocurre un error con la función es porque la celebridad no existe u ocurrió un error inesperado.
// Los resultados se piden por páginas de cierto tamaño.
func (service *CelebrityService) GetMovieCelebritiesList(ctx context.Context, movie int, page int, rows int) ([]Celebrity, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListMovieCelebrities(ctx, sqlc.ListMovieCelebritiesParams{
		ID:     int64(movie),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if len(results) == 0 {
		return nil, ErrCelebrityNotFound
	}
	if len(results) == 1 && !results[0].ID.Valid {
		return []Celebrity{}, nil
	}
	if err != nil {
		return nil, err
	}
	return service.adaptNullableGenreCelebrities(results), nil
}
