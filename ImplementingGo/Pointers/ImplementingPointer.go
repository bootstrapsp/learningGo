package main

import (
	"fmt"
	"log"
)

func main() {

	pointers()
}

func pointers() {

	a := "some string value"
	b := 45

	fmt.Println("memory address storing value of a", &a)
	fmt.Println("memory address storing value of b", &b)

	c := &a
	d := &b

	fmt.Println("Assining a's memory address to c", c)
	fmt.Println("Assining a's memory address to c", d)

	fmt.Printf("Checking the type of c variable %T \n", c)
	fmt.Printf("Checking the type of d variable %T \n", d)

	e := *c
	fmt.Println("fetching the actual value by dereferencing c", e)
	f := *d
	fmt.Println("fetching the actual value by dereferncing d", f)

	*c = "new string val"
	fmt.Println("Assigning new value to *c")

	*d = 6
	fmt.Println("Assigning new value to *d")

	log.Print("THis is from the log!!")
}
