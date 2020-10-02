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
	kata := scanner.Text()
	fmt.Println(`Jumlah huruf vokal nya adalah :`, jumlahKata(kata))
}

func jumlahKata(apa string) int {
	proses := strings.Split(apa, "")
	counter := 0
	for _, isinya := range proses {
		if isinya == "a" || isinya == "i" || isinya == "u" || isinya == "e" || isinya == "o" {
			counter++
		}
	}
	return counter
}
