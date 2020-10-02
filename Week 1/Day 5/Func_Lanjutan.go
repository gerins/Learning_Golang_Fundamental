package main

import (
	"fmt"
)

func main() {
	angka := 10
	angka2 := 5
	cetak(jumlah(angka, angka2))
	func() {
		var jumlah int
		jumlah = angka + angka2
		fmt.Println(jumlah)
	}()
	cetakBanyak(1, 2, 3, 4, 5)
}

var jumlah = func(a, b int) int {
	fmt.Println(`Masuk jumlah`)
	return a * b
}

func cetak(ops int) {
	fmt.Println(`cetak`)
	fmt.Println(ops)
	fmt.Println(jumlah(ops, 5))
}

func cetakBanyak(s ...int) {
	for _, isinya := range s {
		defer fmt.Println(isinya)
	}
}
