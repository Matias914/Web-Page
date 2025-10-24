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

func (service *MovieService) adaptNullablePoster(posterUrl string) sql.NullString {
	var nullableposterUrl = sql.NullString{Valid: false}
	if posterUrl != "" {
		nullableposterUrl.String = posterUrl
		nullableposterUrl.Valid = true
	}
	return nullableposterUrl
}

func (service *MovieService) adaptNullableGenreMovies(results []sqlc.ListGenreMoviesRow) []Movie {
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

// GetMoviesList retorna una lista de peliculas limitada. Se pueden pedir por páginas de cierto tamaño.
// Si ocurre un error con la función es porque hubo un error inesperado.
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
func (service *MovieService) AddMovie(ctx context.Context, movie AddMovieInput) (Movie, error) {
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
		case "23505":
			return Movie{}, ErrMovieDuplicated
		case "23514":
			return Movie{}, ErrInvalidMovie
		default:
			return Movie{}, err
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
func (service *MovieService) UpdateMovie(ctx context.Context, id int, movie UpdateMovieInput) (Movie, error) {
	result, err := service.Queries.UpdateMovie(ctx, sqlc.UpdateMovieParams{
		ID:              int64(id),
		Title:           movie.Title,
		Synopsis:        movie.Synopsis,
		DurationMinutes: int32(movie.DurationMinutes),
		ReleasedAt:      movie.ReleasedAt,
		PosterUrl:       service.adaptNullablePoster(movie.PosterUrl),
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return Movie{}, ErrMovieDuplicated
		case "23514":
			return Movie{}, ErrInvalidMovie
		default:
			return Movie{}, err
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

// DeleteMovie elimina una película a la base de datos dado su ID. Si ocurre un error con
// la función es porque hay películas no existe u ocurrió un error inesperado.
func (service *MovieService) DeleteMovie(ctx context.Context, id int) error {
	_, err := service.Queries.DeleteMovie(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return ErrMovieNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23503":
			return ErrMovieReferenced
		default:
			return err
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// GetGenreMoviesList retorna la lista de películas de un género limitada. Se pueden pedir por páginas
// de cierto tamaño. Si ocurre un error con la función es porque el género no existe, el género u ocurrió
// un error inesperado.
func (service *MovieService) GetGenreMoviesList(ctx context.Context, genre int, page int, rows int) ([]Movie, error) {
	offset := (page - 1) * rows
	result, err := service.Queries.ListGenreMovies(ctx, sqlc.ListGenreMoviesParams{
		ID:     int32(genre),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if len(result) == 0 {
		return nil, ErrMovieNotFound
	}
	if len(result) == 1 && !result[0].ID.Valid {
		return []Movie{}, nil
	}
	if err != nil {
		return nil, err
	}
	return service.adaptNullableGenreMovies(result), nil
}

// GetCelebrityMoviesList agrega una película a la base de datos. La URL del póster puede omitirse.
// Si ocurre un error con la función es porque hay películas duplicadas, algún dato de la misma
// no cumple con alguna de las restricciones de integridad u ocurrió un error inesperado.
func (service *MovieService) GetCelebrityMoviesList(ctx context.Context, celebrity int, page int, rows int) ([]Movie, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListCelebrityMovies(ctx, sqlc.ListCelebrityMoviesParams{
		CelebrityID: int64(celebrity),
		Limit:       int32(rows),
		Offset:      int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(results), nil
}
