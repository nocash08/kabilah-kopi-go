package usersrepository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UsersRepositoryImpl struct {
}

func NewUsersRepository() UsersRepository {
	return &UsersRepositoryImpl{}
}

func (repository *UsersRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "INSERT INTO users (username, password, is_admin) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.IsAdmin)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = uint(id)
	return user
}

func (repository *UsersRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.Users, string, error) {
	SQL := "SELECT id, username, password, is_admin FROM users WHERE username = ?"

	var user domain.Users
	var password string
	err := tx.QueryRowContext(ctx, SQL, username).Scan(&user.Id, &user.Username, &password, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Users{}, "", errors.New("user not found")
		}
		panic(err)
	}

	return user, password, nil
}

func (repository *UsersRepositoryImpl) CountUsers(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(*) FROM users"
	var count int
	err := tx.QueryRowContext(ctx, SQL).Scan(&count)
	helper.PanicIfError(err)
	return count
}
