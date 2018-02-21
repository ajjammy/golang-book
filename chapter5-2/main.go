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
	ln := [4]int{15, 3, 5, 7}
	str := [4]string{"FizzBuzz", "Fizz", "Buzz", "Aha"}

	for i := 0; i < len(ln); i++ {
		if number%ln[i] == 0 {
			return str[i]
		}
	}
	return strconv.Itoa(number)
}
