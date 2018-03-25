package main

import "fmt"

func main() {
	// creating a channel with make, without buffer info
	c := make(chan int)
	// this a go routine that starts a function which passes the value to the channel
	go func() {

		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// this function gets executed sequentially as this is not a go routine; there's no "go" defined in front of it
	func() {
		for n := range c {
			// n receives the values from the the channel above and prints it out
			fmt.Println(n)
		}
	}()
}
