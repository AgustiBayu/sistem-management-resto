package repository

import (
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
)

type TransaksiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Transaksi, map[int]domain.Pesanan, map[int]domain.User)
	FindById(ctx context.Context, tx *sql.Tx, transaksiId int) (domain.Transaksi, domain.Pesanan, domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi
	Delete(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi)
}
