package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := Sqrt(-10.234234)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {

		// Errorf does formating, using "%v" while returning the error

		return 0, fmt.Errorf("invalid number: %v", f)
	}
	return 42, nil
}
