package data

import "fmt"

// struct untuk menyimpan kategori jenis sampah (nama_anggota)
type JenisSampah struct {
	ID        int
	NamaJenis string
	Satuan    string
}

// array global untuk menyimpan seluruh jenis sampah (nama_anggota)
var DaftarJenis [MAKS_JENIS]JenisSampah

// counter jumlah jenis sampah yang terdaftar (nama_anggota)
var JumlahJenis int

// menginisialisasi jenis sampah default saat aplikasi pertama dijalankan (nama_anggota)
func InitJenisSampah() {
	jenisDefault := [4][2]string{
		{"Plastik", "kg"},
		{"Kertas", "kg"},
		{"Logam", "kg"},
		{"Organik", "kg"},
	}
	// mengisi array dengan data jenis sampah default (nama_anggota)
	for i := 0; i < 4; i++ {
		DaftarJenis[i] = JenisSampah{
			ID:        i + 1,
			NamaJenis: jenisDefault[i][0],
			Satuan:    jenisDefault[i][1],
		}
	}
	JumlahJenis = 4
}

// mencari nama jenis sampah berdasarkan ID yang diberikan (nama_anggota)
func GetNamaJenis(id int) string {
	// iterasi seluruh daftar jenis untuk mencocokkan ID (nama_anggota)
	for i := 0; i < JumlahJenis; i++ {
		if DaftarJenis[i].ID == id {
			return DaftarJenis[i].NamaJenis
		}
	}
	return "Tidak Diketahui"
}

// menampilkan seluruh daftar jenis sampah yang tersedia (nama_anggota)
func TampilJenisSampah() {
	fmt.Println("Daftar Jenis Sampah:")
	// mencetak setiap jenis sampah beserta ID dan satuannya (nama_anggota)
	for i := 0; i < JumlahJenis; i++ {
		fmt.Printf("  [%d] %-10s (%s)\n",
			DaftarJenis[i].ID,
			DaftarJenis[i].NamaJenis,
			DaftarJenis[i].Satuan,
		)
	}
}

// memvalidasi apakah ID jenis sampah terdaftar atau tidak (nama_anggota)
func ValidasiIDJenis(id int) bool {
	// melakukan pengecekan satu per satu terhadap seluruh jenis (nama_anggota)
	for i := 0; i < JumlahJenis; i++ {
		if DaftarJenis[i].ID == id {
			return true
		}
	}
	return false
}
