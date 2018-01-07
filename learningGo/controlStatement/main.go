package main

import "fmt"

/*
Since there is no while / do while in Go, For with condition statement is used below
Using "%" to calculate remainder
*/
func main() {
	i := 0

	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
