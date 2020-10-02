package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dataLog("Daftar Kontak.txt", "Menulis file")
	readFile()
}

func dataLog(filePath string, kata string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(f, kata)

	return nil
}

func readFile() {
	file, _ := os.Open("Daftar Kontak.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(string(scanner.Text()))
	}
}
