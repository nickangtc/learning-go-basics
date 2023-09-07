package main

import "fmt"

// NOTE: the inline comment has special powers, shows up on hover in VSCode
type Duration float64 // in hours

// CONCEPT: this is a structure, aka a struct
// you need to use this when your type consists of multiple fields
// NOTE: User is exported, but if another package needs to call the constructor,
// the field names need to be exported as well (i.e. Id instead of id)
type User struct {
	id   int
	name string
	duration Duration
}

type UserKaputt struct {
	id   int
	name string
	// 🚫 will not work - methods on structs need to be defined outside of struct definition
	greet() string {
		return "Hello, " + u.name
	}
}
// ✅ this will work as type method
func (u User) greet() string {
	return "Hello, " + u.name
}

/**
* NOTE: curly braces = invoking the constructor
 */
func main() {
	// option 0: construct an empty struct
	// field values will be set to `nil`
	var u User
	u0 := User{}

	// option 1: construct a struct w/ field names
	// this is the preferred way
	u1 := User{id: 1, name: "John"}

	// option 2: construct a struct w/o field names
	// don't do this, it's fragile,
	// as it relies on the order of the fields in struct definition
	u2 := User{2, "Jane"}

	fmt.Println("u", u)   // {0 }
	fmt.Println("u0", u0) // {0 }
	fmt.Println("u1", u1) // {1 John}
	fmt.Println("u2", u2) // {2 Jane}
}
