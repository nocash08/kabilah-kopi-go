package usersrepository

import (
	"backend/model/dto/usersdto"
	"context"
	"database/sql"
)

type UsersRepository interface {
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (usersdto.UsersResponse, string, error)
}
