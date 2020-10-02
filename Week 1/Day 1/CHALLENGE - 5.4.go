// Nama : Garin Prakoso

package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string
	var cek1 bool = false
	var cek2 bool = false
	var cek3 bool = true
	var cek4 bool = true
	var cek5 bool = true

	fmt.Print("Masukan email : ")
	fmt.Scanln(&email)

	var emailDipecah = strings.Split(email, "")

	for i := 0; i < len(emailDipecah); i++ {
		if emailDipecah[i] != "@" { // tidak ada @
			cek1 = true
		}
		if emailDipecah[i] != "." { // tidak ada .
			cek2 = true
		}
		if emailDipecah[i] != " " { // tidak kosong
			cek3 = false

		}
		if emailDipecah[i] == "*" {
			cek4 = false

		}
		if emailDipecah[i] == "#" {
			cek5 = false
		}
	}

	if cek1 {
		fmt.Println(`Penulisan email harus mengandung "@"`)
	}
	if cek2 {
		fmt.Println(`Penulisan email harus mengandung "."`)
	}
	if cek3 {
		fmt.Println(`Penulisan email tidak boleh kosong`)
	}
	if !cek4 {
		fmt.Println(`Penulisan email tidak boleh mengandung "*"`)
	}
	if !cek5 {
		fmt.Println(`Penulisan email tidak boleh mengandung "#"`)
	}
}
