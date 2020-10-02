//[CHALLENGE - 13.3] Palindrome AngkaAssignment
// 13 Juni 2020
// Nama : Garin Prakoso

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var hasil = new([]int)
	cariPalidrome(9876543123123, hasil)
	fmt.Println(*hasil)
}

func cariPalidrome(value int, hasil *[]int) {
	var cekValue []string
	if value < 11 {
		value++
		*hasil = append(*hasil, value)
	} else {
		for {
			var cekNumber bool = true
			cekValue = strings.Split(strconv.Itoa(value), "")
			for i := 0; i < len(cekValue)/2; i++ {
				if cekValue[i] != cekValue[len(cekValue)-i-1] {
					cekNumber = false
				}
			}
			if !cekNumber {
				value++
			} else {
				break
			}
		}
		for _, isi := range cekValue {
			n, _ := strconv.Atoi(isi)
			*hasil = append(*hasil, n)
		}
	}
}
