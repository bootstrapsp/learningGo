package main

import "fmt"

func createMap() {

	// creating map key being int datatype and value stored being string
	var newMap = make(map[int]string)

	// setting value to the key
	newMap[1] = "Yo"

	fmt.Println(newMap)
}

func createMap2() {

	var newMap2 = map[string]int{}

	newMap2["Sushant"] = 1

	fmt.Println(newMap2)
}

func main() {
	createMap()
	createMap2()
}
