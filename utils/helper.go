package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reader global untuk membaca input dari stdin (Rasya Dika Pratama)
var reader = bufio.NewReader(os.Stdin)

// membaca satu baris input string dari pengguna (Rasya Dika Pratama)
func InputString(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	// menghapus karakter newline dan spasi di awal/akhir (Rasya Dika Pratama)
	return strings.TrimSpace(input)
}

// membaca input integer dari pengguna dengan validasi tipe data (Rasya Dika Pratama)
func InputInt(prompt string) int {
	// melakukan perulangan sampai input yang valid diterima (Rasya Dika Pratama)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err == nil {
			return val
		}
		fmt.Println("Input harus berupa angka bulat. Coba lagi.")
	}
}

// membaca input bilangan desimal dari pengguna dengan validasi (Rasya Dika Pratama)
func InputFloat(prompt string) float64 {
	// melakukan perulangan sampai input angka positif yang valid diterima (Rasya Dika Pratama)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.ParseFloat(input, 64)
		if err == nil && val >= 0 {
			return val
		}
		fmt.Println("Input harus berupa angka positif. Coba lagi.")
	}
}

// memvalidasi format string tanggal YYYY-MM-DD secara struktural (Rasya Dika Pratama)
func ValidasiTanggal(tanggal string) bool {
	// memeriksa panjang string dan pemisah yang benar (Rasya Dika Pratama)
	if len(tanggal) != 10 {
		return false
	}
	parts := strings.Split(tanggal, "-")
	if len(parts) != 3 || len(parts[0]) != 4 || len(parts[1]) != 2 || len(parts[2]) != 2 {
		return false
	}
	tahun, errY := strconv.Atoi(parts[0])
	bulan, errM := strconv.Atoi(parts[1])
	hari, errD := strconv.Atoi(parts[2])
	if errY != nil || errM != nil || errD != nil {
		return false
	}
	// memeriksa rentang nilai bulan dan hari (Rasya Dika Pratama)
	return tahun > 2000 && bulan >= 1 && bulan <= 12 && hari >= 1 && hari <= 31
}

// membaca input tanggal dari pengguna dengan validasi format YYYY-MM-DD (Rasya Dika Pratama)
func InputTanggal(prompt string) string {
	// mengulang permintaan input hingga format tanggal yang benar diterima (Rasya Dika Pratama)
	for {
		input := InputString(prompt + " (format YYYY-MM-DD): ")
		if ValidasiTanggal(input) {
			return input
		}
		fmt.Println("Format tanggal tidak valid. Gunakan YYYY-MM-DD (contoh: 2025-05-01)")
	}
}

// meminta konfirmasi ya/tidak dari pengguna dan mengembalikan boolean (Rasya Dika Pratama)
func Konfirmasi(prompt string) bool {
	input := InputString(prompt + " (y/n): ")
	return strings.ToLower(input) == "y"
}
