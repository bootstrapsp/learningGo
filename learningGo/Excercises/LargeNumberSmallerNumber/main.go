package main

import "fmt"

func main() {
	var snum int64
	var lnum int64

	fmt.Println("Enter large number")
	fmt.Scanln(&lnum)

	fmt.Println("Enter small number")
	fmt.Scanln(&snum)

	//var remainder float64
	//remainder = lnum % snum

	fmt.Println("The reminder is ... ", (lnum % snum))
}
