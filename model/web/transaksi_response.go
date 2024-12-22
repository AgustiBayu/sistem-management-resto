package web

type TransaksiResponse struct {
	Id               int             `json:"id" `
	Pesanan          PesananResponse `json:"pesanan" `
	MetodePembayaran string          `json:"metode_pembayaran" `
	JumlahBayar      int             `json:"jumlah_bayar" `
	Kembalian        int             `json:"kembalian" `
	TanggalTransaksi string          `json:"tanggal_transaksi" `
}
