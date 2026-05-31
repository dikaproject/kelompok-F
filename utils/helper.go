package utils

// TODO: tambahkan import (bufio, fmt, os, strconv, strings) saat implementasi (nama_anggota)
// TODO: tambahkan: var reader = bufio.NewReader(os.Stdin)

// membaca satu baris input string dari pengguna (nama_anggota)
func InputString(prompt string) string {
	// TODO: tampilkan prompt, baca satu baris, trim newline & spasi (nama_anggota)
	return ""
}

// membaca input integer dari pengguna dengan validasi tipe data (nama_anggota)
func InputInt(prompt string) int {
	// TODO: loop baca input, konversi ke int dengan strconv.Atoi, ulangi jika error (nama_anggota)
	return 0
}

// membaca input bilangan desimal dari pengguna dengan validasi (nama_anggota)
func InputFloat(prompt string) float64 {
	// TODO: loop baca input, konversi dengan strconv.ParseFloat, validasi >= 0 (nama_anggota)
	return 0
}

// memvalidasi format string tanggal YYYY-MM-DD secara struktural (nama_anggota)
func ValidasiTanggal(tanggal string) bool {
	// TODO: cek panjang=10, split "-" jadi 3 bagian, validasi rentang tahun/bulan/hari (nama_anggota)
	return false
}

// membaca input tanggal dari pengguna dengan validasi format YYYY-MM-DD (nama_anggota)
func InputTanggal(prompt string) string {
	// TODO: loop InputString sampai ValidasiTanggal() mengembalikan true (nama_anggota)
	return ""
}

// meminta konfirmasi ya/tidak dari pengguna dan mengembalikan boolean (nama_anggota)
func Konfirmasi(prompt string) bool {
	// TODO: tampilkan prompt + " (y/n): ", return true jika input == "y" (nama_anggota)
	return false
}
