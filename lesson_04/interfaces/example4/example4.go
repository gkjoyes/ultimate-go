// Sample program to show you can't always get the address of a value.
package main

import "fmt"

// duration is a named type with a base type of int.
type duration int

func (d *duration) notify() {
	fmt.Println("Sending notification in", *d)
}

func main() {
	// NOTE:  cannot call pointer method on duration(42)
	// NOTE: cannot take the address of duration(42)
	// duration(42).notify()
}
