// this is intentional example highlighting the error
// concerning the attempt to mix the typed and the untyped variables

package main

import "fmt"

func main() {
	x := 0303.3
	y := int64(3)

	fmt.Printf(x + y)
}
