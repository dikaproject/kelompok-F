package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"wastetrack/data"
)

// nama file JSON untuk menyimpan data warga (nama_anggota)
const fileWarga = "data_warga.json"

// nama file JSON untuk menyimpan data setoran (nama_anggota)
const fileSetoran = "data_setoran.json"

// menyimpan seluruh data warga ke file JSON (nama_anggota)
func simpanWarga() {
	// mengambil slice dari array global sesuai jumlah warga yang terdaftar (nama_anggota)
	slice := data.DaftarWarga[:data.JumlahWarga]
	bytes, err := json.MarshalIndent(slice, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode data warga:", err)
		return
	}
	// menulis bytes JSON ke file di direktori yang sama dengan executable (nama_anggota)
	if err = os.WriteFile(fileWarga, bytes, 0644); err != nil {
		fmt.Println("Gagal menulis file warga:", err)
	}
}

// menyimpan seluruh data setoran ke file JSON (nama_anggota)
func simpanSetoran() {
	// mengambil slice dari array global sesuai jumlah setoran yang tercatat (nama_anggota)
	slice := data.DaftarSetoran[:data.JumlahSetoran]
	bytes, err := json.MarshalIndent(slice, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode data setoran:", err)
		return
	}
	// menulis bytes JSON ke file (nama_anggota)
	if err = os.WriteFile(fileSetoran, bytes, 0644); err != nil {
		fmt.Println("Gagal menulis file setoran:", err)
	}
}

// memuat data warga dari file JSON ke array global (nama_anggota)
func muatWarga() {
	bytes, err := os.ReadFile(fileWarga)
	if err != nil {
		// file belum ada saat pertama kali dijalankan, lewati saja (nama_anggota)
		return
	}
	var slice []data.Warga
	if err = json.Unmarshal(bytes, &slice); err != nil {
		fmt.Println("Gagal membaca data warga:", err)
		return
	}
	// menyalin data dari slice hasil JSON ke array fixed-size global (nama_anggota)
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

// memuat data setoran dari file JSON ke array global (nama_anggota)
func muatSetoran() {
	bytes, err := os.ReadFile(fileSetoran)
	if err != nil {
		// file belum ada saat pertama kali dijalankan, lewati saja (nama_anggota)
		return
	}
	var slice []data.LogSetoran
	if err = json.Unmarshal(bytes, &slice); err != nil {
		fmt.Println("Gagal membaca data setoran:", err)
		return
	}
	// menyalin data dari slice hasil JSON ke array fixed-size global (nama_anggota)
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

// SimpanData menyimpan semua data warga dan setoran ke file JSON (nama_anggota)
func SimpanData() {
	simpanWarga()
	simpanSetoran()
}

// MuatData memuat semua data dari file JSON ke memori saat aplikasi dimulai (nama_anggota)
func MuatData() {
	muatWarga()
	muatSetoran()
}
