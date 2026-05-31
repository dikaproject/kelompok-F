package data

import "fmt"

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
	maxID := 0
	// mencari ID terbesar dari seluruh data warga (nama_anggota)
	for i := 0; i < JumlahWarga; i++ {
		if DaftarWarga[i].ID > maxID {
			maxID = DaftarWarga[i].ID
		}
	}
	return maxID + 1
}

// menambahkan warga baru ke dalam array dengan ID otomatis (nama_anggota)
func TambahWarga(nama, alamat, noHP string) bool {
	// memeriksa apakah array masih memiliki kapasitas (nama_anggota)
	if JumlahWarga >= MAKS_WARGA {
		fmt.Println("Data warga sudah mencapai kapasitas maksimum!")
		return false
	}
	id := GenerateIDWarga()
	DaftarWarga[JumlahWarga] = Warga{
		ID:     id,
		Nama:   nama,
		Alamat: alamat,
		NoHP:   noHP,
	}
	JumlahWarga++
	return true
}

// mencari posisi index warga berdasarkan ID, mengembalikan -1 jika tidak ditemukan (nama_anggota)
func CariIndexWargaByID(id int) int {
	// melakukan pencarian linear dari awal hingga akhir array (nama_anggota)
	for i := 0; i < JumlahWarga; i++ {
		if DaftarWarga[i].ID == id {
			return i
		}
	}
	return -1
}

// mengubah data warga berdasarkan ID, field kosong tidak diubah (nama_anggota)
func UbahWarga(id int, nama, alamat, noHP string) bool {
	idx := CariIndexWargaByID(id)
	// memastikan warga dengan ID yang diberikan ditemukan (nama_anggota)
	if idx == -1 {
		return false
	}
	// hanya mengubah field yang tidak kosong (nama_anggota)
	if nama != "" {
		DaftarWarga[idx].Nama = nama
	}
	if alamat != "" {
		DaftarWarga[idx].Alamat = alamat
	}
	if noHP != "" {
		DaftarWarga[idx].NoHP = noHP
	}
	return true
}

// menghapus warga dari array dengan menggeser elemen ke kiri (nama_anggota)
func HapusWarga(id int) bool {
	idx := CariIndexWargaByID(id)
	if idx == -1 {
		return false
	}
	// menggeser semua elemen setelah index yang dihapus satu posisi ke kiri (nama_anggota)
	for i := idx; i < JumlahWarga-1; i++ {
		DaftarWarga[i] = DaftarWarga[i+1]
	}
	// mengosongkan elemen terakhir setelah pergeseran (nama_anggota)
	DaftarWarga[JumlahWarga-1] = Warga{}
	JumlahWarga--
	return true
}

// menghitung total berat sampah yang pernah disetor oleh warga tertentu (nama_anggota)
func GetTotalBeratWarga(wargaID int) float64 {
	total := 0.0
	// menjumlahkan berat seluruh setoran yang dimiliki warga tersebut (nama_anggota)
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].WargaID == wargaID {
			total += DaftarSetoran[i].BeratKg
		}
	}
	return total
}
