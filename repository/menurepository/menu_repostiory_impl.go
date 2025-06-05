package menurepository

import (
	"backend/helper"
	"backend/model/domain"
	"context"
	"database/sql"
	"errors"
)

type MenuRepositoryImpl struct {
}

func NewMenuRepository() MenuRepository {
	return &MenuRepositoryImpl{}
}

func (repository *MenuRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, menu domain.Menu) domain.Menu {
	query := "INSERT INTO menu (heading, subheading, thumbnail) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, menu.Heading, menu.Subheading, menu.Thumbnail)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	menu.Id = uint(id)
	return menu
}

func (repository *MenuRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, menu domain.Menu) domain.Menu {
	query := "UPDATE menu SET heading = ?, subheading = ?, thumbnail = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, menu.Heading, menu.Subheading, menu.Thumbnail, menu.Id)
	helper.PanicIfError(err)

	return menu
}

func (repository *MenuRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, menu domain.Menu) {
	query := "DELETE FROM menu WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, menu.Id)
	helper.PanicIfError(err)
}

func (repository *MenuRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Menu {
	query := "SELECT id, heading, subheading, thumbnail FROM menu"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var menus []domain.Menu
	for rows.Next() {
		menu := domain.Menu{}
		err := rows.Scan(&menu.Id, &menu.Heading, &menu.Subheading, &menu.Thumbnail)
		helper.PanicIfError(err)
		menus = append(menus, menu)
	}
	return menus
}

func (repository *MenuRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, menuId uint) (domain.Menu, error) {
	query := "SELECT id, heading, subheading, thumbnail FROM menu WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, menuId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Menu{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Heading, &category.Subheading, &category.Thumbnail)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("menu not found")
	}
}
