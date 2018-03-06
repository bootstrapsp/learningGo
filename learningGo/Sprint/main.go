package main

import "fmt"

func main() {
	// calling the func to accept the arguments and  printing out the return of the function to standard out
	fmt.Println(supGreeting("Boot", "strap"))

	// calling dualReturnFunc()
	fmt.Println(dualReturnFunc("Bla1", "blah2"))
}

// function with parameters and type string returning string

func supGreeting(fname, lname string) string {
	// uses default formatting; it is not printing out to the standard out (monitor) rather concatinating and returning the string
	return fmt.Sprint(fname, lname)
}

// func accepting two parameters, and is also doing two returns
func dualReturnFunc(string1, string2 string) (string, string) {
	return fmt.Sprint(string1, string2), fmt.Sprint(string2, string1)
}
