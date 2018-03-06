package main

import "fmt"

// struct allows users to create own type, which is based of Go's own type
// struct is an AGGREGATE TYPE, because we agregate fields into it
// this corresponds to states in object oriented programming
type structIntro struct {
	my     int64
	fun    float64
	goTime string
}

type structIntroII struct {
	structIntro
	field string
}

// func for using struct by creating 2 variables that are using struct defined above with 3 custom types

func main() {

	// class is not required to be initialized to use struct
	// Go use TYPE which is struct and with this variables are instantiated
	entry1 := structIntro{42, 00000.23, "Random value"}

	entry2 := structIntro{43, 0000234.234, "New Random Value"}

	fmt.Println(entry1)
	fmt.Println(entry2)

	fmt.Println(entry1.fun)
	fmt.Println(entry1.goTime)
	fmt.Println(entry1.my)
}
