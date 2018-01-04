package main

import "fmt"

// dereferencing whatever is being passed to variable y
// here the dereferencing is done with the notation "*" in front of the variable
// go lang is pass by value
// "*int" is a type in itself namely pointer to int
func derefThePointer(y *int) {
	*y = 0
	fmt.Println(*y)
	fmt.Printf("%p\n", &y)
}

/*
x is getting assgined value 5 but then deferencing functions is called on x to remove whatever was assigned to the x and set to 0

"%p" Pointer formating part of the fmt pacakage; to format and display base 16 notiation with leading 0x
*/
func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	derefThePointer(&x)
	fmt.Println(x)
	fmt.Printf("%p\n", &x)
}
