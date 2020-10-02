package main

import (
	"fmt"
	"sort"
)

func main() {
	var array1 = []int{1, 2, 3, 4, 6, 7, 2, 123, 32, 123, 213, 5}

	hasil := cariMax(array1)
	fmt.Println(`Min:`, hasil[0], `Max :`, hasil[1], `Average :`, hasil[2])
}

func cariMax(array []int) []int {
	sort.Ints(array)
	var simpan int
	for _, isi := range array {
		simpan += isi
	}
	average := simpan / len(array)

	hasil := []int{array[0], array[len(array)-1], average}
	return hasil
}
