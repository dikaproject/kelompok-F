package algoritma

import (
	"strings"
	"wastetrack/data"
)

// melakukan sequential search pada array warga berdasarkan substring nama (nama_anggota)
// kompleksitas waktu: O(n) — memeriksa setiap elemen satu per satu (nama_anggota)
func SequentialSearchNama(keyword string) ([data.MAKS_WARGA]data.Warga, int) {
	var hasil [data.MAKS_WARGA]data.Warga
	count := 0
	keyword = strings.ToLower(keyword)

	// iterasi dari index 0 hingga akhir untuk mencocokkan nama (nama_anggota)
	for i := 0; i < data.JumlahWarga; i++ {
		if strings.Contains(strings.ToLower(data.DaftarWarga[i].Nama), keyword) {
			hasil[count] = data.DaftarWarga[i]
			count++
		}
	}
	return hasil, count
}

// melakukan sequential search pada array warga berdasarkan ID yang tepat (nama_anggota)
// kompleksitas waktu: O(n) — cocok untuk data yang belum terurut (nama_anggota)
func SequentialSearchID(id int) int {
	// memeriksa setiap elemen dari awal hingga akhir array (nama_anggota)
	for i := 0; i < data.JumlahWarga; i++ {
		if data.DaftarWarga[i].ID == id {
			return i
		}
	}
	return -1
}

// melakukan binary search pada array warga berdasarkan ID (nama_anggota)
// prasyarat: array DaftarWarga HARUS terurut ascending berdasarkan ID (nama_anggota)
// kompleksitas waktu: O(log n) — lebih efisien dari sequential search (nama_anggota)
func BinarySearchID(id int) int {
	low := 0
	high := data.JumlahWarga - 1

	// membagi rentang pencarian menjadi dua bagian di setiap iterasi (nama_anggota)
	for low <= high {
		mid := (low + high) / 2

		if data.DaftarWarga[mid].ID == id {
			// elemen ditemukan tepat di posisi tengah (nama_anggota)
			return mid
		} else if id < data.DaftarWarga[mid].ID {
			// target berada di setengah kiri, pindahkan batas atas (nama_anggota)
			high = mid - 1
		} else {
			// target berada di setengah kanan, pindahkan batas bawah (nama_anggota)
			low = mid + 1
		}
	}
	return -1
}
