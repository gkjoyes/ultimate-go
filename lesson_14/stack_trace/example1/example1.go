// Sample program to show how to read a stack trace when it packs values.
package main

func main() {
	example(make([]string, 4), "hello", 10)
}

//go:noinline
func example(slice []string, str string, i int) error {
	panic("Want stack trace")
}

// Use go build -gcflags=all="-N -l" to turn off inlining and optimizations so the stack trace has more accurate info.

/*
This shows a call to:

main.example(
    {0x14000054700, 0x4, 0x4},
    {0x1020b3b69, 0x5},
    0xa
)

Argument 1: {0x14000054700, 0x4, 0x4}
This is very likely a slice, represented as:
type slice struct {
    ptr *T     // 0x14000054700
    len int    // 0x4
    cap int    // 0x4
}

Argument 2: {0x1020b3b69, 0x5}
This is the structure of a Go string, which is:
type stringStruct struct {
    ptr *byte  // 0x1020b3b69
    len int    // 0x5
}

Argument 3: 0xa
Hex 0xa = 10 (decimal) — likely an int, uint, byte, etc.
*/
