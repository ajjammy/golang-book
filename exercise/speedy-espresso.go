package main

import (
	"fmt"
	"time"
)

func main() {
	volumn := 100
	start := time.Now()

	orders := make(chan string)
	coffees := make(chan string)
	container := make(chan string)

	go reciveOrder(volumn, orders)

	barista := 6
	for i := 1; i < barista; i++ {
		go brewCoffee(orders, coffees)
	}

	go serveCoffee(coffees, container)
	print(container)

	end := time.Now()
	fmt.Println(end.Sub(start))
}

func reciveOrder(volumn int, orders chan<- string) {
	for i := 1; i < volumn; i++ {
		// cashier receive order
		time.Sleep(5 * time.Millisecond)
		orders <- fmt.Sprintf("order: %d", i)
	}
	close(orders)
}

func brewCoffee(orders <-chan string, coffees chan<- string) {
	// barista brew coffee
	for coffee := range orders {
		time.Sleep(100 * time.Millisecond)
		coffees <- fmt.Sprintf("%s %s", coffee, "espresso")
	}
	close(coffees)
}

func serveCoffee(coffees <-chan string, container chan<- string) {
	// waiter serve coffee
	for coffee := range coffees {
		time.Sleep(5 * time.Millisecond)
		container <- fmt.Sprintf("%s %s", coffee, "ready :)")
	}
	close(container)
}

func print(container <-chan string) {
	for cup := range container {
		fmt.Println(cup)
	}
}
