package main

import "fmt"

func main() {
	mymap := make(map[int]int)
	mymap[1] = 1
	mymap[2] = 2

	fmt.Println(mymap[3])

	if mymap[3] != 0 {
		fmt.Println(mymap[3])
	}

	// ok?
	value, ok := mymap[3]
	fmt.Println(ok)
	if ok {
		fmt.Println(value)
	}

	if value, ok := mymap[3]; ok {
		fmt.Println(value)
	}
}
