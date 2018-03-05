package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(n int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(n, ":", i)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Finished")
}
