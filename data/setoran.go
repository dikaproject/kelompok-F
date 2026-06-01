package data

import (
	"fmt"
	"time"
)

type LogSetoran struct {
	ID      int
	WargaID int
	JenisID int
	BeratKg float64
	Tanggal string
	Minggu  int
}

var DaftarSetoran [MAKS_SETORAN]LogSetoran

var JumlahSetoran int

func GenerateIDSetoran() int {
	maxID := 0
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].ID > maxID {
			maxID = DaftarSetoran[i].ID
		}
	}
	return maxID + 1
}

func HitungNomorMinggu(tanggal string) int {
	t, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return 0
	}
	_, week := t.ISOWeek()
	return week
}

func MingguKeRentang(minggu, tahun int) string {
	t := time.Date(tahun, 1, 4, 0, 0, 0, 0, time.UTC)
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	t = t.AddDate(0, 0, (minggu-1)*7)
	end := t.AddDate(0, 0, 6)
	return t.Format("02 Jan") + " - " + end.Format("02 Jan 2006")
}

func CatatSetoran(wargaID, jenisID int, beratKg float64, tanggal string) bool {
	if JumlahSetoran >= MAKS_SETORAN {
		fmt.Println("Data setoran sudah mencapai kapasitas maksimum!")
		return false
	}
	if CariIndexWargaByID(wargaID) == -1 {
		fmt.Println("ID warga tidak ditemukan!")
		return false
	}
	if !ValidasiIDJenis(jenisID) {
		fmt.Println("ID jenis sampah tidak ditemukan!")
		return false
	}
	minggu := HitungNomorMinggu(tanggal)
	id := GenerateIDSetoran()
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

func HapusSetoran(id int) bool {
	idx := -1
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false
	}
	for i := idx; i < JumlahSetoran-1; i++ {
		DaftarSetoran[i] = DaftarSetoran[i+1]
	}
	DaftarSetoran[JumlahSetoran-1] = LogSetoran{}
	JumlahSetoran--
	return true
}

func GetSetoranByWarga(wargaID int) ([MAKS_SETORAN]LogSetoran, int) {
	var hasil [MAKS_SETORAN]LogSetoran
	count := 0
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].WargaID == wargaID {
			hasil[count] = DaftarSetoran[i]
			count++
		}
	}
	return hasil, count
}

func GetSetoranByMinggu(minggu int) ([MAKS_SETORAN]LogSetoran, int) {
	var hasil [MAKS_SETORAN]LogSetoran
	count := 0
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].Minggu == minggu {
			hasil[count] = DaftarSetoran[i]
			count++
		}
	}
	return hasil, count
}
