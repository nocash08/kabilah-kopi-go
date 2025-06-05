package aboutrepository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"
	"errors"
)

type AboutRepositoryImpl struct {
}

func NewAboutRepository() AboutRepository {
	return &AboutRepositoryImpl{}
}

func (repository *AboutRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, about domain.About) domain.About {
	query := "INSERT INTO about (heading, subheading) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, query, about.Heading, about.Subheading)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	about.Id = uint(id)
	return about
}

func (repository *AboutRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, about domain.About) domain.About {
	query := "UPDATE about SET heading = ?, subheading = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, about.Heading, about.Subheading, about.Id)
	helper.PanicIfError(err)

	return about
}

func (repository *AboutRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.About {
	query := "SELECT id, heading, subheading FROM about"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var abouts []domain.About
	for rows.Next() {
		about := domain.About{}
		err := rows.Scan(&about.Id, &about.Heading, &about.Subheading)
		helper.PanicIfError(err)
		abouts = append(abouts, about)
	}
	return abouts
}

func (repository *AboutRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, aboutId uint) (domain.About, error) {
	query := "SELECT id, heading, subheading FROM about WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, aboutId)
	helper.PanicIfError(err)
	defer rows.Close()

	about := domain.About{}
	if rows.Next() {
		err := rows.Scan(&about.Id, &about.Heading, &about.Subheading)
		helper.PanicIfError(err)
		return about, nil
	} else {
		return about, errors.New("about not found")
	}
}
