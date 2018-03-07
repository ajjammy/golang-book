package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(16)
	go increment(1)
	go increment(2)
	go increment(3)
	go increment(4)
	go increment(5)
	go increment(6)
	go increment(7)
	go increment(8)
	go increment(9)
	go increment(10)
	go increment(11)
	go increment(12)
	go increment(13)
	go increment(14)
	go increment(15)
	go increment(16)
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func increment(n int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
	}
}
