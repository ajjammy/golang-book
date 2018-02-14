package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fizzBuzz(i)
	}
}

func fizzBuzz(i int) {
	if isFizz(i) && isBuzz(i) {
		fmt.Println(i, "FizzBuzz")
		return
	}
	if isFizz(i) {
		fmt.Println(i, "Fizz")
		return
	}
	if isBuzz(i) {
		fmt.Println(i, "Buzz")
		return
	}
	fmt.Println(i)
}

func isFizz(i int) bool {
	return i%3 == 0
}

func isBuzz(i int) bool {
	return i%5 == 0
}
