package main

import "fmt"

func func1(b int64, c int64) bool {
	var a bool
	a = b > c
	fmt.Println(a)
	return a
}

func main() {
	func1(2, 1)
}
