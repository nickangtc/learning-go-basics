package main

import "fmt"

type location string

func (origin location) distanceTo(destination location) int {
	distance := 123

	log := fmt.Sprintf("%s is %dkm from %s", origin, distance, destination)
	fmt.Println(log)

	return distance
}

func main() {
	nyc := location("NYC")
	london := location("London")

	// 🤯 we effectively create type methods, like with JS prototypes
	nyc.distanceTo(london)
}
