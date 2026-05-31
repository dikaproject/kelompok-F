package data

import (
	"fmt"
	"time"
)

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
	maxID := 0
	// mencari ID setoran terbesar dari seluruh log (nama_anggota)
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].ID > maxID {
			maxID = DaftarSetoran[i].ID
		}
	}
	return maxID + 1
}

// menghitung nomor minggu ISO dari string tanggal format YYYY-MM-DD (nama_anggota)
func HitungNomorMinggu(tanggal string) int {
	t, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return 0
	}
	// mengambil nomor minggu ISO dari objek waktu (nama_anggota)
	_, week := t.ISOWeek()
	return week
}

// menghitung rentang tanggal dari nomor minggu ISO dan tahun yang diberikan (nama_anggota)
func MingguKeRentang(minggu, tahun int) string {
	// Jan 4 selalu berada di minggu ISO ke-1 setiap tahun (nama_anggota)
	t := time.Date(tahun, 1, 4, 0, 0, 0, 0, time.UTC)
	// mundur ke hari Senin di awal minggu ke-1 (nama_anggota)
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	// maju ke minggu yang diminta kemudian hitung hari Minggu-nya (nama_anggota)
	t = t.AddDate(0, 0, (minggu-1)*7)
	end := t.AddDate(0, 0, 6)
	return t.Format("02 Jan") + " - " + end.Format("02 Jan 2006")
}

// mencatat setoran sampah baru ke dalam array setelah validasi (nama_anggota)
func CatatSetoran(wargaID, jenisID int, beratKg float64, tanggal string) bool {
	// memeriksa kapasitas array setoran (nama_anggota)
	if JumlahSetoran >= MAKS_SETORAN {
		fmt.Println("Data setoran sudah mencapai kapasitas maksimum!")
		return false
	}
	// memvalidasi ID warga terdaftar (nama_anggota)
	if CariIndexWargaByID(wargaID) == -1 {
		fmt.Println("ID warga tidak ditemukan!")
		return false
	}
	// memvalidasi ID jenis sampah terdaftar (nama_anggota)
	if !ValidasiIDJenis(jenisID) {
		fmt.Println("ID jenis sampah tidak ditemukan!")
		return false
	}
	minggu := HitungNomorMinggu(tanggal)
	id := GenerateIDSetoran()
	// menyimpan data setoran baru ke posisi berikutnya dalam array (nama_anggota)
	DaftarSetoran[JumlahSetoran] = LogSetoran{
		ID:      id,
		WargaID: wargaID,
		JenisID: jenisID,
		BeratKg: beratKg,
		Tanggal: tanggal,
		Minggu:  minggu,
	}
	JumlahSetoran++
	return true
}

// menghapus setoran berdasarkan ID dengan menggeser elemen ke kiri (nama_anggota)
func HapusSetoran(id int) bool {
	idx := -1
	// mencari posisi index setoran yang ingin dihapus (nama_anggota)
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false
	}
	// menggeser semua elemen setelah index ke kiri (nama_anggota)
	for i := idx; i < JumlahSetoran-1; i++ {
		DaftarSetoran[i] = DaftarSetoran[i+1]
	}
	// mengosongkan elemen terakhir (nama_anggota)
	DaftarSetoran[JumlahSetoran-1] = LogSetoran{}
	JumlahSetoran--
	return true
}

// mengambil salinan seluruh setoran milik warga tertentu beserta jumlahnya (nama_anggota)
func GetSetoranByWarga(wargaID int) ([MAKS_SETORAN]LogSetoran, int) {
	var hasil [MAKS_SETORAN]LogSetoran
	count := 0
	// menyaring setoran yang sesuai dengan wargaID (nama_anggota)
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].WargaID == wargaID {
			hasil[count] = DaftarSetoran[i]
			count++
		}
	}
	return hasil, count
}

// mengambil salinan seluruh setoran pada minggu tertentu beserta jumlahnya (nama_anggota)
func GetSetoranByMinggu(minggu int) ([MAKS_SETORAN]LogSetoran, int) {
	var hasil [MAKS_SETORAN]LogSetoran
	count := 0
	// menyaring setoran yang sesuai dengan nomor minggu (nama_anggota)
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].Minggu == minggu {
			hasil[count] = DaftarSetoran[i]
			count++
		}
	}
	return hasil, count
}
