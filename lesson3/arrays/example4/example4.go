// Sample progarm to show how the for range has both value and pointer semantics.
package main

import "fmt"

func main() {
	users := [5]string{"x1", "x2", "x3", "x4", "x5"}

	// Using the pointer semantic form of the for range.
	fmt.Printf("Before: [%s]\n", users[1])
	for i := range users {
		users[1] = "m2"

		if i == 1 {
			fmt.Printf("After: [%s]\n", users[1])
		}
	}

	users = [5]string{"x1", "x2", "x3", "x4", "x5"}

	// Using the value semantic form of the for range.
	fmt.Printf("Before: [%s]\n", users[1])
	for i, u := range users {
		users[1] = "m2"

		if i == 1 {
			fmt.Printf("After: [%s]\n", u)
		}
	}

	users = [5]string{"x1", "x2", "x3", "x4", "x5"}

	// Using the value semantic form of the for range but with pointer semantic access. DON'T DO THIS!!!
	fmt.Printf("Before: [%s]\n", users[1])
	for i, u := range &users {
		users[1] = "m2"

		if i == 1 {
			fmt.Printf("After: [%s]\n", u)
		}
	}
}
