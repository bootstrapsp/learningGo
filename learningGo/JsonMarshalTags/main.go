package main

import (
	"encoding/json"
	"fmt"
)

//Thing a simple struct type, with tags defined under ``
type Thing struct {
	ID          int
	Name        string `json:"-"`
	IsConnected bool   `json:"is connected value"`
	description string
}

func main() {
	t1 := Thing{101, "DemoThing", false, "Demo thing"}
	bs, err := json.Marshal(t1)
	fmt.Printf("%T, /n", bs)
	fmt.Println(bs)
	// note when printing out the bs, value for description is not shown
	// as it is not exported (defined with "lower case" in the Struct)
	fmt.Println(string(bs))
	fmt.Println(err)
}
