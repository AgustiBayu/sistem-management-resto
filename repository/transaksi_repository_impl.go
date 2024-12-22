package repository

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
	"errors"
)

type TransaksiRepositoryImpl struct{}

func NewTransaksiRepository() TransaksiRepository {
	return &TransaksiRepositoryImpl{}
}

func (t *TransaksiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi {
	SQL := "INSERT INTO transaksis(pesanan_id, metode_pembayaran,jumlah_bayar,kembalian, tanggal_transaksi) VALUES($1,$2,$3,$4,$5) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, transaksi.PesananId, transaksi.MetodePembayaran, transaksi.JumlahBayar, transaksi.Kembalian, transaksi.TanggalTransaksi).Scan(&id)
	helper.PanicIfError(err)
	transaksi.Id = id
	return transaksi
}
func (t *TransaksiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Transaksi, map[int]domain.Pesanan, map[int]domain.User) {
	SQL := "select t.id , t.pesanan_id , p.id , p.user_id , u.id ,u.name, u.email , u.pengguna , u.tanggal_buat_akun ,p.total_harga ,p.status ,p.tanggal_pembuatan_pesanan ,t.metode_pembayaran ,t.jumlah_bayar ,t.kembalian ,t.tanggal_transaksi from transaksis t inner join pesanans p on t.pesanan_id = p.id inner join users u on p.user_id = u.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var transaksis []domain.Transaksi
	pesananMap := make(map[int]domain.Pesanan)
	userMap := make(map[int]domain.User)
	for rows.Next() {
		transaksi := domain.Transaksi{}
		pesanan := domain.Pesanan{}
		user := domain.User{}
		err := rows.Scan(&transaksi.Id, &transaksi.PesananId, &pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga, &pesanan.Status, &pesanan.TanggalPembuatanPesanan, &transaksi.MetodePembayaran, &transaksi.JumlahBayar, &transaksi.Kembalian, &transaksi.TanggalTransaksi)
		helper.PanicIfError(err)
		transaksis = append(transaksis, transaksi)
		pesananMap[pesanan.Id] = pesanan
		userMap[user.Id] = user
	}
	return transaksis, pesananMap, userMap
}
func (t *TransaksiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transaksiId int) (domain.Transaksi, domain.Pesanan, domain.User, error) {
	SQL := "select t.id , t.pesanan_id , p.id , p.user_id , u.id ,u.name, u.email , u.pengguna , u.tanggal_buat_akun ,p.total_harga ,p.status ,p.tanggal_pembuatan_pesanan ,t.metode_pembayaran ,t.jumlah_bayar ,t.kembalian ,t.tanggal_transaksi from transaksis t inner join pesanans p on t.pesanan_id = p.id inner join users u on p.user_id = u.id WHERE t.id = $1"
	rows, err := tx.QueryContext(ctx, SQL, transaksiId)
	helper.PanicIfError(err)
	defer rows.Close()

	transaksi := domain.Transaksi{}
	pesanan := domain.Pesanan{}
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&transaksi.Id, &transaksi.PesananId, &pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga, &pesanan.Status, &pesanan.TanggalPembuatanPesanan, &transaksi.MetodePembayaran, &transaksi.JumlahBayar, &transaksi.Kembalian, &transaksi.TanggalTransaksi)
		helper.PanicIfError(err)
		return transaksi, pesanan, user, nil
	} else {
		return transaksi, pesanan, user, errors.New("transaksi id not found")
	}
}
func (t *TransaksiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi {
	SQL := "UPDATE transaksis SET pesanan_id = $1, metode_pembayaran = $2,jumlah_bayar = $3,kembalian = $4, tanggal_transaksi = $5 WHERE id = $6"
	_, err := tx.ExecContext(ctx, SQL, transaksi.PesananId, transaksi.MetodePembayaran, transaksi.JumlahBayar, transaksi.Kembalian, transaksi.TanggalTransaksi, transaksi.Id)
	helper.PanicIfError(err)
	return transaksi
}
func (t *TransaksiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) {
	SQL := "DELETE FROM transaksis WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, transaksi.Id)
	helper.PanicIfError(err)
}
