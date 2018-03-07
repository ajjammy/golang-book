package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	c <- 2 // Overfilled the buffer
	fmt.Println(<-c)
	fmt.Println(<-c)
}
