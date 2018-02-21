package main

import "fmt"

type Zipcode string

func main() {
	zipcode := Zipcode("11000")
	if zipcode.valid() {
		fmt.Println(zipcode)
		fmt.Println(zipcode)
	}
}
func (z Zipcode) valid() bool {
	return true
}
