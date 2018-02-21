package main

import "fmt"

func main() {
	for number := 1; number <= 100; number++ {
		if number%15 == 0 {
			fmt.Println(number, "FizzBuzz")
		} else if number%3 == 0 {
			fmt.Println(number, "Fizz")
		} else if number%5 == 0 {
			fmt.Println(number, "Buzz")
		} else {
			fmt.Println(number)
		}
	}
}
