// Sample program to show how the capacity of the slice is not available for use.
package main

import "fmt"

func main() {
	// Create a slice with a length 5 and capacity 8.
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// You can't access an index of a slice byond its length.
	fruits[5] = "Runtime error"

	fmt.Println(fruits)
}
