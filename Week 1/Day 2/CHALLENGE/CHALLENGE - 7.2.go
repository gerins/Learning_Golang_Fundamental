// Nama : Garin Prakoso

package main

import "fmt"

var array1 = []int{6, 2, 4, 10, 5}
var array2 = []int{3, 9, 10, 13, 1, 2}
var terbesar, terkecil int

func main() {
	hitung(array2)
}

func hitung(array []int) {
	for _, q := range array {
		if q > terbesar {
			terbesar = q
		} else if terkecil > q || terkecil == 0 {
			terkecil = q
		}
	}
	fmt.Println(`Terbesar :`, terbesar)
	fmt.Println(`Terkecil :`, terkecil)
}
