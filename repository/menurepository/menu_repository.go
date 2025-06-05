package menurepository

import (
	"backend/model/domain"
	"context"
	"database/sql"
)

type MenuRepository interface {
	Create(ctx context.Context, tx *sql.Tx, menu domain.Menu) domain.Menu
	Update(ctx context.Context, tx *sql.Tx, menu domain.Menu) domain.Menu
	Delete(ctx context.Context, tx *sql.Tx, menu domain.Menu)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Menu
	FindById(ctx context.Context, tx *sql.Tx, menuId uint) (domain.Menu, error)
}
