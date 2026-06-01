package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"wastetrack/data"
)

const fileWarga = "data_warga.json"

const fileSetoran = "data_setoran.json"

func simpanWarga() {
	slice := data.DaftarWarga[:data.JumlahWarga]
	bytes, err := json.MarshalIndent(slice, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode data warga:", err)
		return
	}
	if err = os.WriteFile(fileWarga, bytes, 0644); err != nil {
		fmt.Println("Gagal menulis file warga:", err)
	}
}

func simpanSetoran() {
	slice := data.DaftarSetoran[:data.JumlahSetoran]
	bytes, err := json.MarshalIndent(slice, "", "  ")
	if err != nil {
		fmt.Println("Gagal encode data setoran:", err)
		return
	}
	if err = os.WriteFile(fileSetoran, bytes, 0644); err != nil {
		fmt.Println("Gagal menulis file setoran:", err)
	}
}

func muatWarga() {
	bytes, err := os.ReadFile(fileWarga)
	if err != nil {
		return
	}
	var slice []data.Warga
	if err = json.Unmarshal(bytes, &slice); err != nil {
		fmt.Println("Gagal membaca data warga:", err)
		return
	}
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

func muatSetoran() {
	bytes, err := os.ReadFile(fileSetoran)
	if err != nil {
		return
	}
	var slice []data.LogSetoran
	if err = json.Unmarshal(bytes, &slice); err != nil {
		fmt.Println("Gagal membaca data setoran:", err)
		return
	}
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

func SimpanData() {
	simpanWarga()
	simpanSetoran()
}

func MuatData() {
	muatWarga()
	muatSetoran()
}
