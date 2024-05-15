package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const NMAX = 5000

type tabDesa [NMAX]dataDesa
type tabPenduduk [NMAX]dataPenduduk

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
var data tabDesa
var penduduk tabPenduduk
var nData, pData int
var shouldExit bool
var namaDicari string
var pilih, nomorNIK int

func main() {
	//var input Login
	//login(&input)
	loading()
	for !shouldExit {
		menu()
		switch  {
		case pilih == 1:
			inputDesa(&data, &nData)
		case pilih == 2:
			inputPenduduk(&penduduk, data, &pData)
		case pilih == 3:
			cariData(penduduk, pData)
		}
		if pilih == 4{
			shouldExit = true
			fmt.Println("================================")
			fmt.Printf("%14s\n", "Terima Kasih!")
			fmt.Println("================================")
		} 
	}
}

// Tampilan Menu Utama
func menu() {
		fmt.Println("================================")
		fmt.Printf("%18s\n", "SI DESA")
		fmt.Println(" Aplikasi Sistem Informasi Desa")
		fmt.Println("================================")
		fmt.Println("Pilih menu:")
		fmt.Println("1. Input Data Desa")
		fmt.Println("2. Input Data Penduduk")
		fmt.Println("3. Cari Data")
		fmt.Println("4. Edit Data")
		fmt.Println("5. Exit")
		fmt.Println()
		fmt.Println("================================")
		fmt.Print("Pilih:")
		fmt.Scan(&pilih)
		fmt.Println()
		clear()
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
//func login(input *Login) {
// 	durasi := 3
// 	percobaanMaks := 2
// 	percobaan := 0
// 	input.username = "Admin"
// 	input.email = "Admin"
// 	input.password = "Admin"

// 	time.Sleep(time.Duration(durasi) * time.Second)
// 	fmt.Println("================================")
// 	fmt.Println("SELAMAT DATANG DI SIDESA!!")
// 	fmt.Println("================================")

// 	for percobaan <= percobaanMaks {
// 		fmt.Print("Username: ")
// 		fmt.Scan(&input.username)
// 		fmt.Print("Email: ")
// 		fmt.Scan(&input.email)
// 		fmt.Print("Password: ")
// 		fmt.Scan(&input.password)
// 		if input.username == "Admin" && input.email == "Admin" && input.password == "Admin" {
// 			break
// 		} else {
// 			fmt.Println("Data tidak sesuai")
// 			percobaan++
// 		}
// 		if percobaan == percobaanMaks {
// 			fmt.Println("================================")
// 			fmt.Println("MOHON MAAF DATA KAMU SALAH")
// 			fmt.Println("================================")
// 			fmt.Println("SILAHKAN TUNGGU DALAM WAKTU 1 MENIT")
// 			fmt.Println("================================")
// 			time.Sleep(time.Duration(durasi) * time.Second)
// 			login(input)
// 			return
// 		}
// 	}
// //}

// Function to Input Data
func inputDesa(T *tabDesa, n *int) {
	fmt.Println("Menu >> Input data")
	fmt.Println("================================")
	fmt.Println("SILAHKAN MASUKKAN DATA DESA")
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
		i++
	}
}

func inputPenduduk(P *tabPenduduk,t tabDesa, n *int){
	var namaDesa string
	fmt.Println("Menu >> Input Data Penduduk")
	fmt.Println("================================")
	fmt.Printf("%10s", "data penduduk")
	fmt.Println("================================")
	fmt.Print("Masukkan nama desa:")
	fmt.Scan(&namaDesa)
	for i := 0; i < *n; i++{
		if namaDesa == t[i].namaDesa{
			fmt.Printf("Masukkan penduduk ke-%d\n", i)
			fmt.Print("Nama: ")
			fmt.Scan(&P[i].namaPenduduk)
			fmt.Print("Umur: ")
			fmt.Scan(&P[i].umurPenduduk)
			fmt.Print("NIK: ")
			fmt.Scan(&P[i].noNIK)
			fmt.Print("Alamat Rumah: ")
			fmt.Scan(&P[i].alamatRumah)
			fmt.Print("Status Perkawinan: (Sudah/Belum)")
			fmt.Scan(&P[i].statusPerkawinan)
		}
		
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
			if T[i].namaPenduduk == namaDicari {
				fmt.Printf("Data Ditemukan: %+v\n", T[i].namaPenduduk)
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
				fmt.Printf("Data Ditemukan: %+v\n", T[i].noNIK)
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
func cetakData(T tabDesa, p tabPenduduk, n int) {
	fmt.Println("================================")
	fmt.Println("DATA PENDUDUK DESA")
	fmt.Println("================================")
	for i := 0; i < n; i++ {
		fmt.Printf("Data Desa Ke-%d\n", i+1)
		fmt.Printf("Nama Desa: %s\n", T[i].namaDesa)
		fmt.Printf("Alamat Desa: %s\n", T[i].alamatDesa)
		fmt.Printf("Jumlah RT: %d\n", T[i].jumlahRT)
		fmt.Printf("Jumlah RW: %d\n", T[i].jumlahRW)
		fmt.Printf("Nama Penduduk: %s\n", p[i].namaPenduduk)
		fmt.Printf("Alamat Penduduk: %s\n", p[i].alamatRumah)
		fmt.Printf("Umur: %d\n", p[i].umurPenduduk)
		fmt.Printf("RT: %d\n", p[i].noRT)
		fmt.Printf("RW: %d\n", p[i].noRW)
		fmt.Printf("Nomor NIK: %d\n", p[i].noNIK)
		fmt.Printf("Status Kawin: %s\n", p[i].statusPerkawinan)
		fmt.Println("================================")
	}
}

func clear() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}