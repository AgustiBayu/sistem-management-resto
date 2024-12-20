package web

type MenuItemResponse struct {
	Id                        int    `json:"id"`
	Name                      string `json:"name"`
	Deskripsi                 string `json:"deskripsi"`
	Harga                     int    `json:"harga"`
	TanggalPenambahanItemMenu string `json:"tanggal_penambahan_item_menu"`
}
