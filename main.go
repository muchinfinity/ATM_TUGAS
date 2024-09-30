package main

import (
	"fmt"
	"os"
)

// Deklarasi Variabel
	var Username, Password string
	var Pilih int
	var SALDO = [1]int{3500000}
	var HistoriTransaksi = []string{}
	var Jumlah int

// func Lihat Akun
	func LihatAkun() {
		fmt.Println("=== INFORMASI AKUN ===")
		fmt.Println("Nama          : Mutia Zahra")
		fmt.Println("NPM           : 2406359430")
		fmt.Println("Username      : Mutia")
		fmt.Println("Jenis Kelamin : Perempuan")
		fmt.Println("Alamat        : Asrama Mahasiswa UI Depok")
		fmt.Print("Klik ENTER untuk kembali ke Menu\n")
		fmt.Scanln()
	}

// func Lihat Saldo
	func LihatSaldo() {
		fmt.Println("=== SALDO AKUN ===")
		fmt.Println("Saldo Anda : ", SALDO[0])
		fmt.Print("Klik ENTER untuk kembali ke Menu\n")
		fmt.Scanln()
	}

// func Tambah Saldo
	func TambahSaldo() {
		fmt.Println("=== TAMBAH SALDO AKUN ===")
		fmt.Println("Masukkan Jumlah Saldo yang Ingin Anda Tambahkan : ")
		fmt.Scanln(&Jumlah)

		SALDO[0] += Jumlah 
		HistoriTransaksi = append(HistoriTransaksi, fmt.Sprintf("SALDO BERTAMBAH %d\n", Jumlah))
		fmt.Println("SALDO ANDA BERHASIL DITAMBAHKAN")
		fmt.Printf("Saldo sekarang :Rp %d\n", SALDO[0])
		fmt.Print("Klik ENTER untuk kembali ke Menu\n")
		fmt.Scanln()
	}

// func Tarik Saldo
	func TarikSaldo() {
		fmt.Println("=== TARIK SALDO AKUN ===")
		fmt.Println("Masukkan Jumlah Saldo yang Ingin Anda Tarik : ")
		fmt.Scanln(&Jumlah)
	if 	Jumlah < SALDO[0] && Jumlah > 0 {
		SALDO[0] -= Jumlah
		HistoriTransaksi = append(HistoriTransaksi, fmt.Sprintf("SALDO DITARIK %d\n", Jumlah))
		fmt.Printf("Saldo anda sekarang : %d\n", SALDO[0])
		fmt.Println("=== SALDO ANDA BERHASIL DITARIK ===")
		fmt.Print("klik ENTER untuk kembali ke menu\n")
		fmt.Scanln()
	} else {
		fmt.Println("=== TIDAK DAPAT MENARIK ===")
		fmt.Print("klik ENTER untuk kembali ke menu\n")
		fmt.Scanln()
	} 
}

// func History Transaksi
	func HistoryTransaksi() {
		fmt.Println("=== HISTORI TRANSAKSI ===")
	if len(HistoriTransaksi) == 0 {
		fmt.Println("Tidak ada riwayat transaksi")
		fmt.Println("=== Terima Kasih, Anda Telah Keluar ===")
	} else {
		for _, Transaksi := range HistoriTransaksi {
			fmt.Printf(Transaksi)
		}
	}
		fmt.Print("Klik ENTER untuk kembali ke Menu\n")
		fmt.Scanln()
	}


// func main
	func main() {

// Banner & Masukkan Username
		fmt.Print("============= SELAMAT DATANG =============\n")
		fmt.Print("================ BANK BNI ================\n")
		fmt.Print("Masukkan Username Anda lalu Klik ENTER untuk melanjutkan : ")
		fmt.Scanln(&Username)
	
// Kalo Username benar lanjut, kalo salah keluar
	if Username == "Mutia" {
		fmt.Print("Username Anda Benar, Silakan Masukkan Password Anda lalu Klik ENTER Untuk Melanjutkan : " )
		fmt.Scanln(&Password) 
	} else { 
		fmt.Print("Username Anda Salah ! ")
		os.Exit(1) 
	}
// Kalo Password benar Pilih Menu, kalo salah keluar
	if Password == "2406359430" {
		for {
			
		fmt.Print("======= MENU =======\n")
		fmt.Println("1. Lihat Akun")
		fmt.Println("2. Lihat Saldo")
		fmt.Println("3. Tambah Saldo")
		fmt.Println("4. Tarik Saldo")
		fmt.Println("5. Histori Transaksi")
		fmt.Println("6. Keluar")
		fmt.Print("Silakan Pilih Menu lalu klik ENTER untuk melanjutkan : ") 
		fmt.Scanln(&Pilih)
		
	
			switch Pilih {
			case 1:
				LihatAkun()
			case 2:
				LihatSaldo()
			case 3:
				TambahSaldo()
			case 4:
				TarikSaldo()
			case 5:
				HistoryTransaksi()
			case 6: 
				fmt.Println("=== Anda telah keluar, Terima Kasih ! ===")
				os.Exit(0)
			default:
				fmt.Println("PILIHAN TIDAK VALID! SILAKAN PILIH NOMOR SESUAI MENU")
			}
			}
	}  else {
		fmt.Print("Password Anda Salah ! ")
		os.Exit(1) 
	}
}