package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	log.Print("Logging dimulai")

	var awal, akhir int

	fmt.Print("Masukan angka awal :")
	fmt.Scan(&awal)
	fmt.Print("Masukan angka akhir :")
	fmt.Scan(&akhir)

	hasil, err := deret(awal, akhir) // tampung hasil function

	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(hasil)
		fmt.Println(`average :`, hitungAverage(hasil))
	}
}

func deret(awalan, akhiran int) ([]int, error) {
	var hasil []int

	if awalan > akhiran {
		return hasil, errors.New("Tidak dapat menjalankan")
	}

	for i := awalan; i <= akhiran; i++ {
		hasil = append(hasil, i)
	}

	return hasil, nil
}

func hitungAverage(array []int) int {
	tampung := 0
	for _, isinya := range array {
		tampung += isinya
	}
	return tampung / len(array)
}
