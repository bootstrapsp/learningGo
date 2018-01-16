package main

import "fmt"

// function with 2 returns performed under anonymous function
func main() {

	var userIn int
	// anonymous func expression to do 2 returns
	half := func(n int) (int, bool) {
		// taking using input
		fmt.Scanln(&userIn)
		// doing 2 returns 1 diving userinput value and other's evaluating bool expression
		return userIn / 2, (userIn / 2) == 0
	}

	fmt.Println(half(userIn))
}
