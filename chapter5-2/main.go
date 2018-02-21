package main

import (
	"fmt"
	"strconv"
)

func main() {
	for number := 1; number <= 100; number++ {
		fmt.Println(number, fizzBuzz(number))
	}
}

func fizzBuzz(number int) string {
	if number%15 == 0 {
		return "FizzBuzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(number)
}
