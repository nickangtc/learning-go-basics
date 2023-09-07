package main

import "fmt"

// this is a type *alias*
// it is not a new type, but merely points to an existing type
type json = map[string]string

// here is a new type
// benefits:
// 1. it provides *semantics* on top of an existing type
// 2. it allows us to add *methods* to an existing type
type distance float64
type distanceMiles float64

// this is a type method
func (d distance) toMiles() distanceMiles {
	return distanceMiles(d * 0.621371)
}

func main() {
	// instantiate a new type
	d := distance(123)
	// call a type method
	d.toMiles()

	fmt.Println(d)
}
