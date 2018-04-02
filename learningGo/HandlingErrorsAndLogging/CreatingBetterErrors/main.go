package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// doing init to create the log file

func init() {
	myLog, err := os.Create("myLog.log")
	if err != nil {
		fmt.Println(err)
	}
	// this is what directs the log to the log file
	log.SetOutput(myLog)
}

/*
Sqrt is custom method
*/
func Sqrt(f float64) (float64, error) {

	if f < 0 {
		return 0, errors.New("attempting to sqrt of negative number")
	}
	return 42, nil
}

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

	// out, err := creatingError(-123)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(out)

	// calling another func

	_, fullError := anotherAttempt(-1.2342)
	if err != nil {
		log.Fatalln(fullError)
	}
}

// ErrSampleIdiomaticErrorHandling is a variable containing specific error msg
var ErrSampleIdiomaticErrorHandling = errors.New("This is a way to define variable that highlights certain error and can be used anywhere")

func creatingError(n int) (string, error) {
	if n < 0 {
		log.Fatal(ErrSampleIdiomaticErrorHandling)
		return "Didn't work", ErrSampleIdiomaticErrorHandling
	}

	// nil here defines that there is no error
	return "worked", nil
}

func anotherAttempt(f float64) (float64, error) {
	if f < 0 {
		ErrNorGate := fmt.Errorf("Some error: %v", f)
		return 0, ErrNorGate
	}
	return 42, nil

}
