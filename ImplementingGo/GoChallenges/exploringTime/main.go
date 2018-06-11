package main

import (
	"fmt"
	"time"
)

func main() {
	whatTimeIsNow()
	checkDuration()
}

func whatTimeIsNow() {

	currMin := time.Now() // timestamp of the system

	fmt.Println(currMin)
}

func checkDuration() {

	sec := time.Second

	fmt.Print(int64(sec / time.Millisecond)) // converting seconds into miliseconds
}
