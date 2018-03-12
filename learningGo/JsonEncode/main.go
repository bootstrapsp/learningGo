package main

import (
	"encoding/json"
	"os"
)

// ThingTemplate defined with struct
type ThingTemplate struct {
	ID          int
	Name        string
	IsConnected bool
}

func main() {

	ThingDemo := ThingTemplate{101, "Thing101", false}

	json.NewEncoder(os.Stdout).Encode(ThingDemo)

}

/*
* New encoder takes a writer and returns a pointer to the encoder
* If we get pointer to the endcoder so it can be Encoded
 */
