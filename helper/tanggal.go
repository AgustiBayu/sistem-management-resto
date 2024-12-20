package helper

import (
	"fmt"
	"time"
)

func FormatTanggal(t time.Time) string {
	return t.Format("02-01-2006")
}

func ValidateTanggalBaru(tanggalLama, tanggalBaru time.Time) error {
	if !tanggalBaru.After(tanggalLama) {
		return fmt.Errorf("tanggal baru (%s) tidak sesuai dengan tanggal sebelumnya (%s)",
			tanggalBaru.Format("02-01-2006"), tanggalLama.Format("02-01-2006"))
	}
	return nil
}
