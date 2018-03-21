package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("foo ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	wg.Done()
}

func bar() {
	for j := 0; j < 10; j++ {
		fmt.Println("bar ", j)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func main() {
	// wg.Add which is a function in Wait group to use semaphore, i.e. a temp value used as counter
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

}
