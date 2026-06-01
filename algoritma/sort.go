package algoritma

import "wastetrack/data"

func SelectionSort() {
	n := data.JumlahWarga

	var totalBerat [data.MAKS_WARGA]float64
	for i := 0; i < n; i++ {
		totalBerat[i] = data.GetTotalBeratWarga(data.DaftarWarga[i].ID)
	}

	for i := 0; i < n-1; i++ {
		maxIdx := i

		for j := i + 1; j < n; j++ {
			if totalBerat[j] > totalBerat[maxIdx] {
				maxIdx = j
			}
		}

		if maxIdx != i {
			data.DaftarWarga[i], data.DaftarWarga[maxIdx] = data.DaftarWarga[maxIdx], data.DaftarWarga[i]
			totalBerat[i], totalBerat[maxIdx] = totalBerat[maxIdx], totalBerat[i]
		}
	}
}

func InsertionSort() {
	n := data.JumlahWarga

	var totalBerat [data.MAKS_WARGA]float64
	for i := 0; i < n; i++ {
		totalBerat[i] = data.GetTotalBeratWarga(data.DaftarWarga[i].ID)
	}

	for i := 1; i < n; i++ {
		keyWarga := data.DaftarWarga[i]
		keyBerat := totalBerat[i]
		j := i - 1

		for j >= 0 && totalBerat[j] < keyBerat {
			data.DaftarWarga[j+1] = data.DaftarWarga[j]
			totalBerat[j+1] = totalBerat[j]
			j--
		}

		data.DaftarWarga[j+1] = keyWarga
		totalBerat[j+1] = keyBerat
	}
}
