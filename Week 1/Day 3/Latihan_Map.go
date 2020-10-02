package main

import "fmt"

func main() {
	var karyawan = map[int]string{
		13410107: "Garin Prakoso",
		13410106: "Azzalia",
		12345678: "Budi",
	}

	for index, key := range karyawan {
		fmt.Println(`NIK :`, index, `Nama :`, key)
	}

	if key, hasil := karyawan[13410107]; hasil { // find key di map
		fmt.Println(`Ditemukan dengan nama`, key)
	} else {
		fmt.Print(`Karyawan tidak ditemukan`)
	}
}
