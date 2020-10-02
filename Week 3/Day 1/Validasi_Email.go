package main

import (
	"fmt"
	"strings"
)

var harusMengandung = []string{"@", ".", "com"}
var tidakMengandung = []string{"#", "$"}

func main() {
	var inputEmail string
	fmt.Println("Silahkan Masukan Email :")
	fmt.Scan(&inputEmail)
	cekEmail(inputEmail, harusMengandung, tidakMengandung)
}

func cekEmail(input string, cek, cek2 []string) {
	var val1 bool = true
	var val2 bool = true
	for _, value := range cek {
		if cekEmail := strings.Contains(input, value); !cekEmail {
			fmt.Println(`Email harus mengandung`, value)
			val1 = false
		}
	}
	for _, value := range cek2 {
		if cekEmail2 := strings.Contains(input, value); cekEmail2 {
			fmt.Println(`Email tidak boleh mengandung`, value)
			val2 = false
		}
	}
	if val1 && val2 {
		fmt.Println("Email Sudah Benar")
	}
}
