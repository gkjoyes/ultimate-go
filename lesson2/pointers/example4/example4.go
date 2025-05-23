// Sample program to show how stacks grow/change.
package main

// Number of elements to grow each stack frame.
const size = 1024

func main() {
	s := "HELLO"

	stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size of the stack.
//
//go:noinline
func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)

	c++
	if c == 10 {
		return
	}

	stackCopy(s, c, a)
}
