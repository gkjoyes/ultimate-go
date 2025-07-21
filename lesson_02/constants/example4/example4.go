/*
// Common durations. There is no definition for units of Day or larger to avoid confusion across
// daylight savings time zone transitions.

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
	now := time.Now()

	// Subtract 5 nanoseconds from now using a literal constant.
	nanoSeconds := now.Add(-5)

	// Subtract 5 seconds from now using a declared constants.
	const timeout = 5 * time.Second // time.Duration(5) * time.Duration(1000 * MilliSecond)
	seconds := now.Add(-timeout)

	// cannot use minsFive (variable of type int64) as time.Duration value in argument to now.Add
	// minsFive := int64(-5)
	// variable := now.Add(minsFive)

	fmt.Printf("Now: %v\n", now)
	fmt.Printf("Subtract 5 NanoSeconds: %v\n", nanoSeconds)
	fmt.Printf("Subtract 5 Seconds: %v\n", seconds)
}
