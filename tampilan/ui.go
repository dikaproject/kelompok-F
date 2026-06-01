package tampilan

import (
	"fmt"
	"strings"
	"time"
	"wastetrack/data"
)

func CetakGaris(panjang int) {
	fmt.Println(strings.Repeat("=", panjang))
}

func CetakPemisah(panjang int) {
	fmt.Println(strings.Repeat("-", panjang))
}

func CetakHeader() {
	fmt.Println()
	CetakGaris(44)
	fmt.Println("          WASTE-TRACK")
	fmt.Println("   Manajemen Bank Sampah Lingkungan")
	CetakGaris(44)
}

func TampilTabelWarga() {
	if data.JumlahWarga == 0 {
		fmt.Println("Belum ada data warga terdaftar.")
		return
	}
	fmt.Println()
	fmt.Printf("%-5s  %-20s  %-22s  %-13s  %s\n",
		"ID", "Nama", "Alamat", "No. HP", "Total Berat")
	CetakPemisah(76)
	for i := 0; i < data.JumlahWarga; i++ {
		w := data.DaftarWarga[i]
		total := data.GetTotalBeratWarga(w.ID)
		fmt.Printf("%-5d  %-20s  %-22s  %-13s  %.2f kg\n",
			w.ID, w.Nama, w.Alamat, w.NoHP, total)
	}
	CetakPemisah(76)
	fmt.Printf("Total terdaftar: %d warga\n", data.JumlahWarga)
}

func TampilTabelSetoran() {
	if data.JumlahSetoran == 0 {
		fmt.Println("Belum ada data setoran.")
		return
	}
	fmt.Println()
	fmt.Printf("%-5s  %-20s  %-10s  %-9s  %-12s  %s\n",
		"ID", "Nama Warga", "Jenis", "Berat(kg)", "Tanggal", "Minggu")
	CetakPemisah(72)
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
	for i := 0; i < count; i++ {
		s := setoran[i]
		fmt.Printf("%-5d  %-10s  %-9.2f  %s\n",
			s.ID, data.GetNamaJenis(s.JenisID), s.BeratKg, s.Tanggal)
		total += s.BeratKg
	}
	CetakPemisah(42)
	fmt.Printf("Total berat: %.2f kg\n", total)
}

func buatBar(nilai, maks float64, panjangMaks int) string {
	if maks == 0 {
		return ""
	}
	panjang := int((nilai / maks) * float64(panjangMaks))
	return strings.Repeat("█", panjang)
}

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

	var totalPerJenis [data.MAKS_JENIS]float64
	for i := 0; i < count; i++ {
		for j := 0; j < data.JumlahJenis; j++ {
			if data.DaftarJenis[j].ID == setoran[i].JenisID {
				totalPerJenis[j] += setoran[i].BeratKg
				break
			}
		}
	}

	maxBerat := 0.0
	for j := 0; j < data.JumlahJenis; j++ {
		if totalPerJenis[j] > maxBerat {
			maxBerat = totalPerJenis[j]
		}
	}

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

	var beratPerWarga [data.MAKS_WARGA]float64
	var idWarga [data.MAKS_WARGA]int
	nWarga := 0

	for i := 0; i < count; i++ {
		wargaID := setoran[i].WargaID
		found := false
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
