package web

type MenuItemCreateRequest struct {
	Name                      string `json:"name" validate:"required"`
	Deskripsi                 string `json:"deskripsi" validate:"required"`
	Harga                     int    `json:"harga" validate:"required"`
	TanggalPenambahanItemMenu string `json:"tanggal_penambahan_item_menu" validate:"required"`
}
