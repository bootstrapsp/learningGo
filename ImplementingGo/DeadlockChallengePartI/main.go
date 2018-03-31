package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()
	fmt.Println(<-c)

	// commenting out as its not working yet
	//deadlockchannelII()

	redoingdeadlockchannelII()

}

func deadlockchannelII() {
	// this function is not fully working yet
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		wg.Wait()
	}()

	go func() {
		for j := range c {
			fmt.Println(j)
		}
		wg.Wait()
	}()

	go func() {
		wg.Done()

	}()
}

func redoingdeadlockchannelII() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
