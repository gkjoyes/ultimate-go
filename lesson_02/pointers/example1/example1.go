// Sample program to show the basic concept of pass by value.
package main

import "fmt"

func main() {
	count := 10

	// Pass value of the count.
	increment(count)

	fmt.Printf("value of count: [%d]\t address of count: [%p]\n", count, &count)
}

// increment delcares count as a variable that hold actual value.
//
//go:noinline
func increment(c int) {
	c++
	fmt.Printf("value of c: [%d]\t address of c: [%p]\n", c, &c)
}
