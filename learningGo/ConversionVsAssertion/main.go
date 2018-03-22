package main

import "fmt"

func floattointConversion() {

	var x float64
	var y int64
	x = 12.31432
	fmt.Println("Prior to conversion", x)
	// narrowing the scope of the value, due to losing decimal places
	fmt.Println(int64(x) + y)
}

func main() {
	floattointConversion()
}
