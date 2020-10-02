package main

import "fmt"

func main() {
	fmt.Println(bayarHutang(1000, 10000, false, "apin"))
}

func bayarHutang(uangKita, bebanhidup int, status bool, nama string) (string, string, int) {
	var a int = bebanhidup

	if status == true {
		a = bebanhidup - uangKita
	} else {
		peringatan = "tolong bayar hutang kamu coeg    " + nama
	}

	return peringatan, "Sisa hutang anda :", a // return 9000
}

var peringatan string
