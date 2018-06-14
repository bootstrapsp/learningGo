package main

import (
	"fmt"
	"time"
)

func letsLoop() {
	defer fmt.Println("This comes previous to the switch loop but prints in the end because of defer")

	const i int = 10
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Time to work")
	}

}

func main() {
	letsLoop()
}
