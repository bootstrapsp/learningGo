package main

import "fmt"

func returnTheBiggest(findBig ...int) int {
	var i int
	for _, j := range findBig {
		if j > i {
			i = j
		}
	}
	return i
}

func main() {
	data := returnTheBiggest(2, 4, 19, 4747)

	fmt.Println(data)
}
