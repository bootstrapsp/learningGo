package main

import "fmt"

// illustrating the use of recursion where the function calls itself within itself
func recursionDemo(x int) int {
	if x == 0 {
		return 1
	}
	return x * recursionDemo(x-1)
}

func main() {
	var userIn int
	fmt.Scanln(&userIn)
	fmt.Println(recursionDemo(userIn))
}
