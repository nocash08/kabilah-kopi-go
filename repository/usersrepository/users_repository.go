package usersrepository

import (
	"backend/model/domain"
	"context"
	"database/sql"
)

type UsersRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.Users, string, error)
	CountUsers(ctx context.Context, tx *sql.Tx) int
}
