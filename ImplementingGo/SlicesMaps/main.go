package main

import "fmt"

func main() {

	// method 1
	creatingFunnySlices(19)

	// method 2
	slice2d()

	// method 3
	slice3d()
}

// creating 1 dimensional slice
func creatingFunnySlices(n int) {

	s := make([]int, n)

	for i := 1; i < n; i++ {

		s = append(s, i)

	}
	fmt.Println("Finally getting the total insets in the slice of int", s)
}

// creating 2 dimensional slice
func slice2d() [][]string {

	myStringSlice := []string{"Demo Val insert"}

	s := make([][]string, 100) // this is a slice of slice of string

	s[27] = myStringSlice

	fmt.Println("checking value for S, after insert of slice", s)

	return s
}

// inserting 2 dimensional slice into 3 dimensional slice

func slice3d() {

	threeDSlice := make([][][]string, 100)

	threeDSlice[38] = slice2d()

	//fmt.Println("Printing out the val for the 3d slice containing function returning 2d slices", threeDSlice)

	fmt.Println("printing out 38", threeDSlice[38])

}
