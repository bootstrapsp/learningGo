package main

import "fmt"

func main() {

	if return1() != 10 {
		fmt.Println("Return1 is greater than 10")
	} else if return2() > 10 {
		fmt.Println("Return2 is greater than 10 ")
	}
}

func return1() int {
	var newVal int
	fmt.Scanln(&newVal)
	return newVal
}

func return2() int {
	var newVal2 int
	fmt.Scanln(&newVal2)
	return newVal2
}
