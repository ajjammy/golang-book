package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
	mu      sync.Mutex
)

func main() {
	wgnum := 16
	wg.Add(wgnum)
	for i := 1; i <= wgnum; i++ {
		go increment(i)
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func increment(n int) {
	defer wg.Done()
	mu.Lock()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
	}
	mu.Unlock()
}
