package data

import "fmt"

type Warga struct {
	ID     int
	Nama   string
	Alamat string
	NoHP   string
}

var DaftarWarga [MAKS_WARGA]Warga

var JumlahWarga int

func GenerateIDWarga() int {
	maxID := 0
	for i := 0; i < JumlahWarga; i++ {
		if DaftarWarga[i].ID > maxID {
			maxID = DaftarWarga[i].ID
		}
	}
	return maxID + 1
}

func TambahWarga(nama, alamat, noHP string) bool {
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

func CariIndexWargaByID(id int) int {
	for i := 0; i < JumlahWarga; i++ {
		if DaftarWarga[i].ID == id {
			return i
		}
	}
	return -1
}

func UbahWarga(id int, nama, alamat, noHP string) bool {
	idx := CariIndexWargaByID(id)
	if idx == -1 {
		return false
	}
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

func HapusWarga(id int) bool {
	idx := CariIndexWargaByID(id)
	if idx == -1 {
		return false
	}
	for i := idx; i < JumlahWarga-1; i++ {
		DaftarWarga[i] = DaftarWarga[i+1]
	}
	DaftarWarga[JumlahWarga-1] = Warga{}
	JumlahWarga--
	return true
}

func GetTotalBeratWarga(wargaID int) float64 {
	total := 0.0
	for i := 0; i < JumlahSetoran; i++ {
		if DaftarSetoran[i].WargaID == wargaID {
			total += DaftarSetoran[i].BeratKg
		}
	}
	return total
}
