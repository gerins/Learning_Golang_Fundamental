package main

import (
	"fmt"
)

func main() {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(all)
	all = RemoveIndex(all, 5)
	fmt.Println(all)

}

func RemoveIndex(s []int, index int) []int {
	arr1 := []int{16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	fmt.Println(s[:index])
	fmt.Println(s[index+1:])
	return append(s[:index], arr1...)
}
