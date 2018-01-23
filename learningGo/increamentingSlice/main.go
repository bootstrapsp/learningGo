package main

import "fmt"

/*
	This function aims to highlight several way a value on an index in a slice gets appended
	Covers different possible notations on appending
*/
func increamentSliceFunc() {

	// creating the slice with length and cpacity 1
	mySlice := make([]int, 3)

	fmt.Println("before appending the value in the slice index", mySlice[0])

	// following is increamenting the slice value by 1 on index 0
	mySlice[0] = 1

	fmt.Println("after appending the value in the slice index", mySlice[0])

	// value can be appended like so as well for the slice at index 0
	mySlice[0] = mySlice[0] + 10

	fmt.Println("after new append", mySlice[0])

	// value can also be appended like so

	mySlice[0] += 1

	fmt.Println("final append", mySlice[0])

	mySlice[0]++

	fmt.Println("Finally", mySlice[0])
}

func main() {
	increamentSliceFunc()
}
