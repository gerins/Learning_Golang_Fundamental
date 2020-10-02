// Nama : Garin Prakoso
package main

import "fmt"

func main() {
	var panjang, lebar, tinggi int
	fmt.Println("Masukan Nilai Panjang Balok : ")
	fmt.Scan(&panjang)
	fmt.Println("Masukan Nilai Lebar Balok : ")
	fmt.Scan(&lebar)
	fmt.Println("Masukan Nilai Tinggi Balok : ")
	fmt.Scan(&tinggi)

	var volume = panjang * lebar * tinggi
	var permukaan = ((panjang * lebar) + (panjang * tinggi) + (lebar * tinggi)) * 2

	fmt.Println(`Volume balok adalah : `, volume)
	fmt.Println(`Luas permukaan adalah : `, permukaan)
}
