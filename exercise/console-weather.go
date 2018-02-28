package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak sunny"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))
}

func weatherCelsius(degree int, weatherMessage string) string {
	header := []string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
	middle := []string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
	bottom := []string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}

	numbers := make([]string, 3)

	for _, strNumber := range strconv.Itoa(degree) {
		digitNumber, err := strconv.Atoi(string(strNumber))
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		numbers[0] += header[digitNumber]
		numbers[1] += middle[digitNumber]
		numbers[2] += bottom[digitNumber]
	}

	return numbers[0] + "\n" + numbers[1] + "\n" + numbers[2] + " c\n" + weatherMessage + "\n"
}
