package eventrepository

import (
	"backend/model/domain"
	"context"
	"database/sql"
)

type EventRepository interface {
	Create(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event
	Update(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event
	Delete(ctx context.Context, tx *sql.Tx, event domain.Event)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Event
	FindById(ctx context.Context, tx *sql.Tx, eventId uint) (domain.Event, error)
}
