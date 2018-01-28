package main

import "fmt"

func main() {
	mapToRangeOver()
}

// creating the map
func mapToRangeOver() {

	// creating and initializing values in map
	createMap := map[int]string{
		0: "This",
		1: "is",
		2: "fun",
	}

	// range is used for iterating through any collection which is iterable
	for key, val := range createMap {
		fmt.Println(key, "is the key for the value", val)
	}
}
