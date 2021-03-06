package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	// this is opening multiple channels
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- 1
			}
			done <- true
		}()
	}

	// this is handling multiple channels and closing them
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
