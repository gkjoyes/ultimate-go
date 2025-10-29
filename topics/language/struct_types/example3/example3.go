// Sample program to show how variable of an unnamed type can be assigned to variables of a named type, when they are identical.
package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of an anonymous type and initialize using a struct literal.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	// Create a variable of type example.
	var ex example

	// Assign the value of the unnamed struct type to the named struct type value.
	ex = e

	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", ex)
}
