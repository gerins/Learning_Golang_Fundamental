// Nama : Garin Prakoso

package main

import "fmt"

func main() {
	var input, ganjil, genap int
	fmt.Print("Masukan nilai input = ")
	fmt.Scan(&input) // masukan input

	fmt.Print(`Ganjil : `)
	for i := 1; i < input; i++ { // perulangan i sebanyak input kalau input 5 diiulang 5 kali
		if i%2 != 0 { // kalau bukan 0 (!=) berarti ganjil
			if i != 1 {
				fmt.Print(` + `)
			}
			fmt.Print(i)
			ganjil += i // 0 + 1  = 1    1 + 3 = 4     4+5 = 9
		}
	}
	fmt.Println(` = `, ganjil)

	fmt.Print(`Genap : `)
	for c := 1; c <= input*2; c++ { // input nya 5 berarti perulangan 10 x
		if c%2 == 0 {
			if c != 2 {
				fmt.Print(` + `)
			}
			fmt.Print(c)
			genap += c
		}
	}
	fmt.Println(` = `, genap)
}
