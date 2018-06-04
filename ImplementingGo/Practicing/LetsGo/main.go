package main

import (
	"fmt"
)

func main() {

	cv := creatingValues() // creating cv of tye creatingValue

	// values access individually
	fmt.Println("ID :", cv.ID)
	fmt.Println("ID :", cv.Name)
	fmt.Println("ID :", cv.class)
	fmt.Println("ID :", cv.demo)
	fmt.Println("ID :", cv.prime)

}

//ThingShape a struct
type ThingShape struct {
	Name string
	ID   int
	demo
}

// Demo is a struct
type demo struct {
	prime float64
	class string
}

func creatingValues() ThingShape {

	ts := ThingShape{"mockVal", 1, demo{2342.34, "C"}} // creating values for the struct

	return ts
}
