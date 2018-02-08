package main

import "fmt"

func main() {
	fmt.Print("Enter a degree of fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Degree of celsius is %.2f\n", celsius)
}
