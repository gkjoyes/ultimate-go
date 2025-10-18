// Sample program to show how to declare constants and their implementation in Go.
package main

import "fmt"

func main() {
	// Untyped Constants.
	{
		const ui = 1234    // kind: integer
		const uf = 3.14159 // kind: floating-point
	}

	// Typed Constants still use the constant type system but their precision is restricted.
	{
		const ti int = 1234        // type: int
		const tf float64 = 3.14159 // type: float64
	}

	// Kind promotion is used to  determine kind in these scenarios.
	{
		var answer = 3 * 0.333 // float64(answer) = KindFloat(3) * KindFloat(0.333)
		fmt.Printf("var answer type [%T] value [%v]\n", answer, answer)

		const third = 1 / 3.0 // KindFloat(third) = KindFloat(1) / KindFloat(3.0)
		fmt.Printf("const third type [%T] value [%v]\n", third, third)

		const zero = 1 / 3 // KindInt(zero) = KindInt(1) / KindInt(3)
		fmt.Printf("const zero type [%T] value [%v]\n", zero, zero)
	}

	// This is an example of constant arithmetic between typed and untyped constants.
	{
		const one int8 = 1
		const two = 2 * one // int8(2) * int8(1)
	}
}
