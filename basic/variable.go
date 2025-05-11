package main

import (
	"fmt"
)

type SpesifikasiDasar struct {
	merk   string
	tipe   string
	varian string
}

type SpesifikasiMesin struct {
	tipeMesin          string
	kodeMesin          string
	bahanBakar         string
	kubikasi           string
	silinder           string
	tenaga             string
	tenagaMaksimum     string
	tenagaMaksimumRpm  string
	torsi              string
	torsiRpm           string
	rasioKompresi      string
}

// Struct gabungan
type Mobil struct {
	spesifikasiDasar SpesifikasiDasar
	spesifikasiMesin SpesifikasiMesin
}

func tampilkanMobil(m Mobil) {
	fmt.Println("Spesifikasi Dasar")
	fmt.Println("Merk: ", m.spesifikasiDasar.merk)
	fmt.Println("Tipe: ", m.spesifikasiDasar.tipe)
	fmt.Println("Varian: ", m.spesifikasiDasar.varian)
	fmt.Println()
	fmt.Println("Spesifikasi Mesin")
	fmt.Println("Tipe Mesin: ", m.spesifikasiMesin.tipeMesin)
	fmt.Println("Kode Mesin: ", m.spesifikasiMesin.kodeMesin)
	fmt.Println("Bahan Bakar: ", m.spesifikasiMesin.bahanBakar)
	fmt.Println("Kubikasi: ", m.spesifikasiMesin.kubikasi)
	fmt.Println("Silinder: ", m.spesifikasiMesin.silinder)
	fmt.Println("Tenaga: ", m.spesifikasiMesin.tenaga)
	fmt.Println("Tenaga Maksimum: ", m.spesifikasiMesin.tenagaMaksimum)
	fmt.Println("Tenaga Maksimum RPM: ", m.spesifikasiMesin.tenagaMaksimumRpm)
	fmt.Println("Torsi: ", m.spesifikasiMesin.torsi)
	fmt.Println("Torsi RPM: ", m.spesifikasiMesin.torsiRpm)
	fmt.Println("Rasio Kompresi: ", m.spesifikasiMesin.rasioKompresi)
	fmt.Println()
}

func main() {
	// Daftar mobil
	for {
			mobilList := []Mobil{
		{
			spesifikasiDasar: SpesifikasiDasar{"Toyota", "Avanza", "1.3 E M/T"},
			spesifikasiMesin: SpesifikasiMesin{
				"K3-VE", "K3-VE", "Bensin", "1298 cc",
				"4 Silinder Segaris, DOHC, 16 Katup, VVT-i",
				"96 PS / 6000 rpm", "96 PS", "6000 rpm",
				"12.3 kgm / 4200 rpm", "4200 rpm", "10.5 : 1",
			},
		},
		{
			spesifikasiDasar: SpesifikasiDasar{"Honda", "Brio", "Satya E CVT"},
			spesifikasiMesin: SpesifikasiMesin{
				"L12B", "L12B", "Bensin", "1199 cc",
				"4 Silinder SOHC, i-VTEC",
				"90 PS / 6000 rpm", "90 PS", "6000 rpm",
				"110 Nm / 4800 rpm", "4800 rpm", "10.1 : 1",
			},
		},
	}

	// Tampilkan pilihan
	fmt.Println("Pilih mobil yang ingin ditampilkan:")
	for i, m := range mobilList {
		fmt.Printf("%d. %s %s %s\n", i+1, m.spesifikasiDasar.merk, m.spesifikasiDasar.tipe, m.spesifikasiDasar.varian)
	}

	// Input pilihan
	var pilihan int
	fmt.Print("Masukkan nomor pilihan: ")
	fmt.Scanln(&pilihan)

	if pilihan != 0 {
			// Validasi input
		if pilihan >= 1 && pilihan <= len(mobilList) {
			fmt.Println()
			tampilkanMobil(mobilList[pilihan-1])
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
	if pilihan == 0 {
		break
	}
}
}