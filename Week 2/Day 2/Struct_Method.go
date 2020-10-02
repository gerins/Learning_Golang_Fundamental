package main

import (
	"fmt"
)

type Siswa struct {
	nama string
	umur int
}

func (n *Siswa) tambahUmur() {
	n.umur++
}

func (n *Siswa) bagiUmur() {
	n.umur /= 2
}

func main() {
	var gerin = Siswa{`Gerin`, 25}

	// var azza Siswa
	// azza = Siswa{`Azza`, 25}

	azza := Siswa{`Azza`, 25} // di deklarasikan pake :=
	fmt.Println(gerin, azza)

	gerin.umur = 29
	azza.umur = 27
	fmt.Println(gerin, azza)

	gerin.tambahUmur()
	azza.tambahUmur()
	fmt.Println(gerin, azza)

	gerin.bagiUmur()
	azza.bagiUmur()
	fmt.Println(gerin, azza, `Setelah menjalankan fungsi bagi`)

	cetak(gerin)

}

func cetak(n Siswa) {
	fmt.Println(`cetak pake function`, n)
}
