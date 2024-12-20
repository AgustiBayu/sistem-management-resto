package domain

import "time"

type StatusPesanan string

const (
	Panding   StatusPesanan = "PANDING"
	Paid      StatusPesanan = "PAID"
	Cancelled StatusPesanan = "CANCCELLED"
)

type Pesanan struct {
	Id                      int
	UserId                  int
	TotalHarga              int
	Status                  StatusPesanan
	TanggalPembuatanPesanan time.Time
}
