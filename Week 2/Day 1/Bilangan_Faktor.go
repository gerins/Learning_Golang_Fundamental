package main

import (
	"fmt"
)

func main() {
	var input int = 99
	cariFaktor(input)

}

func cariFaktor(number int) {
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
		}
	}
}
