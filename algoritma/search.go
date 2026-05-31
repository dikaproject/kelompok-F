package algoritma

import "wastetrack/data"

// melakukan sequential search pada array warga berdasarkan substring nama (nama_anggota)
// kompleksitas waktu: O(n) — memeriksa setiap elemen satu per satu (nama_anggota)
func SequentialSearchNama(keyword string) ([data.MAKS_WARGA]data.Warga, int) {
	var hasil [data.MAKS_WARGA]data.Warga
	// TODO: implementasi sequential search berdasarkan substring nama (nama_anggota)
	// kompleksitas waktu: O(n) — iterasi setiap elemen dan cocokkan nama
	return hasil, 0
}

// melakukan sequential search pada array warga berdasarkan ID yang tepat (nama_anggota)
// kompleksitas waktu: O(n) — cocok untuk data yang belum terurut (nama_anggota)
func SequentialSearchID(id int) int {
	// TODO: implementasi sequential search berdasarkan ID exact match (nama_anggota)
	// kompleksitas waktu: O(n) — cocok untuk data tidak terurut
	return -1
}

// melakukan binary search pada array warga berdasarkan ID (nama_anggota)
// prasyarat: array DaftarWarga HARUS terurut ascending berdasarkan ID (nama_anggota)
// kompleksitas waktu: O(log n) — lebih efisien dari sequential search (nama_anggota)
func BinarySearchID(id int) int {
	// TODO: implementasi binary search berdasarkan ID (nama_anggota)
	// prasyarat: DaftarWarga HARUS terurut ascending by ID
	// kompleksitas waktu: O(log n) — gunakan low, high, mid
	return -1
}
