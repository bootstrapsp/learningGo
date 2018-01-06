package main

import "fmt"

/*
- Complete program is showing usage of iota how its value resets when its used again in different function
- Also how const is used for defining constant values
- Also showing the bit shifting done using "<<"
*/
func one() {

	const (
		_ = iota
		a = 1 << (iota * 10)
		b = 1 << (iota * 10)
		c = 1 << (iota * 10)
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func two() {

	const (
		_ = iota
		c = 1 << (iota * 10)
		d = 1 << (iota * 10)
	)
	fmt.Println(c)
	fmt.Println(d)
}

func main() {
	one()
	two()

}
