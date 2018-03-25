package main

// this is to show the usage of mutex package, that allows for functions to lock and unlock the execution of the parallel threads

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// used for waiting for goroutines to finish
var wg sync.WaitGroup

// counter used to verifying the sequence execution of the loops
var counter int

// used for locking and unlocking threads
var mutex sync.Mutex

func main() {
	// setting the semaphore in WaitGroup
	wg.Add(2)
	increament("Foo:")
	increament("bar")

}

func increament(s string) {
	for i := 0; i < 20; i++ {
		// pausing for certain miliseconds
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		// locking
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter", counter)
		mutex.Unlock()
	}
	wg.Done()
}
