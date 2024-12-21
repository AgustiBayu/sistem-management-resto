package repository

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
	"errors"
)

type PesananRepositoryImpl struct{}

func NewPesananRepository() PesananRepository {
	return &PesananRepositoryImpl{}
}

func (p *PesananRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan {
	SQL := "INSERT INTO pesanans(user_id,total_harga,status,tanggal_pembuatan_pesanan) VALUES($1,$2,$3,$4) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, pesanan.UserId, pesanan.TotalHarga, pesanan.Status, pesanan.TanggalPembuatanPesanan).Scan(&id)
	helper.PanicIfError(err)
	pesanan.Id = id
	return pesanan
}
func (p *PesananRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pesanan, map[int]domain.User) {
	SQL := "select p.id, p.user_id, u.id ,u.name, u.email, u.pengguna ,u.tanggal_buat_akun,p.total_harga ,p.status ,p.tanggal_pembuatan_pesanan from pesanans p inner join users u on p.user_id = u.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var pesanans []domain.Pesanan
	userMap := make(map[int]domain.User)
	for rows.Next() {
		pesanan := domain.Pesanan{}
		user := domain.User{}
		err := rows.Scan(&pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga, &pesanan.Status, &pesanan.TanggalPembuatanPesanan)
		helper.PanicIfError(err)
		pesanans = append(pesanans, pesanan)
		userMap[user.Id] = user
	}
	return pesanans, userMap
}
func (p *PesananRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, pesananId int) (domain.Pesanan, domain.User, error) {
	SQL := "select p.id, p.user_id, u.id ,u.name, u.email, u.pengguna ,u.tanggal_buat_akun,p.total_harga ,p.status ,p.tanggal_pembuatan_pesanan from pesanans p inner join users u on p.user_id = u.id WHERE p.id = $1"
	rows, err := tx.QueryContext(ctx, SQL, pesananId)
	helper.PanicIfError(err)
	defer rows.Close()

	pesanan := domain.Pesanan{}
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga, &pesanan.Status, &pesanan.TanggalPembuatanPesanan)
		helper.PanicIfError(err)
		return pesanan, user, nil
	} else {
		return pesanan, user, errors.New("pesanan id not found")
	}
}
func (p *PesananRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan {
	SQL := "UPDATE pesanans SET user_id= $1,total_harga= $2,status = $3,tanggal_pembuatan_pesanan= $4 WHERE id = $5"
	_, err := tx.ExecContext(ctx, SQL, pesanan.UserId, pesanan.TotalHarga, pesanan.Status, pesanan.TanggalPembuatanPesanan, pesanan.Id)
	helper.PanicIfError(err)
	return pesanan
}
func (p *PesananRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) {
	SQL := "DELETE FROM pesanans WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, pesanan.Id)
	helper.PanicIfError(err)
}
