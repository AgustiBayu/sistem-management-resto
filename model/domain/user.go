package domain

import "time"

type StatusPengguna string

const (
	Pengguna StatusPengguna = "PENGGUNA"
	Admin    StatusPengguna = "ADMIN"
)

type User struct {
	Id              int
	Name            string
	Email           string
	Password        string
	Pengguna        StatusPengguna
	TanggalBuatAkun time.Time
}
