// Sample program to show how the concrete value assigned to the interface is what is stored inside the interface.
package main

import "fmt"

// printer displays information.
type printer interface {
	print()
}

// cannon defines a cannon printer.
type cannon struct {
	name string
}

// print displays the printer's name.
func (c cannon) print() {
	fmt.Println("Printer Name:", c.name)
}

// epson defines a epson printer.
type epson struct {
	name string
}

// print displays the printer's name.
func (e *epson) print() {
	fmt.Println("Printer Name:", e.name)
}

func main() {
	// Create a cannon and epson printer.
	c := cannon{name: "c1"}
	e := epson{name: "e1"}

	// Add the printers to the collection using both value and pointer semantics.
	printers := []printer{
		c,  // Store a copy of the cannon printer value.
		&e, // Store a copy of the epson printer values's address.
	}

	// Change the name field for both printers.
	c.name = "c2"
	e.name = "e2"

	// Iterate over the slice of printers and call print against the copied interface value.
	for _, p := range printers {
		p.print()
	}

	// NOTE: When we store a value, the interface value has its own copy of the value. Changes to the original value will not be seen.
	// NOTE: When we store a pointer, the interface value has its own copy of the address. Changes to the original value will be seen.
}
