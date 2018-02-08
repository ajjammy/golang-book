package main

import "fmt"

func main() {
	fmt.Print("Enter a number of feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	fmt.Printf("%.4f meters\n", meters)
}
