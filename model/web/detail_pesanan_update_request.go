package web

type DetailPesananUpdateRequest struct {
	Id                int `json:"id" validate:"required"`
	PesananId         int `json:"pesanan_id" validate:"required"`
	MenuItemId        int `json:"menu_item_id" validate:"required"`
	JumlahItemPesanan int `json:"jumlah_item_pesanan" validate:"required"`
	HargaItem         int `json:"harga_item" validate:"required"`
}
