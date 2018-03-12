package main

import (
	"encoding/json"
	"fmt"
)

// Thing is an exported struct type defining somde properties for a Thing
type Thing struct {
	ID          int
	Name        string
	IsConnected bool
	description string
}

/* main function to perform following
* create a Thing using the struct type defined for it
* use Json Marshal() function to print out the byte slice
* and also with string() we can covnert the byte slice to see actual json value
 */
func main() {
	t1 := Thing{101, "DemoThing101", false, "this is fun Thing."}

	bs, err := json.Marshal(t1)
	fmt.Println(bs)
	fmt.Println(err)
	// for checking the type for the bs value defined to store the byte slice
	fmt.Printf("%T", bs)

	fmt.Println(string(bs))
}
