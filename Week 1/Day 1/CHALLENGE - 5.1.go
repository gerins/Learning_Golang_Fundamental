// Nama : Garin Prakoso

package main

import "fmt"

func main() {
	var nama string = "saitama"
	var peran string = "superhero"

	if nama != "" && peran == "superhero" {
		fmt.Println("Selamat Datang Superhero", nama, "Kalahkan Semua Monster Di Muka Bumi")
	} else if nama != "" && peran == "monster" {
		fmt.Println("Selamat Datang Monster", nama, "Hancurkan Semua Superhero Yang Ada")
	} else if nama != "" && peran != "superhero" && peran != "monster" {
		fmt.Println("Selamat Datang", nama, "Pilih Peranmu Untuk Melanjutkan Game Ini")
	} else if nama != "" && peran == "" {
		fmt.Println("Peran Harus Di Isi")
	} else {
		fmt.Println("Nama dan Peran Harus Di Isi")
	}
}
