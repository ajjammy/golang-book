package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
