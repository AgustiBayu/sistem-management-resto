package web

type MenuItemUpdateRequest struct {
	Id                        int    `json:"id" validate:"required"`
	Name                      string `json:"name" validate:"required"`
	Deskripsi                 string `json:"deskripsi" validate:"required"`
	Harga                     int    `json:"harga" validate:"required"`
	TanggalPenambahanItemMenu string `json:"tanggal_penambahan_item_menu" validate:"required"`
}
