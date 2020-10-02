package main

import (
	"fmt"
	"strconv"
)

type student struct {
	firstName, lastName string
	age                 int
}

func (n *student) tambahUmur(c string) {
	fmt.Println(c)
	n.age++
	fmt.Println(`Kemudian bertambah menjadi`, n.age, `, selamat`)
}

func (n *student) fullName() string {
	n.sayHello()
	age := strconv.Itoa(n.age)
	return n.firstName + ` ` + n.lastName + ` umurnya adalah ` + age
}

func (n *student) sayHello() {
	fmt.Println(`hallo`, n.firstName)
}

func main() {
	var budi = student{`Budi`, `Anduk`, 30}

	budi.tambahUmur(budi.fullName())
	// fmt.Println(budi)
}
