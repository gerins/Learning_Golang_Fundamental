package main

import (
	"fmt"
)

type interfaceKosong interface{}

type buku struct {
	judul string
}

func (a buku) cetakJudul() string {
	return `Judul bukunya adalah : ` + a.judul
}

type mobil struct {
	merek string
}

func (a mobil) cetakMerk() string {
	return `Merk mobilnya adalah : ` + a.merek
}

type motor struct {
	merek string
}

func (a motor) cetakMerk() string {
	return `Merk motor adalah : ` + a.merek
}

func main() {
	var bukuApa = buku{"Coding"}
	var merkMobil = mobil{"Honda"}
	var merkMotor = motor{"Yamaha"}

	cetak(bukuApa)
	cetak(merkMobil)
	cetak(merkMotor)
}

func cetak(a interfaceKosong) {

	switch val := a.(type) {
	case buku:
		fmt.Println(val.cetakJudul())
	case mobil:
		fmt.Println(val.cetakMerk())
	default:
		fmt.Println(`Structur`, a, `tidak ada`)
	}

}
