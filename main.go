package main

import (
	"fmt"
	"time"
)

const NMAX = 5000

type tabPenduduk[NMAX] dataDesa

type dataDesa struct{
	namaDesa string
	alamatDesa string
	jumlahRT int
	jumlahRW int
	penduduk dataPenduduk
}

type dataPenduduk struct{
	namaPenduduk string
	umurPenduduk string
	alamatRumah string
	noRT int
	noRW int
	noNIK int
	statusPerkawinan string
}

type Login struct{
	username string 
	email string
	password string
}

var data tabPenduduk
var n int
var shouldExit bool

func main() {
	var input Login
	var pilih int
	login(&input)
	loading()
	menu(pilih)
}

//Tampilan Menu Utama
func menu(pilihMenu int) {

	for !shouldExit{
		fmt.Println("================================")
	    fmt.Printf("%18s\n","SI DESA")
		fmt.Println(" Aplikasi Sistem Informasi Desa")
		fmt.Println("================================")
		fmt.Println("Pilih menu:")
		fmt.Println("1. Input Data")
		fmt.Println("2. Hapus Data")
		fmt.Println("3. Ubah Data")
		fmt.Println("4. Exit")
		fmt.Println()
		fmt.Print("Pilih:")
		fmt.Scan(&pilihMenu)
		fmt.Println("================================")
		fmt.Println()

		switch pilihMenu{
		case 1:
			inputData(&data, n)
		case 4:
			shouldExit = true
			fmt.Println("================================")
			fmt.Println("TERIMA KASIH")
			fmt.Println("================================")
		}
	}

}

func loading(){
	for i := 0; i < 10; i++{
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

	for percobaan <= percobaanMaks{
		fmt.Print("Username: ",)
		fmt.Scan(&input.username)
		fmt.Print("Email: ",)
	    fmt.Scan( &input.email)
		fmt.Print("Password: ")
	    fmt.Scan(&input.password)
		if input.username == "Admin" && input.email == "Admin" && input.password == "Admin"{
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

	func inputData(T *tabPenduduk, n int){
		fmt.Println("================================")
		fmt.Println("================================")
		fmt.Println("SILAHKAN MASUKKAN DATA PENDUDUK")
		fmt.Println("================================")
		fmt.Println("Masukkan jumlah data:")
		fmt.Scan(&n)
		i := 0
		data := 1
		for i < n {
			fmt.Printf("Data Desa Ke-%d", data)
			fmt.Println()
			fmt.Print("Nama Desa:")
			fmt.Scan(&T[i].namaDesa)
			fmt.Print("Alamat Desa:")
			fmt.Scan(&T[i].alamatDesa)
			fmt.Print("Jumlah RT:")
			fmt.Scan(&T[i].jumlahRT)
			fmt.Print("Jumlah RW:")
			fmt.Scan(&T[i].jumlahRW)
			fmt.Print("Nama Penduduk:")
			fmt.Scan(&T[i].penduduk.namaPenduduk)
			fmt.Print("Alamat Penduduk:")
			fmt.Scan(&T[i].penduduk.alamatRumah)
			fmt.Print("Umur:")
			fmt.Scan(&T[i].penduduk.umurPenduduk)
			fmt.Print("RT:")
			fmt.Scan(&T[i].penduduk.noRT)
			fmt.Print("RW:")
			fmt.Scan(&T[i].penduduk.noRW)
			fmt.Print("Masukkan Nomor NIK:")
			fmt.Scan(&T[i].penduduk.noNIK)
			fmt.Print("Status Kawin?:")
			fmt.Scan(&T[i].penduduk.statusPerkawinan)
			i++
			data++
		}

	}
	
	

	
