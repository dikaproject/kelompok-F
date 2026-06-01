package data

import "fmt"

type JenisSampah struct {
	ID        int
	NamaJenis string
	Satuan    string
}

var DaftarJenis [MAKS_JENIS]JenisSampah

var JumlahJenis int

func InitJenisSampah() {
	jenisDefault := [4][2]string{
		{"Plastik", "kg"},
		{"Kertas", "kg"},
		{"Logam", "kg"},
		{"Organik", "kg"},
	}
	for i := 0; i < 4; i++ {
		DaftarJenis[i] = JenisSampah{
			ID:        i + 1,
			NamaJenis: jenisDefault[i][0],
			Satuan:    jenisDefault[i][1],
		}
	}
	JumlahJenis = 4
}

func GetNamaJenis(id int) string {
	for i := 0; i < JumlahJenis; i++ {
		if DaftarJenis[i].ID == id {
			return DaftarJenis[i].NamaJenis
		}
	}
	return "Tidak Diketahui"
}

func TampilJenisSampah() {
	fmt.Println("Daftar Jenis Sampah:")
	for i := 0; i < JumlahJenis; i++ {
		fmt.Printf("  [%d] %-10s (%s)\n",
			DaftarJenis[i].ID,
			DaftarJenis[i].NamaJenis,
			DaftarJenis[i].Satuan,
		)
	}
}

func ValidasiIDJenis(id int) bool {
	for i := 0; i < JumlahJenis; i++ {
		if DaftarJenis[i].ID == id {
			return true
		}
	}
	return false
}
