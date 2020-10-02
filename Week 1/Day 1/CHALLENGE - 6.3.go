// Nama : Garin Prakoso

package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukan nilai bilangan = ")
	fmt.Scan(&bilangan)

	for i := bilangan; i >= 1; i-- {
		if bilangan%i == 0 {
			defer fmt.Println(i)
		}
	}

}
