package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
)

func (service *UserService) adaptQueryResult(result sqlc.User) User {
	return User{
		ID:        int(result.ID),
		Username:  result.Username,
		Password:  result.Password,
		Mail:      result.Mail,
		CreatedAt: result.CreatedAt,
	}
}

func (service *UserService) adaptQueryResults(results []sqlc.User) []User {
	var users = make([]User, len(results))
	for i, user := range results {
		users[i] = User{
			ID:        int(user.ID),
			Username:  user.Username,
			Password:  user.Password,
			Mail:      user.Mail,
			CreatedAt: user.CreatedAt,
		}
	}
	return users
}

func (service *UserService) GetUsersList(ctx context.Context, page int, rows int) ([]User, error) {
	offset := (page - 1) * rows

	result, err := service.Queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  int32(rows),
		Offset: int32(offset),
	})

	if err != nil {
		return nil, err
	}
	return service.adaptQueryResults(result), nil
}

func (service *UserService) AddUser(ctx context.Context, user User) (User, error) {
	result, err := service.Queries.AddUser(ctx, sqlc.AddUserParams{
		Username: user.Username,
		Mail:     user.Mail,
	})
	if err != nil {
		return User{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *UserService) GetUser(ctx context.Context, user int) (User, error) {
	result, err := service.Queries.GetUser(ctx, int64(user))
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, errors.New("invalid celebrity identifier")
	}
	if err != nil {
		return User{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *UserService) UpdateUser(ctx context.Context, user User) (User, error) {
	result, err := service.Queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:       int64(user.ID),
		Username: user.Username,
		Mail:     user.Mail,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, errors.New("invalid celebrity identifier")
	}
	if err != nil {
		return User{}, err
	}
	return service.adaptQueryResult(result), nil
}

func (service *UserService) DeleteUser(ctx context.Context, id int) error {
	if err := service.Queries.DeleteUser(ctx, int64(id)); err != nil {
		return err
	}
	return nil
}
