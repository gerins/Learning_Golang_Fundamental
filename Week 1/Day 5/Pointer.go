package main

import (
	"fmt"
)

func main() {
	var angka = []int{}
	// var alamat *[]int

	// alamat = &angka
	// fmt.Println(alamat)
	fmt.Println(`Mencoba print alamat :`, &angka)
	cetak(&angka)
	fmt.Println(`setelah diubah :`, angka)
}

func cetak(apa *[]int) {
	*apa = append(*apa, 1, 2, 3)
	fmt.Println(`berhasil diubah ke :`, *apa)
}
