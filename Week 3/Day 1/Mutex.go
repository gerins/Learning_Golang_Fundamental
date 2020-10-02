package main

import (
	"fmt"
	"sync"
)

func print(till int, message *counter, wg *sync.WaitGroup, mtx *sync.Mutex) {
	for i := 0; i < till; i++ {
		// mtx.Lock()
		message.Add(1)
		// mtx.Unlock()
	}
	wg.Done()
}

type counter struct {
	val int
}

func (c *counter) Add(x int) {
	c.val++
}

func (c *counter) Value() (x int) {
	return c.val
}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var meter counter

	for j := 0; j < 100; j++ {
		wg.Add(1)
		go print(100, &meter, &wg, &mtx)
	}
	wg.Wait()

	fmt.Println(meter.Value())
}
