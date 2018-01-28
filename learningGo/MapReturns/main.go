package main

import "fmt"

func mapCreator() {

	// way to create map
	bootstrap := map[int]int64{}

	fmt.Println(bootstrap)

	if bootstrap == nil {
		fmt.Println("nothing in the map")
	} else {
		fmt.Println("key value pair in the map is .. ", bootstrap)
	}

}

// another way to create map
func mapCreator2() {

	var bootstrap2 = make(map[string]float64)

	bootstrap2["this is fun"] = 123e9

	fmt.Println(bootstrap2)
}

// revising some slicing the bread ... haha

func breadSliceCreator() {

	// one way to create slice
	slice1 := make([]int, 2)

	slice1[0] = 1

	slice1[1] = 2

	fmt.Println("lenght and capacity of the slice1...", len(slice1), cap(slice1))

	slice1 = append(slice1, 3)

	fmt.Println("lenght and capacity of the slice1...", len(slice1), cap(slice1))

	slice1 = append(slice1, 4)

	fmt.Println("lenght and capacity of the slice1...", len(slice1), cap(slice1))

	slice1 = append(slice1, 55)

	fmt.Println("lenght and capacity of the slice1...", len(slice1), cap(slice1))

	fmt.Println(slice1)

	// adding above slice1 to slice of slice2

	slice2 := make([][]int, 10)

	slice2[0] = slice1

	fmt.Println("Putting slice into slice", slice2)

}

func main() {
	mapCreator()
	mapCreator2()
	breadSliceCreator()
}
