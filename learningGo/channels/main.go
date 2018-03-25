package main

import (
	"fmt"
	"time"
)

func main() {
	// channel needs to be made using "make"
	// below channel is not a buffered channel, so it can only take 1 item at a time
	// to define buffered channel something like c := make(chan int, <someNUmber here>)
	c := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {

			// this is the important part as here the channel passes the value due to the "<-"
			c <- i
		}
	}() // since i am creating anonymous functions this "()" calls the function on its own

	go func() {
		for {
			// this receives the item from the channel
			fmt.Println(<-c)
		}

	}()

	// this sleep is must to have to allow the time to complete the channel work above
	time.Sleep(time.Second)
}
