package main

import (
	"fmt"
)

func main() {
	cetakBanyak("budi", 55, 1, 2, 3, 4, 5)
}

func cetakBanyak(nama string, umur int, s ...int) {
	fmt.Println(s)
	for index, isinya := range s {
		fmt.Println(index, isinya)
	}
}
