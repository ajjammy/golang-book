package main

import "fmt"

func main() {
	// Long Declaration
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	// Short Declaration
	// Type Inference
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

	const xx string = "Hello, World"
	//xx = "Other string"

	fmt.Println("\nDefine Multiple Variables")
	var (
		a = 5
		b = 10
		c = 15
	)
	println(a, b, c)

	fmt.Println("\nSwapping variable")
	foo, bar := 1, 2
	fmt.Println(foo, bar)
	foo, bar = bar, foo
	fmt.Println(foo, bar)
}
