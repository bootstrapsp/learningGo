package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ThingTemplate struct for all the base things :::: This also serves as the meta structure for consuming JSON data
type ThingTemplate struct {
	ID          int
	Name        string
	IsConnected bool
}

func main() {
	// creating an object of type ThingTemplate
	var T1 ThingTemplate
	// creating a NewReader object, which can take JSON data, in below case JSON data is simulated with string
	rdr := strings.NewReader(`{"ID":101,"Name":"DemoThing","IsConnected":false}`)
	// is decoding/parsing the value received from the reader above; values are parsed in the form of ThingTemplate struct defined above to consume the JSON data
	json.NewDecoder(rdr).Decode(&T1)

	fmt.Println(T1.ID)
	fmt.Println(T1.Name)
	fmt.Println(T1.IsConnected)
}
