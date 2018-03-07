package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 1; i <= 100; i++ {

		fmt.Println(
			func(i int) string {
				fbTmp := func(number int, str string) func(i int) (string, bool) {
					return func(i int) (string, bool) {
						if i%number == 0 {
							return str, true
						}
						return "", false
					}
				}

				fbArray := [...]func(n int) (string, bool){
					fbTmp(15, "FizzBuzz"),
					fbTmp(5, "Buzz"),
					fbTmp(3, "Fizz"),
				}
				for x := 0; x < len(fbArray); x++ {
					if str, ok := fbArray[x](i); ok {
						return str
					}
				}
				return strconv.Itoa(i)
			}(i))
	}
}
