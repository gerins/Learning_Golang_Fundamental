// Nama : Garin Prakoso

package main

import "fmt"

func main() {
	var value = []int{1, 2, 3, 4, 5}
	var target int = 3

	fmt.Println(totalPasangan(value, target))
}

func totalPasangan(value []int, target int) int {
	var counter int
	for index, isi := range value {
		for i := index + 1; i < len(value); i++ {
			if isi+value[i] == target {
				counter++
			}
		}
	}
	return counter
}
