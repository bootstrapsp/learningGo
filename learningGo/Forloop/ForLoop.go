package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		//%d for decimal; %b is for binary; %x for hexadecimal; %q is for UTF8; /t is for tab; \n is for new line
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
