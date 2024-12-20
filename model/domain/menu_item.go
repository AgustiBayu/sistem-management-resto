package domain

import "time"

type MenuItem struct {
	Id                        int
	Name                      string
	Deskripsi                 string
	Harga                     int
	TanggalPenambahanItemMenu time.Time
}
