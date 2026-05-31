package algoritma

import "wastetrack/data"

// melakukan selection sort pada DaftarWarga berdasarkan total berat sampah secara descending (nama_anggota)
// kompleksitas waktu: O(n²) — cocok untuk dataset kecil hingga menengah (nama_anggota)
func SelectionSort() {
	n := data.JumlahWarga

	// menghitung total berat setiap warga terlebih dahulu untuk efisiensi (nama_anggota)
	var totalBerat [data.MAKS_WARGA]float64
	for i := 0; i < n; i++ {
		totalBerat[i] = data.GetTotalBeratWarga(data.DaftarWarga[i].ID)
	}

	// melakukan seleksi elemen terbesar dan memindahkannya ke posisi i (nama_anggota)
	for i := 0; i < n-1; i++ {
		maxIdx := i

		// mencari index dengan total berat terbesar di sisa array (nama_anggota)
		for j := i + 1; j < n; j++ {
			if totalBerat[j] > totalBerat[maxIdx] {
				maxIdx = j
			}
		}

		// menukar posisi warga terbesar ke posisi i jika berbeda (nama_anggota)
		if maxIdx != i {
			data.DaftarWarga[i], data.DaftarWarga[maxIdx] = data.DaftarWarga[maxIdx], data.DaftarWarga[i]
			totalBerat[i], totalBerat[maxIdx] = totalBerat[maxIdx], totalBerat[i]
		}
	}
}

// melakukan insertion sort pada DaftarWarga berdasarkan total berat sampah secara descending (nama_anggota)
// kompleksitas waktu: O(n²) worst case, O(n) best case saat data sudah hampir terurut (nama_anggota)
func InsertionSort() {
	n := data.JumlahWarga

	// menghitung total berat setiap warga terlebih dahulu untuk efisiensi (nama_anggota)
	var totalBerat [data.MAKS_WARGA]float64
	for i := 0; i < n; i++ {
		totalBerat[i] = data.GetTotalBeratWarga(data.DaftarWarga[i].ID)
	}

	// menyisipkan elemen ke posisi yang tepat satu per satu dari kiri ke kanan (nama_anggota)
	for i := 1; i < n; i++ {
		keyWarga := data.DaftarWarga[i]
		keyBerat := totalBerat[i]
		j := i - 1

		// menggeser elemen yang lebih kecil ke kanan sampai posisi yang tepat ditemukan (nama_anggota)
		for j >= 0 && totalBerat[j] < keyBerat {
			data.DaftarWarga[j+1] = data.DaftarWarga[j]
			totalBerat[j+1] = totalBerat[j]
			j--
		}

		// menempatkan elemen key di posisi yang sudah tepat (nama_anggota)
		data.DaftarWarga[j+1] = keyWarga
		totalBerat[j+1] = keyBerat
	}
}
