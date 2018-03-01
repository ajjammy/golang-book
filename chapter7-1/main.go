package main

import (
	"fmt"
)

func main() {
	var x [5]int
	x[3] = 4
	fmt.Println(x)

	x = [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)

	y := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(y)
}
