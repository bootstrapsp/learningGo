package main

import "fmt"

func main() {
	// "[]" this is called slice, which is nothing but the List of n number of inputs
	data := []float64{42, 32, 56, 23, 55, 555, 212}
	
	// here the "..." is also part of the variadic function which is basically reading and inputting that value to the variadicFunc created below to process.

	// it would fail if we will put directly the data as its only representing 1 slice
	n := variadicFunc(data...)
	fmt.Println(n)
}

// variadic functions are the functions that can have n number of parameters of a type, this is done by putting "..." in front of the type
func variadicFunc(demo2 ...float64) float64 {
	total := 0.0
	for _, v := range demo2 {
		total += v
	}
	return total / float64(len(demo2))
}
