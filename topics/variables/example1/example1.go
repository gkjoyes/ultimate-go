// Sample program to show how to declare variables.
package main

import "fmt"

func main() {
	// Declare variables that are set to their zero value.
	var (
		a int
		b string
		c float64
		d bool
	)

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)

	// Declare variable and initialize.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("\n")
	fmt.Printf("var aa int \t %T [%v]\n", aa, aa)
	fmt.Printf("var bb string \t %T [%v]\n", bb, bb)
	fmt.Printf("var cc float64 \t %T [%v]\n", cc, cc)
	fmt.Printf("var dd bool \t %T [%v]\n", dd, dd)

	// Specify type and perfrom conversion.
	aaa := int32(10)

	fmt.Printf("\n")
	fmt.Printf("var aaa int32 \t %T [%v]\n", aaa, aaa)
}
