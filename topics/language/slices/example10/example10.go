// Sample program to show how to use a third index slice.
package main

import "fmt"

func main() {
	// Create a slice of strings with different types of fruits.
	slice := []string{"x1", "x2", "x3", "x4", "x5"}

	// Take a slice of slice. We want just index 2.
	slice1 := slice[2:3]
	inspectSlice(slice1)

	fmt.Println("-----------------------------")

	// Take a slice of just index 2 with a length and capacity of 1.
	slice2 := slice[2:3:3]
	inspectSlice(slice2)

	fmt.Println("----------------------------")

	// Append a new element which will create a new underlying array to increase capacity.
	slice2 = append(slice2, "y3")
	inspectSlice(slice2)
}

func inspectSlice(slice []string) {
	fmt.Printf("Length: [%d]\nCapacity: [%d]\n", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("[%d]\t[%p]\t[%s]\n", i, &slice[i], s)
	}
}
