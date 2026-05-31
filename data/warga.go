package data


// struct untuk menyimpan data warga pengguna bank sampah (nama_anggota)
type Warga struct {
	ID     int
	Nama   string
	Alamat string
	NoHP   string
}

// array global untuk menyimpan seluruh data warga (nama_anggota)
var DaftarWarga [MAKS_WARGA]Warga

// counter jumlah warga yang terdaftar saat ini (nama_anggota)
var JumlahWarga int

// menghasilkan ID baru secara otomatis berdasarkan ID terbesar yang sudah ada (nama_anggota)
func GenerateIDWarga() int {
	// TODO: cari ID terbesar dari DaftarWarga, kembalikan maxID + 1 (nama_anggota)
	return 0
}

// menambahkan warga baru ke dalam array dengan ID otomatis (nama_anggota)
func TambahWarga(nama, alamat, noHP string) bool {
	// TODO: cek kapasitas, generate ID, isi DaftarWarga[JumlahWarga], increment JumlahWarga (nama_anggota)
	return false
}

// mencari posisi index warga berdasarkan ID, mengembalikan -1 jika tidak ditemukan (nama_anggota)
func CariIndexWargaByID(id int) int {
	// TODO: linear search ID di DaftarWarga, return index atau -1 (nama_anggota)
	return -1
}

// mengubah data warga berdasarkan ID, field kosong tidak diubah (nama_anggota)
func UbahWarga(id int, nama, alamat, noHP string) bool {
	// TODO: cari index warga, update field yang tidak kosong saja (nama_anggota)
	return false
}

// menghapus warga dari array dengan menggeser elemen ke kiri (nama_anggota)
func HapusWarga(id int) bool {
	// TODO: cari index, geser elemen ke kiri, kosongkan elemen terakhir, kurangi counter (nama_anggota)
	return false
}

// menghitung total berat sampah yang pernah disetor oleh warga tertentu (nama_anggota)
func GetTotalBeratWarga(wargaID int) float64 {
	// TODO: jumlahkan BeratKg semua setoran yang WargaID-nya cocok (nama_anggota)
	return 0
}
