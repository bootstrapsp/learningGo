package main

import "fmt"

func getUserInput() {
	// switch condition will look through the cases to return for the appropriate response
	// if not found it will just go for the default one.
	switch "DEF" {
	case "ABC":
		fmt.Println("howdy ABC")
	case "DEF":
		fmt.Println("howdy DEF")
	default:
		fmt.Println("This is strange!")
	}

}

func main() {
	getUserInput()
}
