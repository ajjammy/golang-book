package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=====Floating point=====")
	third := 1.0 / 3.0
	fmt.Printf("third = %v\n", third)
	fmt.Printf("third + third + third = %v\n", third+third+third)

	fmt.Println("\n=====Comparing floating point=====")
	fmt.Println("0.1 + 0.2 == 0.3 is", 0.1+0.2 == 0.3)
	num := 0.1
	num += 0.2
	fmt.Println("num == 0.3 is", num == 0.3)
	fmt.Println("num is", num)
	// ensure the difference is not too big
	fmt.Println(math.Abs(num-0.3) < 0.0001)
}
