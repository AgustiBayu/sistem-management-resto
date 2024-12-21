package repository

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
	"errors"
)

type MenuItemRepositoryImpl struct{}

func NewMenuItemRepository() MenuItemRepository {
	return &MenuItemRepositoryImpl{}
}

func (m *MenuItemRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) domain.MenuItem {
	SQL := "INSERT INTO menu_items(name, deskripsi, harga ,tanggal_penambahan_item_menu) VALUES($1,$2,$3,$4) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, menuItem.Name, menuItem.Deskripsi, menuItem.Harga, menuItem.TanggalPenambahanItemMenu).Scan(&id)
	helper.PanicIfError(err)
	menuItem.Id = id
	return menuItem
}
func (m *MenuItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MenuItem {
	SQL := "SELECT id, name, deskripsi,harga,tanggal_penambahan_item_menu FROM menu_items"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	var menuItems []domain.MenuItem
	for rows.Next() {
		menuItem := domain.MenuItem{}
		err := rows.Scan(&menuItem.Id, &menuItem.Name, &menuItem.Deskripsi, &menuItem.Harga, &menuItem.TanggalPenambahanItemMenu)
		helper.PanicIfError(err)
		menuItems = append(menuItems, menuItem)
	}
	return menuItems
}
func (m *MenuItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, menuItemId int) (domain.MenuItem, error) {
	SQL := "SELECT id, name, deskripsi,harga,tanggal_penambahan_item_menu FROM menu_items WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, menuItemId)
	helper.PanicIfError(err)
	defer rows.Close()
	menuItem := domain.MenuItem{}
	if rows.Next() {
		err := rows.Scan(&menuItem.Id, &menuItem.Name, &menuItem.Deskripsi, &menuItem.Harga, &menuItem.TanggalPenambahanItemMenu)
		helper.PanicIfError(err)
		return menuItem, nil
	} else {
		return menuItem, errors.New("menu item id not found")
	}
}
func (m *MenuItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) domain.MenuItem {
	SQL := "UPDATE menu_items SET name = $1, deskripsi = $2,harga = $3,tanggal_penambahan_item_menu = $4 WHERE id = $5"
	_, err := tx.QueryContext(ctx, SQL, menuItem.Name, menuItem.Deskripsi, menuItem.Harga, menuItem.TanggalPenambahanItemMenu, menuItem.Id)
	helper.PanicIfError(err)
	return menuItem
}
func (m *MenuItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) {
	SQL := "DELETE FROM menu_items WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, menuItem.Id)
	helper.PanicIfError(err)
}
