package main

import (
	"fmt"
	"sort"
)

// perfomring reversal of the slice of string
func main() {
	s := []string{"asd", "ozjit", "SVDF", "yzhebg"}
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

}
