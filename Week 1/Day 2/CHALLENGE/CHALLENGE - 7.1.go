// Nama : Garin Prakoso

package main

import "fmt"

var Numbs = []int{3, 5, 2, -4, 8, 11}
var sum int = 7
var values = [][]int{}

func main() {
	hitung(sum, Numbs)
}

func hitung(sum int, number []int) {
	for i, _ := range number {
		for a := 1; a < len(number); a++ {
			if number[i]+number[a] == sum {
				row1 := []int{number[i], number[a]}
				values = append(values, row1)
				number[i] = 0
				number[a] = 0
			}
		}
	}
	fmt.Println(values)
}
