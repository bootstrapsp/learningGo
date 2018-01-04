package main

import "fmt"

func main() {
	calRemainder()
}

// function created to cal the remainder out of the equation
// loop printing out ODD/Even based on the remainder
// As among different coding lang "%" is used for calculating remainder
func calRemainder() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}

	}

}
