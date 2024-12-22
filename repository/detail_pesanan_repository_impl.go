package repository

import (
	"SistemManagementResto/helper"
	"SistemManagementResto/model/domain"
	"context"
	"database/sql"
	"errors"
)

type DetailPesananRespositoryImpl struct{}

func NewDetailPesananRepository() DetailPesananRespository {
	return &DetailPesananRespositoryImpl{}
}

func (d *DetailPesananRespositoryImpl) Save(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan {
	SQL := "INSERT INTO detail_pesanans(pesanan_id, menu_item_id, jumlah_item_pesanan,harga_item) VALUES($1,$2,$3,$4) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, detailPesanan.PesananId, detailPesanan.MenuItemId, detailPesanan.JumlahItemPesanan, detailPesanan.HargaItem).Scan(&id)
	helper.PanicIfError(err)
	detailPesanan.Id = id
	return detailPesanan
}
func (d *DetailPesananRespositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.DetailPesanan, map[int]domain.Pesanan, map[int]domain.User, map[int]domain.MenuItem) {
	SQL := "select dp.id , dp.pesanan_id ,dp.menu_item_id ,p.id ,p.user_id , u.id,u.name, u.email,u.pengguna,u.tanggal_buat_akun ,p.total_harga ,p.status ,p.tanggal_pembuatan_pesanan ,mi.id ,mi.name,mi.deskripsi ,mi.harga ,mi.tanggal_penambahan_item_menu ,dp.jumlah_item_pesanan ,dp.harga_item from detail_pesanans dp inner join pesanans p on dp.pesanan_id = p.id inner join menu_items mi on dp.menu_item_id = mi.id inner join users u on p.user_id = u.id "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var detailPesanans []domain.DetailPesanan
	pesananMap := make(map[int]domain.Pesanan)
	userMap := make(map[int]domain.User)
	menuItemMap := make(map[int]domain.MenuItem)
	for rows.Next() {
		detailPesanan := domain.DetailPesanan{}
		pesanan := domain.Pesanan{}
		user := domain.User{}
		menuItem := domain.MenuItem{}
		err := rows.Scan(&detailPesanan.Id, &detailPesanan.PesananId, &detailPesanan.MenuItemId, &pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga, &pesanan.Status, &pesanan.TanggalPembuatanPesanan, &menuItem.Id, &menuItem.Name, &menuItem.Deskripsi, &menuItem.Harga, &menuItem.TanggalPenambahanItemMenu, &detailPesanan.JumlahItemPesanan, &detailPesanan.HargaItem)
		helper.PanicIfError(err)
		detailPesanans = append(detailPesanans, detailPesanan)
		pesananMap[pesanan.Id] = pesanan
		userMap[user.Id] = user
		menuItemMap[menuItem.Id] = menuItem
	}
	return detailPesanans, pesananMap, userMap, menuItemMap
}
func (d *DetailPesananRespositoryImpl) FindById(ctx context.Context, tx *sql.Tx, detailPesananId int) (domain.DetailPesanan, domain.Pesanan, domain.User, domain.MenuItem, error) {
	SQL := "select dp.id , dp.pesanan_id ,dp.menu_item_id ,p.id ,p.user_id , u.id,u.name, u.email,u.pengguna,u.tanggal_buat_akun ,p.total_harga ,p.status ,p.tanggal_pembuatan_pesanan ,mi.id ,mi.name,mi.deskripsi ,mi.harga ,mi.tanggal_penambahan_item_menu ,dp.jumlah_item_pesanan ,dp.harga_item from detail_pesanans dp inner join pesanans p on dp.pesanan_id = p.id inner join menu_items mi on dp.menu_item_id = mi.id inner join users u on p.user_id = u.id WHERE dp.id = $1"
	rows, err := tx.QueryContext(ctx, SQL, detailPesananId)
	helper.PanicIfError(err)
	defer rows.Close()

	detailPesanan := domain.DetailPesanan{}
	pesanan := domain.Pesanan{}
	user := domain.User{}
	menuItem := domain.MenuItem{}
	if rows.Next() {
		err := rows.Scan(&detailPesanan.Id, &detailPesanan.PesananId, &detailPesanan.MenuItemId, &pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga, &pesanan.Status, &pesanan.TanggalPembuatanPesanan, &menuItem.Id, &menuItem.Name, &menuItem.Deskripsi, &menuItem.Harga, &menuItem.TanggalPenambahanItemMenu, &detailPesanan.JumlahItemPesanan, &detailPesanan.HargaItem)
		helper.PanicIfError(err)
		return detailPesanan, pesanan, user, menuItem, nil
	} else {
		return detailPesanan, pesanan, user, menuItem, errors.New("detail pesanan id not found")
	}
}
func (d *DetailPesananRespositoryImpl) Update(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan {
	SQL := "UPDATE detail_pesanans SET pesanan_id = $1, menu_item_id = $2, jumlah_item_pesanan = $3,harga_item = $4 WHERE id = $5"
	_, err := tx.ExecContext(ctx, SQL, detailPesanan.PesananId, detailPesanan.MenuItemId, detailPesanan.JumlahItemPesanan, detailPesanan.HargaItem, detailPesanan.Id)
	helper.PanicIfError(err)
	return detailPesanan
}
func (d *DetailPesananRespositoryImpl) Delete(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) {
	SQL := "DELETE FROM detail_pesanans WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, detailPesanan.Id)
	helper.PanicIfError(err)
}
