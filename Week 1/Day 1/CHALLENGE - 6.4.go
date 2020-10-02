// Nama : Garin Prakoso

package main

import (
	"fmt"
	"strings"
)

func main() {
	var kata string

	fmt.Print("Masukan kata : ")
	fmt.Scan(&kata)

	var kataDipecah = strings.Split(kata, "")

	var kataDigabung string
	i := len(kataDipecah)
	for c := i - 1; c >= 0; c-- {
		kataDigabung += kataDipecah[c]
	}

	var palindrome bool
	if kataDigabung == kata {
		palindrome = true
	} else {
		palindrome = false
	}
	fmt.Println(`palindrome :`, palindrome)
}
