package main

import "fmt"

func main() {
	// %d is for decimal ; \t is for Tab for formatting; # is to put the representation of hexadecimal in front of the hexadecimle value generated out of the %X
	fmt.Printf("%d \t \t %b \t %#X", 42, 42, 42)
}
