// Sample program to show how to declare and initialize struct types.
package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of type example and set to its zero value.
	var e1 example
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and initialize using a struct literals.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}
	fmt.Printf("%+v\n", e2)
}
