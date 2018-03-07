package main

import (
	"fmt"
)

func main() {
	var addVar func(int, int) int
	addVar = func(a, b int) int {
		return a + b
	}
	fmt.Println(addVar(2, 3))
}
