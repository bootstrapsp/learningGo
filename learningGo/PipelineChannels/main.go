package main

import "fmt"

func main() {

	// putting two values on to the channel
	c := gen(2, 3)

	// calling  sq function to perform square root on to the values that are put into the channel by gen()
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
}

// gen() is a variadic function as its taking multiple int values, returs <-chan of type, this can only be put on to another channel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq() can only receive the channels thus the notation <-chan, returns int

func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
