package main

import "fmt"

func assertVal() {
	// creating a variable with empty interface{} as type
	var entityName interface{} = "String value"

	// performing assertion with .(string)
	anyVal, ok := entityName.(string)

	if ok {
		fmt.Printf("%T\n", anyVal)
	} else {
		fmt.Printf("Value is not a string")
	}
}

func anotherAssertion() {

	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	fmt.Printf("%T\n", val.(int))
}

func main() {
	assertVal()
	anotherAssertion()
}
