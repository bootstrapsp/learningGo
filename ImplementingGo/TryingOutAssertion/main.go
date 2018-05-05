package main

import (
	"fmt"
)

func main() {
	assertCheck()
	Value()
	tryingAssertt()
	insertSlice()
	insertInLoop()
}

func assertCheck() {

	var demoVal interface{} = `dsfsvsdfsdfvsd`

	fmt.Printf("%T", demoVal)

}

// an empty interface
var assert1 interface{}

func tryingAssertt() {

	// check for empty interface
	if assert1 == nil {
		assert1 = 7
		i := assert1.(int)
		fmt.Println("assert was nil now assigned with", i)
	} else {
		fmt.Println("Assert isn't nil")
	}

}

func insertSlice() {

	x := make([]int, 8)

	x[0] = 0
	x[2] = 2

	x = append(x) // appending to slice
	fmt.Println(len(x))

	fmt.Println("check for the slice", x)
}

func insertInLoop() {

	x := make([]int, 100)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			x[i] = j
			x = append(x)

			fmt.Println("inserting in loop for x in inserInLoop function", x)
		}
	}
}

// const increamenting with iota
const (
	a = iota
	b = 2 << iota
	c = iota
)

const (
	a1 = 1 << iota
	a2 = 1 << iota
	a3 = iota
)

// Value is exported
func Value() {

	b := 2

	fmt.Println("fetching iota", a3)
	fmt.Println("fetching another iota", b)
}
