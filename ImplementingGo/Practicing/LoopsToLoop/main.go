package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	// simple for loop
	var i float64
	for i = 0; i < 10; i++ {
		randomSqrt(i)

	}

}

var wg sync.WaitGroup

func randomSqrt(f float64) {

	wg.Add(1)
	go func() float64 { // go routine for cal sqrt

		m := math.Sqrt(f) // built in math func
		fmt.Println("square root :", m)
		return m

	}()
	wg.Done()
	// wait group to let go routine finish
	go func() {
		wg.Wait()

	}()

}
