package main

import "fmt"

/*
This example shows usage of only channels no WaitGroup involved, just to understand how with just milti channels we can control the execution flow of multiple go routines
*/
func main() {

	// creating 2 channels
	c := make(chan int)

	// this second channel is there as a semaphore, that is raising true as bool value once the go routine is done with its work
	done := make(chan bool)

	// go routine 1 to put value in channel c
	go func() {

		for i := 1; i < 10; i++ {
			c <- i

		}
		// semaphore channel setting true to the channeö
		done <- true
	}()

	go func() {
		for i := 11; i < 20; i++ {
			c <- i

		}
		// semaphore channel setting true to the channeö
		done <- true
	}()

	go func() {
		// like we did with WaitGroup Add 2 to manage 2 go routines at once, here we are handling the same with semaphore channel value with bool value set to true
		<-done
		<-done

		// closing the other channel
		close(c)
	}()

	func() {
		for n := range c {
			fmt.Println(n)
		}
	}()
}
