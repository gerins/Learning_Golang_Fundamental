package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- `Every 500ms`
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- `Every 2 Seconds`
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

	// c := make(chan string, 3)
	// c <- "hello"
	// c <- "world"

	// msg := <-c
	// fmt.Println(msg)

	// msg = <-c
	// fmt.Println(msg)

	// go ulang(`kucing`, c)

	// for msg := range c { // hanya jalankan ketika c masih mengirim
	// 	fmt.Println(msg)
	// }
}

// func ulang(apa string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- apa
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	fmt.Println(`closed channel C`)
// 	close(c)
// }
