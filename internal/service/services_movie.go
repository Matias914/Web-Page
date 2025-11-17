package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
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

func (service *MovieService) adaptGenreMoviesResults(results []sqlc.ListGenreMoviesRow) []Movie {
	var movies = make([]Movie, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, movie := range results {
		posterUrl := ""
		if movie.PosterUrl.Valid {
			posterUrl = movie.PosterUrl.String
		}
		movies[i] = Movie{
			ID:              int(movie.ID.Int64),
			Title:           movie.Title.String,
			Synopsis:        movie.Synopsis.String,
			DurationMinutes: int(movie.DurationMinutes.Int32),
			ReleasedAt:      movie.ReleasedAt.Time,
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

// GetMoviesList retorna una lista de peliculas limitada. Los resultados se piden por páginas de
// cierto tamaño. Si ocurre un error con la función es porque hubo un error inesperado.
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

// AddMovie agrega una película a la base de datos. La URL del póster puede omitirse.
// Si ocurre un error con la función es porque hay películas duplicadas, algún dato de la misma
// no cumple con alguna de las restricciones de integridad u ocurrió un error inesperado.
func (service *MovieService) AddMovie(ctx context.Context, movie MovieData) (Movie, error) {
	result, err := service.Queries.AddMovie(ctx, sqlc.AddMovieParams{
		Title:           movie.Title,
		Synopsis:        movie.Synopsis,
		DurationMinutes: int32(movie.DurationMinutes),
		ReleasedAt:      movie.ReleasedAt,
		PosterUrl:       service.adaptNullablePoster(movie.PosterUrl),
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// Por los claves alternativas
		case ErrCodeConstraintPK:

			return Movie{}, ErrMovieDuplicated
		case ErrCodeConstraintCHK:
			return Movie{}, ErrInvalidMovie
		}
	}
	if err != nil {
		return Movie{}, err
	}
	return service.adaptQueryResult(result), nil
}

// GetMovie obtiene una película dado su ID. Si ocurre un error con la función es porque
// la película no existe u ocurrió un error inesperado.
func (service *MovieService) GetMovie(ctx context.Context, id int) (Movie, error) {
	result, err := service.Queries.GetMovie(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return Movie{}, ErrMovieNotFound
	}
	if err != nil {
		return Movie{}, err
	}
	return service.adaptQueryResult(result), nil
}

// UpdateMovie actualiza todos los datos de una película. Si ocurre un error con la función
// es porque hay películas duplicadas, algún dato de la misma no cumple con alguna de las
// restricciones de integridad u ocurrió un error inesperado.
func (service *MovieService) UpdateMovie(ctx context.Context, id int, movie MovieData) (Movie, error) {
	result, err := service.Queries.UpdateMovie(ctx, sqlc.UpdateMovieParams{
		Title:           movie.Title,
		Synopsis:        movie.Synopsis,
		ReleasedAt:      movie.ReleasedAt,
		DurationMinutes: int32(movie.DurationMinutes),
		PosterUrl:       service.adaptNullablePoster(movie.PosterUrl),
		ID:              int64(id),
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// Por los claves alternativas
		case ErrCodeConstraintPK:
			return Movie{}, ErrMovieDuplicated
		case ErrCodeConstraintCHK:
			return Movie{}, ErrInvalidMovie
		}
	}
	if errors.Is(err, sql.ErrNoRows) {
		return Movie{}, ErrMovieNotFound
	}
	if err != nil {
		return Movie{}, err
	}
	return service.adaptQueryResult(result), nil
}

// DeleteMovie elimina una película de la base de datos dado su ID. Si ocurre un error con
// la función es porque la película no existe u ocurrió un error inesperado.
func (service *MovieService) DeleteMovie(ctx context.Context, id int) error {
	_, err := service.Queries.DeleteMovie(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return ErrMovieNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintFK:
			return ErrMovieReferenced
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// GetGenreMoviesList retorna la lista de películas de un género limitada. Se pueden pedir por páginas
// de cierto tamaño. Si ocurre un error con la función es porque el género no existe u ocurrió un error
// inesperado. Los resultados se piden por páginas de cierto tamaño.
func (service *MovieService) GetGenreMoviesList(ctx context.Context, genre int, page int, rows int) ([]Movie, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListGenreMovies(ctx, sqlc.ListGenreMoviesParams{
		ID:     int32(genre),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if len(results) == 0 {
		return nil, ErrMovieNotFound
	}
	if len(results) == 1 && !results[0].ID.Valid {
		return []Movie{}, nil
	}
	if err != nil {
		return nil, err
	}
	return service.adaptGenreMoviesResults(results), nil
}
