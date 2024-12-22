package web

type DetailPesananResponse struct {
	Id                int              `json:"id"`
	Pesanan           PesananResponse  `json:"pesanan"`
	MenuItem          MenuItemResponse `json:"menu_item"`
	JumlahItemPesanan int              `json:"jumlah_item_pesanan"`
	HargaItem         int              `json:"harga_item"`
}
