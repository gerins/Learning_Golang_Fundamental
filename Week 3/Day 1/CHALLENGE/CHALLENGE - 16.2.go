//[CHALLENGE - 16.2] Restoran Cepat Saji
// 15 Juni 2020
// Nama : Garin Prakoso

package main

import (
	"fmt"
	"sync"
	"time"
)

var daftarMenu = map[string]time.Duration{
	"Nasi Goreng": 4,
	"Spaghetti":   3,
	"Ayam Goreng": 5,
	"Jus Alpukat": 2,
	"Air Mineral": 1}

var wd sync.WaitGroup

func main() {
	Budi := []time.Duration{daftarMenu["Air Mineral"], daftarMenu["Ayam Goreng"]}
	go buatMakanan("Budi", Budi)
	wd.Add(1)
	Tono := []time.Duration{daftarMenu["Nasi Goreng"]}
	go buatMakanan("Tono", Tono)
	wd.Add(1)
	fmt.Println("Pesanan Sedang Dibuat")
	cetak := true
	go func() {
		for cetak {
			time.Sleep(time.Millisecond * 500)
			fmt.Print(`.`)
			if !cetak {
				break
			}
		}
	}()

	wd.Wait()
	cetak = false
}

func buatMakanan(nama string, waktu []time.Duration) {
	var totalDuration time.Duration
	for _, value := range waktu {
		time.Sleep(time.Second * value)
		totalDuration += value
	}
	var daftarPesanan []string
	for _, value1 := range waktu {
		for key, value2 := range daftarMenu {
			if value1 == value2 {
				daftarPesanan = append(daftarPesanan, key)
			}
		}
	}
	fmt.Println("")
	fmt.Println(`Pesanan`, nama, daftarPesanan, `Berhasil dibuat dengan durasi`, time.Second*totalDuration)
	wd.Done()
}
