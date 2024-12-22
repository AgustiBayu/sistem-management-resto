package web

type TransaksiCreateRequest struct {
	PesananId        int    `json:"pesanan_id" validate:"required"`
	MetodePembayaran string `json:"metode_pembayaran" validate:"required"`
	JumlahBayar      int    `json:"jumlah_bayar" validate:"required"`
	TanggalTransaksi string `json:"tanggal_transaksi" validate:"required"`
}
