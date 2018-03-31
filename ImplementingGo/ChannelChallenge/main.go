package main

import "fmt"

func main() {
	//c := factorial(99)
	//for n := range c {
	//	fmt.Println(n)

	anotherFactorialMethod(10)

}

func factorial(n int) chan int {
	out := make(chan int)

	go func() {
		total := 1
		for j := n; j < 0; j-- {
			total *= j
		}
		out <- total
		close(out)
	}()

	return out
}

func anotherFactorialMethod(n int) int {

	total := 1

	for i := n; i > 0; i-- {

		total = total * i

		fmt.Println(total)
	}
	return total
}
