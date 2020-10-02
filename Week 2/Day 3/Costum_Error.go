package main

import (
	"fmt"
)

func main() {
	if hasil, err := fungsiTambah(1, 2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hasil)
	}
}

type StrukturErrorBaru struct {
	pesanError string
}

func (n StrukturErrorBaru) Error() string {
	return ERROR_MESSAGE
}

func fungsiTambah(a, b int) (int, error) {
	if a < 5 && b < 5 {
		return 0, StrukturErrorBaru{}
	}

	return a + b, nil
}

const (
	ERROR_MESSAGE = "Tidak dapat menjalankan Program"
)
