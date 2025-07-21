// Sample program to show how arrays of different sizes are not of the same type.
package main

import "fmt"

func main() {

	// Declare an array of 5 integers that is initialized to its zero value.
	var five [5]int

	// Declare an array of 4 integers that is initialized with some values.
	four := [4]int{10, 20, 30, 40}

	// five = four
	// ./example5.go:14:9: cannot use four (variable of type [4]int) as [5]int value in assignment

	fmt.Println(five)
	fmt.Println(four)
}
