package main

import (
	"fmt"
	"time"
)

const NMAX = 5000

type tabPenduduk [NMAX]dataPenduduk
type tabDesa [NMAX]dataDesa

type dataDesa struct {
	namaDesa   string
	alamatDesa string
	jumlahRT   int
	jumlahRW   int
}

type dataPenduduk struct {
	namaPenduduk     string
	umurPenduduk     int
	alamatRumah      string
	noRT             int
	noRW             int
	noNIK            int
	statusPerkawinan string
}

type Login struct {
	username string
	email    string
	password string
}

// KAMUS GLOBAL
var dataP tabPenduduk
var dataD tabDesa
var nDesa int
var nPenduduk int
var shouldExit bool
var namaDicari string
var pilih, nomorNIK int

func main() {
	// var input Login
	// login(&input)
	// loading()
	menu()
}

// Tampilan Menu Utama
func menu() {
	for !shouldExit {
		fmt.Println("================================")
		fmt.Printf("%18s\n", "SI DESA")
		fmt.Println(" Aplikasi Sistem Informasi Desa")
		fmt.Println("================================")
		fmt.Println("Pilih menu:")
		fmt.Println("1. Input Data")
		fmt.Println("2. Cari Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Edit Data")
		fmt.Println("5. Exit")
		fmt.Println()
		fmt.Print("Pilih:")
		fmt.Scan(&pilih)
		fmt.Println("================================")
		fmt.Println()

		switch pilih {
		case 1:
			inputData(&dataD, &dataP, &nDesa, &nPenduduk)
		case 2:
			cariData(dataP, nPenduduk)
		// case 3:
		// 	Implement Hapus Data
		// case 4:
		// 	Implement Edit Data
		case 5:
			shouldExit = true
			fmt.Println("================================")
			fmt.Println("TERIMA KASIH")
			fmt.Println("================================")
		}
	}
}

// Loading Animation
func loading() {
	for i := 0; i < 10; i++ {
		fmt.Print("\rMOHON TUNGGU SEBENTAR YA SAYANG!! .")
		time.Sleep(200 * time.Millisecond)
		fmt.Print("\rMOHON TUNGGU SEBENTAR YA SAYANG!!  .")
		time.Sleep(200 * time.Millisecond)
		fmt.Print("\rMOHON TUNGGU SEBENTAR YA SAYANG!!   .")
		time.Sleep(200 * time.Millisecond)
		fmt.Print("\rMOHON TUNGGU SEBENTAR YA SAYANG!!    .")
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()
}

// Login Function
func login(input *Login) {
	durasi := 3
	percobaanMaks := 2
	percobaan := 0
	input.username = "Admin"
	input.email = "Admin"
	input.password = "Admin"

	time.Sleep(time.Duration(durasi) * time.Second)
	fmt.Println("================================")
	fmt.Println("SELAMAT DATANG DI SIDESA!!")
	fmt.Println("================================")

	for percobaan <= percobaanMaks {
		fmt.Print("Username: ")
		fmt.Scan(&input.username)
		fmt.Print("Email: ")
		fmt.Scan(&input.email)
		fmt.Print("Password: ")
		fmt.Scan(&input.password)
		if input.username == "Admin" && input.email == "Admin" && input.password == "Admin" {
			break
		} else {
			fmt.Println("Data tidak sesuai")
			percobaan++
		}
		if percobaan == percobaanMaks {
			fmt.Println("================================")
			fmt.Println("MOHON MAAF DATA KAMU SALAH")
			fmt.Println("================================")
			fmt.Println("SILAHKAN TUNGGU DALAM WAKTU 1 MENIT")
			fmt.Println("================================")
			time.Sleep(time.Duration(durasi) * time.Second)
			login(input)
			return
		}
	}
}

// Function to Input Data
func inputData(K *tabDesa, T *tabPenduduk, nDesa, nPenduduk *int) {
	fmt.Println("================================")
	fmt.Println("SILAHKAN MASUKKAN DATA PENDUDUK")
	fmt.Println("================================")
	fmt.Print("Masukkan jumlah Desa: ")
	fmt.Scan(nDesa)

	totalPenduduk := 0
	for i := 0; i < *nDesa; i++ {
		fmt.Printf("Data Desa Ke-%d\n", i+1)
		fmt.Print("Nama Desa: ")
		fmt.Scan(&(*K)[i].namaDesa)
		fmt.Print("Alamat Desa: ")
		fmt.Scan(&(*K)[i].alamatDesa)
		fmt.Print("Jumlah RT: ")
		fmt.Scan(&(*K)[i].jumlahRT)
		fmt.Print("Jumlah RW: ")
		fmt.Scan(&(*K)[i].jumlahRW)

		var nPendudukDesa int
		fmt.Print("Masukkan jumlah penduduk: ")
		fmt.Scan(&nPendudukDesa)

		for j := 0; j < nPendudukDesa; j++ {
			fmt.Printf("Data Penduduk Ke-%d Desa %s\n", j+1, (*K)[i].namaDesa)
			fmt.Print("Nama Penduduk: ")
			fmt.Scan(&(*T)[totalPenduduk].namaPenduduk)
			fmt.Print("Alamat Penduduk: ")
			fmt.Scan(&(*T)[totalPenduduk].alamatRumah)
			fmt.Print("Umur: ")
			fmt.Scan(&(*T)[totalPenduduk].umurPenduduk)
			fmt.Print("RT: ")
			fmt.Scan(&(*T)[totalPenduduk].noRT)
			fmt.Print("RW: ")
			fmt.Scan(&(*T)[totalPenduduk].noRW)
			fmt.Print("Masukkan Nomor NIK: ")
			fmt.Scan(&(*T)[totalPenduduk].noNIK)
			fmt.Print("Status Kawin?: ")
			fmt.Scan(&(*T)[totalPenduduk].statusPerkawinan)
			totalPenduduk++
		}
	}
	*nPenduduk = totalPenduduk
}

// Function to Search Data
func cariData(T tabPenduduk, n int) {
	fmt.Println("================================")
	fmt.Println("PENCARIAN DATA PENDUDUK DESA")
	fmt.Println("================================")
	fmt.Println("MENU PILIHAN PENCARIAN")
	fmt.Println("1. Nama")
	fmt.Println("2. NIK")
	fmt.Println("3. EXIT")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		fmt.Println("================================")
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&namaDicari)
		i := 0
		found := false
		for i < n {
			if T[i].namaPenduduk == namaDicari {
				fmt.Printf("Data Ditemukan: %+v\n", T[i])
				found = true
			}
			i++
		}
		if !found {
			fmt.Println("Data Tidak Ada")
		}
	case 2:
		fmt.Println("================================")
		fmt.Print("Masukkan NIK: ")
		fmt.Scan(&nomorNIK)
		i := 0
		found := false
		for i < n {
			if T[i].noNIK == nomorNIK {
				fmt.Printf("Data Ditemukan: %+v\n", T[i])
				found = true
			}
			i++
		}
		if !found {
			fmt.Println("Data Tidak Ada")
		}
	case 3:
		return
	}
}

// Function to Print Data
func cetakData(T tabPenduduk, K tabDesa, nDesa, nPenduduk int) {
	fmt.Println("================================")
	fmt.Println("DATA PENDUDUK DESA")
	fmt.Println("================================")
	for i := 0; i < nDesa; i++ {
		fmt.Printf("Data Desa Ke-%d\n", i+1)
		fmt.Printf("Nama Desa: %s\n", K[i].namaDesa)
		fmt.Printf("Alamat Desa: %s\n", K[i].alamatDesa)
		fmt.Printf("Jumlah RT: %d\n", K[i].jumlahRT)
		fmt.Printf("Jumlah RW: %d\n", K[i].jumlahRW)
		fmt.Println("================================")
		for j := 0; j < nPenduduk; j++ {
			if T[j].alamatRumah == K[i].alamatDesa { // Assuming the address matches the village
				fmt.Printf("Nama Penduduk: %s\n", T[j].namaPenduduk)
				fmt.Printf("Alamat Penduduk: %s\n", T[j].alamatRumah)
				fmt.Printf("Umur: %d\n", T[j])
			}
		}
	}
}