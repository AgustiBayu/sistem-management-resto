package helper

import (
	"SistemManagementResto/model/domain"
	"SistemManagementResto/model/web"
)

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
