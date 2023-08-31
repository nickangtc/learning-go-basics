# Types in Go

There's type, then there's type alias.

Type:

```go
type location string
```

Type alias:

```go
type location = string
type sliceOfStrings = []string
```

## Types can have methods like classes

There are no classes in Go. No objects, no classes.

But similar functionality can be achieved by creating type methods:

```go
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

```
