package main

import "fmt"

/*
* Calling all functions created in this
 */
func main() {
	revSlice()
}

func revSlice() {
	var defineSliceSize int64
	fmt.Println("enter the size of slice")
	fmt.Scanln(&defineSliceSize)
	// initialized slice
	s := make([]int64, defineSliceSize)

	var userans string
	if userans == "No" {
		fmt.Println("Not gonna insert anymore")
	} else {
		var usersInt int64
		fmt.Println("enter the value")
		fmt.Scan(&usersInt)
		s[10] = usersInt
		fmt.Println(s[10])
	}

	// ranging over slice and checking for values over or below 0
	for i := range s {
		if s[i] < 1 {
			fmt.Println("slice value is smaller than 0")
		} else {
			fmt.Println("Values are larger than 0", s[i])
		}
	}

}
