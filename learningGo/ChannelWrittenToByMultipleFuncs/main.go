package main

import (
	"fmt"
	"sync"
)

func main() {

	// creating channel
	c := make(chan int)

	// using WaitGroup from Sync pacakge to define how many routines we have and to wait for each of them to finish
	var wg sync.WaitGroup

	// define go number of routines
	wg.Add(2)

	// go routine 1 with Done to tell WaitGroup overwatch that its done executing
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()

	}()

	// go routine 2 with Done to tell WaitGroup overwatch that its done executing
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()

	}()

	// specific anonymous function to wait and close the channels

	go func() {
		wg.Wait()
		close(c)
	}()

	func() {

		// ranging using range over the channel to print out the items
		for n := range c {
			fmt.Println(n)
		}
	}()
}
