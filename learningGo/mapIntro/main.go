package main

import "fmt"

/*
below are following ways how maps are could be created
Unline slice map can't be created just by writing something like below

"var myGreetings map[string]string"

this above mentioned procedure will point the map to a nil reference and since there are no append function to add to this like the slice assignment of the value will fail+
*/

func createMap() {

	// creating map key being int datatype and value stored being string
	var newMap = make(map[int]string)

	// setting value to the key
	newMap[1] = "Yo"

	fmt.Println(newMap)
}

func createMap2() {
	// another way to define map
	var newMap2 = map[string]int{}

	newMap2["Sushant"] = 1

	fmt.Println(newMap2)
}

func createMap3() {
	// another way on how to create map without using "var"
	newMap3 := map[int]float64{
		1: 2334.234,
		2: 3423.43,
		3: 0.0003,
	}
	fmt.Println(newMap3)

	// this is how we delete from map
	delete(newMap3, 3)

	fmt.Println(newMap3)
}

func main() {
	createMap()
	createMap2()
	createMap3()
}
