package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fizzBuzz(i)
	}
}

func fizzBuzz(i int) {
	if i%15 == 0 {
		fmt.Println(i, "FizzBuzz")
	} else if i%3 == 0 {
		fmt.Println(i, "Fizz")
	} else if i%5 == 0 {
		fmt.Println(i, "Buzz")
	} else {
		fmt.Println(i)
	}
}
