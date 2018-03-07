package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		func(a, b int) int {
			return a + b
		}(2, 3))
}
