package main

import "fmt"

func half(inputVal int) {
	var inpujjtVal = (inputVal / 2)

	if inpujjtVal == 0 {
		fmt.Println("Provided variable is even", inpujjtVal)
		fmt.Println(inpujjtVal)
	} else {
		fmt.Println("Provided variable is odd", inpujjtVal)
	}
}

func main() {
	half(4)
}
