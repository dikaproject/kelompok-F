package algoritma

import (
	"strings"
	"wastetrack/data"
)


func SequentialSearchNama(keyword string) ([data.MAKS_WARGA]data.Warga, int) {
	var hasil [data.MAKS_WARGA]data.Warga
	count := 0
	keyword = strings.ToLower(keyword)

	for i := 0; i < data.JumlahWarga; i++ {
		if strings.Contains(strings.ToLower(data.DaftarWarga[i].Nama), keyword) {
			hasil[count] = data.DaftarWarga[i]
			count++
		}
	}
	return hasil, count
}

func SequentialSearchID(id int) int {
	for i := 0; i < data.JumlahWarga; i++ {
		if data.DaftarWarga[i].ID == id {
			return i
		}
	}
	return -1
}

func BinarySearchID(id int) int {
	low := 0
	high := data.JumlahWarga - 1

	for low <= high {
		mid := (low + high) / 2

		if data.DaftarWarga[mid].ID == id {
			return mid
		} else if id < data.DaftarWarga[mid].ID {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
