package main

import "fmt"

// defining struct with some fields
type person struct {
	first string
	last  string
	age   int
}

// defining struct which is including the struct created above, this is called embedded type
type doubleZero struct {
	person
	licenseToKill bool
}

func main() {

	// this is one idiomatic way of creating values for the struct by listing out the fields and their corresponding values
	p1 := doubleZero{
		person: person{
			first: "S",
			last:  "P",
			age:   333333,
		},
		licenseToKill: true,
	}

	// this is another idiomatic way of creating values without listing out the fields
	p2 := doubleZero{person{"P", "S", 232}, false}

	fmt.Println(p1)
	fmt.Println(p2)

	// another example of accessing the individual fields and also seeing the promotion how struct doubleZero has rights to access Person struct
	fmt.Println(p1.licenseToKill)
	fmt.Println(p2.licenseToKill)
}
