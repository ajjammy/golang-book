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

	fmt.Println("=====Comparing floating point=====")
	fmt.Println("0.1 + 0.2 == 0.3 is", 0.1+0.2 == 0.3)
	num := 0.1
	num += 0.2
	fmt.Println("num == 0.3 is", num == 0.3)
	fmt.Println("num is", num)
	// ensure the difference is not too big
	fmt.Println(math.Abs(num-0.3) < 0.0001)

	fmt.Println("\n=====String=====")
	backticks := `hello world!, 
today's good day.`
	fmt.Println(backticks)

	doubleQuotes := "hello world!,\ntoday's good day."
	fmt.Println(doubleQuotes)

	fmt.Println("\n=====Zero Value=====")
	var number int
	var str string
	var boolean bool
	fmt.Printf("number: %v\n", number)
	fmt.Printf("str: '%v'\n", str)
	fmt.Printf("boolean: %v\n", boolean)
}
