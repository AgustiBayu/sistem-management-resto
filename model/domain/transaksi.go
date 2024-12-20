package domain

import "time"

type StatusMetodePembayaran string

const (
	Cash    StatusMetodePembayaran = "CASH"
	Card    StatusMetodePembayaran = "CARD"
	Ewalled StatusMetodePembayaran = "EWALLET"
)

type Transaksi struct {
	Id               int
	PesananId        int
	MetodePembayaran StatusMetodePembayaran
	JumlahBayar      int
	Kembalian        int
	TanggalTransaksi time.Time
}
