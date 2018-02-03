package main

import "fmt"

// creating map and assining a value to the key, since map is key:value pair
func useMapForStorage() {

	m := map[int]string{}

	m[0] = "fun"
	fmt.Println(m)
}

// creating slice of slice of a certain datatype and then using append to insert the values

func useSliceForStorage() {

	a := []int{2, 3, 4, 5, 5}

	s := make([][]int, 0)

	s = append(s, a)

	fmt.Println(s)
}

func main() {

	useMapForStorage()
	useSliceForStorage()
}
