package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var anyNumberVal float64
	fmt.Scan(&anyNumberVal)
	fmt.Println(roundingOff(anyNumberVal))
	fmt.Println(anyNumberVal)

	arrayToStoreUserInput()
}

/*
roundingOff takes float and
*/
func roundingOff(anynumber float64) float64 {
	n := rand.Float64()
	return n
}

func arrayToStoreUserInput() {

	var userInput string
	fmt.Println("Enter the user strings")
	fmt.Scan(&userInput)
	//l := len(userInput)

	var x [100]string

	for i := range x {
		fmt.Println(i)
	}
}
