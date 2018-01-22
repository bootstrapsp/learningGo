package main

import "fmt"

// all the functions for creating and using slice in this example uses "make"

// this function creates slice to store slices of string
func sliceAppender() {
	// double brackets i.e. "[][]<datatype>" denotes that there this makes a slice which stores string slice
	sliceCollector := make([][]string, 0)

	record1 := make([]string, 4)
	record1[0] = "YO"
	record1[1] = "Man"
	record1[2] = "Blah"
	record1[3] = "This"

	// this is appending the string slice to the slice storing slice of string
	sliceCollector = append(sliceCollector, record1)
	fmt.Println(sliceCollector)

	record2 := make([]string, 20)

	record2[0] = "this"
	record2[1] = "is"
	record2[2] = "the"
	record2[3] = "new"
	record2[4] = "fun"

	sliceCollector = append(sliceCollector, record2)
	fmt.Println(sliceCollector)
}

func sliceAppender2() {
	sliceCollector2 := make([][]int, 0, 4)

	for i := 0; i < 4; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {

			transaction = append(transaction, j)
		}
		sliceCollector2 = append(sliceCollector2, transaction)
	}
	fmt.Println(sliceCollector2)
}

func main() {
	sliceAppender()
	sliceAppender2()
}
