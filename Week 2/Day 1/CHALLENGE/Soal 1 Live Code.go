// Nama : Garin Prakoso

package main

import "fmt"

var koin = []int{1, 5, 7, 9, 11}
var input int = 17
var values = [][]int{}

func main() {
	hasil, cek := hitung(input, koin)
	if cek {
		fmt.Println(hasil)
	} else {
		fmt.Println(`Input diatas 30`)
	}
}

func hitung(sum int, number []int) ([][]int, bool) {
	if sum > 30 {
		return values, false
	}

	for i, _ := range number {
		for a := 1; a < len(number); a++ {
			for c := 2; c < len(number); c++ {
				var kombinasi2 bool = false
				if number[i]+number[a] == sum {
					row1 := []int{number[i], number[a]}
					values = append(values, row1)
					kombinasi2 = true
					number[i] = 0
					number[a] = 0
					// fmt.Println(`masuk kombinasi 2`)
				}
				if !kombinasi2 && number[i]+number[a]+number[c] == sum {
					row2 := []int{number[i], number[a], number[c]}
					values = append(values, row2)
					// fmt.Println(`masuk kombinasi 3`)
				}
			}
		}
	}
	return values, true
}
