package tampilan

import (
	"fmt"
	"strings"
	"time"
	"wastetrack/data"
)

// mencetak baris garis horizontal dengan panjang yang ditentukan (Rasya Dika Pratama)
func CetakGaris(panjang int) {
	fmt.Println(strings.Repeat("=", panjang))
}

// mencetak garis pemisah dengan karakter minus (Rasya Dika Pratama)
func CetakPemisah(panjang int) {
	fmt.Println(strings.Repeat("-", panjang))
}

// mencetak header utama aplikasi Waste-Track (Rasya Dika Pratama)
func CetakHeader() {
	fmt.Println()
	CetakGaris(44)
	fmt.Println("          WASTE-TRACK v1.0")
	fmt.Println("   Manajemen Bank Sampah Lingkungan")
	CetakGaris(44)
}

// mencetak tabel seluruh data warga beserta total berat sampah mereka (Rasya Dika Pratama)
func TampilTabelWarga() {
	if data.JumlahWarga == 0 {
		fmt.Println("Belum ada data warga terdaftar.")
		return
	}
	fmt.Println()
	fmt.Printf("%-5s  %-20s  %-22s  %-13s  %s\n",
		"ID", "Nama", "Alamat", "No. HP", "Total Berat")
	CetakPemisah(76)
	// mencetak setiap baris data warga beserta total berat akumulasinya (Rasya Dika Pratama)
	for i := 0; i < data.JumlahWarga; i++ {
		w := data.DaftarWarga[i]
		total := data.GetTotalBeratWarga(w.ID)
		fmt.Printf("%-5d  %-20s  %-22s  %-13s  %.2f kg\n",
			w.ID, w.Nama, w.Alamat, w.NoHP, total)
	}
	CetakPemisah(76)
	fmt.Printf("Total terdaftar: %d warga\n", data.JumlahWarga)
}

// mencetak tabel seluruh log setoran yang pernah dicatat (Rasya Dika Pratama)
func TampilTabelSetoran() {
	if data.JumlahSetoran == 0 {
		fmt.Println("Belum ada data setoran.")
		return
	}
	fmt.Println()
	fmt.Printf("%-5s  %-20s  %-10s  %-9s  %-12s  %s\n",
		"ID", "Nama Warga", "Jenis", "Berat(kg)", "Tanggal", "Minggu")
	CetakPemisah(72)
	// mencetak setiap entri log setoran dengan nama warga dan jenis sampah (Rasya Dika Pratama)
	for i := 0; i < data.JumlahSetoran; i++ {
		s := data.DaftarSetoran[i]
		namaWarga := "?"
		idx := data.CariIndexWargaByID(s.WargaID)
		if idx != -1 {
			namaWarga = data.DaftarWarga[idx].Nama
		}
		namaJenis := data.GetNamaJenis(s.JenisID)
		fmt.Printf("%-5d  %-20s  %-10s  %-9.2f  %-12s  %d\n",
			s.ID, namaWarga, namaJenis, s.BeratKg, s.Tanggal, s.Minggu)
	}
	CetakPemisah(72)
	fmt.Printf("Total: %d setoran\n", data.JumlahSetoran)
}

// mencetak riwayat setoran sampah untuk satu warga tertentu (Rasya Dika Pratama)
func TampilSetoranByWarga(wargaID int) {
	idx := data.CariIndexWargaByID(wargaID)
	if idx == -1 {
		fmt.Println("ID warga tidak ditemukan!")
		return
	}
	warga := data.DaftarWarga[idx]
	fmt.Printf("\nRiwayat Setoran - %s (ID: %d)\n", warga.Nama, warga.ID)
	CetakPemisah(50)

	setoran, count := data.GetSetoranByWarga(wargaID)
	if count == 0 {
		fmt.Println("Belum ada setoran dari warga ini.")
		return
	}

	total := 0.0
	fmt.Printf("%-5s  %-10s  %-9s  %s\n", "ID", "Jenis", "Berat(kg)", "Tanggal")
	CetakPemisah(42)
	// mencetak setiap entri setoran milik warga yang dicari (Rasya Dika Pratama)
	for i := 0; i < count; i++ {
		s := setoran[i]
		fmt.Printf("%-5d  %-10s  %-9.2f  %s\n",
			s.ID, data.GetNamaJenis(s.JenisID), s.BeratKg, s.Tanggal)
		total += s.BeratKg
	}
	CetakPemisah(42)
	fmt.Printf("Total berat: %.2f kg\n", total)
}

// menghasilkan bar chart ASCII berdasarkan nilai relatif terhadap nilai maksimum (Rasya Dika Pratama)
func buatBar(nilai, maks float64, panjangMaks int) string {
	if maks == 0 {
		return ""
	}
	// menghitung panjang bar secara proporsional (Rasya Dika Pratama)
	panjang := int((nilai / maks) * float64(panjangMaks))
	return strings.Repeat("█", panjang)
}

// mencetak statistik akumulasi setoran sampah berdasarkan nomor minggu (Rasya Dika Pratama)
func TampilStatistikMingguan(minggu int) {
	tahun := time.Now().Year()
	setoran, count := data.GetSetoranByMinggu(minggu)

	fmt.Println()
	CetakGaris(52)
	fmt.Printf("   STATISTIK SAMPAH MINGGU KE-%d (%d)\n", minggu, tahun)
	rentang := data.MingguKeRentang(minggu, tahun)
	fmt.Printf("   Periode : %s\n", rentang)
	CetakGaris(52)

	if count == 0 {
		fmt.Println("Tidak ada data setoran pada minggu ini.")
		CetakGaris(52)
		return
	}

	// akumulasi total berat per jenis sampah dari setoran minggu ini (Rasya Dika Pratama)
	var totalPerJenis [data.MAKS_JENIS]float64
	for i := 0; i < count; i++ {
		for j := 0; j < data.JumlahJenis; j++ {
			if data.DaftarJenis[j].ID == setoran[i].JenisID {
				totalPerJenis[j] += setoran[i].BeratKg
				break
			}
		}
	}

	// mencari nilai terbesar sebagai acuan skala bar chart (Rasya Dika Pratama)
	maxBerat := 0.0
	for j := 0; j < data.JumlahJenis; j++ {
		if totalPerJenis[j] > maxBerat {
			maxBerat = totalPerJenis[j]
		}
	}

	// mencetak tabel statistik per jenis sampah beserta bar chart (Rasya Dika Pratama)
	totalKeseluruhan := 0.0
	fmt.Printf("%-14s | %-11s | %s\n", "Jenis Sampah", "Total Berat", "Bar")
	CetakPemisah(52)
	for j := 0; j < data.JumlahJenis; j++ {
		bar := buatBar(totalPerJenis[j], maxBerat, 18)
		fmt.Printf("%-14s | %8.2f kg | %s\n",
			data.DaftarJenis[j].NamaJenis, totalPerJenis[j], bar)
		totalKeseluruhan += totalPerJenis[j]
	}
	CetakPemisah(52)
	fmt.Printf("%-14s | %8.2f kg |\n", "TOTAL", totalKeseluruhan)
	CetakGaris(52)

	// akumulasi total berat per warga untuk minggu ini (Rasya Dika Pratama)
	var beratPerWarga [data.MAKS_WARGA]float64
	var idWarga [data.MAKS_WARGA]int
	nWarga := 0

	for i := 0; i < count; i++ {
		wargaID := setoran[i].WargaID
		found := false
		// mencari apakah warga sudah ada di daftar akumulasi (Rasya Dika Pratama)
		for k := 0; k < nWarga; k++ {
			if idWarga[k] == wargaID {
				beratPerWarga[k] += setoran[i].BeratKg
				found = true
				break
			}
		}
		if !found {
			idWarga[nWarga] = wargaID
			beratPerWarga[nWarga] = setoran[i].BeratKg
			nWarga++
		}
	}

	// mengurutkan warga berdasarkan total berat descending menggunakan selection sort (Rasya Dika Pratama)
	for i := 0; i < nWarga-1; i++ {
		maxIdx := i
		for j := i + 1; j < nWarga; j++ {
			if beratPerWarga[j] > beratPerWarga[maxIdx] {
				maxIdx = j
			}
		}
		beratPerWarga[i], beratPerWarga[maxIdx] = beratPerWarga[maxIdx], beratPerWarga[i]
		idWarga[i], idWarga[maxIdx] = idWarga[maxIdx], idWarga[i]
	}

	// menampilkan daftar warga pengumpul terbanyak minggu ini (Rasya Dika Pratama)
	tampil := nWarga
	if tampil > 3 {
		tampil = 3
	}
	fmt.Println("Top Warga Pengumpul Terbanyak:")
	for i := 0; i < tampil; i++ {
		namaWarga := "?"
		idx := data.CariIndexWargaByID(idWarga[i])
		if idx != -1 {
			namaWarga = data.DaftarWarga[idx].Nama
		}
		fmt.Printf("  %d. %-20s  —  %.2f kg\n", i+1, namaWarga, beratPerWarga[i])
	}
	CetakGaris(52)
}
