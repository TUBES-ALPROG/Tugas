// package main

// import (
//     "fmt"
//     "time"
// )

// const NMAX = 100

// type tabDesa [NMAX]dataDesa

// type dataDesa struct {
//     namaDesa   string
//     alamatDesa string
//     jumlahRT   int
//     jumlahRW   int
//     penduduk   [NMAX]dataPenduduk // Penambahan slice untuk menyimpan penduduk
// }

// type dataPenduduk struct {
//     namaPenduduk     string
//     umurPenduduk     int
//     alamatRumah      string
//     noRT             int
//     noRW             int
//     noNIK            int
//     statusPerkawinan string
// }

// 	type Login struct {
// 		username string
// 		email    string
// 		password string
// 	}

// // ----------KAMUS GLOBAL----------
// var dataD tabDesa
// var nDesa int
// var nPenduduk int
// var shouldExit bool
// var namaDicari string
// var pilih, nomorNIK int

// func main() {
// 	var input Login
// 	login(&input)
// 	loading()
//     menu()
// }

// // Tampilan Menu Utama
// func menu() {
//     for !shouldExit {
//         fmt.Println("================================")
//         fmt.Printf("%18s\n", "SI DESA")
//         fmt.Println(" Aplikasi Sistem Informasi Desa")
//         fmt.Println("================================")
//         fmt.Println("Pilih menu:")
//         fmt.Println("1. Input Data")
//         fmt.Println("2. Cari Data")
//         fmt.Println("3. Hapus Data")
//         fmt.Println("4. Edit Data")
//         fmt.Println("5. Cetak Data")
//         fmt.Println("6. Exit")
//         fmt.Println()
//         fmt.Print("Pilih:")
//         fmt.Scan(&pilih)
//         fmt.Println("================================")
//         fmt.Println()

//         switch pilih {
//         case 1:
//             inputData(&dataD, &nDesa, &nPenduduk)
//         case 2:
//             cariData(dataD, nDesa, nPenduduk)
//         // case 3:
//         //  Implement Hapus Data
//         // case 4:
//         //  Implement Edit Data
//         case 5:
//             cetakData(dataD, nDesa ,nPenduduk)
//         case 6:
//             shouldExit = true
//             fmt.Println("================================")
//             fmt.Println("TERIMA KASIH")
//             fmt.Println("================================")
//         }
//     }
// }

// // -------LOADING---------
// func loading() {
//     fmt.Println("Loading...")

//     frames := []string{
//         "⠋",
//         "⠙",
//         "⠹",
//         "⠸",
//         "⠼",
//         "⠴",
//         "⠦",
//         "⠧",
//         "⠇",
//         "⠏",
//     }

//     for i := 0; i < 10; i++ {
//         fmt.Printf("\r%s Processing... ", frames[i])
//         time.Sleep(200 * time.Millisecond)
//     }

//     fmt.Println("\nDone!")
// }


// // -------FUNGSI UNTUK LOGIN---------
// func login(input *Login) {
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
// }

// // -------FUNGSI UNTUK MASUKKAN DATA DESA & PENDUDUK---------
// func inputData(K *tabDesa, nDesa, nPendudukDesa *int) {
// 	fmt.Println("================================")
// 	fmt.Println("SILAHKAN MASUKKAN DATA PENDUDUK")
// 	fmt.Println("================================")
// 	fmt.Print("Masukkan jumlah Desa: ")
// 	fmt.Scan(nDesa)

// 	for i := 0; i < *nDesa; i++ {
// 		fmt.Printf("Data Desa Ke-%d\n", i+1)
// 		fmt.Print("Nama Desa: ")
// 		fmt.Scan(&(*K)[i].namaDesa)
// 		fmt.Print("Alamat Desa: ")
// 		fmt.Scan(&(*K)[i].alamatDesa)
// 		fmt.Print("Jumlah RT: ")
// 		fmt.Scan(&(*K)[i].jumlahRT)
// 		fmt.Print("Jumlah RW: ")
// 		fmt.Scan(&(*K)[i].jumlahRW)

// 		fmt.Print("Masukkan jumlah penduduk: ")
// 		fmt.Scan(&(*nPendudukDesa))

// 		for j := 0; j < *nPendudukDesa; j++ {
// 			fmt.Printf("Data Penduduk Ke-%d Desa %s\n", j+1, (*K)[i].namaDesa)
// 			fmt.Print("Nama Penduduk: ")
// 			fmt.Scan(&(*K)[i].penduduk[j].namaPenduduk)
// 			fmt.Print("Alamat Penduduk: ")
// 			fmt.Scan(&(*K)[i].penduduk[j].alamatRumah)
// 			fmt.Print("Umur: ")
// 			fmt.Scan(&(*K)[i].penduduk[j].umurPenduduk)
// 			fmt.Print("RT: ")
// 			fmt.Scan(&(*K)[i].penduduk[j].noRT)
// 			fmt.Print("RW: ")
// 			fmt.Scan(&(*K)[i].penduduk[j].noRW)
// 			fmt.Print("Masukkan Nomor NIK: ")
// 			fmt.Scan(&(*K)[i].penduduk[j].noNIK)
// 			fmt.Print("Status Kawin?: ")
// 			fmt.Scan(&(*K)[i].penduduk[j].statusPerkawinan)
// 		}
// 	}
// }

// // -------FUNGSI UNTUK MENCETAK DATA DESA & PENDUDUK---------
// func cetakData(K tabDesa, nDesa, nPenduduk int) {
//     fmt.Println("================================")
//     fmt.Println("CETAK DATA PENDUDUK DESA")
//     fmt.Println("================================")

//     for i := 0; i < nDesa; i++ {
//         fmt.Printf("Data Desa ke-%d\n", i+1)
//         fmt.Printf("Nama Desa: %s\n", K[i].namaDesa)
//         fmt.Printf("Alamat Desa: %s\n", K[i].alamatDesa)
//         fmt.Printf("Jumlah RT: %d\n", K[i].jumlahRT)
//         fmt.Printf("Jumlah RW: %d\n", K[i].jumlahRW)
//         fmt.Println("================================")

//         fmt.Printf("%-20s%-20s%-10s%-5s%-5s%-15s%-20s\n", "Nama Penduduk", "Umur", "Alamat", "RT", "RW", "NIK", "Status Perkawinan")
//         for j := 0; j < nPenduduk; j++ {
//             fmt.Printf("%-20s%-20s%-10d%-5d%-5d%-15d%-20s\n", K[i].penduduk[j].namaPenduduk, K[i].penduduk[j].alamatRumah, K[i].penduduk[j].umurPenduduk, K[i].penduduk[j].noRT, K[i].penduduk[j].noRW, K[i].penduduk[j].noNIK, K[i].penduduk[j].statusPerkawinan)
//         }
//         fmt.Println("================================")
//     }
// }


// // -------FUNGSI UNTUK MENCARI PENDUDUK---------
// func cariData(K tabDesa, nDesa, nPenduduk int) {
//     fmt.Println("================================")
//     fmt.Println("PENCARIAN DATA PENDUDUK DESA")
//     fmt.Println("================================")
//     fmt.Println("MENU PILIHAN PENCARIAN")
//     fmt.Println("1. Nama")
//     fmt.Println("2. NIK")
//     fmt.Println("3. EXIT")
//     fmt.Print("Pilih: ")
//     fmt.Scan(&pilih)

//     var searchField string
//     switch pilih {
//     case 1:
//         searchField = "Nama"
//     case 2:
//         searchField = "NIK"
//     case 3:
//         return
//     }

//     fmt.Println("================================")
//     fmt.Printf("Masukkan %s: ", searchField)
//     var searchQuery string
//     fmt.Scan(&searchQuery)

//     found := false
//     for i := 0; i < nDesa; i++ {
//         for j := 0; j < K[i].jumlahRT; j++ {
//             for k := 0; k < K[i].jumlahRW; k++ {
//                 var data dataPenduduk
//                 if pilih == 1 {
//                     data = K[i].penduduk[j*K[i].jumlahRW+k]
//                 } else {
//                     data = K[i].penduduk[j*K[i].jumlahRW+k]
//                 }

//                 var match bool
//                 if pilih == 1 {
//                     match = data.namaPenduduk == searchQuery
//                 } else {
//                     match = data.noNIK == nomorNIK
//                 }

//                 if match {
//                     fmt.Printf("Data Ditemukan: %+v\n", data)
//                     found = true
//                 }
//             }
//         }
//     }
//     if !found {
//         fmt.Println("Data Tidak Ada")
//     }
// }
