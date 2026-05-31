package main

// eel | @jebb_24

import (
	"fmt"
	"strings"
	"wastetrack/algoritma"
	"wastetrack/data"
	"wastetrack/storage"
	"wastetrack/tampilan"
	"wastetrack/utils"
)

// =============================================
//  MENU: MANAJEMEN DATA WARGA
// =============================================

// menampilkan sub-menu manajemen data warga (nama_anggota)
func menuWarga() {
	for {
		fmt.Println("\n-- Manajemen Data Warga --")
		fmt.Println("1. Tambah Warga")
		fmt.Println("2. Tampil Semua Warga")
		fmt.Println("3. Ubah Data Warga")
		fmt.Println("4. Hapus Warga")
		fmt.Println("0. Kembali ke Menu Utama")
		pilihan := utils.InputInt("Pilih menu: ")
		switch pilihan {
		case 1:
			formTambahWarga()
		case 2:
			tampilan.TampilTabelWarga()
		case 3:
			formUbahWarga()
		case 4:
			formHapusWarga()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// menangani input form tambah warga baru (nama_anggota)
func formTambahWarga() {
	fmt.Println("\n-- Tambah Warga Baru --")
	nama := utils.InputString("Nama    : ")
	alamat := utils.InputString("Alamat  : ")
	noHP := utils.InputString("No. HP  : ")
	if nama == "" {
		fmt.Println("Nama tidak boleh kosong.")
		return
	}
	if data.TambahWarga(nama, alamat, noHP) {
		storage.SimpanData()
		fmt.Println("Warga berhasil ditambahkan.")
	}
}

// menangani input form ubah data warga yang sudah ada (nama_anggota)
func formUbahWarga() {
	fmt.Println("\n-- Ubah Data Warga --")
	tampilan.TampilTabelWarga()
	if data.JumlahWarga == 0 {
		return
	}
	id := utils.InputInt("Masukkan ID warga yang akan diubah: ")
	fmt.Println("(Kosongkan field yang tidak ingin diubah)")
	nama := utils.InputString("Nama baru   : ")
	alamat := utils.InputString("Alamat baru : ")
	noHP := utils.InputString("No. HP baru : ")
	if data.UbahWarga(id, nama, alamat, noHP) {
		storage.SimpanData()
		fmt.Println("Data warga berhasil diubah.")
	} else {
		fmt.Println("ID warga tidak ditemukan.")
	}
}

// menangani konfirmasi dan eksekusi hapus warga (nama_anggota)
func formHapusWarga() {
	fmt.Println("\n-- Hapus Warga --")
	tampilan.TampilTabelWarga()
	if data.JumlahWarga == 0 {
		return
	}
	id := utils.InputInt("Masukkan ID warga yang akan dihapus: ")
	if utils.Konfirmasi("Yakin ingin menghapus warga ini?") {
		if data.HapusWarga(id) {
			storage.SimpanData()
			fmt.Println("Warga berhasil dihapus.")
		} else {
			fmt.Println("ID warga tidak ditemukan.")
		}
	}
}

// =============================================
//  MENU: PENCATATAN SETORAN SAMPAH
// =============================================

// menampilkan sub-menu pencatatan setoran sampah (nama_anggota)
func menuSetoran() {
	for {
		fmt.Println("\n-- Pencatatan Setoran Sampah --")
		fmt.Println("1. Catat Setoran Baru")
		fmt.Println("2. Tampil Semua Setoran")
		fmt.Println("3. Tampil Setoran per Warga")
		fmt.Println("4. Hapus Setoran")
		fmt.Println("0. Kembali ke Menu Utama")
		pilihan := utils.InputInt("Pilih menu: ")
		switch pilihan {
		case 1:
			formCatatSetoran()
		case 2:
			tampilan.TampilTabelSetoran()
		case 3:
			id := utils.InputInt("Masukkan ID warga: ")
			tampilan.TampilSetoranByWarga(id)
		case 4:
			formHapusSetoran()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// menangani input form pencatatan setoran sampah baru (nama_anggota)
func formCatatSetoran() {
	fmt.Println("\n-- Catat Setoran Baru --")
	tampilan.TampilTabelWarga()
	if data.JumlahWarga == 0 {
		return
	}
	wargaID := utils.InputInt("ID Warga   : ")
	fmt.Println()
	data.TampilJenisSampah()
	jenisID := utils.InputInt("ID Jenis   : ")
	berat := utils.InputFloat("Berat (kg) : ")
	tanggal := utils.InputTanggal("Tanggal    ")
	if data.CatatSetoran(wargaID, jenisID, berat, tanggal) {
		storage.SimpanData()
		fmt.Println("Setoran berhasil dicatat.")
	}
}

// menangani konfirmasi dan eksekusi hapus setoran (nama_anggota)
func formHapusSetoran() {
	fmt.Println("\n-- Hapus Setoran --")
	tampilan.TampilTabelSetoran()
	if data.JumlahSetoran == 0 {
		return
	}
	id := utils.InputInt("Masukkan ID setoran yang akan dihapus: ")
	if utils.Konfirmasi("Yakin ingin menghapus setoran ini?") {
		if data.HapusSetoran(id) {
			storage.SimpanData()
			fmt.Println("Setoran berhasil dihapus.")
		} else {
			fmt.Println("ID setoran tidak ditemukan.")
		}
	}
}

// =============================================
//  MENU: PENCARIAN WARGA
// =============================================

// menampilkan sub-menu pencarian warga dengan pilihan algoritma (nama_anggota)
func menuCari() {
	for {
		fmt.Println("\n-- Pencarian Warga --")
		fmt.Println("1. Cari by Nama     (Sequential Search)")
		fmt.Println("2. Cari by ID       (Sequential Search)")
		fmt.Println("3. Cari by ID       (Binary Search)")
		fmt.Println("0. Kembali ke Menu Utama")
		pilihan := utils.InputInt("Pilih menu: ")
		switch pilihan {
		case 1:
			// sequential search tidak memerlukan data terurut (nama_anggota)
			keyword := utils.InputString("Masukkan nama yang dicari: ")
			hasil, count := algoritma.SequentialSearchNama(keyword)
			if count == 0 {
				fmt.Println("Tidak ditemukan warga dengan nama tersebut.")
			} else {
				fmt.Printf("\nDitemukan %d hasil:\n", count)
				fmt.Printf("%-5s  %-20s  %-22s  %s\n", "ID", "Nama", "Alamat", "No. HP")
				fmt.Println(strings.Repeat("-", 64))
				// mencetak setiap hasil pencarian yang ditemukan (nama_anggota)
				for i := 0; i < count; i++ {
					w := hasil[i]
					fmt.Printf("%-5d  %-20s  %-22s  %s\n", w.ID, w.Nama, w.Alamat, w.NoHP)
				}
			}
		case 2:
			// sequential search by ID bekerja pada data tidak terurut (nama_anggota)
			id := utils.InputInt("Masukkan ID warga: ")
			idx := algoritma.SequentialSearchID(id)
			if idx == -1 {
				fmt.Println("Warga tidak ditemukan.")
			} else {
				w := data.DaftarWarga[idx]
				fmt.Printf("\nDitemukan: [%d] %s | %s | %s\n", w.ID, w.Nama, w.Alamat, w.NoHP)
			}
		case 3:
			// binary search membutuhkan data terurut berdasarkan ID (nama_anggota)
			fmt.Println("Catatan: Binary Search membutuhkan data yang terurut by ID.")
			fmt.Println("         Gunakan menu Pengurutan jika data belum terurut by ID.")
			id := utils.InputInt("Masukkan ID warga: ")
			idx := algoritma.BinarySearchID(id)
			if idx == -1 {
				fmt.Println("Warga tidak ditemukan (pastikan data sudah terurut by ID).")
			} else {
				w := data.DaftarWarga[idx]
				fmt.Printf("\nDitemukan: [%d] %s | %s | %s\n", w.ID, w.Nama, w.Alamat, w.NoHP)
			}
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// =============================================
//  MENU: PENGURUTAN WARGA
// =============================================

// menampilkan sub-menu pengurutan data warga dengan pilihan algoritma (nama_anggota)
func menuUrut() {
	for {
		fmt.Println("\n-- Pengurutan Warga (by Total Berat Sampah) --")
		fmt.Println("1. Selection Sort  (descending)")
		fmt.Println("2. Insertion Sort  (descending)")
		fmt.Println("0. Kembali ke Menu Utama")
		pilihan := utils.InputInt("Pilih menu: ")
		switch pilihan {
		case 1:
			// mengurutkan dengan selection sort dan menampilkan hasilnya (nama_anggota)
			algoritma.SelectionSort()
			fmt.Println("Data diurutkan dengan Selection Sort (terbanyak di atas).")
			tampilan.TampilTabelWarga()
		case 2:
			// mengurutkan dengan insertion sort dan menampilkan hasilnya (nama_anggota)
			algoritma.InsertionSort()
			fmt.Println("Data diurutkan dengan Insertion Sort (terbanyak di atas).")
			tampilan.TampilTabelWarga()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// =============================================
//  MENU UTAMA & ENTRY POINT
// =============================================

// fungsi utama yang menjadi titik masuk program Waste-Track (nama_anggota)
func main() {
	// menginisialisasi data jenis sampah default saat aplikasi dimulai (nama_anggota)
	data.InitJenisSampah()
	// memuat data warga dan setoran dari file JSON jika sudah ada (nama_anggota)
	storage.MuatData()

	// melakukan perulangan menu utama hingga pengguna memilih keluar (nama_anggota)
	for {
		tampilan.CetakHeader()
		fmt.Println("1. Manajemen Data Warga")
		fmt.Println("2. Pencatatan Setoran Sampah")
		fmt.Println("3. Pencarian Warga")
		fmt.Println("4. Pengurutan Warga")
		fmt.Println("5. Statistik Mingguan")
		fmt.Println("0. Keluar")
		tampilan.CetakGaris(44)

		pilihan := utils.InputInt("Pilih menu: ")
		switch pilihan {
		case 1:
			menuWarga()
		case 2:
			menuSetoran()
		case 3:
			menuCari()
		case 4:
			menuUrut()
		case 5:
			minggu := utils.InputInt("Masukkan nomor minggu (1-52): ")
			if minggu < 1 || minggu > 52 {
				fmt.Println("Nomor minggu tidak valid (1-52).")
			} else {
				tampilan.TampilStatistikMingguan(minggu)
			}
		case 0:
			fmt.Println("\nTerima kasih telah menggunakan Waste-Track. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
