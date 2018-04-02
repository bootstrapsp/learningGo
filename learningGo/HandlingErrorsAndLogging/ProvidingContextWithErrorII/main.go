package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.23)

	if err != nil {
		log.Println(err)
	}
}

// ThingShapeError us a custom of type struct, implementing error
type ThingShapeError struct {
	name string
	id   int
	err  error
}

// here Error() function is attached to the type pointer to ThingShapeError type
func (t *ThingShapeError) Error() string {
	return fmt.Sprintf("Some error: %v %v %v", t.name, t.id, t.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {

		// apart from normal Errorf, also returning &ThingShapeError
		thing := fmt.Errorf("Some random error: %v", f)
		return 0, &ThingShapeError{"DemoName", 101, thing}
	}

	return 42, nil
}
