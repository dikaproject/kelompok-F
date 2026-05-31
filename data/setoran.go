package data


// struct untuk menyimpan log setoran sampah setiap transaksi (nama_anggota)
type LogSetoran struct {
	ID      int
	WargaID int
	JenisID int
	BeratKg float64
	Tanggal string
	Minggu  int
}

// array global untuk menyimpan seluruh log setoran sampah (nama_anggota)
var DaftarSetoran [MAKS_SETORAN]LogSetoran

// counter jumlah setoran yang telah dicatat (nama_anggota)
var JumlahSetoran int

// menghasilkan ID setoran otomatis berdasarkan ID terbesar yang sudah ada (nama_anggota)
func GenerateIDSetoran() int {
	// TODO: cari ID setoran terbesar, kembalikan maxID + 1 (nama_anggota)
	return 0
}

// menghitung nomor minggu ISO dari string tanggal format YYYY-MM-DD (nama_anggota)
func HitungNomorMinggu(tanggal string) int {
	// TODO: parse tanggal "2006-01-02" lalu ambil ISOWeek() (nama_anggota)
	// hint: import "time", gunakan time.Parse lalu t.ISOWeek()
	return 0
}

// menghitung rentang tanggal dari nomor minggu ISO dan tahun yang diberikan (nama_anggota)
func MingguKeRentang(minggu, tahun int) string {
	// TODO: konversi nomor minggu ISO ke rentang tanggal format "DD Mon - DD Mon YYYY" (nama_anggota)
	// hint: Jan 4 selalu di minggu ISO ke-1, mundur ke Senin, maju (minggu-1)*7 hari
	return ""
}

// mencatat setoran sampah baru ke dalam array setelah validasi (nama_anggota)
func CatatSetoran(wargaID, jenisID int, beratKg float64, tanggal string) bool {
	// TODO: validasi kapasitas + wargaID + jenisID, lalu simpan LogSetoran baru (nama_anggota)
	return false
}

// menghapus setoran berdasarkan ID dengan menggeser elemen ke kiri (nama_anggota)
func HapusSetoran(id int) bool {
	// TODO: cari index setoran, geser elemen ke kiri, kosongkan terakhir (nama_anggota)
	return false
}

// mengambil salinan seluruh setoran milik warga tertentu beserta jumlahnya (nama_anggota)
func GetSetoranByWarga(wargaID int) ([MAKS_SETORAN]LogSetoran, int) {
	var hasil [MAKS_SETORAN]LogSetoran
	// TODO: filter DaftarSetoran yang WargaID-nya cocok (nama_anggota)
	return hasil, 0
}

// mengambil salinan seluruh setoran pada minggu tertentu beserta jumlahnya (nama_anggota)
func GetSetoranByMinggu(minggu int) ([MAKS_SETORAN]LogSetoran, int) {
	var hasil [MAKS_SETORAN]LogSetoran
	// TODO: filter DaftarSetoran yang nomor Minggu-nya cocok (nama_anggota)
	return hasil, 0
}
