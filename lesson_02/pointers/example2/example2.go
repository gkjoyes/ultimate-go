// Sample program to show the basic concept of using a pointer to share data.
package main

import "fmt"

func main() {
	count := 10

	// Pass the address of count.
	increment(&count)

	fmt.Printf("Value of count: [%d] \t Address of count: [%p]\n", count, &count)
}

// increament delcares count as a pointer variable whose value is always and address.
//
//go:noinline
func increment(c *int) {
	*c++
	fmt.Printf("Value of c: [%d] \t Address of c: [%p]\n", *c, c)
}
