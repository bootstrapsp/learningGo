package main

import (
	"fmt"
	"regexp"
)

func main() {
	createAcronym()
	funnyRegex()
}

func createAcronym() {
	var userVal string
	fmt.Scan(&userVal)
	fmt.Println(len(userVal)) // printing out the length
}

func funnyRegex() {
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("a(b", "seafood")
	fmt.Println(matched, err)
}
