package helper

import (
	"SistemManagementResto/model/domain"
	"SistemManagementResto/model/web"
)

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
