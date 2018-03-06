package main

import "fmt"

func main() {
	// variables are not initialized, only defined
	var a int32
	var b float64
	var c bool
	var d int64
	var e int8
	var f string
	var g byte

	// using %v to fetch default Go value assigned to the uninitialized variables.

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)

	// %T will print out the datatype of a variable.
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", g)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)

}
