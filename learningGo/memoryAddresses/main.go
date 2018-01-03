package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a -", a)
	/*
		the &a used in the statement below prints
		out the memory address where the value for the
		variable a is stored.
	*/
	fmt.Println("a's memory address - ", &a)
	/*
		Using Print format with %d we can see the
		decimal value of the hexadecimal value shown
		above with &a
	*/
	fmt.Printf("%d", &a)
}
