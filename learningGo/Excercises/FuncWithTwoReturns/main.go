package main

import "fmt"

// func returning 2 values, while accepting only 1 input
func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

// func taking userinput and calling half function
func main() {
	var userInput int
	fmt.Scanln(&userInput)
	h, even := half(userInput)
	fmt.Println(h, even)
}
