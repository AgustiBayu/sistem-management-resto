package web

type PesananResponse struct {
	Id                      int          `json:"id"`
	User                    UserResponse `json:"user"`
	TotalHarga              int          `json:"total_harga"`
	Status                  string       `json:"status"`
	TanggalPembuatanPesanan string       `json:"tanggal_pembuatan_pesanan"`
}
