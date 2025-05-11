package main

import "fmt"

const maxData = 100

type Aktivitas struct {
	nama         string
	jamPerMinggu int
	penghasilan  int
}

var data [maxData]Aktivitas
var nData int

func main() {
	var pilihan int
	for {
		fmt.Println("==== SIDE HUSTLE INCOME PLANNER (SHIP) ====")
		fmt.Println("1. Tambah aktivitas baru")
		fmt.Println("2. Lihat semua aktivitas")
		fmt.Println("3. Cari aktivitas")
		fmt.Println("4. Urutkan aktivitas")
		fmt.Println("5. Lihat aktivitas paling efisien")
		fmt.Println("6. Analisis ketahanan finansial Side Hustle")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahAktivitas()
		case 2:
			lihatSemuaAktivitas()
		case 3:
			cariAktivitas()
		case 4:
			urutkanAktivitas()
		case 5:
			aktivitasPalingEfisien()
		case 6:
			analisisKetahanan()
		case 7:
			fmt.Println("Terima kasih telah menggunakan SHIP!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi-fungsi di bawah ini akan kita lengkapi satu per satu:
func tambahAktivitas() {
	// Implementasi input dan simpan aktivitas
}

func lihatSemuaAktivitas() {
	// Tampilkan semua data
}

func cariAktivitas() {
	// Linear search berdasarkan nama
}

func urutkanAktivitas() {
	// Selection sort berdasarkan efisiensi (penghasilan/jam)
}

func aktivitasPalingEfisien() {
	// Cari aktivitas dengan rasio penghasilan/jam tertinggi
}

func analisisKetahanan() {
	// Total penghasilan per minggu dan estimasi penghematan
}
