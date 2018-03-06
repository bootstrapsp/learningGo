package main

import "fmt"

// defining the struct with certain fields
type person struct {
	first string
	last  string
	age   int
}

// creating function with a receiver which here is instantiating the object Person created above, fullName is the name of the function and it returns string
func (p person) fullName() string {
	return p.first + p.last
}

// putting in some values for the objects created out of the struct Person
// and then calling the method assigned to the struct, notice that now the objects created out of the Person can call the method fullName()
func main() {

	// fields aren't specified when supplying the value
	p1 := person{"S", "P", 31}
	p2 := person{"P", "S", 29}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

}
