// Sample program to show how to read a stack trace when it packs values.
package main

func main() {
	example(true, false, true, 25)
}

func example(b1, b2, b3 bool, i uint8) error {
	panic("Want stack trace")
}

// Use go build -gcflags=all="-N -l" to turn off inlining and optimizations so the stack trace has more accurate info.

/*
panic: Want stack trace

goroutine 1 [running]:
main.example(0x1, 0x0, 0x1, 0x19)
        /Users/georgekuttyjoyes/work/go/src/github.com/gkjoyes/ultimate-go/lesson_14/stack_trace/example2/example2.go:11 +0x40
main.main()
        /Users/georgekuttyjoyes/work/go/src/github.com/gkjoyes/ultimate-go/lesson_14/stack_trace/example2/example2.go:7 +0x2c
*/

/*
Interpretation:

The current running goroutine is executing the function main.example.

It has been called with 4 raw argument values: 0x1, 0x0, 0x1, 0x19

These are raw values, not composite structs or pointers — so likely simple types like:
- int
- bool
- uintptr

Possible meanings of each value:

Argument 1: 0x1
Could be:
 - An int with value 1
 - A bool with value true

Argument 2: 0x0
Could be:
 - A false boolean
 - An int/uint with value 0

Argument 3: 0x1
 - Same as argument 1: probably true or 1

Argument 4: 0x19
 - Hex 0x19 = decimal 25
Could be:
 - An int, uint, or byte value
*/
