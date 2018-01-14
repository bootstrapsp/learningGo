package main

import "fmt"

// callback: passing a func as an argument
/*
following function is taking slice of int, which is defined with "[]int" and a function
*/
func visit(numbers []int, callbackDemoFunc func(int)) {
	for _, n := range numbers {
		callbackDemoFunc(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}
