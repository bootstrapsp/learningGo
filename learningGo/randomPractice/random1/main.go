package main

import "fmt"

func annonymousFuncDemo() {

	var userInput string

	fmt.Scanln("Enter your response", &userInput)

	if userInput == "y" || userInput == "true" {

		var getInt64Val int64
		fmt.Scanln(&getInt64Val)

		doForMethod = getInt64Val
	}

}

func main() {
	annonymousFuncDemo()
}
