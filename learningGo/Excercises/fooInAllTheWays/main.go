package main

import "fmt"

func fooInAllTheWays(inputval ...int) {

	fmt.Println(inputval)
}

func main() {
	fooInAllTheWays(1, 2)

	strangeNumber := []int{1, 2, 3, 4, 5, 5, 5, 6, 6, 6, 6, 6, 7, 87, 6, 53, 44, 53, 45, 32, 2345, 234, 52, 345, 234, 5234}

	fooInAllTheWays(strangeNumber...)

}
