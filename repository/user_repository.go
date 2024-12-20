package repository

import (
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
}
