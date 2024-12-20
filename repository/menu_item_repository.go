package repository

import (
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
)

type MenuItemRepository interface {
	Save(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) domain.MenuItem
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MenuItem
	FindById(ctx context.Context, tx *sql.Tx, menuItemId int) (domain.MenuItem, error)
	Update(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem) domain.MenuItem
	Delete(ctx context.Context, tx *sql.Tx, menuItem domain.MenuItem)
}
