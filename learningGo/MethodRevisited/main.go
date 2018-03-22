package main

import (
	"fmt"
	"math"
)

// creating circle type as struct
type circle struct {
	radius float64
}

// creating shape interface which implements area()
type shape interface {
	area() float64
}

// attaching the area() method to the circle
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

// info function performing the cal using area function
func info(s shape) {
	fmt.Println("area", s.area())

}

func main() {
	c := circle{444}
	info(&c)
}
