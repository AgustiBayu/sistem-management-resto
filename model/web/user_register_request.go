package web

type UserRegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	Pengguna        string `json:"pengguna" validate:"required,oneof=ADMIN PENGGUNA"`
	TanggalBuatAkun string `json:"tanggal_buat_akun" validate:"required"`
}
