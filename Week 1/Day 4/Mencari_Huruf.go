package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`Masukan Kalimat :`)
	scanner.Scan()
	kata := strings.Split(scanner.Text(), "")

	var parameter string
	fmt.Println(`Cari huruf apa :`)
	fmt.Scan(&parameter)

	fmt.Println(cariHuruf(parameter, kata))
}

func cariHuruf(apa string, dimana []string) bool {
	for _, isinya := range dimana {
		if isinya == apa {
			return true
		}
	}
	return false
}
