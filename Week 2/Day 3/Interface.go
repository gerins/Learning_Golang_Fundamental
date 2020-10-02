package main

import (
	"fmt"
)

type Kendaraan interface {
	GetWarna()
}

type Mobil struct {
	Depan    string
	Belakang string
	Atas     string
	Dalam    string
}

type Motor struct {
	Depan    string
	Belakang string
}

func (m Mobil) GetWarna() {
	fmt.Println(m)
}

func (m Motor) GetWarna() {
	fmt.Println(m)
}

func main() {
	var k Kendaraan
	mobil := Mobil{
		Depan:    "Biru",
		Belakang: "Biru",
		Atas:     "Pink",
		Dalam:    "Putih",
	}

	k = mobil
	k.GetWarna()

	motor := Motor{
		Depan:    "Green",
		Belakang: "Black",
	}

	k = motor

	k.GetWarna()
}
