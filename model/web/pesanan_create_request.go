package web

type PesananCreateRequest struct {
	UserId                  int    `json:"user_id" validate:"required"`
	TotalHarga              int    `json:"total_harga" validate:"required"`
	Status                  string `json:"status" validate:"required"`
	TanggalPembuatanPesanan string `json:"tanggal_pembuatan_pesanan" validate:"required"`
}
