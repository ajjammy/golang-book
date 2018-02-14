package main

import "fmt"

func main() {
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Printf("%v\n", string(c))
	}
}
