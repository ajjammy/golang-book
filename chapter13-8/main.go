package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var a, b value
	wg.Add(2)
	go printSum(&a, &b, &wg)
	go printSum(&b, &a, &wg)
	wg.Wait()
}

func printSum(a, b *value, wg *sync.WaitGroup) {
	defer wg.Done()
	a.mu.Lock()
	defer a.mu.Unlock() // introduce deadlock

	time.Sleep(2 * time.Second)
	b.mu.Lock()
	defer b.mu.Unlock() // introduce deadlock

	fmt.Printf("sum=%v\n", a.value+b.value)
}
