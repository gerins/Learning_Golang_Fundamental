//[CHALLENGE - 14.2] Deret Aritmatika
// 13 Juni 2020
// Nama : Garin Prakoso

package main

import "fmt"

func main() {
	var input = []int{2, 4, 6, 8, 10, 12}
	fmt.Println(cekKonsistensi(input))
}

func cekKonsistensi(value []int) bool {
	var cekAngka bool = true
	parameter := value[1] - value[0]
	for i := 1; i < len(value); i++ {
		if value[i-1]+parameter != value[i] {
			cekAngka = false
		}
	}

	return cekAngka
}
