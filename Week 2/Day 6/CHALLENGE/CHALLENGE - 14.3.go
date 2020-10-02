//[CHALLENGE - 14.3] Kelompok Angka
// 13 Juni 2020
// Nama : Garin Prakoso

package main

import (
	"fmt"
)

func main() {
	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var hasil [][]int
	kelompokanAngka(input, &hasil)
	fmt.Println(hasil)
}

func kelompokanAngka(value []int, hasil *[][]int) {
	var angkaGanjil, angkaGenap, angkaKelipatan3 []int

	for _, isi := range value {
		if isi%3 == 0 {
			angkaKelipatan3 = append(angkaKelipatan3, isi)
		} else if isi%2 == 0 {
			angkaGenap = append(angkaGenap, isi)
		} else {
			angkaGanjil = append(angkaGanjil, isi)
		}
	}

	var kumpulanArray = [][]int{angkaGenap, angkaGanjil, angkaKelipatan3}
	*hasil = append(*hasil, kumpulanArray...)
}
