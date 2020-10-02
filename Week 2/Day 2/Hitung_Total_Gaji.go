package main

import "fmt"

type gaji struct {
	basic, tunjangan, lembur float64
}

type karyawan struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []gaji
}

func (e *karyawan) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.basic)
		fmt.Println(info.tunjangan)
		fmt.Println(info.lembur)
	}
	return "----------------------"
}

func (n *karyawan) totalGaji() {
	var totalSeluruhnya float64
	for index, isi := range n.MonthlySalary {
		fmt.Println(`index ke `, index, isi)
		total := isi.basic + isi.tunjangan + isi.lembur
		totalSeluruhnya += total
	}
	fmt.Println(`Total gajinya adalah :`, totalSeluruhnya)
}

func main() {
	e := karyawan{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []gaji{
			gaji{
				basic:     15000.00,
				tunjangan: 5000.00,
				lembur:    2000.00,
			},
			gaji{
				basic:     16000.00,
				tunjangan: 5000.00,
				lembur:    2100.00,
			},
			gaji{
				basic:     17000.00,
				tunjangan: 5000.00,
				lembur:    2200.00,
			},
		},
	}
	e.EmpInfo()
	e.totalGaji()
}
