// Sample program to show how strings have a UTF-8 encoded byte array.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Declare a string with both chinese and english characters.
	s := "世界 means world"

	// UTFMax is 4, upto 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	for i, r := range s {

		// Capture the number of byte for this rune.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated with this rune.
		si := i + rl

		// Copy rune from the string to our buffer.
		copy(buf[:], s[i:si])

		fmt.Printf("%2d: %q\tcodepoint: [%#6x]\tencoded bytes: [%#v]\n", i, r, r, buf[:rl])
	}
}
