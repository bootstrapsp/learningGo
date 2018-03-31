package main

import (
	"fmt"
	"strings"
)

func main() {
	factorialCal(8)
	combinationFactorial("demo")
	uniqueStringFactorial("demo")
}

// function for calculating factorial
func factorialCal(usrIn int) int {
	total := 1

	for j := usrIn; j > 0; j-- {

		total = total * j

		fmt.Println(total)

	}

	return total
}

// combinationFactorial takes strings and calculate in how many combinations can the letters in that word be arranged
func combinationFactorial(usrString string) int {

	var myVal int

	myVal = len(usrString)

	myVal = factorialCal(myVal)

	return myVal
}

// func takes strings and claculate in how many unique combinations can the letters in that word be arranged

func uniqueStringFactorial(userIn string) int {

	// using strings.Split function to split the word
	letters := strings.Split(userIn, "")
	fmt.Printf("%T \n", letters)

	var duplicate int

	// outer loop to run through the slice of letters
	for i := 0; i < len(letters); i++ {

		// inner loop to run through the slice of letters
		for j := i + 1; j < len(letters); j++ {

			// checking for duplicate by comparing outer to inner loop
			if letters[i] == letters[j] {
				duplicate = duplicate + 1
				fmt.Println("duplicate number of letters", duplicate)

			}

		}

	}

	if duplicate != 0 {
		resultWithDuplicate := (factorialCal(len(letters)) / factorialCal(duplicate))
		fmt.Println("Word found to have duplicate so factoring it by dividing it with the duplicate number of letters", resultWithDuplicate)
		return resultWithDuplicate

	}
	resultWithOutDuplicate := factorialCal(len(letters))
	fmt.Println("Word was not found to have any duplicates, returning its factor", resultWithOutDuplicate)
	return resultWithOutDuplicate

}

/*
* Note: The problem with this code is that it's counting the number of duplicates rather only counting the that many letters are duplicate e.g. if the number of times s is appears in word physics this code only counts 1 towards s appears twice, since of counting 2 since s appears twice

 */
