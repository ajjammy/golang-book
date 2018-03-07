package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	go func() { ch <- 1 }()
	ch <- 2
	fmt.Println("cap:", cap(ch))
	fmt.Println("len:", len(ch))
}
