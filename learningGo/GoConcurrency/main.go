package main

import "fmt"

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("foo ", i)
	}

}

func bar() {
	for j := 0; j < 100; j++ {
		fmt.Println("bar ", j)
	}

}

// having go in front starts a separate thread of go routine for concurrent execution of threads
func main() {
	go foo()
	go bar()
}
