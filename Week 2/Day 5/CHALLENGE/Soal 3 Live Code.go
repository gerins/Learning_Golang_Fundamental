// Nama : Garin Prakoso

package main

import (
	"fmt"
	"os"
	"sort"
)

type member struct {
	id       string
	uang     int
	barang   []string
	sisaUang int
}

func main() {
	result := tokoEnigma("123132ssd666", 700000)
	fmt.Println(`id:`, result.id, `uang: Rp`, result.uang, `barang:`, result.barang, `sisaUang: Rp`, result.sisaUang)
}

func tokoEnigma(id string, uang int) member {
	if id == "" {
		fmt.Println("Mohon maaf, toko Enigma Jaya hanya berlaku untuk member saja")
		return member{}
	} else if uang < 50000 {
		fmt.Println("Mohon maaf, uang tidak cukup")
		return member{}
	}

	var kerangjangBelanja []string
	var Budget int = uang
	var daftarBarang = map[int]string{
		1500000: "Sepatu Ceebok",
		500000:  "Baju Zoro",
		250000:  "Baju H&N",
		175000:  "Sweater Uniklooh",
		50000:   "Casing Handphone",
	}

	var keys []int
	for key := range daftarBarang {
		keys = append(keys, key)
	}
	// fmt.Println(keys)
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	// fmt.Println(keys)

	for _, value := range keys {
		if Budget-value >= 0 {
			kerangjangBelanja = append(kerangjangBelanja, daftarBarang[value]+",")
			Budget -= value
		}
	}
	dataLog(member{id: id, uang: uang, barang: kerangjangBelanja, sisaUang: Budget})
	return member{id: id, uang: uang, barang: kerangjangBelanja, sisaUang: Budget}
}

func dataLog(data member) {
	f, err := os.Create("dataPembeli.txt")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	d := data
	fmt.Fprintln(f, d)

	err = f.Close()
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
