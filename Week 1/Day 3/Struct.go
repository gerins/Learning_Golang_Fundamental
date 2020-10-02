package main

import (
	"fmt"
)

type peserta struct {
	nama string
	asal string
	umur int
}

func main() {
	var inputNama, inputAsal string
	var inputUmur int
	for {
		fmt.Print("Masukan Nama :")
		fmt.Scan(&inputNama)
		fmt.Print("Masukan Asal :")
		fmt.Scan(&inputAsal)
		fmt.Print("Masukan Umur :")
		fmt.Scan(&inputUmur)

		tambahPeserta(inputNama, inputAsal, inputUmur)
	}
}

func tambahPeserta(name string, from string, age int) {
	x := peserta{nama: name, asal: from, umur: age}
	fmt.Println(`Peserta berhasil ditambahkan`)
	fmt.Println(` Nama :`, x.nama, "\n", `Asal :`, x.asal, "\n", `Umur :`, x.umur)
}
