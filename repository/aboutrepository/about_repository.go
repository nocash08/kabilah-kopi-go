package aboutrepository

import (
	"backend/model/domain"
	"context"
	"database/sql"
)

type AboutRepository interface {
	Create(ctx context.Context, tx *sql.Tx, about domain.About) domain.About
	Update(ctx context.Context, tx *sql.Tx, about domain.About) domain.About
	FindAll(ctx context.Context, tx *sql.Tx) []domain.About
	FindById(ctx context.Context, tx *sql.Tx, aboutId uint) (domain.About, error)
}
