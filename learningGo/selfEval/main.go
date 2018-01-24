package main

import "fmt"

func main() {
	arrCreation()
}

// function for creating the array and assigning two values to it
// one hardcoded and other just as userINput
func arrCreation() {

	var dem [10]int

	dem[0] = 1

	fmt.Println(dem[0])

	var userInput int
	fmt.Scanln(&userInput)

	dem[1] = userInput

	fmt.Println(dem[1])
}
