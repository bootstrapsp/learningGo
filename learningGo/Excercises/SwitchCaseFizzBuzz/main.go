package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Println("This number is multiple of 3", "Fizz")
		} else if i%5 == 0 {

			fmt.Println("This number is multiple of 5", "Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
