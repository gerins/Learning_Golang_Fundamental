package main

import (
	"fmt"
)

type Siswa struct {
	nama string
	uang int
}

func main() {
	var gerin = Siswa{`gerin`, 25000}
	var budi = Siswa{`budi`, 30000}
	// 	var gerin = daftarMahasiswa{`Gerin`, 25, `Psikologi`}
	// var budi = daftarMahasiswa{`Budi`, 30, `Matematika`}
	// var azza = daftarMahasiswa{`Azzalia`, 25, `Akutansi`}

	var daftar = []Siswa{gerin, budi} // buat slice of struct

	for i := 10000; i <= 50000; i += 10000 {
		daftar = append(daftar, Siswa{"nomor", i})

	}

	for _, isi := range daftar {
		// fmt.Println(`index ke : `, index, isi)
		if isi.uang > 30000 {
			fmt.Println(`ditemukan dengan uang diatas 30 ribu:`, isi)
		}
	}
}
