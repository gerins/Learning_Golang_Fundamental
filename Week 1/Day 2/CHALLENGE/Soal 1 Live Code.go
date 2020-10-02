// Nama : Garin Prakoso

package main

import "fmt"

func main() {
	var pertama, akhir int

	fmt.Print("Masukan Angka Pertama : ")
	fmt.Scan(&pertama)

	fmt.Print("Masukan Angka Terakhir : ")
	fmt.Scan(&akhir)

	if pertama%2 == 0 && akhir%2 != 0 {
		fmt.Println(`One Zombie Down!`)
	} else if pertama%2 == 0 && akhir%2 == 0 {
		fmt.Println(`Try Again To Attack`)
	} else if pertama%2 != 0 && akhir%2 != 0 {
		fmt.Println(`Try Again To Attack`)
	} else if pertama%2 != 0 && akhir%2 == 0 {
		fmt.Println(`You Are Dead!`)
	}
}
