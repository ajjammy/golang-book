package main

import "fmt"

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	var c Circle
	fmt.Printf(" c type: %T\n", c)
	fmt.Println(c.x, c.y, c.r)

	c1 := new(Circle)
	fmt.Printf(" c type: %T\n", c1)
	fmt.Println(c1.x, c1.y, c1.r)

	c2 := Circle{x: 0, y: 0, r: 5}
	fmt.Printf(" c type: %T\n", c2)
	fmt.Println(c2.x, c2.y, c2.r)

	c3 := NewCircle(1, 2, 3)
	fmt.Printf(" c type: %T\n", c3)
	fmt.Println(c3.x, c3.y, c3.r)
}

func NewCircle(x, y, r float64) *Circle {
	return &Circle{x, y, r}
}
