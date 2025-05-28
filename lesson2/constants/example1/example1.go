// Sample program to show how to declare contants and their implementation in Go
package main

func main() {
	// Untyped Contants.
	const ui = 1234    // kind: integer
	const uf = 3.14159 // kind: floating-point

	// overflows uint8
	// const myUint8 uint8 = 1000

	// Typed Contants still use the constant type system but their precision is restricted.
	const ti int = 1234        // type: int
	const tf float64 = 3.14159 // type: float64

	// Kind promotion is used to  determine kind in these senarios.
	const answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	const third = 1 / 3.0    // KindFloat(1) / KindFloat(3.0)
	const zero = 1 / 3       // KindInt(1) / KindInt(3)

	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
}
