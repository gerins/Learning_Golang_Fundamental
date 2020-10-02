// Nama : Garin Prakoso

package main

import "fmt"

func main() {
	var kumpulanUmur = []int{15, 25, 30, 100, 200, 300, 40, 50}
	fmt.Println(kumpulanUmur, `tampilkan array slice`)
	fmt.Println(len(kumpulanUmur), `Panjang index`)
	fmt.Println(kumpulanUmur[0], `index ke 0`)
	fmt.Println(kumpulanUmur[7], `index ke 7`)

	kumpulanUmur[7] = 12 // mengubah index ke 7
	fmt.Println(kumpulanUmur[7], `index ke 7 setelah di ubah`)

	kumpulanUmur = append(kumpulanUmur, 2030, 2020, 2340)
	fmt.Println(kumpulanUmur, `setelah di append`)

	fmt.Println("/////////////////////////////////////////////////////////////////////////////////////////")

	kumpulanNomor := make([]string, 5) // membuat slices type string menggunakan make dengan panjang index 3
	kumpulanNomor = append(kumpulanNomor, "satu", "dua")
	fmt.Println(`kumpulan nomor`, kumpulanNomor)

	fmt.Println("/////////////////////////////////////////////////////////////////////////////////////////")

	var array1 = [...]string{"satu", "dua", "tiga"}
	fmt.Println(array1, `array`)
	var slice1 = array1[:] // merubah array menjadi slice dengan append [:] = dari ujung ke ujung
	// || [1:] = dari index 1 sampai ke ujung kanan

	slice1 = append(slice1, "empat")
	slice1 = append(slice1, "lima")
	fmt.Println(slice1, `slice setelah ditambah append`)
	fmt.Println(cap(slice1), `kapasitas slice`)
	fmt.Println(len(slice1), `panjang slice`)
}
