package helper

import (
	"SistemManagementResto/model/domain"
	"SistemManagementResto/model/web"
)

func ToTransaksiResponses(transaksis []domain.Transaksi, pesananMap map[int]domain.Pesanan, userMap map[int]domain.User) []web.TransaksiResponse {
	var transaksiResponse []web.TransaksiResponse
	for _, transaksi := range transaksis {
		pesanan, exits := pesananMap[transaksi.PesananId]
		if !exits {
			pesanan = domain.Pesanan{}
		}
		user, exits := userMap[pesanan.UserId]
		if !exits {
			user = domain.User{}
		}
		transaksiResponse = append(transaksiResponse, ToTransaksiResponse(transaksi, pesanan, user))
	}
	return transaksiResponse
}

func ToTransaksiResponse(transaksi domain.Transaksi, pesanan domain.Pesanan, user domain.User) web.TransaksiResponse {
	return web.TransaksiResponse{
		Id: transaksi.Id,
		Pesanan: web.PesananResponse{
			Id: pesanan.Id,
			User: web.UserResponse{
				Id:              user.Id,
				Name:            user.Name,
				Email:           user.Email,
				Pengguna:        string(user.Pengguna),
				TanggalBuatAkun: FormatTanggal(user.TanggalBuatAkun),
			},
			TotalHarga:              pesanan.TotalHarga,
			Status:                  string(pesanan.Status),
			TanggalPembuatanPesanan: FormatTanggal(pesanan.TanggalPembuatanPesanan),
		},
		MetodePembayaran: string(transaksi.MetodePembayaran),
		JumlahBayar:      transaksi.JumlahBayar,
		Kembalian:        transaksi.Kembalian,
		TanggalTransaksi: FormatTanggal(transaksi.TanggalTransaksi),
	}
}

func ToDetailPesananResponses(detailPesanans []domain.DetailPesanan, pesananMap map[int]domain.Pesanan, userMap map[int]domain.User, menuIteMap map[int]domain.MenuItem) []web.DetailPesananResponse {
	var detailPesananResponse []web.DetailPesananResponse
	for _, detailPesanan := range detailPesanans {
		pesanan, exits := pesananMap[detailPesanan.PesananId]
		if !exits {
			pesanan = domain.Pesanan{}
		}
		user, exits := userMap[pesanan.UserId]
		if !exits {
			user = domain.User{}
		}
		menuItem, exits := menuIteMap[detailPesanan.MenuItemId]
		if !exits {
			menuItem = domain.MenuItem{}
		}
		detailPesananResponse = append(detailPesananResponse, ToDetailPesananResponse(detailPesanan, pesanan, user, menuItem))
	}
	return detailPesananResponse
}

func ToDetailPesananResponse(detailPesanan domain.DetailPesanan, pesanan domain.Pesanan, user domain.User, menuItem domain.MenuItem) web.DetailPesananResponse {
	return web.DetailPesananResponse{
		Id: detailPesanan.Id,
		Pesanan: web.PesananResponse{
			Id: pesanan.Id,
			User: web.UserResponse{
				Id:              user.Id,
				Name:            user.Name,
				Email:           user.Email,
				Pengguna:        string(user.Pengguna),
				TanggalBuatAkun: FormatTanggal(user.TanggalBuatAkun),
			},
			TotalHarga:              pesanan.TotalHarga,
			Status:                  string(pesanan.Status),
			TanggalPembuatanPesanan: FormatTanggal(pesanan.TanggalPembuatanPesanan),
		},
		MenuItem: web.MenuItemResponse{
			Id:                        menuItem.Id,
			Name:                      menuItem.Name,
			Deskripsi:                 menuItem.Deskripsi,
			Harga:                     menuItem.Harga,
			TanggalPenambahanItemMenu: FormatTanggal(menuItem.TanggalPenambahanItemMenu),
		},
		JumlahItemPesanan: detailPesanan.JumlahItemPesanan,
		HargaItem:         detailPesanan.HargaItem,
	}
}

func ToPesananResponses(pesanans []domain.Pesanan, userMap map[int]domain.User) []web.PesananResponse {
	var pesananResponse []web.PesananResponse
	for _, pesanan := range pesanans {
		user, exits := userMap[pesanan.UserId]
		if !exits {
			user = domain.User{}
		}
		pesananResponse = append(pesananResponse, ToPesananResponse(pesanan, user))
	}
	return pesananResponse
}

func ToPesananResponse(pesanan domain.Pesanan, user domain.User) web.PesananResponse {
	return web.PesananResponse{
		Id: pesanan.Id,
		User: web.UserResponse{
			Id:              user.Id,
			Name:            user.Name,
			Email:           user.Email,
			Pengguna:        string(user.Pengguna),
			TanggalBuatAkun: FormatTanggal(user.TanggalBuatAkun),
		},
		TotalHarga:              pesanan.TotalHarga,
		Status:                  string(pesanan.Status),
		TanggalPembuatanPesanan: FormatTanggal(pesanan.TanggalPembuatanPesanan),
	}
}

func ToMenuItemResponses(menuItems []domain.MenuItem) []web.MenuItemResponse {
	var menuItemsResponse []web.MenuItemResponse
	for _, menuItem := range menuItems {
		menuItemsResponse = append(menuItemsResponse, ToMenuItemResponse(menuItem))
	}
	return menuItemsResponse
}

func ToMenuItemResponse(menuItem domain.MenuItem) web.MenuItemResponse {
	return web.MenuItemResponse{
		Id:                        menuItem.Id,
		Name:                      menuItem.Name,
		Deskripsi:                 menuItem.Deskripsi,
		Harga:                     menuItem.Harga,
		TanggalPenambahanItemMenu: FormatTanggal(menuItem.TanggalPenambahanItemMenu),
	}
}
func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:              user.Id,
		Name:            user.Name,
		Email:           user.Email,
		Pengguna:        string(user.Pengguna),
		TanggalBuatAkun: FormatTanggal(user.TanggalBuatAkun),
	}
}
