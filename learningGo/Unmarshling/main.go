package main

import (
	"encoding/json"
	"fmt"
)

// ThingTemplate is a struct consuming the Json
type ThingTemplate struct {
	ID          int
	Name        string
	IsConnected bool
}

func main() {
	var Thing1 ThingTemplate

	fmt.Println(Thing1.ID)
	fmt.Println(Thing1.Name)
	fmt.Println(Thing1.IsConnected)

	bs := []byte(`{"ID":"101", "Name":"ThingDemo","IsConnected":"true"}`)

	json.Unmarshal(bs, &Thing1)

	fmt.Println("----------")

	fmt.Println(Thing1.ID)
	fmt.Println(Thing1.Name)
	fmt.Println(Thing1.IsConnected)
	fmt.Printf("%T \n", Thing1)

}
