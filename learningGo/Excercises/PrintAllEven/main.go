package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {

		if j := i % 2; j == 0 {
			fmt.Println("Even numbers are \n", i)
		}

	}

}
