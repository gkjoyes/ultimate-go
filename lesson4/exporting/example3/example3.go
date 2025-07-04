// Sample program to show how the program can access a value of an unexported identifier from another package.
package main

import (
	"fmt"

	"github.com/gkjoyes/ultimate-go/lesson4/exporting/example3/counters"
)

func main() {
	// Create a variable of the unexported type using the exported New function from the package counters.
	counter := counters.New(10)

	fmt.Println("Counter:", counter)
}
