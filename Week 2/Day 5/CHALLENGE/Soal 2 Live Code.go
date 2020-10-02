// Nama : Garin Prakoso

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var kredit int = 5000
	var slotMachine = []int{7, 8, 9}

	pesan1, pesan2 := mesinJudi(&kredit, slotMachine)
	fmt.Println(pesan1)
	fmt.Println(pesan2)
	fmt.Println(kredit)
}

func mesinJudi(balance *int, input []int) (string, string) {
	var pesan1, pesan2 string

	if input[0] == input[1] && input[1] == input[2] {
		winBalance := (input[0] + input[1] + input[2]) * 200
		*balance += winBalance
		pesan1 = `You win ` + strconv.Itoa(winBalance) + ` Dollars`
		pesan2 = `Your Total Credit Balance Is ` + strconv.Itoa(*balance) + ` Dollar`
	} else if input[0] == input[1] || input[1] == input[2] {
		winBalance := (input[0] + input[1] + input[2]) * 100
		*balance += winBalance
		pesan1 = `You win ` + strconv.Itoa(winBalance) + ` Dollars`
		pesan2 = `Your Total Credit Balance Is ` + strconv.Itoa(*balance) + ` Dollar`
	} else {
		winBalance := (input[0] + input[1] + input[2]) * 50
		*balance -= winBalance
		pesan1 = `You Lose ` + strconv.Itoa(winBalance) + ` Dollars`
		pesan2 = `Your Total Credit Balance Is ` + strconv.Itoa(*balance) + ` Dollar`
	}

	return pesan1, pesan2
}
