package eventrepository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"
	"errors"
)

type EventRepositoryImpl struct {
}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{}
}

func (repository *EventRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event {
	query := "INSERT INTO event (heading, subheading, thumbnail) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, event.Heading, event.Subheading, event.Thumbnail)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	event.Id = uint(id)
	return event
}

func (repository *EventRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event {
	query := "UPDATE event SET heading = ?, subheading = ?, thumbnail = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, event.Heading, event.Subheading, event.Thumbnail, event.Id)
	helper.PanicIfError(err)

	return event
}

func (repository *EventRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, event domain.Event) {
	query := "DELETE FROM event WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, event.Id)
	helper.PanicIfError(err)
}

func (repository *EventRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Event {
	query := "SELECT id, heading, subheading, thumbnail FROM event"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		event := domain.Event{}
		err := rows.Scan(&event.Id, &event.Heading, &event.Subheading, &event.Thumbnail)
		helper.PanicIfError(err)
		events = append(events, event)
	}
	return events
}

func (repository *EventRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, eventId uint) (domain.Event, error) {
	query := "SELECT id, heading, subheading, thumbnail FROM event WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, eventId)
	helper.PanicIfError(err)
	defer rows.Close()

	event := domain.Event{}
	if rows.Next() {
		err := rows.Scan(&event.Id, &event.Heading, &event.Subheading, &event.Thumbnail)
		helper.PanicIfError(err)
		return event, nil
	} else {
		return event, errors.New("event not found")
	}
}
