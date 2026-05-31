package data


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
	// TODO: isi DaftarJenis dengan 4 jenis default: Plastik, Kertas, Logam, Organik (semua satuan "kg") (nama_anggota)
	// set JumlahJenis = 4
}

// mencari nama jenis sampah berdasarkan ID yang diberikan (nama_anggota)
func GetNamaJenis(id int) string {
	// TODO: cari NamaJenis berdasarkan ID, return "Tidak Diketahui" jika tidak ada (nama_anggota)
	return ""
}

// menampilkan seluruh daftar jenis sampah yang tersedia (nama_anggota)
func TampilJenisSampah() {
	// TODO: cetak seluruh DaftarJenis dalam format "  [ID] NamaJenis (Satuan)" (nama_anggota)
}

// memvalidasi apakah ID jenis sampah terdaftar atau tidak (nama_anggota)
func ValidasiIDJenis(id int) bool {
	// TODO: return true jika ID ditemukan di DaftarJenis (nama_anggota)
	return false
}
