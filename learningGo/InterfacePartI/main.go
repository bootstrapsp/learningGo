package main

import (
	"fmt"
	"sort"
)

// defining a custom type of datatype slice of string
type people []string

// implementing first function to attach to the type required for performing sort
func (p people) Len() int {
	return len(p)
}

// implementing second function to attach to the type required for performing sort
func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

// implementing third function to attach to the type required for performing sort

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Another example with string of slices
type thingNames []string

// function using the slice of string type defined above,
// initializing with some values and then sorting it with the sort's String method
func simpleSliceOfStringSorter() {
	things := thingNames{"Thing3", "Thing1", "Thing2"}

	fmt.Println(things)

	sort.Strings(things)

	fmt.Println(things)
}

// Another another example with the slice of int

type thingNumbers []int

func (t thingNumbers) Len() int {
	return len(t)
}

// function using sort pacakge to sort through the slice of int values
func simpleSliceofIntSorter() {
	tNumbers := thingNumbers{234, 23454, 3, 3, 3, 332, 34, 2345, 2456, 2345, 7, 8, 8, 55, 32, 45, 23, 524, 7457, 867, 9, 8, 435, 52, 36, 245, 425, 246, 6, 235, 235, 24, 57, 5, 624, 655673, 235, 6245, 634, 356, 5, 65, 565, 6, 456, 2, 662456, 45, 645, 6245, 66, 532, 245, 234, 2, 23, 53, 52, 345, 2345, 234, 256, 8, 8, 897, 98, 567, 6, 6, 5, 43, 34, 5, 67, 87654}

	fmt.Println(tNumbers)

	sort.Ints(tNumbers)

	fmt.Println(tNumbers)

	// doing reverse sort with this
	sort.Sort(sort.Reverse(sort.IntSlice(tNumbers)))

	fmt.Println(tNumbers)

}

// main function which instantiate values for slice of strings and then sorts using sort function
func main() {

	funGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(funGroup)
	sort.Sort(funGroup)
	fmt.Println(funGroup)

	// sorting slice of strings
	simpleSliceOfStringSorter()

	// sorting slice of ints
	simpleSliceofIntSorter()
}
