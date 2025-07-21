// Sample program to show how to access an exported identifier.
package main

import (
	"fmt"

	"github.com/gkjoyes/ultimate-go/lesson4/exporting/example1/counters"
)

func main() {
	// Create a variable of the exported type and initialize the value to 10.
	counter := counters.AlertCounter(10)
	fmt.Println("Counter:", counter)
}
