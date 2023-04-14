package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func sumArea(shapes ...Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.area()
	}
	return total
}

func main() {
	total := sumArea(
		Circle{radius: 5},
		Circle{radius: 10},
		Rectangle{width: 4, height: 7},
	)
	fmt.Println(total)
}
