package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {

		if j := i % 3; j == 0 {
			fmt.Println("This number is multiple of 3", "Fizz")
		}
		if k := i % 5; k == 0 {

			fmt.Println("This number is multiple of 5", "sBuzz")
		}

	}
}
