package main

import (
	"fmt"
	"time"
)

func main() {
	b := make(chan string)
	c := make(chan string)

	go cetakKata("kucing", b)
	go cetakWord("burung", c)

	for msg1 := range b {
		fmt.Println(msg1)
	}
	for msg2 := range c {
		fmt.Println(msg2)
	}
	time.Sleep(time.Millisecond * 10000)
	fmt.Println(`End Loop`)
}

func cetakKata(apa string, tunnel chan string) {
	for i := 0; i < 5; i++ {
		tunnel <- apa
		time.Sleep(time.Millisecond * 300)
	}
	close(tunnel)
	fmt.Println(`Server Closed`)
}

func cetakWord(apa string, tunnel chan string) {
	for i := 0; i < 5; i++ {
		tunnel <- apa
		time.Sleep(time.Millisecond * 1000)
	}
	close(tunnel)
	fmt.Println(`Server Closed`)
}
