package main

import "fmt"

func main() {
	var userInputVal string

	fmt.Scanln(&userInputVal)
	fmt.Println("Hello", userInputVal)
}
