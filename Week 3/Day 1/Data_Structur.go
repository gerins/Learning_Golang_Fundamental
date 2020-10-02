package main

import "fmt"

type Kontak struct {
	group string
	nama  []string
}

func (n *Kontak) addKontak(nama string) {
	n.nama = append(n.nama, nama)
}

var daftarKontak []Kontak

func main() {
	newGroup("Rumah")
	newGroup("Sekolah")
	newKontak("Rumah", "Gerin")
	newKontak("Rumah", "Azza")
	newKontak("Rumah", "Budi")
	newKontak("Sekolah", "Dodi")
	newKontak("Sekolah", "Dodi")
	newKontak("Sekolah", "Dodi")
	fmt.Println(daftarKontak)
	tampilkanGroup(daftarKontak)
}

func newGroup(namaGroup string) {
	var tempName []string
	if len(daftarKontak) == 0 {
		daftarKontak = append(daftarKontak, Kontak{namaGroup, tempName})
		return
	} else {
		for _, value := range daftarKontak {
			if value.group == namaGroup {
				fmt.Println("Group Sudah Ada")
				return
			} else {
				daftarKontak = append(daftarKontak, Kontak{namaGroup, tempName})
			}
		}
	}
}

func newKontak(namaGroup string, namaKontak string) {
	if len(daftarKontak) == 0 {
		fmt.Println("Group Belum Dibuat")
		return
	} else {
		for index, value := range daftarKontak {
			if value.group == namaGroup {
				daftarKontak[index].addKontak(namaKontak)
				return
			}
		}
		fmt.Println("Group", namaGroup, `Tidak ditemukan`)
		return
	}
}

func tampilkanGroup(list []Kontak) {
	totalLen := 0
	for _, value := range list {
		totalLen += len(value.nama)
		for index, isi := range value.nama {
			fmt.Println(index, isi, value.group)
		}
	}
	fmt.Println(`Total len nya adalah`, totalLen)
}
