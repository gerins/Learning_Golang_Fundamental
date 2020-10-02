package main

import "fmt"

var array1 = []int{1, 2, 3, 4, 5, 6}

func main() {

	for kucing := 0; kucing < 5; kucing++ {
		array1 = append(array1, kucing)
	}
	fmt.Println(array1)

}

// ceklis semua file > preferences > setting >
//run in terminal
//save file before
//clear previous output
