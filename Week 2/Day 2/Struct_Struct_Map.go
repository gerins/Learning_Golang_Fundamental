package main

import (
	"fmt"
)

type sliceMap []map[string]string
type mapString map[string]string

type student struct {
	firstName, lastName string
	age                 int
	alamat              address
}

type address struct {
	home    sliceMap
	poscode int
}

func main() {
	budi := student{
		firstName: "budi",
		lastName:  "anduk",
		age:       18,
		alamat: address{
			home: sliceMap{
				mapString{"address1": "Bekasi", "Address2": "Jakarta"},
				mapString{"Office1": "Jakarta", "Office2": "Bandung"}},
			poscode: 17115},
	}

	fmt.Println(budi)
	fmt.Println(budi.alamat.home[0]["address1"])
	fmt.Println(budi.alamat.home[1])
	fmt.Println(budi.alamat.poscode)

	gerin := student{
		firstName: "Gerin",
		lastName:  "Prakoso",
		age:       25,
		alamat: address{
			home:    sliceMap{mapString{"alamat1": "Bekasi"}},
			poscode: 17115,
		},
	}

	azza := student{
		firstName: `azza`,
		lastName:  `Kurnianingrum`,
		age:       25,
		alamat: address{
			home: sliceMap{
				mapString{`kota`: `madiun`, `kantor`: `jakarta`},
			},
			poscode: 17115,
		},
	}

	fmt.Println(gerin)
	fmt.Println(azza)
}
