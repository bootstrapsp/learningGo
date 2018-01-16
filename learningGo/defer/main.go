package main

import "fmt"

func blah1() {
	fmt.Println("Yo this is going on like ... ")
}

func blah2() {
	fmt.Println("forever")
}

func main() {

	blah2()

	// defer pushes the execution of the function to the very end of the function under which its called; basically right before that function is exited
	defer blah1()

}
