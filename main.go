package main

import (
	"fmt"
	"time"
)

const NMAX = 5000

type tabPenduduk [NMAX]dataDesa

type dataDesa struct {
	namaDesa   string
	alamatDesa string
	jumlahRT   int
	jumlahRW   int
	penduduk   dataPenduduk
}

type dataPenduduk struct {
	namaPenduduk     string
	umurPenduduk     int
	alamatRumah      string
	jumlahRT         int
	jumlahRW         int
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
var data tabPenduduk
var nData int
var shouldExit bool
var namaDicari string
var pilih, nomorNIK int

func main() {
	var input Login
	login(&input)
	loading()
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
			inputData(&data, &nData)
		case 2:
			cariData(data, nData)
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
func inputData(T *tabPenduduk, n *int) {
	fmt.Println("================================")
	fmt.Println("================================")
	fmt.Println("SILAHKAN MASUKKAN DATA PENDUDUK")
	fmt.Println("================================")
	fmt.Println("Masukkan jumlah data:")
	fmt.Scan(n)
	i := 0
	for i < *n {
		fmt.Printf("Data Desa Ke-%d\n", i+1)
		fmt.Print("Nama Desa: ")
		fmt.Scan(&T[i].namaDesa)
		fmt.Print("Alamat Desa: ")
		fmt.Scan(&T[i].alamatDesa)
		fmt.Print("Jumlah RT: ")
		fmt.Scan(&T[i].jumlahRT)
		fmt.Print("Jumlah RW: ")
		fmt.Scan(&T[i].jumlahRW)
		fmt.Print("Nama Penduduk: ")
		fmt.Scan(&T[i].penduduk.namaPenduduk)
		fmt.Print("Alamat Penduduk: ")
		fmt.Scan(&T[i].penduduk.alamatRumah)
		fmt.Print("Umur: ")
		fmt.Scan(&T[i].penduduk.umurPenduduk)
		fmt.Print("RT: ")
		fmt.Scan(&T[i].penduduk.noRT)
		fmt.Print("RW: ")
		fmt.Scan(&T[i].penduduk.noRW)
		fmt.Print("Masukkan Nomor NIK: ")
		fmt.Scan(&T[i].penduduk.noNIK)
		fmt.Print("Status Kawin?: ")
		fmt.Scan(&T[i].penduduk.statusPerkawinan)
		i++
	}
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
			if T[i].penduduk.namaPenduduk == namaDicari {
				fmt.Printf("Data Ditemukan: %+v\n", T[i].penduduk)
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
			if T[i].penduduk.noNIK == nomorNIK {
				fmt.Printf("Data Ditemukan: %+v\n", T[i].penduduk)
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
func cetakData(T tabPenduduk, n int) {
	fmt.Println("================================")
	fmt.Println("DATA PENDUDUK DESA")
	fmt.Println("================================")
	for i := 0; i < n; i++ {
		fmt.Printf("Data Desa Ke-%d\n", i+1)
		fmt.Printf("Nama Desa: %s\n", T[i].namaDesa)
		fmt.Printf("Alamat Desa: %s\n", T[i].alamatDesa)
		fmt.Printf("Jumlah RT: %d\n", T[i].jumlahRT)
		fmt.Printf("Jumlah RW: %d\n", T[i].jumlahRW)
		fmt.Printf("Nama Penduduk: %s\n", T[i].penduduk.namaPenduduk)
		fmt.Printf("Alamat Penduduk: %s\n", T[i].penduduk.alamatRumah)
		fmt.Printf("Umur: %d\n", T[i].penduduk.umurPenduduk)
		fmt.Printf("RT: %d\n", T[i].penduduk.noRT)
		fmt.Printf("RW: %d\n", T[i].penduduk.noRW)
		fmt.Printf("Nomor NIK: %d\n", T[i].penduduk.noNIK)
		fmt.Printf("Status Kawin: %s\n", T[i].penduduk.statusPerkawinan)
		fmt.Println("================================")
	}
}
