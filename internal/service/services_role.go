package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
)

func (service *RoleService) adaptQueryResult(result sqlc.Role) Role {
	return Role{
		MovieID:     int(result.MovieID),
		CelebrityID: int(result.CelebrityID),
		Role:        result.Role,
	}
}

func (service *RoleService) adaptMovieRolesResults(results []sqlc.ListMovieRolesRow) []Role {
	var roles = make([]Role, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, role := range results {
		roles[i] = Role{
			MovieID:     int(role.MovieID.Int64),
			CelebrityID: int(role.CelebrityID.Int64),
			Role:        role.Role.String,
		}
	}
	return roles
}

func (service *RoleService) adaptUserRolesResults(results []sqlc.ListCelebrityRolesRow) []Role {
	var roles = make([]Role, len(results))
	// Se sabe de antes que no pueden ser vacíos, de otra manera, habría una fila
	for i, role := range results {
		roles[i] = Role{
			MovieID:     int(role.MovieID.Int64),
			CelebrityID: int(role.CelebrityID.Int64),
			Role:        role.Role.String,
		}
	}
	return roles
}

// AddRole agrega un rol de una celebridad a una película a la base de datos. Para ello necesita del ID
// de la película a la cual agregarle el rol y los campos de datos. Si ocurre un error con la función
// es porque ya existe un rol de la celebridad para la película, algún dato de la misma no cumple con
// alguna de las restricciones de integridad u ocurrió un error inesperado.
func (service *RoleService) AddRole(ctx context.Context, movieID int, role RoleData) (Role, error) {
	result, err := service.Queries.AddRole(ctx, sqlc.AddRoleParams{
		MovieID:     int64(movieID),
		CelebrityID: int64(role.CelebrityID),
		Role:        role.Role,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintPK:
			return Role{}, ErrRoleDuplicated
		case ErrCodeConstraintCHK:
			return Role{}, ErrInvalidRole
		case ErrCodeConstraintFK:
			if strings.Contains(pgErr.ConstraintName, "movie") {
				return Role{}, ErrMovieNotFound
			}
			return Role{}, ErrCelebrityNotFound
		}
	}
	if err != nil {
		return Role{}, err
	}
	return service.adaptQueryResult(result), nil
}

// GetRole obtiene un rol dado el ID de la celebridad y la película. Si ocurre un
// error con la función es porque la celebridad no tiene un rol en la película u
// ocurrió un error inesperado.
func (service *RoleService) GetRole(ctx context.Context, movieID int, celebrityID int) (Role, error) {
	result, err := service.Queries.GetRole(ctx, sqlc.GetRoleParams{
		MovieID:     int64(movieID),
		CelebrityID: int64(celebrityID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Role{}, ErrRoleNotFound
	}
	if err != nil {
		return Role{}, err
	}
	return service.adaptQueryResult(result), nil
}

// UpdateRole actualiza todos los datos de un rol. Si ocurre un error con la función
// es porque algún dato de la misma no cumple con alguna de las restricciones de integridad
// u ocurrió un error inesperado.
func (service *RoleService) UpdateRole(ctx context.Context, movieID int, celebrityID int, data UpdatableRoleData) (Role, error) {
	result, err := service.Queries.UpdateRole(ctx, sqlc.UpdateRoleParams{
		MovieID:     int64(movieID),
		CelebrityID: int64(celebrityID),
		Role:        data.Role,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintCHK:
			return Role{}, ErrInvalidRole
		}
	}
	if err != nil {
		return Role{}, err
	}
	return service.adaptQueryResult(result), nil
}

// DeleteRole borra un rol dado el ID de la celebridad y de la película. Si ocurre
// un error es porque la celebridad no tenia un rol en la película u ocurrió un error
// inesperado.
func (service *RoleService) DeleteRole(ctx context.Context, movieID int, celebrityID int) error {
	_, err := service.Queries.DeleteRole(ctx, sqlc.DeleteRoleParams{
		MovieID:     int64(movieID),
		CelebrityID: int64(celebrityID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return ErrRoleNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

// GetMovieRolesList obtiene una lista con los roles de una película dada por ID. Si ocurre un
// error es porque la película no existe u ocurrió un error inesperado. Los resultados se piden
// por páginas de cierto tamaño.
func (service *RoleService) GetMovieRolesList(ctx context.Context, movieID int, page int, rows int) ([]Role, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListMovieRoles(ctx, sqlc.ListMovieRolesParams{
		ID:     int64(movieID),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return []Role{}, ErrMovieNotFound
	}
	if len(results) == 1 && !results[0].MovieID.Valid {
		return []Role{}, nil
	}
	return service.adaptMovieRolesResults(results), nil
}

// GetCelebrityRolesList obtiene una lista con los roles de una celebridad dada por ID. Si ocurre un
// error es porque la celebridad no existe u ocurrió un error inesperado. Los resultados se piden por
// páginas de cierto tamaño.
func (service *RoleService) GetCelebrityRolesList(ctx context.Context, celebrityID int, page int, rows int) ([]Role, error) {
	offset := (page - 1) * rows
	results, err := service.Queries.ListCelebrityRoles(ctx, sqlc.ListCelebrityRolesParams{
		ID:     int64(celebrityID),
		Limit:  int32(rows),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return []Role{}, ErrCelebrityNotFound
	}
	if len(results) == 1 && !results[0].CelebrityID.Valid {
		return []Role{}, nil
	}
	return service.adaptUserRolesResults(results), nil
}
