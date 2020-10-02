package main

import (
	"fmt"
)

type bangunRuang interface {
	hitungLuas()
	hitungKeliling()
}

type squareBox struct {
	panjang, lebar int
}

type triangle struct {
	alas, tinggi int
}

func (n squareBox) hitungLuas() {
	luas := n.panjang * n.lebar
	fmt.Println(`Luas kotak =`, luas)
}
func (n squareBox) hitungKeliling() {
	keliling := n.panjang * 4
	fmt.Println(`keliling kotak =`, keliling)
}

func (b triangle) hitungLuas() {
	luas := (b.alas / 2) * b.tinggi
	fmt.Println(`Luas segitiga =`, luas)
}
func (b triangle) hitungKeliling() {
	keliling := b.alas * 3
	fmt.Println(`keliling segitiga =`, keliling)
}

func main() {
	kotak := squareBox{panjang: 10, lebar: 10}
	segiTiga := triangle{alas: 10, tinggi: 10}

	hitungBangunRuang(kotak)
	hitungBangunRuang(segiTiga)
}

func hitungBangunRuang(n bangunRuang) {
	n.hitungLuas()
	n.hitungKeliling()
}
