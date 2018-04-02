package main

import (
	"fmt"
	"log"
	"os"
)

// init function is called only once, before main is executed
//here with init we are creating a log file using OS package from go to create a file
func init() {
	nf, err := os.Create("generalLog.log")
	if err != nil {
		fmt.Println(err)
	}
	// having this pushes the log to the created log file else it goes to std out which default to terminal
	log.SetOutput(nf)
}

func main() {

	// thiis is bound to throw error since there is no such file
	_, err := os.Open("no-File.txt")

	// with this we are handling the error by adding some custom text and togther with err
	if err != nil {
		// log.Println logs to the log created in the init() above
		log.Println("err happened", err)
	}

}
