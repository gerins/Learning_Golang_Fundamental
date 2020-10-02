// Nama : Garin Prakoso

package main

import "fmt"

func main() {
	// Array
	var kumpulanNama = [3]string{"kucing", "monyet", "macan"} // array string
	fmt.Println(kumpulanNama)

	z := [5]int{1: 10, 3: 30} // mengisi array pada index 1 dan 3
	fmt.Println(z)

	x := [...]int{30, 20, 40} // [...] nilai index menyesuaikan isi dari array
	fmt.Println(x)

	var y [5]int = [5]int{10, 20, 30} // mengisi array Y dengan array int 10 20 30, pengisian dimulai dari depan
	fmt.Println(y)

	x[0] = 100
	fmt.Println(x)
}
