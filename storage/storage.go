package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"wastetrack/data"
)

// nama file JSON untuk menyimpan data warga (Rasya Dika Pratama)
const fileWarga = "data_warga.json"

// nama file JSON untuk menyimpan data setoran (Rasya Dika Pratama)
const fileSetoran = "data_setoran.json"

// menyimpan seluruh data warga ke file JSON (Rasya Dika Pratama)
func simpanWarga() {
	// mengambil slice dari array global sesuai jumlah warga yang terdaftar (Rasya Dika Pratama)
	slice := data.DaftarWarga[:data.JumlahWarga]
	bytes, err := json.MarshalIndent(slice, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode data warga:", err)
		return
	}
	// menulis bytes JSON ke file di direktori yang sama dengan executable (Rasya Dika Pratama)
	if err = os.WriteFile(fileWarga, bytes, 0644); err != nil {
		fmt.Println("Gagal menulis file warga:", err)
	}
}

// menyimpan seluruh data setoran ke file JSON (Rasya Dika Pratama)
func simpanSetoran() {
	// mengambil slice dari array global sesuai jumlah setoran yang tercatat (Rasya Dika Pratama)
	slice := data.DaftarSetoran[:data.JumlahSetoran]
	bytes, err := json.MarshalIndent(slice, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode data setoran:", err)
		return
	}
	// menulis bytes JSON ke file (Rasya Dika Pratama)
	if err = os.WriteFile(fileSetoran, bytes, 0644); err != nil {
		fmt.Println("Gagal menulis file setoran:", err)
	}
}

// memuat data warga dari file JSON ke array global (Rasya Dika Pratama)
func muatWarga() {
	bytes, err := os.ReadFile(fileWarga)
	if err != nil {
		// file belum ada saat pertama kali dijalankan, lewati saja (Rasya Dika Pratama)
		return
	}
	var slice []data.Warga
	if err = json.Unmarshal(bytes, &slice); err != nil {
		fmt.Println("Gagal membaca data warga:", err)
		return
	}
	// menyalin data dari slice hasil JSON ke array fixed-size global (Rasya Dika Pratama)
	for i, w := range slice {
		if i >= data.MAKS_WARGA {
			break
		}
		data.DaftarWarga[i] = w
	}
	data.JumlahWarga = len(slice)
	if data.JumlahWarga > data.MAKS_WARGA {
		data.JumlahWarga = data.MAKS_WARGA
	}
}

// memuat data setoran dari file JSON ke array global (Rasya Dika Pratama)
func muatSetoran() {
	bytes, err := os.ReadFile(fileSetoran)
	if err != nil {
		// file belum ada saat pertama kali dijalankan, lewati saja (Rasya Dika Pratama)
		return
	}
	var slice []data.LogSetoran
	if err = json.Unmarshal(bytes, &slice); err != nil {
		fmt.Println("Gagal membaca data setoran:", err)
		return
	}
	// menyalin data dari slice hasil JSON ke array fixed-size global (Rasya Dika Pratama)
	for i, s := range slice {
		if i >= data.MAKS_SETORAN {
			break
		}
		data.DaftarSetoran[i] = s
	}
	data.JumlahSetoran = len(slice)
	if data.JumlahSetoran > data.MAKS_SETORAN {
		data.JumlahSetoran = data.MAKS_SETORAN
	}
}

// SimpanData menyimpan semua data warga dan setoran ke file JSON (Rasya Dika Pratama)
func SimpanData() {
	simpanWarga()
	simpanSetoran()
}

// MuatData memuat semua data dari file JSON ke memori saat aplikasi dimulai (Rasya Dika Pratama)
func MuatData() {
	muatWarga()
	muatSetoran()
}
