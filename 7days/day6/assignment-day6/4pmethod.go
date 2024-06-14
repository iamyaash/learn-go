package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {
	circ := Circle{radius:8}
	circ.Scale()
	fmt.Println(circ.radius)
}

func (c *Circle) Scale() {
	c.radius *= 2
}
