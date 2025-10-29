// Sample program to show the syntax and mechanics of type switches and the empty interface.
package main

import "fmt"

func main() {
	// fmt.Println can be called with values of any type.
	fmt.Println("Hello, world")
	fmt.Println(1234)
	fmt.Println(3.14159)
	fmt.Println(true)
	fmt.Println("------------------------")

	// How can we do the same?
	myPrintln("Hello, world")
	myPrintln(1234)
	myPrintln(3.14159)
	myPrintln(true)
}

func myPrintln(val interface{}) {
	switch v := val.(type) {
	case string:
		fmt.Printf("Is string: type(%T): value(%s)\n", v, v)
	case int:
		fmt.Printf("Is int: type(%T): value(%d)\n", v, v)
	case float64:
		fmt.Printf("Is float64: type(%T): value(%f)\n", v, v)
	case bool:
		fmt.Printf("Is bool: type(%T): value(%t)\n", v, v)
	default:
		fmt.Printf("Is unknown: type(%T): value(%v)\n", v, v)
	}
}
