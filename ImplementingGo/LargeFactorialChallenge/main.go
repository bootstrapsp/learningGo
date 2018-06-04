package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		largeFactorialCalc(i)
	}
}

func largeFactorialCalc(userInput int) int {
	total := 1

	func() {

		for i := 0; i > userInput; i-- {
			total = total * i
			fmt.Println(total)
		}

	}()
	return total
}
