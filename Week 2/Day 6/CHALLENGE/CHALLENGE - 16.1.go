//[CHALLENGE - 16.1] Change Me!
// 13 Juni 2020
// Nama : Garin Prakoso

package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       string
}

func main() {
	var input = [][]string{{"budi", "tono", "1993"}, {"tini", "nurlela", "2021"}}
	inputPerson(input)
}

func inputPerson(input [][]string) {
	var outputPerson []Person
	for i := 0; i < len(input); i++ {
		birthYear, _ := strconv.Atoi(input[i][2])
		if birthYear > 2020 || input[i][2] == "" {
			input[i][2] = "Invalid Birth Year"
		} else {
			input[i][2] = strconv.Itoa(2020 - birthYear)
		}
	}

	for _, value := range input {
		outputPerson = append(outputPerson, Person{firstName: value[0], lastName: value[1], age: value[2]})
	}

	fmt.Println(outputPerson)
	for _, isiPerson := range outputPerson {
		fmt.Println(`firstName:`, isiPerson.firstName)
		fmt.Println(`lastname:`, isiPerson.lastName)
		fmt.Println(`age:`, isiPerson.age)
		fmt.Println("")
	}
}
