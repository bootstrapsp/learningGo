package main

import "fmt"

/*
 functions in this example uses the "var" approach to create slice

 slice created with var approach has no underlying datastructure

 therefore values can't simply be referenced using the index number, since its nil

 any attempts to create do such a referencing via the index number directly, will lead to index out of range error

 to insert "append" needs to be used

 this whole example shows short hand on how to create slice and append to it without using make
*/

func sliceWithoutMake() {

	student := []string{}
	students := [][]string{}
	student = append(student, "Sushant")
	fmt.Println(student)
	// appending the "[]string" which is "student" to the "[][]string" which is students
	students = append(students, student)
	fmt.Println(students)

	// just a variant on how above slice creation is working just that its created with var
	// class is a slice "[]" which stores slice "[]" that stores slice of string "[]string"
	var class [][][]string

	class = append(class, students)
	fmt.Println(class)
	fmt.Println("capacity for class slice", cap(class))
	fmt.Println("capacity for students slice", cap(students))
	fmt.Println("capacity for the student", cap(student))

	fmt.Println("length for class slice", len(class))
	fmt.Println("length for students slice", len(students))
	fmt.Println("length for the student", len(student))
}

func sliceWithIntandNoMake() {

	newSliceCollectorWithInt := []float64{}

	var userInput float64
	fmt.Println("value prior to the append", cap(newSliceCollectorWithInt))
	fmt.Println("value prior to the append", len(newSliceCollectorWithInt))

	fmt.Scanln(&userInput)
	newSliceCollectorWithInt = append(newSliceCollectorWithInt, userInput)

	fmt.Println("value after the append", cap(newSliceCollectorWithInt))
	fmt.Println("value after the append", len(newSliceCollectorWithInt))

	fmt.Println(newSliceCollectorWithInt)

}

func main() {
	sliceWithoutMake()
	sliceWithIntandNoMake()
}
