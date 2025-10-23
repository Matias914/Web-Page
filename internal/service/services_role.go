package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

func (service *RoleService) adaptQueryResult(result sqlc.Role) Role {
	return Role{
		CelebrityID: int(result.CelebrityID),
		MovieID:     int(result.MovieID),
		Role:        result.Role,
	}
}

func (service *RoleService) adaptQueryResults(results []sqlc.Role) []Role {
	var roles = make([]Role, len(results))
	for i, role := range results {
		roles[i] = Role{
			CelebrityID: int(role.CelebrityID),
			MovieID:     int(role.MovieID),
			Role:        role.Role,
		}
	}
	return roles
}

func (service *RoleService) ValidateRole(role Role) error {
	if role.CelebrityID <= 0 {
		return errors.New("invalid celebrity identifier")
	}
	if role.MovieID <= 0 {
		return errors.New("invalid movie identifier")
	}
	//	if len(role.Role) <= 0 || len(role.Role) > MaxRoleLengthLength {
	//		text := fmt.Sprintf("role length is not between 0 and %d", MaxRoleLength)
	//		return errors.New(text)
	//	}
	return nil
}

func (service *RoleService) AddRole(ctx context.Context, role Role) (Role, error) {
	result, err := service.Queries.AddRole(ctx, sqlc.AddRoleParams{
		CelebrityID: int64(role.CelebrityID),
		MovieID:     int64(role.MovieID),
		Role:        role.Role,
	})
	if err != nil {
		return Role{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *RoleService) GetRole(ctx context.Context, role Role) (Role, error) {
	result, err := service.Queries.GetRole(ctx, sqlc.GetRoleParams{
		CelebrityID: int64(role.CelebrityID),
		MovieID:     int64(role.MovieID),
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Role{}, errors.New("invalid role identifier")
	}
	if err != nil {
		return Role{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *RoleService) UpdateRole(ctx context.Context, role Role) (Role, error) {
	result, err := service.Queries.UpdateRole(ctx, sqlc.UpdateRoleParams{
		CelebrityID: int64(role.CelebrityID),
		MovieID:     int64(role.MovieID),
		Role:        role.Role,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return Role{}, errors.New("invalid role identifier")
	}
	if err != nil {
		return Role{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *RoleService) DeleteRole(ctx context.Context, role Role) error {
	if err := service.Queries.DeleteRole(ctx, sqlc.DeleteRoleParams{
		CelebrityID: int64(role.CelebrityID),
		MovieID:     int64(role.MovieID),
	}); err != nil {
		return err
	}
	return nil
}
