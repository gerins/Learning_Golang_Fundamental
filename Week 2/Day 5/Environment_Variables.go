package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	config := os.Args
	fmt.Println(`File Config:`)
	dbServer := strings.Split(config[1], `=`)[1]
	ipAddress := strings.Split(config[2], `=`)[1]

	for index, isi := range config {
		fmt.Println(`index ke`, index, `berisi`, isi)
	}

	fmt.Println(`dbServer is `, dbServer)
	fmt.Println(`ipAddress is `, ipAddress)
}
