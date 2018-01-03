package main

import "fmt"

const (
	a = iota // to ignore 0
	b = iota
)

func main() {
	fmt.Println(a, b)
}
