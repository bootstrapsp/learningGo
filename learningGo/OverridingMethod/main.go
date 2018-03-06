package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	licenseToKill bool
}

// function assigned to the person struct
func (p person) Greetings() {
	fmt.Println("I'm just a regular peprson.")

}

// this function overrides the greetings() function also defined above for struct person
func (dz doubleZero) Greetings() {
	fmt.Println("Miss moneypenny, so good to see you")

}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := doubleZero{
		person: person{
			first: "Sushant",
			last:  "pandey",
			age:   23423,
		},
		licenseToKill: true,
	}
	// actual greetings function
	p1.Greetings()
	p2.Greetings()
	//executing the overriden greetings()
	p2.person.Greetings()
}
