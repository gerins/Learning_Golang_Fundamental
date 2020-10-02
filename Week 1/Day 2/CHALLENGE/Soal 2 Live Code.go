// Nama : Garin Prakoso

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var nilaiString string = "+4+3+7+1"
	var convertString = strings.Split(nilaiString, "") // Ubah String menjadi Array

	var total int

	for i, v := range convertString {
		if v == "-" {
			var a, _ = strconv.Atoi(convertString[i+1])
			total = total - a
		} else if v == "+" {
			var a, _ = strconv.Atoi(convertString[i+1])
			total = total + a
		}
		fmt.Println(`index ke [`, i, `] berisi =`, v)
	}
	fmt.Println(total)
}
