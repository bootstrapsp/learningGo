package main

import "fmt"

func main() {

	array1()
	array2()
	array3()
}

func array1() {

	// uninizialized array so value is 0 for all places
	var x [10]int

	// printing out the lenth of the array
	fmt.Println(len(x))
	fmt.Println(x)
	// assigning value to the 3rd index in the array
	x[3] = 1999
	fmt.Println(x[3])

}

func array2() {
	var y [58]string

	// logic :
	/*
		starting with index 0 we assign string i which is string 0, and in the array
	*/
	for i := 65; i <= 122; i++ {
		y[i-65] = string(i)
	}

	fmt.Println(y)
	fmt.Println(y[42])

}

func array3() {
	var x [256]byte

	// printing the
	for i, v := range x {
		fmt.Println(i, v)
		fmt.Printf("%v", v)
	}

}
