// Sample program to show how to decalre and initialize anonymous struct types.
package main

import "fmt"

func main() {
	// Declare a variable of an anonymous type and set to its zero value.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymouse type and init using a struct literal.
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}
	fmt.Printf("%+v\n", e2)
}
