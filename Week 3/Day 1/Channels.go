package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channelA := make(chan string)
	go printNama(`Budi`, 5, channelA)
	for msg := range channelA {
		fmt.Println(msg)
	}
	fmt.Println("Finish main")
}

func printNama(name string, times int, a chan string) {
	fmt.Println(`channel a`, a)
	for i := 0; i <= times; i++ {
		a <- `Selamat Pagi ` + name + strconv.Itoa(i)
	}
	return
	close(a)
}
