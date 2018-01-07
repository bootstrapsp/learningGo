package main

import "fmt"

func main() {
	// foo is a string variable with interpreted string literals defined within " "
	foo := "Demo"

	// foo1 is a string variable with raw string literals defined with ``
	// this prints out exactly the way string value has been assigned to the variable i.e. with line breaks and spaces
	foo1 := `dfsdfsdhfkj 
	sdfsdfsdf`

	fmt.Println(foo)
	fmt.Println(foo1)
}
