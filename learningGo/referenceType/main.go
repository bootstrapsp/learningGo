package main

import "fmt"

// golang is pass by value as represented below
func main() {
	age := 42
	fmt.Println(&age)

	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 24
	fmt.Println(z)
	fmt.Println(*z)

}
