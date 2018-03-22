package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// WaitGroup waits for a collection of goroutines to finish
var wg sync.WaitGroup

//init function executes at the very beginning of the program, good for setting up the playground
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

}

// added sleep function
func foo() {

	for i := 0; i < 100000000000; i++ {
		fmt.Println("foo ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	wg.Done()
}

func bar() {
	for j := 0; j < 10000000; j++ {
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
