//[CHALLENGE - 16.1] Change Me!
// 13 Juni 2020
// Nama : Garin Prakoso

package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       string
}

func main() {
	var input = new([][]string)
	*input = [][]string{{"budi", "tono", "1993"}, {"tini", "nurlela", "2021"}}

	// var pointInput *[][]string
	// pointInput = &input
	// fmt.Println(&pointInput)

	inputPerson(input)
	fmt.Printf("%p", input)
}

func inputPerson(input *[][]string) {
	temp := []string{"Kucing"}
	*input = append(*input, temp)

	fmt.Printf("%p\n", input)
}
