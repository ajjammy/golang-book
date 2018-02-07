package main

import "fmt"

func main() {
	fmt.Println("=====String=====")
	backticks := `hello world!, 
today's good day.`
	fmt.Println(backticks)

	doubleQuotes := "hello world!,\ntoday's good day."
	fmt.Println(doubleQuotes)
}
