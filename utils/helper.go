package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func InputString(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func InputInt(prompt string) int {
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

func InputFloat(prompt string) float64 {
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

func ValidasiTanggal(tanggal string) bool {
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
	return tahun > 2000 && bulan >= 1 && bulan <= 12 && hari >= 1 && hari <= 31
}

func InputTanggal(prompt string) string {
	for {
		input := InputString(prompt + " (format YYYY-MM-DD): ")
		if ValidasiTanggal(input) {
			return input
		}
		fmt.Println("Format tanggal tidak valid. Gunakan YYYY-MM-DD (contoh: 2025-05-01)")
	}
}

func Konfirmasi(prompt string) bool {
	input := InputString(prompt + " (y/n): ")
	return strings.ToLower(input) == "y"
}
