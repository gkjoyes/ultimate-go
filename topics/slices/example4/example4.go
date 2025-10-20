// Sample program to show how to grow a slice using the built-in function append and
// how append grows the capacity of the underlying array.
package main

import "fmt"

func main() {
	// Declare a nil slice of strings.
	var data []string

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~100k strings to the slice.
	for record := 0; record < 1e5; record++ {

		// Add values to the slice using append function.
		value := fmt.Sprintf("Record: %d", record)
		data = append(data, value)

		// When the capacity of the slice changes, calculate and display the percentage of changes.
		if lastCap != cap(data) {
			capChange := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)

			fmt.Printf("Address: [%p]\tIndex: [%d]\t\tCapacity[%d - %2.f%%]\n", &data[0], record, cap(data), capChange)
		}
	}
}
