package main

import "fmt"

type customType struct {
	fname string
	lname string
}

func (p customType) speakFunc() {
	fmt.Println(p.fname, p.lname, "says. Good Morning, James Bond.")
}

func main() {

	// creating slice
	xi := []int{2, 34, 5, 2, 23}
	fmt.Println("Printing out slice", xi)

	// creating map
	m := map[string]int{
		"Sushant": 23,
		"Pandey":  2323,
	}
	fmt.Println("Printing out the map", m)

	p1 := customType{
		"miss",
		"moneypenny",
	}
	fmt.Println(p1)

	p1.speakFunc()
}
