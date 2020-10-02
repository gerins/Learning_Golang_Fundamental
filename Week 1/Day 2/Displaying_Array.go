// // Nama : Garin Prakoso

// package main

// import "fmt"

// func main() {

// 	var arraySatu = []int{}

// 	for i := 1; i <= 1024; i++ {
// 		arraySatu = append(arraySatu, i)
// 		fmt.Println(len(arraySatu), `dan kapasitasnya adalah`, cap(arraySatu))
// 	}
// 	fmt.Println(arraySatu)

// }

////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	var arraySatu = []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(`Dari Depan`)
	for i := 0; i < len(arraySatu); i++ {
		fmt.Println(arraySatu[i])
	}

	fmt.Println(`Dari belakang`)
	for i := len(arraySatu) - 1; i >= 0; i-- {
		fmt.Println(arraySatu[i])
	}

	fmt.Println(`Menggunakan range`)
	for i, v := range arraySatu {
		fmt.Println("Index ke ", i, "=>", v)
		// fmt.Println(arraySatu[len(arraySatu)-i-1]) // daribelakang
	}

	fmt.Println("Array Multi Dimensi")
	var angka = [][]int{{0, 0}, {2, 2}, {3, 4}, {5, 6}, {7, 8}}
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			// fmt.Printf("Data ke [%d][%d] = %d\n", i, j, angka[i][j])
			fmt.Println(`data ke `, i, `dan ke`, j, `adalah`, angka[i][j])
		}
	}
	fmt.Println(angka)
}

////////////////////////////////////////////////////////////////////////////////////////////////
