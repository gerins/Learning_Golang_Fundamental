package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// var susunanAngka string = "42#3 * 188 = 80#204"
	var susunanAngka2 string = "8#61 * 895 = 78410#5"

	var jawaban = new([]int)
	cariPola(susunanAngka2, jawaban)
	fmt.Print(*jawaban)
}

func cariPola(input string, hasil *[]int) {
	sliceString := strings.Split(strings.ReplaceAll(strings.ReplaceAll(input, "=", "&"), "*", "&"), " & ")
	var angka2, _ = strconv.Atoi(sliceString[1])
	for index, isi := range sliceString {
		fmt.Println(index, isi)
	}

	for i := 1; i < 10; i++ {
		var angka1, _ = strconv.Atoi(strings.ReplaceAll(sliceString[0], "#", strconv.Itoa(i)))
		for c := 1; c < 10; c++ {
			var angka3, _ = strconv.Atoi(strings.ReplaceAll(sliceString[2], "#", strconv.Itoa(c)))
			if angka1*angka2 == angka3 {
				*hasil = append(*hasil, i, c)
			}
		}
	}
}
