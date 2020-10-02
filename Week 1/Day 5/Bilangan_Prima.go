package main

import (
	"fmt"
	"math"
)

func main() {
	inputPrima()
}

func inputPrima() {
	var inputMin, inputMax int
	fmt.Print("Masukan nilai Minimum :")
	fmt.Scan(&inputMin)
	fmt.Print("Masukan nilai Maximum :")
	fmt.Scan(&inputMax)
	fmt.Println(cariPrime(inputMin, inputMax))
}

func cariPrime(min, max int) []int {
	var bilanganPrima []int // penampung
	for prime := min; prime <= max; prime++ {
		var sqrt int = int(math.Sqrt(float64(prime)))
		var cek bool = true          // buat cek, kalo tetep true berarti prima
		for i := 2; i <= sqrt; i++ { // pengulangan pengecekan
			if prime%i == 0 { // prime di mod i kalo 0 berarti bukan prima
				cek = false // ubah ke false biar di line 33 ngga ke append
				break       // hentikan loop, balik ke prime ++ ( line 23)
			}
		}
		if cek == true {
			bilanganPrima = append(bilanganPrima, prime)
		}
	}
	return bilanganPrima
}
