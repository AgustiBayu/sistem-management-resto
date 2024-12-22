package repository

import (
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
)

type DetailPesananRespository interface {
	Save(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.DetailPesanan, map[int]domain.Pesanan, map[int]domain.User, map[int]domain.MenuItem)
	FindById(ctx context.Context, tx *sql.Tx, detailPesananId int) (domain.DetailPesanan, domain.Pesanan, domain.User, domain.MenuItem, error)
	Update(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan
	Delete(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan)
}
