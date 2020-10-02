package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(Factorial(uint64(i)))
	}
}

func Factorial(n uint64) uint64 {
	if n == 1 {
		return n
	}
	return n * Factorial(n-1)
}
