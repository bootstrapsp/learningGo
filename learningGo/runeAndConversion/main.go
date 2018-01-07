package main

import "fmt"

func main() {

	for i := 500; i <= 505; i++ {
		//fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
		// defining rune with '' quotation marks around a
		// this prints out the number for the alias for "i" in the UTF8 chart
		foo := 'i' // this is 105 in the UTF8 mapping
		fmt.Println(foo)

		// following prints out the alias for the rune which is int32
		fmt.Printf("%T \n", rune(foo))
	}
}
