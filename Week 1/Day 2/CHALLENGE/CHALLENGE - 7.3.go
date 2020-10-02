// Nama : Garin Prakoso

package main

import (
	"fmt"
	"sort"
)

func main() {
	olehOleh(30000, []int{15000, 12000, 5000, 3000, 10000})
	olehOleh(10000, []int{2000, 2000, 3000, 1000, 2000, 10000})
	olehOleh(4000, []int{7500, 1500, 2000, 3000})
	olehOleh(50000, []int{25000, 25000, 10000, 15000})
	olehOleh(0, []int{10000, 3000})
}

func olehOleh(saldo int, barang []int) {
	var counter int
	var uang int = saldo
	sort.Ints(barang)
	for _, harga := range barang {
		if saldo-harga >= 0 {
			counter++
			saldo -= harga
		}
	}
	fmt.Println(`Dengan uang sebesar Rp`, uang, `barang yang dapat dibeli sebanyak`, counter, `buah`)
}
