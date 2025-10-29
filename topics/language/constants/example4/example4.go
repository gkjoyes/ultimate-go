/*
type Duration int64

const (
	NanoSecond Duration = 1
	MicroSecond  		= 1000 * NanoSecond
	MilliSecond 		= 1000 * MicroSecond
	Second				= 1000 * MilliSecond
	Minute				= 60 * Second
	Hour				= 60 * Minute
)
*/

// Sample program to show how literal, constant and variables work within the scope of implicit conversion.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Use the time package to get the current time/date.
	now := time.Now()

	// Subtract 5 nanoseconds from now using a literal constant.
	nanoSeconds := now.Add(-5)

	// Subtract 5 seconds from now using a declared constants.
	const timeout = 5 * time.Second
	seconds := now.Add(-timeout)

	/*
		// Subtract 5 nanoseconds from now using a variable of type int64.
		minusFive := int64(-5)
		variable := now.Add(minusFive)
	*/

	fmt.Printf("Now: \t%v\n", now)
	fmt.Printf("Subtract 5 NanoSeconds: \t%v\n", nanoSeconds)
	fmt.Printf("Subtract 5 Seconds: \t%v\n", seconds)
}
