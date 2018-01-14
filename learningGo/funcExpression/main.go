package main

import "fmt"

func main() {

	// defining function within the function, only anonymous function can be created within a function. Such as the following which is getting assigned to a parameter.
	greetings := func() {
		fmt.Println("Hellow World")
	}

	greetings()
	fmt.Printf("%T", greetings)
	greet := makeGreeter()
	fmt.Println(greet())
}

// creating a function whose return type is a func of datatype string
func makeGreeter() func() string {
	return func() string {
		return "Hello World"

	}
}
