package web

type UserResponse struct {
	Id              int    `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Pengguna        string `json:"pengguna" validate:"required,oneof=ADMIN PENGGUNA"`
	TanggalBuatAkun string `json:"tanggal_buat_akun" validate:"required"`
}
