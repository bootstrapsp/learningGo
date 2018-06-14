package main

import (
	"fmt"
)

func main() {
	createAcronym()

}

func createAcronym() {
	userVal := "What The Fuck!"
	fmt.Println(userVal)

	fmt.Print(string(userVal[0:1])) // prints out the 1st letter in the 1st word

	var myVal []string
	myVal[100] = userVal
	fmt.Println(myVal)

}

/*
func funnyRegex() {
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("a(b", "seafood")
	fmt.Println(matched, err)

}
*/
