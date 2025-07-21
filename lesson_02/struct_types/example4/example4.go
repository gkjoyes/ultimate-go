// Sample program to show how struct types align on boundaries.
package main

import (
	"fmt"
	"unsafe"
)

// No byte padding.
type nbp struct {
	a bool // 	1 byte
	b bool // 	1 byte
	c bool // 	1 byte
}

// Single byte padding.
type bp1 struct {
	a bool // 	1 byte
	// 			1 byte padding
	b int16 //  2 byte
}

// Three byte padding.
type bp3 struct {
	a bool //	1 byte
	// 			3 byte padding
	b int32 //	4 byte
}

// Seven byte padding
type bp7 struct {
	a bool // 	1 byte
	// 			7 byte padding
	b int64 // 	8 byte
}

// Eight byte padding on 64bit Arch.
type bp8 struct {
	a string //		16 bytes
	b int32  // 	4 bytes
	// 				4 byte padding
	c string // 	16 bytes
	d int32  // 	4 bytes
	// 				4 byte padding
}

// No padding on 32bit Arch.
type np32 struct {
	a string //	 8 bytes
	b int32  //	 4 bytes
	c string //  8 bytes
	d int32  //  4 bytes
}

// No padding.
type np struct {
	a string //	16 byte
	b string // 16 byte
	c int32  // 4 byte
	d int32  //	4 byte
}

func main() {
	var npb nbp
	size := unsafe.Sizeof(npb)
	fmt.Printf("Size Of npb is [%d][%p %p %p]\n", size, &npb.a, &npb.b, &npb.c)

	var bp1 bp1
	size = unsafe.Sizeof(bp1)
	fmt.Printf("Size Of bp1 is [%d][%p %p]\n", size, &bp1.a, &bp1.b)

	var bp3 bp3
	size = unsafe.Sizeof(bp3)
	fmt.Printf("Size Of bp3 is [%d][%p %p]\n", size, &bp3.a, &bp3.b)

	var bp7 bp7
	size = unsafe.Sizeof(bp7)
	fmt.Printf("Size Of bp7 is [%d][%p %p]\n", size, &bp7.a, &bp7.b)

	var bp8 bp8
	size = unsafe.Sizeof(bp8)
	fmt.Printf("Size Of bp8 is [%d][%p %p %p %p]\n", size, &bp8.a, &bp8.b, &bp8.c, &bp8.d)

	var np32 np32
	size = unsafe.Sizeof(np32)
	fmt.Printf("Size Of np32 is [%d][%p %p %p %p]\n", size, &np32.a, &np32.b, &np32.c, &np32.d)

	var np np
	size = unsafe.Sizeof(np)
	fmt.Printf("Size Of np is [%d][%p %p %p %p]\n", size, &np.a, &np.b, &np.c, &np.d)
}
