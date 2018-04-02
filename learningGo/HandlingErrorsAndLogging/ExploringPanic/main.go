package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("anotherlog.log")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)

}
func main() {

	_, err := os.Open("filethatdoesntexist.txt")

	if err != nil {
		// panic logs the stack related to the error together with the exit status, anything non 0 is pointing to an error
		panic(err)
	}
}
