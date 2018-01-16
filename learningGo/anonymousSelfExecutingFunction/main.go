package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is the anonymous self executing function")
	}() // parenthesis at the end of the function denotes that this is a self executing function
}
