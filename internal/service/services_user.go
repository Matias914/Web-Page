package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Matias914/Web-Page/internal/storage/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword retorna el hash de la contraseña que se debe guardar en la base de datos. La
// encriptación no es reversible por motivos de seguridad.
func (service *UserService) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword compara la contraseña en texto plano con el hash guardado en la base de datos.
// Retorna nil si coincide.
func (service *UserService) CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

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

// GetUsersList retorna una lista de usuarios limitada. Los resultados se piden por páginas de
// cierto tamaño. Si ocurre un error con la función es porque hubo un error inesperado.
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

// AddUser agrega un usuario a la base de datos. Si ocurre un error con la función es porque
// hay usuarios duplicados, algún dato de la misma no cumple con alguna de las restricciones
// de integridad u ocurrió un error inesperado.
func (service *UserService) AddUser(ctx context.Context, user UserData) (User, error) {
	hashed, err := service.HashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	result, err := service.Queries.AddUser(ctx, sqlc.AddUserParams{
		Username: user.Username,
		Password: hashed,
		Mail:     user.Mail,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// Por los claves alternativas
		case ErrCodeConstraintPK:

			return User{}, ErrUserDuplicated
		case ErrCodeConstraintCHK:
			return User{}, ErrInvalidUser
		}
	}
	if err != nil {
		return User{}, err
	}
	return service.adaptQueryResult(result), nil
}

// GetUser obtiene un usuario dado su ID. Si ocurre un error con la función es porque el usuario
// no existe u ocurrió un error inesperado.
func (service *UserService) GetUser(ctx context.Context, id int) (User, error) {
	result, err := service.Queries.GetUser(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrUserNotFound
	}
	if err != nil {
		return User{}, err
	}
	return service.adaptQueryResult(result), nil
}

// UpdateUser actualiza todos los datos de un usuario. Si ocurre un error con la función
// es porque hay usuarios duplicados, algún dato de la misma no cumple con alguna de las
// restricciones de integridad u ocurrió un error inesperado.
func (service *UserService) UpdateUser(ctx context.Context, id int, data UpdatableUserData) (User, error) {
	result, err := service.Queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:       int64(id),
		Username: data.Username,
		Mail:     data.Mail,
	})
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// Por sus claves alternativas
		case ErrCodeConstraintPK:
			return User{}, ErrUserDuplicated
		case ErrCodeConstraintCHK:
			return User{}, ErrInvalidUser
		}
	}
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrUserNotFound
	}
	if err != nil {
		return User{}, err
	}
	return service.adaptQueryResult(result), nil
}

// DeleteUser elimina un usuario de la base de datos dado su ID. Si ocurre un error con
// la función es porque el usuario no existe u ocurrió un error inesperado.
func (service *UserService) DeleteUser(ctx context.Context, id int) error {
	_, err := service.Queries.DeleteUser(ctx, int64(id))
	if errors.Is(err, sql.ErrNoRows) {
		return ErrUserNotFound
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case ErrCodeConstraintFK:
			return ErrUserReferenced
		}
	}
	if err != nil {
		return err
	}
	return nil
}
