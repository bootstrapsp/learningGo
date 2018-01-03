package main

import "fmt"

func main() {
	a := 42

	// here the "*int" in itself is a datatype; its a pointer datatype * is denoting the pointer characteristic and int is showing that it's an integer pointer
	var b *int = &a

	fmt.Println(a)
	fmt.Println("Value of the &a memory address assigned to the b variable", b)

	// the *b is fetching the value from the variable b which is storing the memory address from a and a has the actual value of 42
	fmt.Println("Dereferincing the memory address back to actual value", *b)
}
