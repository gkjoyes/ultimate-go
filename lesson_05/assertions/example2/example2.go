// Sample program demonstrating that type assertions are a runtime and not compile time construct.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// car represents something you drive.
type car struct {
	name string
}

// String implements the fmt.Stringer interface.
func (car) String() string {
	return "Hi Honda!"
}

// cloud represents somewhere you store information.
type cloud struct {
	name string
}

func (cloud) String() string {
	return "Big Data!"
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a slice of the Stringer interface values.
	vals := []fmt.Stringer{
		car{name: "x1"},
		cloud{name: "aws"},
	}

	// Let's run this experiment ten times.
	for i := 0; i < 10; i++ {
		// Choose a random number from 0 to 1.
		n := r.Intn(2)

		// Perform a type assertion that we have a concrete type of cloud in the interface value we randomly choose.
		if v, ok := vals[n].(cloud); ok {
			fmt.Println("Got Lucky:", v)
			continue
		}
		fmt.Println("Got Unlucky")
	}
}
