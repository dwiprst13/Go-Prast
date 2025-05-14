package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strings"
)

type identitas struct {
	nama string
	nik string
	alamat string
	pekerjaan string
	status string
}

type dataKeluarga struct {
	anggota []identitas
}

func scanInput() string {
	var input string
	bufioReader := bufio.NewReader(os.Stdin)
	input, _ = bufioReader.ReadString('\n')
	return input
}

func upperCaseFirstLetter(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			word = strings.ToUpper(string(word[0])) + word[1:]
			words[i] = word
		}
	}
	return strings.Join(words, " ")
}

func scanInputNama() string {
	var input string
	bufioReader := bufio.NewReader(os.Stdin)
	input, _ = bufioReader.ReadString('\n')
	upperCaseFirstLetter(input)
	return input
}

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func scanInputNIK() string {
	var input string
	var sb strings.Builder
	bufioReader := bufio.NewReader(os.Stdin)
	input, _ = bufioReader.ReadString('\n')
	input = strings.TrimSpace(input) 
	if input == "\n" {
		fmt.Println("NIK tidak boleh kosong")
		return scanInputNIK()
	}
	if len(input) != 16 || !isNumeric(input) {
		fmt.Println("NIK harus terdiri dari 16 digit angka")
		return scanInputNIK()
	}
	sb.WriteString(input)
	sb.WriteRune('\n')
	input = sb.String()
	return input
}

func addAnggota(daftarKeluarga *dataKeluarga) {
	var nama, nik, alamat, pekerjaan, status string
	fmt.Println("Masukkan nama anggota keluarga:")
	nama = scanInputNama()
	fmt.Println("Masukkan NIK anggota keluarga:")
	nik = scanInputNIK()
	fmt.Println("Masukkan alamat anggota keluarga:")
	alamat = scanInput()
	fmt.Println("Masukkan pekerjaan anggota keluarga:")
	pekerjaan = scanInput()
	fmt.Println("Masukkan status anggota keluarga (Kepala Keluarga/Anggota Keluarga):")
	status = scanInput()
	daftarKeluarga.anggota = append(daftarKeluarga.anggota, identitas{nama: nama, nik: nik, alamat: alamat, pekerjaan: pekerjaan, status: status})
	fmt.Println("Anggota keluarga berhasil ditambahkan")
}

func listAnggota(daftarKeluarga *dataKeluarga) {
	if len(daftarKeluarga.anggota) == 0 {
		fmt.Println("Tidak ada anggota keluarga yang tersedia")
		return
	} else {
		fmt.Println("Daftar Anggota Keluarga:")
		for i, anggota := range daftarKeluarga.anggota {
			fmt.Printf("%d Nama : %s NIK : %s Alamat : %s Pekerjaan : %s Status : %s \n", i+1, anggota.nama, anggota.nik, anggota.alamat, anggota.pekerjaan, anggota.status)
			fmt.Println()
		}
	}
}

func deleteAnggota(daftarKeluarga *dataKeluarga) {
	var index int
	fmt.Println("Masukkan nomor anggota keluarga yang ingin dihapus:")
	fmt.Scanln(&index)
	if index < 1 || index > len(daftarKeluarga.anggota) {
		fmt.Println("Nomor anggota keluarga tidak valid")
		return
	} else {
		daftarKeluarga.anggota = append(daftarKeluarga.anggota[:index-1], daftarKeluarga.anggota[index:]...)
		fmt.Println("Anggota keluarga berhasil dihapus")
	}
}

func main() {
	daftarKeluarga := dataKeluarga{}
	var pilihan int

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Tambah Anggota Keluarga")
		fmt.Println("2. Tampilkan Daftar Anggota Keluarga")
		fmt.Println("3. Hapus Anggota Keluarga")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			addAnggota(&daftarKeluarga)
		case 2:
			listAnggota(&daftarKeluarga)
		case 3:
			deleteAnggota(&daftarKeluarga)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}