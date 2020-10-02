package main

import "fmt"

func main() {
	for i := 1; i < 1000; i++ {
		fmt.Println(cariHasil(i))
	}
}

func cariHasil(input int) string {
	switch {
	case input%2 != 0 && input%3 != 0:
		return "SALAH"
	case input%2 == 0 && input%3 == 0:
		return "TESLA"
	case input%2 == 0:
		return "TES"
	case input%3 == 0:
		return "LA"
	}
	return ""
}
