package main

import "fmt"

func main() {
	var j int
	for i := 0; i <= 1000; i++ {
		fmt.Println("Value of i is", i)
		if i%3 == 0 {
			// += operator increments and assigns the value to the variable
			fmt.Println("Value of previous 3multiple j is", j)
			j += i
			fmt.Println("Summing up the J if multiple of 3", j)

		} else if i%5 == 0 {

			// += operator increments and assigns the value to the variable
			fmt.Println("Value of previous 5multiple j is", j)
			j += i
			fmt.Println("Summing up the J if multiple of 5", j)
		} else {
			fmt.Println("Not summed up", i)
		}

	}

}
