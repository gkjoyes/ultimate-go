// Sample program to show how the for range has both value and pointer semantics.
package main

import "fmt"

func main() {
	// Using the value semantic form of the for range.
	friends := []string{"x1", "x2", "x3", "x4", "x5"}
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("name: [%s]\n", v)
	}
	fmt.Printf("friends: [%v]\n", friends)

	fmt.Println("----------------------------------")

	// Using the pointer semantic for of the for range.
	friends = []string{"a1", "a2", "a3", "a4", "a5"}
	for i := range friends {
		friends = friends[:2]
		fmt.Printf("name: [%s]\n", friends[i])
	}
}
