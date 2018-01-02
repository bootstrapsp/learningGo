package scopes

import "fmt"

func doAdd() int {
	a := 1
	b := 2
	c := a + b

	fmt.Println(a + b)

	return c
}
