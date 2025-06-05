package usersrepository

import (
	"backend/model/dto/usersdto"
	"context"
	"database/sql"
	"errors"
)

type UsersRepositoryImpl struct {
}

func NewUsersRepository() UsersRepository {
	return &UsersRepositoryImpl{}
}

func (repository *UsersRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (usersdto.UsersResponse, string, error) {
	SQL := "SELECT id, username, password, is_admin FROM users WHERE username = $1"
	rows, err := tx.QueryContext(ctx, SQL, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var user usersdto.UsersResponse
		var password string
		err := rows.Scan(&user.Id, &user.Username, &password, &user.IsAdmin)
		if err != nil {
			panic(err)
		}
		return user, password, nil
	} else {
		return usersdto.UsersResponse{}, "", errors.New("user not found")
	}
}
