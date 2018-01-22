package main

import "fmt"

// function creating slice defined with  []<dataType>
func slice1() {
	x := []float64{324234.23423, 232423.2342, 23423.2432, 23423.234, 23423.5556, 435634563.45634563456, 34564356.3456456}

	// this will check the type for the variable with "%T"
	fmt.Printf("%T \n", x)

	// this will start frm index value 2nd and show all the value prior to 4th place (excluding 4th place)
	fmt.Println("slicing the slice... by getting the values from index 2 to 4", x[2:4])

	// printing the length of the slice, Slice is dynamic in memory allocation
	fmt.Println(len(x))

}

func slice2() {

	greeting := []string{
		"yo",
		"this",
		"is",
		"new",
		"shit",
	}

	fmt.Print("[1:2]")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2]")
	fmt.Println(greeting[1:2])
}

func slice3() {

	// making the slice using "make" keyword of len 3
	// when only putting 1 number in there, which is 3 in this case
	// both len and capacity becomes the same

	// length : number of elements referred to by the slice
	// capacity : number of elements in the underlying array

	customer := make([]int, 3)
	customer[0] = 7
	customer[1] = 10
	customer[2] = 15

	fmt.Println(customer[0])
	fmt.Println(customer[1])
	fmt.Println(customer[2])

	// making the slice using "make" keyword of len 3 and capacity 5
	greeting := make([]string, 3, 5)

	greeting[0] = "Good Morning"
	greeting[1] = "You"
	greeting[2] = "MF"
	fmt.Println("Printing the length prior to appending", len(greeting))
	fmt.Println("Printing the capacity prior to appending", cap(greeting))

	fmt.Println(greeting[0])
	fmt.Println(greeting[1])
	fmt.Println(greeting[2])

	// "append" is used for adding to the slice, note if the slice had smaller length and capacity, append will increase it instead of causing out of bound index error
	greeting = append(greeting, "hahahah!!!")
	fmt.Println("Printing the length prior to appending", len(greeting))
	fmt.Println("Printing the capacity prior to appending", cap(greeting))
	fmt.Println(greeting[3])
}

func slice4() {
	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8}

	// this is adding all the data from "myOtherSlice" to "mySlice" using the "..."
	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}
func main() {
	slice1()
	slice2()
	slice3()
	slice4()
}
