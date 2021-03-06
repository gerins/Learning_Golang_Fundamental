package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(fib(i))
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
