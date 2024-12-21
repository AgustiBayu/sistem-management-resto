package web

type PesananUpdateRequest struct {
	Id                      int    `json:"id" validate:"required"`
	UserId                  int    `json:"user" validate:"required"`
	TotalHarga              int    `json:"total_harga" validate:"required"`
	Status                  string `json:"status" validate:"required"`
	TanggalPembuatanPesanan string `json:"tanggal_pembuatan_pesanan" validate:"required"`
}
