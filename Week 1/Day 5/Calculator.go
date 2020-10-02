package main

import (
	"fmt"
)

func main() {
	calculator(5, 10, "*")
}

func calculator(a, b int, operator string) {
	var c int
	switch operator {
	case "+":
		c = b + a
	case "-":
		c = b - a
	case "/":
		c = b / a
	case "*":
		c = b * a
	}
	cetak(c)
}

func cetak(a int) {
	fmt.Println(a)
	// fmt.Println(calculator(5, 5, "*"))
}
