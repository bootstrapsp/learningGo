package main

import "fmt"

func main() {
	var demo string

	/*
		scan is the library in fmt
		that listens for user's input

		&demo is the variable to which
		values inserted by the user is assigned to

	*/
	fmt.Scanln(&demo)
	fmt.Println(demo)
}
