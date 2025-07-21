// Sample program to show how to take slices of slices to create different views and make changes to the underlying array.
package main

import "fmt"

func main() {
	// Create a slice with a length of 5 elements and capacity of 8.
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"
	inspectSlice(slice1)

	// Take a slice of slice1. We want just indexes 2 and 3.
	slice2 := slice1[2:4]
	inspectSlice(slice2)

	// Change the value of the index 0 of slice2.
	slice2[0] = "CHANGED"
	inspectSlice(slice1)
	inspectSlice(slice2)

	// Make a new slice big enough to hold elements of slice1 and copy the values over using the builtin copy function.
	slice3 := make([]string, len(slice1))
	copy(slice3, slice1)
	inspectSlice(slice3)
}

func inspectSlice(slice []string) {
	fmt.Printf("Length: [%d] Capacity: [%d]\n", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("[%d]\t[%p]\t[%s]\n", i, &slice[i], s)
	}
}
